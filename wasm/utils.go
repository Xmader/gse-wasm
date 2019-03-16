package main

import "syscall/js"

func resolveCutArgs(args []js.Value) (string, bool) {
	str := args[0].String()

	hmm := false
	if len(args) >= 2 {
		hmm = args[1].Bool()
	}

	return str, hmm
}

func resolveSegmentArgs(args []js.Value) ([]byte, bool) {
	str := args[0].String()
	bytes := []byte(str)

	var searchMode bool
	if len(args) >= 2 {
		searchMode = args[1].Bool()
	}

	return bytes, searchMode
}

func stringArray(goStringArray []string) js.Value {
	length := len(goStringArray)
	jsStringArray := js.ValueOf(make([]interface{}, length))

	for index, value := range goStringArray {
		jsStringArray.SetIndex(index, value)
	}

	return jsStringArray
}
