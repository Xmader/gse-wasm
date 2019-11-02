// Copyright 2013 Hui Chen
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

package gse

import "github.com/Xmader/gse-wasm/src/cedar"

// Dictionary 结构体实现了一个字串前缀树，
// 一个分词可能出现在叶子节点也有可能出现在非叶节点
type Dictionary struct {
	Trie           *cedar.Cedar // Cedar 前缀树
	MaxTokenLen    int          // 词典中最长的分词
	Tokens         []Token      // 词典中所有的分词，方便遍历
	TotalFrequency int64        // 词典中所有分词的频率之和
}

// NewDict new dictionary
func NewDict() *Dictionary {
	return &Dictionary{Trie: cedar.New()}
}

// NumTokens 词典中分词数目
func (dict *Dictionary) NumTokens() int {
	return len(dict.Tokens)
}

// TotalFreq 词典中所有分词的频率之和
func (dict *Dictionary) TotalFreq() int64 {
	return dict.TotalFrequency
}

// addToken 向词典中加入一个分词
func (dict *Dictionary) addToken(token Token) {
	bytes := textSliceToBytes(token.Texts)
	_, err := dict.Trie.Get(bytes)
	if err == nil {
		return
	}

	dict.Trie.Insert(bytes, dict.NumTokens())
	dict.Tokens = append(dict.Tokens, token)
	dict.TotalFrequency += int64(token.Frequency)

	if len(token.Texts) > dict.MaxTokenLen {
		dict.MaxTokenLen = len(token.Texts)
	}
}

// lookupTokens 在词典中查找和字元组 words 可以前缀匹配的所有分词
// 返回值为找到的分词数
func (dict *Dictionary) lookupTokens(
	words []Text, tokens []*Token) (numOfTokens int) {
	var (
		id, value int
		err       error
	)

	for _, word := range words {
		id, err = dict.Trie.Jump(word, id)
		if err != nil {
			break
		}

		value, err = dict.Trie.Value(id)
		if err == nil {
			tokens[numOfTokens] = &dict.Tokens[value]
			numOfTokens++
		}
	}

	return
}

// Find find word in the dictionary is non-existent
// and the word's frequency
func (dict *Dictionary) Find(word []byte) (int, bool) {
	var (
		id, value, freq int
		err             error
	)

	id, err = dict.Trie.Jump(word, id)
	if err != nil {
		return 0, false
	}

	value, err = dict.Trie.Value(id)
	if err != nil && id != 0 {
		return 0, true
	}

	if err != nil {
		return 0, false
	}

	freq = dict.Tokens[value].Frequency

	return freq, true
}
