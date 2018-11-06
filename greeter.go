package main

import "plugin"
import "os"
import "fmt"



type Greeter interface {
	Greet()
}

func main() {
	// determine plugin to load
	lang := "english"
	if len(os.Args) == 2 {
		lang = os.Args[1]
	}
	var mod string
	switch lang {
	case "english":
		mod = "./eng/eng.so"
	case "chinese":
		mod = "./chi/chi.so"
	default:
		fmt.Println("don't speak that language")
		os.Exit(1)
	}

	// load module
	// 1. open the so file to load the symbols
	fmt.Println("Starting")
	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
  fmt.Println("Opened")
	// 2. look up a symbol (an exported function or variable)
	// in this case, variable Greeter
	symGreeter, err := plug.Lookup("Greeter")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
  fmt.Println("Lookup Complete")
	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type Greeter (defined above)
	var greeter Greeter
	greeter, ok := symGreeter.(Greeter)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}
	fmt.Println("Init Extern")
	// 4. use the module
	greeter.Greet()
	fmt.Println("Extern Called")
}
