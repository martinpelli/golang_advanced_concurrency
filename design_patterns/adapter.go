package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct {
}

func (CashPayment) Pay() {
	fmt.Println("Payment using cash")
}

func ProcessPayment(payment Payment) {
	payment.Pay()
}

type BankPayment struct {
}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying using bank account %d\n", bankAccount)
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (bankPaymentAdapater BankPaymentAdapter) Pay() {
	bankPaymentAdapater.BankPayment.Pay(bankPaymentAdapater.bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)
	//bank := &BankPayment{}
	//ProcessPayment(bank)
	bank := &BankPaymentAdapter{
		bankAccount: 5,
		BankPayment: &BankPayment{},
	}
	ProcessPayment(bank)
}
