// Copyright 2016 ego authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

/*

package gse Go efficient text segmentation, Go 语言高性能分词
*/

package gse

import "github.com/Xmader/gse-wasm/src/hmm"

const (
	version = "v0.1.1"

	// minTokenFrequency = 2 // 仅从字典文件中读取大于等于此频率的分词
)

func init() {
	hmm.LoadModel()
}

// GetVersion get the gse version
func GetVersion() string {
	return version
}

// Prob type hmm model struct
type Prob struct {
	B, E, M, S map[rune]float64
}

// Cut cuts a str into words using accurate mode.
// Parameter hmm controls whether to use the HMM
// or use the user's model.
func (seg *Segmenter) Cut(str string, hmm ...bool) []string {
	if len(hmm) <= 0 || (len(hmm) > 0 && hmm[0] == false) {
		return seg.Slice([]byte(str))
		// return seg.cutDAGNoHMM(str)
	}

	return seg.cutDAG(str)
}

// CutSearch cuts str into words using search engine mode.
func (seg *Segmenter) CutSearch(str string, hmm ...bool) []string {
	if len(hmm) <= 0 || (len(hmm) > 0 && hmm[0] == false) {
		return seg.Slice([]byte(str), true)
	}

	return seg.cutForSearch(str, hmm...)
}

// CutAll cuts a str into words using full mode.
func (seg *Segmenter) CutAll(str string) []string {
	return seg.cutAll(str)
}

// Slice use modeSegment segment retrun []string
// using search mode if searchMode is true
func (seg *Segmenter) Slice(bytes []byte, searchMode ...bool) []string {
	segs := seg.ModeSegment(bytes, searchMode...)
	return ToSlice(segs, searchMode...)
}

// Slice use modeSegment segment retrun string
// using search mode if searchMode is true
func (seg *Segmenter) String(bytes []byte, searchMode ...bool) string {
	segs := seg.ModeSegment(bytes, searchMode...)
	return ToString(segs, searchMode...)
}

// LoadModel load the hmm model
func (seg *Segmenter) LoadModel(prob ...map[rune]float64) {
	hmm.LoadModel(prob...)
}

// HMMCut cut sentence string use HMM with Viterbi
func (seg *Segmenter) HMMCut(str string) []string {
	// hmm.LoadModel(prob...)
	return hmm.Cut(str)
}

// HMMCutMod cut sentence string use HMM with Viterbi
func (seg *Segmenter) HMMCutMod(str string, prob ...map[rune]float64) []string {
	hmm.LoadModel(prob...)
	return hmm.Cut(str)
}
