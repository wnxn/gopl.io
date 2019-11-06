// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraw = make(chan struct {
	amount  int
	isValid chan bool
})

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	isValid := make(chan bool)
	withdraw <- struct {
		amount  int
		isValid chan bool
	}{
		amount:  amount,
		isValid: isValid,
	}
	return <-isValid
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case draw := <-withdraw:
			if balance < draw.amount {
				draw.isValid <- false
			} else {
				balance -= draw.amount
				draw.isValid <- true
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
