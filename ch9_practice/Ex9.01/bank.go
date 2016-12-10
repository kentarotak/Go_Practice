// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

import "fmt"

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var debits = make(chan int)
var transactionresult = make(chan bool)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {
	debits <- amount
	return <-transactionresult
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-debits:
			if amount < balance {
				fmt.Printf("取引に成功しました。\n")
				balance = balance - amount
				fmt.Printf("残高は%dです\n", balance)
				transactionresult <- true
			} else {
				fmt.Printf("残高が足りません\n")
				fmt.Printf("残高は%dです\n", balance)
				transactionresult <- false
			}

		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
