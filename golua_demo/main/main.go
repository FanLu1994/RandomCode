package main

import (
	"bufio"
	"fmt"
	"github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

// TODO: 加载lua代码执行
// TODO: 多线程

var wg sync.WaitGroup

func main() {
	books := []string{
		"活着", "白鹿原", "春秋战国", "兄弟", "许三观卖血记", "丰乳肥臀",
	}

	luaPath := "./main/test.lua"
	luaProto, err := compileFile(luaPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go DoRead(luaProto, uint32(i), "Reader"+strconv.Itoa(i), books)
	}

	wg.Wait()

}

// 机器人主流程
func DoRead(luaProto *lua.FunctionProto, id uint32, name string, books []string) {
	fmt.Println(id)
	L := lua.NewState()
	defer L.Close()
	registerReaderType(L)
	L.PreloadModule("mymod", Loader)          // 注入自己的模块
	lFunc := L.NewFunctionFromProto(luaProto) // 从字节码解析得到
	L.Push(lFunc)
	L.PCall(0, lua.MultRet, nil)

	// init
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("init"),
		NRet:    0,
		Protect: true,
	}, lua.LNil); err != nil {
		fmt.Println(err)
	}

	// 新建机器人
	L.SetGlobal("global_id", lua.LNumber(id))
	L.SetGlobal("global_name", lua.LString(name))
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("newReader"),
		NRet:    0,
		Protect: true,
	}, lua.LNil); err != nil {
		fmt.Println(err)
	}

	// 读书
	for i := 0; i < 3; i++ {
		book := books[rand.Int()%len(books)]
		if err := L.CallByParam(lua.P{
			Fn:      L.GetGlobal("read"),
			NRet:    0,
			Protect: true,
		}, lua.LString(book)); err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second)
	}

	// 结束
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("finish"),
		NRet:    0,
		Protect: true,
	}, lua.LNil); err != nil {
		fmt.Println(err)
	}
	wg.Done()
}

// 解析文件变成lua字节码
func compileFile(filePath string) (*lua.FunctionProto, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	chunk, err := parse.Parse(reader, filePath)
	if err != nil {
		return nil, err
	}
	proto, err := lua.Compile(chunk, filePath)
	if err != nil {
		return nil, err
	}
	return proto, nil
}
