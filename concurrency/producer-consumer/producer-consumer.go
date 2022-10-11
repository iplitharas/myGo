package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

// Producer is a type for structs that holds two channels: one for pizzas, with all
// information for a given pizza order including whether it was made
// successfully,
// and another to handle end of processing (when we quit the channel)
type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

// Close is simply a method of closing the channel when we are done with it (i.e.
// something is pushed to the quit channel)
func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

// PizzaOrder is a type for structs that describes a given pizza order. It has the order
// number, a message indicating what happened to the order, and a boolean
// indicating if the order was successfully completed.
type PizzaOrder struct {
	number  int
	success bool
	message string
}

// makePizza attempts to make a pizza. We generate a random Number from 1-12,
// and put in two cases where we can't make the pizza in time. Otherwise,
// we make the pizza without issue. To make things interesting, each pizza
// will take a different length of time to produce (some pizzas are harder than others).
func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		fmt.Printf("Received order #%d!\n", pizzaNumber)
		delay := rand.Intn(5) + 1
		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		// delay for a bit
		fmt.Printf("Making pizza #%d. It will take %d seconds....\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			number:  pizzaNumber,
			message: msg,
			success: success,
		}

		return &p

	}

	return &PizzaOrder{
		number: pizzaNumber,
	}
}

// openPizzeria is a goroutine that runs in the background and
// calls makePizza to try to make one order each time it iterates through
// the for loop. It executes until it receives something on the quit
// channel. The quit channel does not receive anything until the consumer
// sends it (when the PizzaNumber of orders is greater than or equal to the
// constant NumberOfPizzas).
func openPizzeria(pizzaMaker *Producer) {

	var i = 0
	for {
		newPizza := makePizza(i)
		if newPizza != nil {
			i = newPizza.number
			select {
			case pizzaMaker.data <- *newPizza:
			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return

			}
		}
	}

}

func consumer(pizzeria *Producer) {
	// create and run consumer
	for i := range pizzeria.data {
		if i.number <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!", i.number)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!")
			}
		} else {
			color.Cyan("Done making pizzas...")
			err := pizzeria.Close()
			if err != nil {
				color.Red("*** Error closing channel!", err)
			}
		}
	}

}
func main() {
	color.Cyan("----------------------------------")
	color.Cyan("Welcome to Pizzeria - \nProducer - Consumer example")
	pizzeria := Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}
	go openPizzeria(&pizzeria)
	consumer(&pizzeria)

}
