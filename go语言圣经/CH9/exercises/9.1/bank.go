// 练习 9.1： 给gopl.io/ch9/bank1程序添加一个Withdraw(amount int)取款函数。
// 其返回结果应该要表明事务是成功了还是因为没有足够资金失败了。
// 这条消息会被发送给monitor的goroutine，且消息需要包含取款的额度和一个新的channel，
// 这个新channel会被monitorgoroutine来把boolean结果发回给Withdraw
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var takes = make(chan takemoney)

type takemoney struct {
	number int
	ok     chan bool
}

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	if amount <= 0 {
		return false
	}
	
	ok := make(chan bool)
	t := takemoney{amount, ok}
	takes <- t
	return <-ok
}

func teller() {
	// balance is confined to teller goroutine
	var balance int

	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case t := <-takes:
			if t.number > balance {
				t.ok <- false
			} else {
				balance -= t.number
				t.ok <- true
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
