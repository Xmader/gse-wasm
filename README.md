# gse

Go efficient text segmentation; support english, chinese, japanese and other.

<!--<img align="right" src="https://raw.githubusercontent.com/go-ego/ego/master/logo.jpg">-->
<!--<a href="https://circleci.com/gh/go-ego/ego/tree/dev"><img src="https://img.shields.io/circleci/project/go-ego/ego/dev.svg" alt="Build Status"></a>-->
[![CircleCI Status](https://circleci.com/gh/go-ego/gse.svg?style=shield)](https://circleci.com/gh/go-ego/gse)
[![codecov](https://codecov.io/gh/go-ego/gse/branch/master/graph/badge.svg)](https://codecov.io/gh/go-ego/gse)
[![Build Status](https://travis-ci.org/go-ego/gse.svg)](https://travis-ci.org/go-ego/gse)
[![Go Report Card](https://goreportcard.com/badge/github.com/Xmader/gse-wasm)](https://goreportcard.com/report/github.com/Xmader/gse-wasm)
[![GoDoc](https://godoc.org/github.com/Xmader/gse-wasm?status.svg)](https://godoc.org/github.com/Xmader/gse-wasm)
[![GitHub release](https://img.shields.io/github/release/go-ego/gse.svg)](https://github.com/Xmader/gse-wasm/releases/latest)
[![Join the chat at https://gitter.im/go-ego/ego](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/go-ego/ego?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
<!-- [![Release](https://github-release-version.herokuapp.com/github/go-ego/gse/release.svg?style=flat)](https://github.com/Xmader/gse-wasm/releases/latest) -->
<!--<a href="https://github.com/go-ego/ego/releases"><img src="https://img.shields.io/badge/%20version%20-%206.0.0%20-blue.svg?style=flat-square" alt="Releases"></a>-->

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

## [Build-tools](https://github.com/go-ego/re)
```
go get -u github.com/go-ego/re
```

### re gse
To create a new gse application

```
$ re gse my-gse
```

### re run

To run the application we just created, you can navigate to the application folder and execute:
```
$ cd my-gse && re run
```

## Use

```go
package main

import (
	"fmt"

	"github.com/Xmader/gse-wasm"
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
	// seg.LoadDict("your gopath"+"/src/github.com/Xmader/gse-wasm/data/dict/dictionary.txt")

	cut()

	segCut()
}

```

[Look at an custom dictionary example](/examples/dict/main.go)

```Go
package main

import (
	"fmt"

	"github.com/Xmader/gse-wasm"
)

func main() {
	var seg gse.Segmenter
	seg.LoadDict("zh,testdata/test_dict.txt,testdata/test_dict1.txt")

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
