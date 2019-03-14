# gse

Go efficient text segmentation; support english, chinese, japanese and other.

[![Go Report Card](https://goreportcard.com/badge/github.com/Xmader/gse-wasm)](https://goreportcard.com/report/github.com/Xmader/gse-wasm)
[![GoDoc](https://godoc.org/github.com/Xmader/gse-wasm/src?status.svg)](https://godoc.org/github.com/Xmader/gse-wasm/src)

[简体中文](https://github.com/Xmader/gse-wasm/blob/master/README_zh.md)

<a href="https://github.com/Xmader/gse-wasm/blob/master/dictionary.go">Dictionary </a> with double array trie (Double-Array Trie) to achieve,
<a href="https://github.com/Xmader/gse-wasm/blob/master/segmenter.go">Sender </a> algorithm is the shortest path based on word frequency plus dynamic programming, and DAG and HMM algorithm word segmentation.

Support common, search engine, full mode, precise mode and HMM mode multiple word segmentation modes, support user dictionary, POS tagging, run<a href="https://github.com/Xmader/gse-wasm/blob/master/server/server.go"> JSON RPC service</a>.

Support HMM cut text use Viterbi algorithm.

Text Segmentation speed<a href="https://github.com/Xmader/gse-wasm/blob/master/benchmark/benchmark.go"> single thread</a> 9.2MB/s，<a href="https://github.com/Xmader/gse-wasm/blob/master/benchmark/goroutines/goroutines.go">goroutines concurrent</a> 26.8MB/s. HMM text segmentation single thread 3.2MB/s. (2core 4threads Macbook Pro).

## Install / update

```
go get -u github.com/Xmader/gse-wasm
```

## Use

```go
package main

import (
	"fmt"

	gse "github.com/Xmader/gse-wasm/src"
)

var (
	text = "你好世界, Hello world."

	seg gse.Segmenter
)

func cut() {
	hmm := seg.Cut(text, true)
	fmt.Println("hmm cut: ", hmm)

	hmm = seg.CutSearch(text, true)
	fmt.Println("hmm cut: ", hmm)

	hmm = seg.CutAll(text)
	fmt.Println("cut all: ", hmm)
}

func segCut() {
	// Text Segmentation
	tb := []byte(text)
	fmt.Println(seg.String(tb, true))

	segments := seg.Segment(tb)

	// Handle word segmentation results
	// Support for normal mode and search mode two participle,
	// see the comments in the code ToString function.
	// The search mode is mainly used to provide search engines
	// with as many keywords as possible
	fmt.Println(gse.ToString(segments, true))
}

func main() {
	// Loading the default dictionary
	seg.LoadDict()
	// Load the dictionary
	// seg.LoadDict("迪拜 113 ns\n哈里法 3 n\n哈利法塔 3 nr")

	cut()

	segCut()
}

```

[Look at an custom dictionary example](/examples/dict/main.go)

```Go
package main

import (
	"fmt"

	gse "github.com/Xmader/gse-wasm/src"
	"github.com/Xmader/gse-wasm/testdata"
)

func main() {
	var seg gse.Segmenter
	seg.LoadDict("zh", testdata.TestDict0, testdata.TestDict1)

	text1 := []byte("你好世界, Hello world")
	fmt.Println(seg.String(text1, true))

	segments := seg.Segment(text1)
	fmt.Println(gse.ToString(segments))
}
```

[Look at an Chinese example](https://github.com/Xmader/gse-wasm/blob/master/examples/example.go)

[Look at an Japanese example](https://github.com/Xmader/gse-wasm/blob/master/examples/jp/main.go)

## Authors
* [The author is vz](https://github.com/vcaesar)
* [Maintainers](https://github.com/orgs/go-ego/people)
* [Contributors](https://github.com/Xmader/gse-wasm/graphs/contributors)

## License

Gse is primarily distributed under the terms of both the MIT license and the Apache License (Version 2.0), thanks for [sego](https://github.com/huichen/sego) and [jieba](https://github.com/fxsjy/jieba).
