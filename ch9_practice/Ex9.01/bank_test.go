// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"fmt"
	"testing"

	"github.com/kentarotak/Go_Practice/ch9_practice/Ex9.01"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("Alice Deposit=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	var result bool
	go func() {
		result = bank.Withdraw(100)
		fmt.Println("Bob Withdraw=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		fmt.Println("Bob Deposit=", bank.Balance())
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done
	<-done

	if result == true {
		if got, want := bank.Balance(), 200; got != want {
			t.Errorf("Balance = %d, want %d", got, want)
		}
	} else {
		if got, want := bank.Balance(), 300; got != want {
			t.Errorf("Balance = %d, want %d", got, want)
		}
	}
}
