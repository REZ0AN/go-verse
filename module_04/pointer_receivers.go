package module_04

import (
	"fmt"
	"io"
)

type WalletP struct {
	Balance int
}

func (w WalletP) AddValue(amount int) {
	w.Balance += amount
}

func (w *WalletP) AddPointer(amount int) {
	w.Balance += amount
}

func PointerReceivers(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Value Receiver")
	fmt.Fprintln(w, "--------------")

	wallet := WalletP{Balance: 100}

	wallet.AddValue(50)

	fmt.Fprintf(w, "Balance = %d\n", wallet.Balance)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Pointer Receiver")
	fmt.Fprintln(w, "----------------")

	wallet.AddPointer(50)

	fmt.Fprintf(w, "Balance = %d\n", wallet.Balance)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Methods on Pointer Values")
	fmt.Fprintln(w, "-------------------------")

	ptr := &WalletP{Balance: 200}

	ptr.AddPointer(25)

	fmt.Fprintf(w, "Balance = %d\n", ptr.Balance)

	return nil
}
