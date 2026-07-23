package module_04

import "fmt"

type Wallet struct {
	Balance int
}

func (w Wallet) AddValue(v int) {
	w.Balance += v
}

func (w *Wallet) AddPointer(v int) {
	w.Balance += v
}

func ValueVsPointerReceiver_Examples() {
	w := Wallet{Balance: 100}
	w.AddValue(50)
	fmt.Println(w.Balance)

	w.AddPointer(50)
	fmt.Println(w.Balance)
}
