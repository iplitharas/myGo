package main

import (
	"github.com/fatih/color"
	"sync"
	"time"
)

// The Dining Philosophers problem is well known in computer science circles.
// `Five` philosophers, numbered from 0 through 4, live in a house where the
// table is laid for them; each philosopher has their own place at the table.
// Their only difficulty – besides those of philosophy – is that the dish
// served is a very difficult kind of spaghetti which has to be eaten with
// `two forks`. There are two forks next to each plate, so that presents no
// difficulty. As a consequence, however, this means that no two neighbours
// may be eating simultaneously, since there are five philosophers and five forks.
//

// Philosopher is a struct which stores information about a philosopher.
type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// philosophers is a list of all philosophers
var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

// define some variables
var hunger = 3 // how many times does a person eat?
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

var oderMutex sync.Mutex
var orderFinished []string

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))
	// forks is a map for all 5 forks
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}
	// start the meal.
	for i := 0; i < len(philosophers); i++ {
		// fire off a goroutine for the current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}
	wg.Wait()
}
func diningProblem(philosopher Philosopher, group *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer group.Done()
	// seat the philosopher at the table
	color.Red("%s is seated at the table\n", philosopher.name)
	seated.Done()
	seated.Wait()
	// eat three times
	for i := hunger; i > 0; i-- {
		// get a lock on both forks

		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			color.Yellow("%s takes the right fork\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			color.Red("%s takes the left fork\n", philosopher.name)

		} else {
			forks[philosopher.leftFork].Lock()
			color.Yellow("%s takes the left fork\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			color.Red("%s takes the right fork\n", philosopher.name)

		}

		color.Cyan("\t%s has both forks and is eating...\n", philosopher.name)
		time.Sleep(eatTime)
		color.Cyan("\t%s has both forks and is thinking..\n", philosopher.name)
		time.Sleep(thinkTime)
		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()
		color.Green("\t%s put downs the forks..\n", philosopher.name)

	}
	color.Cyan("%s is satisfied", philosopher.name)
	color.Red("%s left the table", philosopher.name)
	oderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	oderMutex.Unlock()

}

func main() {
	color.Cyan("Dining philosophers problem\n")
	color.Cyan("----------------------------\n")
	color.Cyan("The table is empty\n")
	time.Sleep(sleepTime)
	// start the meal
	dine()
	color.Cyan("The table is empty\n")
	color.Cyan("The order is:%s", orderFinished)
}
