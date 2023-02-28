package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

var modFuncs = map[string]lua.LGFunction{
	"eat":    Eat,
	"drink":  Drink,
	"record": Record,
}

func Eat(L *lua.LState) int {
	msg := L.CheckString(1)
	fmt.Println("eat:", msg)
	return 0
}

func Drink(L *lua.LState) int {
	msg := L.CheckString(1)
	fmt.Println("drink:", msg)
	return 0
}

func Record(L *lua.LState) int {
	r := checkReader(L)
	fmt.Printf("%v读完了！一共%v本书！\n", r.UserName, r.ReaderCount)
	return 1
}

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), modFuncs)
	L.SetField(mod, "mymod", lua.LString("value"))
	L.Push(mod)
	return 1
}
