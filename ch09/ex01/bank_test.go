// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package ex01

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	go func() {
		success := Withdraw(200)
		fmt.Println("Withdraw1: ", success, Balance())
		done <- struct{}{}
	}()
	<-done
	if got, want, success := Balance(), 100, true; got != want || !success {
		t.Errorf("Withdraw1: Balance = %d, want %d", got, want)
	}

	go func() {
		success := Withdraw(200)
		fmt.Println("Withdraw2: ", success, Balance())
		done <- struct{}{}
	}()
	<-done
	if got, want, success := Balance(), 100, false; got != want || success {
		t.Errorf("Withdraw2: Balance = %d, want %d", got, want)
	}
}
