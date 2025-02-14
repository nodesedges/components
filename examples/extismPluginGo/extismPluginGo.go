package main

import (
	"extismPluginGo/internal/pluginmeta"
	"fmt"

	pdk "github.com/extism/go-pdk"
)

//go:wasmexport run_test
func run_test() int32 {
	input := pdk.InputString()
	fmt.Println("this was printed from the plugin", input)
	pdk.Log(pdk.LogInfo, "From Inside of plugin log")
	return 0
}

//go:wasmexport GetDescription
func GetDescription() int32 {
	// return the JSON description of this plugin and all of its functions and config...
	s := pluginmeta.GetDescription()
	pdk.Log(pdk.LogDebug, "HUHU")
	pdk.OutputString(s)

	return 0
}

//go:wasmexport test
func test() int32 {

	s := "test :-)"
	pdk.OutputString(s)

	return 0
}

func main() {

}
