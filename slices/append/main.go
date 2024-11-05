/*
Append

The built-in append function is used to dynamically add elements to a slice:

func append(slice []Type, elems ...Type) []Type

If the underlying array is not large enough, append() will create a new underlying array and point the returned slice to it.

Notice that append() is variadic, the following are all valid:

slice = append(slice, oneThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice...)

# Assignment

We've been asked to total costs per day, in a given period.

Complete the getCostsByDay() function using the append() function. It accepts a slice of cost structs and returns a slice of float64s, where each float64 represents the total cost for a day.

The length of the returned slice should be inferred from the highest numbered day field. Some days may have multiple costs, while others may have none. Include all actual and implied totals in the returned slice, starting with day '0'. Use the append() function to dynamically increase the size of the returned slice when needed.
Example

Given this input:

	[]cost{
	    {0, 4.0},
	    {1, 2.1},
	    {5, 2.5},
	    {1, 3.1},
	}

We expect this result:

	[]float64{
	    4.0, // first day
	    5.2, // 2.1 + 3.1
	    0.0, // intermediate days with no costs
	    0.0, // ...
	    0.0, // ...
	    2.5, // last day
	}
*/
package main

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	// Step 1: Find maximum day
	maxDay := 0
	for _, c := range costs {
			if c.day > maxDay {
					maxDay = c.day
			}
	}

	// Step 2: Create empty slice
	dailyCosts := []float64{}

	// Step 3: Append values up to maxDay + 1
	for i := 0; i <= maxDay; i++ {
			sum := 0.0
			// Sum all costs for current day
			for _, c := range costs {
					if c.day == i {
							sum += c.value
					}
			}
			// Append sum (or 0.0 if no costs for that day)
			dailyCosts = append(dailyCosts, sum)
	}

	return dailyCosts
}
