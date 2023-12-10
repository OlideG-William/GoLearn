package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dop251/goja"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Read the JavaScript file
	jsFilePath := "C:\\Users\\olegv\\Documents\\GoLearn\\fileServer\\public\\main.js"
	jsCode, err := os.ReadFile(jsFilePath)
	if err != nil {
		http.Error(w, "Error reading JavaScript file", http.StatusInternalServerError)
		return
	}

	// Create a new JavaScript runtime
	vm := goja.New()

	// Evaluate the JavaScript code
	result, err := vm.RunString(string(jsCode))
	if err != nil {
		http.Error(w, "Error executing JavaScript", http.StatusInternalServerError)
		return
	}

	// Return the result as the API response
	fmt.Fprintf(w, "JavaScript file executed successfully. Result: %v", result)
}

func main() {
	http.HandleFunc("/execute", handleRequest)
	http.ListenAndServe(":8080", nil)
}

/*tpl, _ = template.ParseGlob("templates/*.html")
http.HandleFunc("/resultjs", helloHandler)
http.ListenAndServe(":7090", nil)*/

/* myDir := http.Dir("C:\\Users\\olegv\\Documents\\GoLearn\\fileServer\\public")
fmt.Printf("myDir type %T", myDir)

myHandler := http.FileServer(myDir)
http.Handle("/", myHandler) */
//http.Handle("/", http.FileServer(http.Dir("C:\\Users\\olegv\\Documents\\GoLearn\\fileServer\\public")))

//http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

/* vm := goja.New()

new(require.Registry).Enable(vm)
console.Enable(vm)

script := `console.log("Hello world - from Javascript inside Go! ")`

fmt.Println("Compiling ...")
prog, err := goja.Compile("", script, true)
if err != nil {
	fmt.Printf("Error compiling the script %v", err)
	return
}
fmt.Println("Running ...")
_, err = vm.RunProgram(prog) */

//http.HandleFunc("/hello", helloHandler)
//http.ListenAndServe(":8090", nil)
