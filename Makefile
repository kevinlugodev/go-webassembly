build:
	GOOS=js GOARCH=wasm go build -o  main.wasm && cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./
