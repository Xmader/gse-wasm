package main

import (
	"bytes"
	"encoding/gob"
	"syscall/js"

	gse "github.com/Xmader/gse-wasm/src"
)

var (
	seg gse.Segmenter
)

// Todo: 通过js包装实现throw error
func LoadDict(this js.Value, args []js.Value) interface{} {

	dictStrList := make([]string, len(args))

	for i := range args {
		dictStrList[i] = args[i].String()
	}

	err := seg.LoadDict(dictStrList...)
	return err
}

func SetDict(this js.Value, args []js.Value) interface{} {

	b := []byte(args[0].String())
	reader := bytes.NewReader(b)

	dict := &gse.Dictionary{}

	dec := gob.NewDecoder(reader)
	dec.Decode(dict)

	var seg gse.Segmenter
	seg.SetDictionary(dict)

	return nil
}

func AddToken(this js.Value, args []js.Value) interface{} {
	text := args[0].String()
	frequency := args[1].Int()

	if len(args) >= 3 {
		po := args[2].String()
		seg.AddToken(text, frequency, po)
	} else {
		seg.AddToken(text, frequency)
	}

	return nil
}

func AddTokenForce(this js.Value, args []js.Value) interface{} {
	AddToken(this, args)
	seg.CalcToken()

	return nil
}

func CalcToken(this js.Value, args []js.Value) interface{} {
	seg.CalcToken()
	return nil
}

func Cut(this js.Value, args []js.Value) interface{} {
	str, hmm := resolveCutArgs(args)
	return stringArray(seg.Cut(str, hmm))
}

func CutSearch(this js.Value, args []js.Value) interface{} {
	str, hmm := resolveCutArgs(args)
	return stringArray(seg.CutSearch(str, hmm))
}

func CutAll(this js.Value, args []js.Value) interface{} {
	str := args[0].String()
	return stringArray(seg.CutAll(str))
}

func String(this js.Value, args []js.Value) interface{} {
	bytes, searchMode := resolveSegmentArgs(args)
	return seg.String(bytes, searchMode)
}

func Slice(this js.Value, args []js.Value) interface{} {
	bytes, searchMode := resolveSegmentArgs(args)
	return stringArray(seg.Slice(bytes, searchMode))
}

func HMMCut(this js.Value, args []js.Value) interface{} {
	str := args[0].String()
	return stringArray(seg.HMMCut(str))
}
