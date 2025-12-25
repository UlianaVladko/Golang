package main

import (
	"fmt"
	// "time"
)

type BankAccount struct {
	Owner   string
	Balance float64
}

// по задаче надо изменять баланс счета, поэтому указатели на структуру
func (b *BankAccount) Deposit(amount float64) {
    b.Balance += amount
    fmt.Printf("Deposit: +%.2f$. New balance: %.2f$\n", amount, b.Balance)
}

func (b *BankAccount) Withdraw(amount float64) {
    if amount > b.Balance {
        fmt.Printf("Not enough money. Requested: %.2f$. Available: %.2f$.\n", amount, b.Balance)
    } else {
        b.Balance -= amount
        fmt.Printf("Withdraw: -%.2f$. New balance: %.2f$\n", amount, b.Balance)
    }
}

func main() {
    bankaccount := &BankAccount{
        Owner:   "Elon Reeve Musk",
        Balance: 400_000_000_000,
    }

	bankaccount.Deposit(2165432.55648)           // удачное пополнение
	bankaccount.Withdraw(564654654465465.679875) // неудачное снятие
	bankaccount.Withdraw(163763763.5643)         // успешное снятие
	bankaccount.Withdraw(537863763.1414)
    
    fmt.Printf("\nВalance: %.2f$\n", bankaccount.Balance)
}

/*
type Operation struct {
	Type    string
	Amount  float64
	Date    time.Time
	Success bool
	Message string
	ChangesBalance float64
}

type BankAccount struct {
	Owner      string
	Balance    float64
	Operations []Operation
}

func (b BankAccount) GetInfo() {
	fmt.Printf("Owner: %s. Balance: %.4f$.\n", b.Owner, b.Balance)
}

// пусть пополнение строго выполняется, а списание ограничено балансом
func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
	op := Operation{
		Type:    "DEPOSIT",
		Amount:  amount,
		Date:    time.Now(),
		Success: true,
		Message: "Success",
		ChangesBalance: b.Balance,
	}
	b.Operations = append(b.Operations, op)
	fmt.Printf("Deposit: +%.2f$. New balance: %.2f$\n", amount, b.Balance)
}

func (b *BankAccount) Withdraw(amount float64) {
	if amount > b.Balance {
		op := Operation{
			Type:    "WITHDRAW",
			Amount:  amount,
			Date:    time.Now(),
			Success: false,
			Message: "Not enough money.",
			ChangesBalance: b.Balance,
		}
		b.Operations = append(b.Operations, op)
		fmt.Printf("Not enough money. Requested: %.2f$. Availabled: %.2f$.\n", amount, b.Balance)
	} else {
		b.Balance -= amount
		op := Operation{
			Type:    "WITHDRAW",
			Amount:  amount,
			Date:    time.Now(),
			Success: true,
			Message: "Success",
			ChangesBalance: b.Balance,
		}
		b.Operations = append(b.Operations, op)
		fmt.Printf("Withdraw: -%.2f$. New balance: %.2f$.\n", amount, b.Balance)
	}
}

func (b *BankAccount) History() {
	fmt.Println("History")
	for index, op := range b.Operations {
		status := "DONE"
		if !op.Success {
			status = "UNDONE"
		}
		var changes string
		if op.Success {
			if op.Type == "DEPOSIT" {
				changes = fmt.Sprintf("Balance %.2f$", op.ChangesBalance)
			} else {
				changes = fmt.Sprintf("Balance %.2f$", op.ChangesBalance)
			}
		}else{
			changes = fmt.Sprintf("Balance %.2f$", op.ChangesBalance)
		}
		fmt.Printf("%d. %s %s %-10s %12.2f | %s | %s\n",
			index+1,
			op.Date.Format("02.01.2006 15:04:05"),
			status,
			op.Type,
			op.Amount,
			op.Message,
			changes,
		)
	}
}

func main() {
	bankaccount := &BankAccount{
		Owner:   "Elon Reeve Musk",
		Balance: 400_000_000_000,
	}

	bankaccount.Deposit(2165432.55648)           // удачное пополнение
	bankaccount.Withdraw(564654654465465.679875) // неудачное снятие
	bankaccount.Withdraw(163763763.5643)         // успешное снятие
	bankaccount.Withdraw(537863763.1414)         // успешное снятие

	bankaccount.History()
}
*/