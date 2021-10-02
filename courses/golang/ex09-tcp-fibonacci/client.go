package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"os"
	"strconv"
	"time"
)

type ReqFib struct {
	Number int
}

type RespFib struct {
	Number int
	Fib    *big.Int
	Time   time.Duration
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
	conn, err := net.Dial("tcp", host+":"+port)
	err = handleError(err)
	if err != nil {
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			err := handleError(err)
			if err != nil {
				return
			}
		}
	}(conn)
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		num, err := strconv.ParseInt(scan.Text(), 10, 64)
		err = handleError(err)
		if err != nil {
			break
		}
		req := ReqFib{Number: int(num)}
		encoder := json.NewEncoder(conn)
		err = handleError(encoder.Encode(req))
		if err != nil {
			break
		}
		var resp RespFib
		decoder := json.NewDecoder(conn)
		err = handleError(decoder.Decode(&resp))
		if err != nil {
			break
		}
		fmt.Printf("%s %d\n", resp.Time, resp.Fib)
	}
}
