package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func getCakeHandler(w http.ResponseWriter, _ *http.Request, u User) {
	_, err := w.Write([]byte(u.FavoriteCake))
	if err != nil {
		return
	}
}

func wrapJWT(
	jwt *JWTService,
	f func(http.ResponseWriter, *http.Request, *JWTService),
) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		f(rw, r, jwt)
	}
}

func main() {
	r := mux.NewRouter()

	userService := UserService{
		repository: NewInMemoryUserStorage(),
	}

	jwtService, err := NewJWTService("pubkey.rsa", "privkey.rsa")
	if err != nil {
		panic(err)
	}

	r.HandleFunc("/user/me", logRequest(jwtService.jwtAuth(userService.repository, getCakeHandler))).Methods(http.MethodGet)
	r.HandleFunc("/user/favorite_cake", logRequest(jwtService.jwtAuth(userService.repository, userService.UpdateCake))).Methods(http.MethodPut)
	r.HandleFunc("/user/email", logRequest(jwtService.jwtAuth(userService.repository, userService.UpdateEmail))).Methods(http.MethodPut)
	r.HandleFunc("/user/password", logRequest(jwtService.jwtAuth(userService.repository, userService.UpdatePassword))).Methods(http.MethodPut)

	r.HandleFunc("/user/register", logRequest(userService.Register)).Methods(http.MethodPost)
	r.HandleFunc("/user/jwt", logRequest(wrapJWT(jwtService, userService.JWT))).Methods(http.MethodPost)

	r.HandleFunc("/admin/ban", logRequest(jwtService.jwtAuth(userService.repository, userService.BanUser))).Methods(http.MethodPost)
	r.HandleFunc("/admin/unban", logRequest(jwtService.jwtAuth(userService.repository, userService.UnBanUser))).Methods(http.MethodPost)
	r.HandleFunc("/admin/inspect", logRequest(jwtService.jwtAuth(userService.repository, userService.History))).Methods(http.MethodGet)

	r.HandleFunc("/admin/promote", logRequest(jwtService.jwtAuth(userService.repository, userService.Promote))).Methods(http.MethodPost)
	r.HandleFunc("/admin/fire", logRequest(jwtService.jwtAuth(userService.repository, userService.Fire))).Methods(http.MethodPost)

	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		<-interrupt
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := srv.Shutdown(ctx)
		if err != nil {
			return
		}
	}()

	log.Println("Server started, hit Ctrl+C to stop")
	err = srv.ListenAndServe()
	if err != nil {
		log.Println("Server exited with error:", err)
	}

	log.Println("Good bye :)")
}
