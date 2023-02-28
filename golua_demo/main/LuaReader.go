package main

import lua "github.com/yuin/gopher-lua"

const luaPersonTypeName = "reader"

var readerMethods = map[string]lua.LGFunction{
	"read":     luaReaderRead,
	"username": readerGetSetUsername,
}

// 注册定义的类成为lua的一个元表
func registerReaderType(L *lua.LState) {
	mt := L.NewTypeMetatable(luaPersonTypeName)
	L.SetGlobal("reader", mt)
	L.SetField(mt, "new", L.NewFunction(luaNewReader))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), readerMethods))
}

// lua创建对象方法
func luaNewReader(L *lua.LState) int {
	reader := &Reader{
		uint32(L.CheckInt(1)),
		L.CheckString(2),
		uint8(L.CheckInt(3)),
	}
	ud := L.NewUserData()
	ud.Value = reader
	L.SetMetatable(ud, L.GetTypeMetatable(luaPersonTypeName))
	L.Push(ud)
	return 1
}

// 在lua中获取对象的重要一步
func checkReader(L *lua.LState) *Reader {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Reader); ok {
		return v
	}
	L.ArgError(1, "reader expected")
	return nil
}

// 方法注册到lua中
func luaReaderRead(L *lua.LState) int {
	r := checkReader(L)
	book := L.ToString(2)
	r.read(book)
	return 1
}

// 属性的get Set方法， 注意方法名必须这样写：结构名GetSet属性名，大小写也要注意
func readerGetSetUsername(L *lua.LState) int {
	r := checkReader(L)
	if L.GetTop() == 2 {
		r.UserName = L.CheckString(2)
		return 0
	}
	L.Push(lua.LString(r.UserName))
	return 1
}
