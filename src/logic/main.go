package main

import(
	"fmt"
	"unsafe"
	"reflect"
)

var (
	modelVersion = 1002
	modelName = "game2"
)

type IEngine interface{
	Get(string)(interface{})
	Set(string, interface{})
	Del(string)
}

type Game struct{
	Version int
	Name string
}

func Exec(e IEngine, name string) string{
	var game *Game
	v := e.Get("game")
	if v == nil{
		game = &Game{}
		e.Set("game", game)
	}else{
		game = (*Game)(unsafe.Pointer(reflect.ValueOf(v).Pointer()))
	}

	ret := fmt.Sprintf("hello %s, this is golang plugin test!, version=%d, name=%s, oldversion=%d, oldName=%s\n", name, modelVersion, modelName, game.Version, game.Name)
	game.Version = modelVersion
	game.Name = modelName

	return ret
}
