/*
	Type assertions in Go

When working with interfaces in Go, every once-in-awhile you'll need access to the underlying type of an interface value. You can cast an interface to its underlying type using a type assertion.

In the example below:

we want to check if s is a circle in order to cast it into its underlying concrete type
we know s is an instance of the shape interface, but we do not know if it's also a circle
c is a new circle struct cast from s
ok is true if s is indeed a circle, or false if s is NOT a circle

	type shape interface {
		area() float64
	}

	type circle struct {
		radius float64
	}

c, ok := s.(circle)

	if !ok {
		// log an error if s isn't a circle
		log.Fatal("s is not a circle")
	}

radius := c.radius
Copy icon
Assignment
Implement the getExpenseReport function.

If the expense is an email then it should return the email's toAddress and the cost of the email.
If the expense is an sms then it should return the sms's toPhoneNumber and its cost.
If the expense has any other underlying type, just return an empty string and 0.0 for the cost.
*/
package main

func getExpenseReport(e expense) (string, float64) {

	email, ok1 := e.(email) // type assertions indicated by i.(T)
	sms, ok2 := e.(sms)

	if ok1 {
		return email.toAddress, email.cost()
	}

	if ok2 {
		return sms.toPhoneNumber, sms.cost()
	}

	return "", 0

}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}