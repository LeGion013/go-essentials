package main

/*

-------     -------     -------
|step1| --> |step2| --> |step3|
-------     -------     -------

if step2 can be broken down into multiple steps,
we can run them in parallel. This process of
breaking down the step is called FAN OUT:

            -------
            |step1|
            -------
               ↓
   -------------------------     ==> FAN OUT (step2 broke down to several steps)
   ↓           ↓           ↓
-------     -------     -------
|step2| --> |step2| --> |step2|
-------     -------     -------
   ↓           ↓           ↓
   -------------------------     ==> FAN IN
               ↓
            -------
            |step3|
            -------

// Explanation from Bill:
// FanOut: you are a manager and you hire one new employee for the exact amount
// of work you have to get done. Each new employee knows immediately what they
// are expected to do and starts their work. You sit waiting for all the results
// are received by you. No given employee needs an immediate guarantee that you
// received their result.

*/

import "fmt"

func main() {
	fmt.Println("app started...")

	fmt.Println("terminated...")
}
