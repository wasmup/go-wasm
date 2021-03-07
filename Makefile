all:
	# tinygo build -o view/main.wasm -target wasm wasm/main.go
	GOOS=js GOARCH=wasm go build -o ./view/main.wasm ./wasm/
	ls -lh ./view/main.wasm
	# 2.5M

linux64:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ${GOTMPDIR}/serve .
	# ls -lh ${GOTMPDIR}/serve
	# 7.0M
	# file ${GOTMPDIR}/serve
	# ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, stripped

win64:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ${GOTMPDIR}/serve.exe .

init:
	cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./view/
	# go get
	# git log --graph --oneline --all
	