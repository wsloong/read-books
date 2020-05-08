package bank

import "testing"

import "time"

func TestDeposit(t *testing.T) {
	go Deposit(10)
	go Deposit(20)
	go Deposit(30)
	go Deposit(40)
}

func TestBalance(t *testing.T) {
	time.Sleep(time.Millisecond * 500)
	if res := Balance(); res != 100 {
		t.Fatalf("expect:%d, get:%d", 100, res)
	}
}

func TestWithdraw(t *testing.T) {
	isok := Withdraw(60)
	if isok != true {
		t.Fatalf("expect:%t, get:%t", true, isok)
	}

	time.Sleep(time.Millisecond * 500)
	if res := Balance(); res != 40 {
		t.Fatalf("expect:%d, get:%d", 40, res)
	}

	isok = Withdraw(60)
	if isok != false {
		t.Fatalf("expect:%t, get:%t", false, isok)
	}
}
