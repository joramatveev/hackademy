package main

import (
	"net/http"
	"strings"
)

type AuthHandler func(rw http.ResponseWriter, r *http.Request, u User)

func (j *JWTService) jwtAuth(ur UserRepository, h AuthHandler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		auth, err := j.ParseJWT(token)
		if err != nil {
			rw.WriteHeader(401)
			_, err := rw.Write([]byte("unauthorized"))
			if err != nil {
				return
			}
			return
		}

		err = ur.IsBanned(auth.Email)
		if err != nil {
			rw.WriteHeader(401)
			_, err := rw.Write([]byte(err.Error()))
			if err != nil {
				return
			}
			return
		}

		err = ur.CheckNotInDB(token)
		if err != nil {
			rw.WriteHeader(401)
			_, err := rw.Write([]byte(err.Error()))
			if err != nil {
				return
			}
			return
		}

		user, err := ur.Get(auth.Email)
		if err != nil {
			rw.WriteHeader(401)
			_, err := rw.Write([]byte("unauthorized"))
			if err != nil {
				return
			}
			return
		}

		h(rw, r, user)
	}
}
