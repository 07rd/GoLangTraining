package main

import "fmt"

type services interface {
	deposit(amount float64)
	withdraw(amount float64)
	checkBalance() float64
	checkInterest() float64
	accountType() string
}

type savingsAcc struct {
	balance float64
	interest float64
}

type currentAcc struct {
	balance float64
}

func (c *currentAcc) accountType() string {
	return "current"
}

func (c *currentAcc) deposit(amount float64) {
	c.balance = c.balance + amount
}

func (c *currentAcc) withdraw(amount float64) {
	c.balance = c.balance - amount
}

func (c *currentAcc) checkBalance()  float64{
	return c.balance
}

func (c *currentAcc) checkInterest()float64  {
	return 0.0
}

func (c *savingsAcc) accountType() string {
	return "savings"
}

func (s *savingsAcc) deposit(amount float64) {
	s.balance = s.balance + amount
}

func (s *savingsAcc) withdraw(amount float64) {
	s.balance = s.balance - amount
}

func (s *savingsAcc) checkBalance()  float64{
	return s.balance
}

func (s *savingsAcc) checkInterest()float64  {
	return s.balance * (s.interest/100)
}

func accountDetails(service services)  {
	fmt.Println("Account type : ",service.accountType())
	fmt.Println("Account Balance : ",service.checkBalance())
	fmt.Println("Total Interest on account : ",service.checkInterest())
	fmt.Println()
}

func main() {
	cuAcc:= currentAcc{0.0}
	saAcc:= savingsAcc{0.0,4.0}
	cuAcc.deposit(50000.00)
	saAcc.deposit(3000.00)
	cuAcc.withdraw(15650)
	saAcc.withdraw(400)
	accountDetails(&saAcc)
	accountDetails(&cuAcc)

}
/* output:

Account type :  savings
Account Balance :  2600
Total Interest on account :  104

Account type :  current
Account Balance :  34350
Total Interest on account :  0
 */