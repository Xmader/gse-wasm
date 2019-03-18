# [gse-wasm](https://github.com/Xmader/gse-wasm)

WebAssembly efficient text segmentation; support english, chinese, japanese and other.

[![Go Report Card](https://goreportcard.com/badge/github.com/Xmader/gse-wasm)](https://goreportcard.com/report/github.com/Xmader/gse-wasm)
[![GoDoc](https://godoc.org/github.com/Xmader/gse-wasm/src?status.svg)](https://godoc.org/github.com/Xmader/gse-wasm/src)
[![npm downloads](https://img.shields.io/npm/dm/gse-wasm.svg)](https://www.npmjs.com/package/gse-wasm)
[![version](https://img.shields.io/github/package-json/v/Xmader/gse-wasm.svg)](https://www.npmjs.com/package/gse-wasm)

[简体中文](https://github.com/Xmader/gse-wasm/blob/master/README_zh.md)

Based on [go-ego/gse](https://github.com/go-ego/gse) 。

<a href="https://github.com/Xmader/gse-wasm/blob/master/dictionary.go">Dictionary </a> with double array trie (Double-Array Trie) to achieve,
<a href="https://github.com/Xmader/gse-wasm/blob/master/segmenter.go">Sender </a> algorithm is the shortest path based on word frequency plus dynamic programming, and DAG and HMM algorithm word segmentation.

Support common, search engine, full mode, precise mode and HMM mode multiple word segmentation modes, support user dictionary, POS tagging.

Support HMM cut text use Viterbi algorithm.

## Install

```
npm install gse-wasm
```

## Usage

[example.js](/example.js)

[Chinese Text Segmentation Demo](https://www.xmader.com/gse-wasm/)

[Japanese Text Segmentation Demo](https://www.xmader.com/gse-wasm/?jp)

## Build 

(Go 1.12 or above should be installed)

```
npm run build:all
```

## Authors

* [Xmader](https://github.com/Xmader)

* [Contributors](https://github.com/Xmader/gse-wasm/graphs/contributors)

## License

Gse Wasm is primarily distributed under the terms of both the MIT license and the Apache License (Version 2.0), thanks to [sego](https://github.com/huichen/sego) and [jieba](https://github.com/fxsjy/jieba).
