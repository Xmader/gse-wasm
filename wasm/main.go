package main

import (
	"syscall/js"
)

type goFunc func(js.Value, []js.Value) interface{}

func main() {

	goFunctions := map[string]goFunc{
		"LoadDict":      LoadDict,
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

	js.Global().Set("__ges_seg", jsFunctions)

	select {}
}
