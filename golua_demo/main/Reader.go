package main

import "fmt"

type Reader struct {
	Uid         uint32
	UserName    string
	ReaderCount uint8
}

func (reader *Reader) read(book string) {
	reader.ReaderCount++
	fmt.Printf("Reader:%v,Name:%v,read book %v\n", reader.Uid, reader.UserName, book)
}
