// Copyright Information.

// Example is provided with help by Gabriel Aszalos.

// This sample program demonstrates how to use a channel to
// monitor the amount of time the program is running and terminate
// the program if it runs too long.
package main

import (
	"log"
	"time"

	"github.com/goinaction/code/chapter7/patterns/runner"
)

// timeout is the number of second the program has to finish.
const timeout = 3 * time.Second

// main is the entry point for the program.
func main() {
	log.Println("Starting work.")

	// Create a new timer value for this run.
	r := runner.New(timeout)

	// Add the tasks to be run and start running.
	r.Add(createTask(1), createTask(2), createTask(1))
	r.Start()

	log.Println("Process Ended")
}

// createTask returns an example task that sleeps for the given
// duration (in seconds).
func createTask(d time.Duration) func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d", id)
		time.Sleep(d * time.Second)
	}
}