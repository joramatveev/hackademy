package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"time"
)

type ReqFib struct {
	Number int
}

type RespFib struct {
	Fib    *big.Int
	Time   time.Duration
	Number int
}

var cache = make(map[int]*big.Int, 0)

func fib(num int) *big.Int {
	if num < 0 {
		panic("num < 0")
	}
	if val, true := cache[num]; true {
		return val
	}
	cacheSize := len(cache)
	f1 := cache[cacheSize-2]
	f2 := cache[cacheSize-1]
	for i := cacheSize; i <= num; i++ {
		f1.Add(f1, f2)
		cache[i] = f1
		f1, f2 = f2, f1
	}
	return f2
}

func handleError(err error) error {
	if err == nil {
		return nil
	}
	fmt.Println("Fail: ", err.Error())
	return err
}

const host string = "127.0.0.1"
const port string = "48999"

func main() {
	cache = make(map[int]*big.Int)
	cache[0] = big.NewInt(0)
	cache[1] = big.NewInt(1)

	fmt.Println("Starting Fibonacci server ...")

	ln, err := net.Listen("tcp", host+":"+port)
	err = handleError(err)
	if err != nil {
		return 
	}
	fmt.Printf("Opened server on %v:%v\n", host, port)
	conn, err := ln.Accept()
	err = handleError(err)
	if err != nil {
		return 
	}

	for {
		var req ReqFib
		decoder := json.NewDecoder(conn)
		if handleError(decoder.Decode(&req)) != nil {
			break
		}

		start := time.Now()
		resp := RespFib{
			Number: req.Number,
			Fib:    fib(req.Number),
			Time:   time.Since(start),
		}

		enc := json.NewEncoder(conn)
		if handleError(enc.Encode(resp))!= nil {
			break
		}
	}
}
