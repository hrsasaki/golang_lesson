package ex01

// type withDrawReq struct {
// 	amount int
// 	result chan <- bool
// }

var deposits = make(chan int)         // send amount to deposit
var balances = make(chan int)         // receive balance
var withdraws = make(chan int)        // send amount to withdraw
var withdrawSuccess = make(chan bool) // send if withdrawing is succeed
// var withdraws = make(chan *withDrawReq)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	withdraws <- amount
	return <-withdrawSuccess
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraws:
			withdrawSuccess <- balance >= amount
			if balance >= amount {
				balance -= amount
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
