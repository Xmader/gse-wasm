package main

import (
	"fmt"

	gse "github.com/Xmader/gse-wasm/src"
)

var (
	text = "纽约时代广场, 纽约帝国大厦"

	seg gse.Segmenter
)

func main() {
	hmm := seg.HMMCutMod(text)
	fmt.Println("hmm cut: ", hmm)

	seg.LoadDict()

	hmm = seg.Cut(text, true)
	fmt.Println("hmm cut: ", hmm)

	hmm = seg.CutSearch(text, true)
	fmt.Println("hmm cut: ", hmm)

	hmm = seg.CutAll(text)
	fmt.Println("cut all: ", hmm)
}
