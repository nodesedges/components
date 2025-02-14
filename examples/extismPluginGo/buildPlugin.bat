@echo off
REM SET GOOS=js
SET GOOS=wasip1
REM SET GOOS=wasi
SET GOARCH=wasm


tinygo build -o extismPlugin.wasm -target wasip1 -buildmode=c-shared extismPluginGo.go

REM tinygo build -target wasi -o extismPlugin.wasm extismPluginGo.go

REM tinygo build --no-debug -o extismPlugin.wasm -target=wasi -panic=trap -scheduler=none extismPluginGo.go

REM golang v1.24... chek out
REM go build -o extismPlugin.wasm extismPluginGo.go