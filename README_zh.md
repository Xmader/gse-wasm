# [gse-wasm](https://github.com/Xmader/gse-wasm)

WebAssembly 高效分词, 支持英文、中文、日文等，可用于生产环境

[![Go Report Card](https://goreportcard.com/badge/github.com/Xmader/gse-wasm)](https://goreportcard.com/report/github.com/Xmader/gse-wasm)
[![GoDoc](https://godoc.org/github.com/Xmader/gse-wasm/src?status.svg)](https://godoc.org/github.com/Xmader/gse-wasm/src)
[![npm downloads](https://img.shields.io/npm/dm/gse-wasm.svg)](https://www.npmjs.com/package/gse-wasm)
[![version](https://img.shields.io/github/package-json/v/Xmader/gse-wasm.svg)](https://www.npmjs.com/package/gse-wasm)

基于 [go-ego/gse](https://github.com/go-ego/gse) 。

<a href="https://github.com/Xmader/gse-wasm/blob/master/dictionary.go">词典</a>用双数组 trie（Double-Array Trie）实现，
<a href="https://github.com/Xmader/gse-wasm/blob/master/segmenter.go">分词器</a>算法为基于词频的最短路径加动态规划, 以及 DAG 和 HMM 算法分词.

支持 HMM 分词, 使用 viterbi 算法.

支持普通、搜索引擎、全模式、精确模式和 HMM 模式多种分词模式，支持用户词典、词性标注。

## 安装

```
npm install gse-wasm
```

## 使用

[example.js](/example.js)

[演示地址](https://www.xmader.com/gse-wasm/)

[日文分词演示](https://www.xmader.com/gse-wasm/?jp)

## 自行构建

(需要安装 Go 1.13 或以上版本)

```
npm run build:all
```

## 对不同构建版本的解释

在 NPM 包的 dist/ 目录你将会找到很多不同的 WebAssembly 构建版本。这里列出了它们之间的差别：

| 版本 | 说明 |
|---|---|
| gse.wasm | 默认使用，包含中文分词词典 |
| gse_full.wasm | 完整版，包含中文和日语分词词典 |
| gse_lite.wasm | **不**包含分词词典 |

## Authors

* [Xmader](https://github.com/Xmader)

* [Contributors](https://github.com/Xmader/gse-wasm/graphs/contributors)

## License

Gse Wasm is primarily distributed under the terms of both the MIT license and the Apache License (Version 2.0), thanks to [sego](https://github.com/huichen/sego) and [jieba](https://github.com/fxsjy/jieba).
