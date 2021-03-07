package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"syscall/js"
)

var (
	window   = js.Global()
	document = window.Get("document")
	location = document.Get("location")
	host     = location.Get("host")
	url      = "http://" + host.String() + "/api/v1/counter"

	inputNumber      = document.Call("getElementById", "inputNumber")
	text1            = document.Call("getElementById", "text1")
	buttonClear      = document.Call("getElementById", "buttonClear")
	buttonIncCounter = document.Call("getElementById", "buttonIncCounter")
	buttonDecCounter = document.Call("getElementById", "buttonDecCounter")
	buttonSetCounter = document.Call("getElementById", "buttonSetCounter")
	buttonGetCounter = document.Call("getElementById", "buttonGetCounter")
)

func logger(s string) {
	text1.Set("value", s+"\n"+text1.Get("value").String())
}

func main() {
	buttonClear.Call("addEventListener", "click", js.FuncOf(buttonClearClick))
	buttonIncCounter.Call("addEventListener", "click", js.FuncOf(buttonIncCounterClick))
	buttonDecCounter.Call("addEventListener", "click", js.FuncOf(buttonDecCounterClick))
	buttonSetCounter.Call("addEventListener", "click", js.FuncOf(buttonSetCounterClick))
	buttonGetCounter.Call("addEventListener", "click", js.FuncOf(buttonGetCounterClick))

	fmt.Println("Hi") // log.console

	text1.Set("value", url)

	document.Set("title", "Go WebAssembly")

	window.Set("add", js.FuncOf(add)) // export window.add

	// window.Call("Inc") // Call JS function window.Inc from WebAssembly

	buttonGetCounterClick(js.ValueOf(nil), nil)

	select {} // stay
}

func buttonSetCounterClick(this js.Value, args []js.Value) interface{} {
	logger("Set")
	// PUT is idempotent.
	// PUT sets the server's counter.
	s := inputNumber.Get("value").String()
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	toJSON := map[string]int{"counter": int(i)}
	bs, err := json.Marshal(toJSON)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	m := map[string]interface{}{
		"method":  "PUT",
		"body":    string(bs), // JSON
		"headers": map[string]interface{}{"Content-type": "application/json; charset=UTF-8"}}

	window.Call("fetch", url, m).Call("then", js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			return args[0].Call("json")
		})).Call("then",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			inputNumber.Set("value", args[0].Get("counter"))
			return nil
		}))

	return nil
}

func buttonGetCounterClick(this js.Value, args []js.Value) interface{} {
	logger("Get")
	// GET:
	window.Call("fetch", url).Call("then", js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			return args[0].Call("json")
		})).Call("then",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			inputNumber.Set("value", args[0].Get("counter"))
			return nil
		}))
	return nil
}

func buttonIncCounterClick(this js.Value, args []js.Value) interface{} {
	logger("Inc")
	// POST is NOT idempotent.
	// POST adds 100 to the server's counter.
	toJSON := map[string]int{"add": 100}
	bs, err := json.Marshal(toJSON)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	m := map[string]interface{}{
		"method":  "POST",
		"body":    string(bs), // JSON
		"headers": map[string]interface{}{"Content-type": "application/json; charset=UTF-8"}}

	window.Call("fetch", url, m).Call("then", js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			return args[0].Call("json")
		})).Call("then",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			inputNumber.Set("value", args[0].Get("counter"))
			return nil
		}))
	return nil
}

func buttonDecCounterClick(this js.Value, args []js.Value) interface{} {
	logger("Inc")
	// POST is NOT idempotent.
	// POST adds 100 to the server's counter.
	toJSON := map[string]int{"add": -100}
	bs, err := json.Marshal(toJSON)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	m := map[string]interface{}{
		"method":  "POST",
		"body":    string(bs), // JSON
		"headers": map[string]interface{}{"Content-type": "application/json; charset=UTF-8"}}

	window.Call("fetch", url, m).Call("then", js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			return args[0].Call("json")
		})).Call("then",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			inputNumber.Set("value", args[0].Get("counter"))
			return nil
		}))
	return nil
}
func buttonClearClick(this js.Value, args []js.Value) interface{} {
	logger("Clear")
	inputNumber.Set("value", "0")
	return nil
}

func add(this js.Value, args []js.Value) interface{} {
	var c int = args[0].Int() + args[1].Int()
	return js.ValueOf(c)
}
