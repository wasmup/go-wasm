.PHONY: init wasm linux64 run

all: init wasm linux64 run

run:
	${GOTMPDIR}/serve

wasm:
	GOOS=js GOARCH=wasm go build -ldflags "-s -w" -o view/main.wasm wasm/main.go
	ls -ll view/main.wasm
# 2_525_255

# tinygo build -o view/main.wasm -target wasm -no-debug wasm/main.go
# 192_868


linux64:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ${GOTMPDIR}/serve .
	ls -lh ${GOTMPDIR}/serve
# 7.0M
	file ${GOTMPDIR}/serve
# ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, stripped

win64:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ${GOTMPDIR}/serve.exe .

init:
	cp /usr/local/go/misc/wasm/wasm_exec.js view/
# cp ${go env GOROOT}/misc/wasm/wasm_exec.js view/
# cp /usr/local/tinygo/targets/wasm_exec.js ./view/

# go get
# git log --graph --oneline --all
	