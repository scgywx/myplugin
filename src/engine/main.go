package main

import (
	"io"
	"net/http"
	"log"
	"plugin"
	"sync"
)

var (
	engine *Engine
	fnExec func(IEngine, string)string
)

type IEngine interface{
	Get(string)interface{}
	Set(string, interface{})
	Del(string)
}

type Engine struct{
	mutex sync.Mutex
	vars map[string]interface{}
}

func (e *Engine)Set(key string, val interface{}){
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.vars[key] = val
}

func (e *Engine)Get(key string)interface{}{
	e.mutex.Lock()
	defer e.mutex.Unlock()
	return e.vars[key]
}

func (e *Engine)Del(key string){
	e.mutex.Lock()
	defer e.mutex.Unlock()
	delete(e.vars, key)
}

func handleLoad(w http.ResponseWriter, req *http.Request){
	req.ParseForm()
	names := req.Form["name"]
	if len(names) > 0 {
		load(names[0])
	}else{
		log.Println("filename error")
	}
	io.WriteString(w, "done")
}

func handleHello(w http.ResponseWriter, req *http.Request){
	str := fnExec(engine, "test")
	io.WriteString(w, str)
}

func main() {
	engine = &Engine{
		vars: make(map[string]interface{}),
	}

	load("plugin1.so")

	http.HandleFunc("/load", handleLoad)
	http.HandleFunc("/hello", handleHello)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

func load(filename string){
	p, err := plugin.Open(filename)
	if err != nil{
		log.Println("open plugin err:", err, filename)
		return
	}

	fn, err := p.Lookup("Exec")
	if err != nil{
		log.Println("not found symbol Exec", err)
		return
	}

	fnExec = fn.(func(IEngine, string)string)
	
	log.Println("loaded plugin successed! file=", filename)
}
