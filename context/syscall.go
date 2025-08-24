package context

import (
	"fmt"
	"syscall/js"
)

func FetchFromGo(url string, onDone func(s string)) {
	fetchPromise := js.Global().Call("fetch", url)

	// When fetch is successful
	thenCallback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resp := args[0]
		textPromise := resp.Call("text")

		textCallback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			text := args[0].String()
			// ✅ Use the result in Go!
			onDone(text)

			return nil
		})

		textPromise.Call("then", textCallback)
		return nil
	})

	// If fetch fails
	catchCallback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		err := args[0]
		fmt.Println(err)
		return nil
	})

	fetchPromise.Call("then", thenCallback).Call("catch", catchCallback)
}

func MakeOutsideReqeust(url string) string {
	ch := make(chan string)
	FetchFromGo(url, func(s string) {
		ch <- s
	})
	result := <-ch
	return result
}
