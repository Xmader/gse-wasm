package main

import (
	"syscall/js"
)

type goFunc func(js.Value, []js.Value) interface{}

// Segmenter - gse.Segmenter 的 js 绑定
func Segmenter() js.Value {

	goFunctions := map[string]goFunc{
		"LoadDict":      LoadDict,
		"SetDict":       SetDict,
		"AddToken":      AddToken,
		"AddTokenForce": AddTokenForce,
		"CalcToken":     CalcToken,
		"Cut":           Cut,
		"CutSearch":     CutSearch,
		"CutAll":        CutAll,
		"String":        String,
		"Slice":         Slice,
		"HMMCut":        HMMCut,
	}

	var jsFunctions = js.ValueOf(make(map[string]interface{}))

	for name, gofunction := range goFunctions {
		jsFunctions.Set(name, js.FuncOf(gofunction))
	}

	return jsFunctions
}

// Gse - gse 的 js 绑定
//export Gse
func Gse() js.Value {
	var gse = js.ValueOf(make(map[string]interface{}))
	gse.Set("Segmenter", Segmenter())

	return gse
}

func main() {

	js.Global().Set("gse", Gse())

	select {}

}
