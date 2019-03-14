# [gse](https://github.com/Xmader/gse-wasm)

Go 语言高效分词, 支持英文、中文、日文等

[![Go Report Card](https://goreportcard.com/badge/github.com/Xmader/gse-wasm)](https://goreportcard.com/report/github.com/Xmader/gse-wasm)
[![GoDoc](https://godoc.org/github.com/Xmader/gse-wasm/src?status.svg)](https://godoc.org/github.com/Xmader/gse-wasm/src)

<a href="https://github.com/Xmader/gse-wasm/blob/master/dictionary.go">词典</a>用双数组 trie（Double-Array Trie）实现，
<a href="https://github.com/Xmader/gse-wasm/blob/master/segmenter.go">分词器</a>算法为基于词频的最短路径加动态规划, 以及 DAG 和 HMM 算法分词.

支持 HMM 分词, 使用 viterbi 算法.

支持普通、搜索引擎、全模式、精确模式和 HMM 模式多种分词模式，支持用户词典、词性标注，可运行<a href="https://github.com/Xmader/gse-wasm/blob/master/server/server.go"> JSON RPC 服务</a>。

分词速度<a href="https://github.com/Xmader/gse-wasm/blob/master/benchmark/benchmark.go">单线程</a> 9.2MB/s，<a href="https://github.com/Xmader/gse-wasm/blob/master/benchmark/goroutines/goroutines.go">goroutines 并发</a> 26.8MB/s. HMM 模式单线程分词速度 3.2MB/s.（ 双核4线程 Macbook Pro）。

## 安装/更新

```
go get -u github.com/Xmader/gse-wasm
```

## 使用

```go
package main

import (
	"fmt"

	gse "github.com/Xmader/gse-wasm/src"
)

var seg gse.Segmenter

func cut() {
	text := "你好世界, Hello world."

	hmm := seg.Cut(text, true)
	fmt.Println("hmm cut: ", hmm)

	hmm = seg.CutSearch(text, true)
	fmt.Println("hmm cut: ", hmm)

	hmm = seg.CutAll(text)
	fmt.Println("cut all: ", hmm)
}

func segCut() {
	// 分词文本
	tb := []byte("山达尔星联邦共和国联邦政府")

	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中 ToString 函数的注释。
	// 搜索模式主要用于给搜索引擎提供尽可能多的关键字
	fmt.Println("输出分词结果, 类型为字符串, 使用搜索模式: ", seg.String(tb, true))
	fmt.Println("输出分词结果, 类型为 slice: ", seg.Slice(tb))

	segments := seg.Segment(tb)
	// 处理分词结果
	fmt.Println(gse.ToString(segments))

	text1 := []byte("上海地标建筑, 东方明珠电视台塔上海中心大厦")
	segments1 := seg.Segment([]byte(text1))
	fmt.Println(gse.ToString(segments1, true))
}

func main() {
	// 加载默认字典
	seg.LoadDict()
	// 载入词典
	// seg.LoadDict("迪拜 113 ns\n哈里法 3 n\n哈利法塔 3 nr")

	cut()

	segCut()
}
```

[自定义词典分词示例](/examples/dict/main.go)

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

	text1 := []byte("所以, 你好, 再见")
	fmt.Println(seg.String(text1, true))

	segments := seg.Segment(text1)
	fmt.Println(gse.ToString(segments))
}
```

[中文分词示例](/examples/example.go)

[日文分词示例](/examples/jp/main.go)

## Authors
* [The author is vz](https://github.com/vcaesar)
* [Maintainers](https://github.com/orgs/go-ego/people)
* [Contributors](https://github.com/Xmader/gse-wasm/graphs/contributors)

## License

Gse is primarily distributed under the terms of both the MIT license and the Apache License (Version 2.0), thanks for [sego](https://github.com/huichen/sego) and [jieba](https://github.com/fxsjy/jieba).
