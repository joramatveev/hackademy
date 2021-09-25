package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Visitor struct {
	name string
}

var wg sync.WaitGroup

var customersAtAll = 10
var numberOfWaitingVisitors = 0
var barberSleep = false
var wasCustomers = 0
var visitorsVisitedBarber []Visitor

func main() {

	fmt.Print("\nWelcome! Barber shop opened!\n")

	fmt.Print("Start preparing seats...\n")
	time.Sleep(time.Second * 3)
	fmt.Print("Waiting seats were prepared.\n")

	visitorsWaiting := make(chan Visitor, customersAtAll)

	wg.Add(1)
	go comeIn(visitorsWaiting)

	wg.Add(1)
	go cutting(visitorsWaiting)

	wg.Wait()

	defer fmt.Print("\nHuh, it's a tough day today! Come back tomorrow...")
}

func comeIn(visitorsWaiting chan Visitor) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < customersAtAll; i++ {
		var nextVisitorAt = rand.Intn(5)
		time.Sleep(time.Second * time.Duration(nextVisitorAt))

		mensNames := map[int]string{
			0:  "Vova",
			1:  "Pash",
			2:  "Vitalic",
			3:  "Dima",
			4:  "Sergo",
			5:  "Nazar",
			6:  "Ivan",
			7:  "Misha",
			8:  "Grisha",
			9:  "Valentin",
			10: "Mark",
			11: "Makar",
			12: "Viсtor",
			13: "Igor",
			14: "Yura",
			15: "Maxim",
			16: "Kiril",
			17: "Daniil",
			18: "Andrey",
			19: "Sasha",
		}

		visitorName := rand.Intn(len(mensNames))

		if customerAlreadyVisited(string(mensNames[visitorName])) {
			i--
		} else if numberOfWaitingVisitors < 5 {
			visitorsWaiting <- Visitor{string(mensNames[visitorName])}
			fmt.Print(string(mensNames[visitorName]), " has gone and waiting\n")
			numberOfWaitingVisitors++
		} else {
			fmt.Print(string(mensNames[visitorName]), " has not found a seat so he/she has gone\n")
		}
		wasCustomers++
	}
}

func customerAlreadyVisited(customerName string) bool {
	for _, v := range visitorsVisitedBarber {
		if v.name == customerName {
			return true
		}
	}
	return false
}

func cutting(visitorsWaiting chan Visitor) {
	defer wg.Done()

	for wasCustomers < customersAtAll {
		if numberOfWaitingVisitors > 0 {
			currentVisitor := <-visitorsWaiting
			if barberSleep {
				barberSleep = false
				fmt.Print(currentVisitor.name, " awakened the barber\n")
			}
			fmt.Println("Barber is cutting", currentVisitor.name)
			time.Sleep(time.Second * 2)
			numberOfWaitingVisitors--
			visitorsVisitedBarber = append(visitorsVisitedBarber, currentVisitor)
		} else if !barberSleep {
			fmt.Print("Barber going sleep...Zzzzz\n")
			barberSleep = true
		}
	}
}
