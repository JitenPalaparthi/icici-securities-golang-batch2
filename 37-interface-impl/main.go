package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "hello world")

	fo := New("data.txt")
	fmt.Fprintln(fo, "Hello ICICI Securities")

	var fo1 *FileOps // here it is nil
	_, err := fo1.Write([]byte("hello world"))
	fmt.Println(err.Error())
}

type FileOps struct {
	FileName  string
	ErrCode   string
	ErrString string
}

func New(fn string) *FileOps {
	return &FileOps{FileName: fn}
}

func (fw *FileOps) Error() string {
	return fmt.Sprintf("Error Code:%v Error Msg:%v", fw.ErrCode, fw.ErrString)
}

func (fw *FileOps) Write(p []byte) (n int, err error) {
	if fw == nil {
		return 0, &FileOps{ErrCode: "101", ErrString: "Nil Error"}
	}
	f, err := os.OpenFile(fw.FileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close() // not to deallocate, it is to tell the os to close gracefully
	return f.Write(p)
}

// func Pay(p IPayment) error {
// 	_, err := p.Transfer()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// type IPayment interface {
// 	Transfer() (float32, error)
// }

// type UPI struct {
// 	UPIID  string
// 	Bank   string
// 	Amount float32
// }

// func (u UPI) Transfer() error {

// 	return nil
// }

// type NetBanking struct {
// 	CID    string
// 	Bank   string
// 	Amount float32
// }

// func (u NetBanking) Transfer() error {
// 	return nil
// }

// type BitCoinPayment struct {
// 	CoinID string
// 	Amount float32
// }

// func (u BitCoinPayment) Transfer() error {
// 	return nil
// }
