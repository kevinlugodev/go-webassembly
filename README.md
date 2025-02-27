# WebAssembly with Go Example

This repository demonstrates the basics of using WebAssembly (Wasm) with the Go programming language.

## Requirements

- Go 1.11 or later

## Setup

1. Ensure you have Go installed. You can download it from https://golang.org/dl/

2. Create a new directory for your project and navigate into it:
   ```bash
   mkdir go-wasm-example
   cd go-wasm-example
   ```

3. Initialize a new Go module:
   ```bash
   go mod init go-wasm-example
   ```

4. Create the main Go file for your Wasm module:
   ```go
    package main

    import (
    	"syscall/js"
    )

    func main() {
    	c := make(chan struct{}, 0)
    	js.Global().Set("add", js.FuncOf(add))
    	<-c
    }

    func add(this js.Value, args []js.Value) interface{}{
    	sum := 0
    	for _, value := range args {
    		sum += value.Int()
    	}
    	return js.ValueOf(sum)
    }
   ```

5. Compile the Go code to WebAssembly:
   ```bash
   GOOS=js GOARCH=wasm go build -o main.wasm
   ```
   ```bash
   cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./
   ```

6. Create an HTML file to load the Wasm module:
   ```html
    <!DOCTYPE html>
    <html lang="en">
        <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
        <script src="./wasm_exec.js"></script>
        <script>
            const go = new Go();
            WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
                go.run(result.instance)
            })

            function calculate(...args) {
                console.log(add.apply(null, args));
            }
        </script>
    </head>
    <body>
        <button onclick="calculate(4, 5, 6, 99)">Calculate</button>
    </body>
    </html>
   ```

7. Start a simple web server to serve the files. You can use http-server:
   ```bash
   npm i -g http-server
   ```
   ```bash
   http-server -p 3000 -o
   ```

8. Open your web browser and navigate to `http://localhost:3000` to see the example in action.

By following these steps, you've created a basic example of using WebAssembly with Go. You can continue to expand and experiment with more complex functionality as needed.