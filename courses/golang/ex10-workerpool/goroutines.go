package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func Run(poolSize int) {

	jobs := make(chan float64, poolSize)
	s := bufio.NewScanner(os.Stdin)

	var wg sync.WaitGroup
	w := 1
	for s.Scan() {
		input := string(s.Bytes())
		f, err := strconv.ParseFloat(input, 64)
		handleError(err)
		jobs <- f

		if w <= poolSize {
			wg.Add(1)
			go worker(w, jobs, &wg)
			w++
		}
	}
	close(jobs)
	wg.Wait()
}

func worker(w int, jobs <-chan float64, wg *sync.WaitGroup) {
	fmt.Printf("worker:%d spawning\n", w)
	for j := range jobs {
		fmt.Printf("worker:%d sleep:%.1f\n", w, j)
		time.Sleep(time.Millisecond * time.Duration(int(1000*j)))
	}
	fmt.Printf("worker:%d stopping\n", w)
	wg.Done()
}

func handleError(err error) {
	if err == nil {
		return
	}
	fmt.Println("Fail: ", err)
}
