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

import (
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/go-ego/gse/src/dict"
)

var (
	// LoadNoFreq load not have freq dict word
	LoadNoFreq bool
	// MinTokenFreq load min freq token
	MinTokenFreq = 2
)

// Dictionary 返回分词器使用的词典
func (seg *Segmenter) Dictionary() *Dictionary {
	return seg.dict
}

// AddToken add new text to token
func (seg *Segmenter) AddToken(text string, frequency int, pos ...string) {
	var po string
	if len(pos) > 0 {
		po = pos[0]
	}

	words := splitTextToWords([]byte(text))
	token := Token{text: words, frequency: frequency, pos: po}

	seg.dict.addToken(token)
}

// AddTokenForce add new text to token and force
func (seg *Segmenter) AddTokenForce(text string, frequency int, pos ...string) {
	seg.AddToken(text, frequency, pos...)
	seg.CalcToken()
}

// LoadDict load the dictionary from the dictionary string
//
// The format of the dictionary is (one for each participle):
//	participle text, frequency, part of speech
//
// Can load multiple dictionary strings,
// the front of the dictionary preferentially load the participle,
//	such as: "上海 2801 ns\n中心 12 n\n", "迪拜 113 ns\n哈里法 3 n\n哈利法塔 3 nr"
//
// 从词典文件字符串中载入词典
//
// 可以载入多个词典文件字符串，排在前面的词典优先载入分词，比如:
// 	"上海 2801 ns\n中心 12 n\n", "迪拜 113 ns\n哈里法 3 n\n哈利法塔 3 nr"
//
// 词典的格式为（每个分词一行）：
//	分词文本 频率 词性
func (seg *Segmenter) LoadDict(dictStrList ...string) error {
	seg.dict = NewDict()

	log.Println("Gse dictionary loading.")

	if len(dictStrList) > 0 {
		for i := 0; i < len(dictStrList); i++ {
			dictStr := dictStrList[i]

			if dictStr == "zh" || dictStr == "jp" {
				dictStr = LoadBuiltinDict(dictStr)
			} else if dictStr == "en" {
				continue
			}

			err := seg.Read(dictStr)
			if err != nil {
				return err
			}
		}
	}

	if len(dictStrList) == 0 {
		dictStr := dict.Dictionary
		err := seg.Read(dictStr)
		if err != nil {
			return err
		}
	}

	seg.CalcToken()
	log.Println("Gse dictionary loaded finished.")

	return nil
}

// Read read the dict string
func (seg *Segmenter) Read(dictStr string) error {

	reader := strings.NewReader(dictStr)
	var (
		text      string
		freqText  string
		frequency int
		pos       string
	)

	// 逐行读入分词
	line := 0
	for {
		line++
		size, fsErr := fmt.Fscanln(reader, &text, &freqText, &pos)
		if fsErr != nil {
			if fsErr == io.EOF {
				// End of file
				break
			}

			if size > 0 {
				log.Printf("Line \"%v\" read error: %v, skip",
					line, fsErr.Error())
			} else {
				log.Printf("Line \"%v\" is empty, read error: %v, skip",
					line, fsErr.Error())
			}
		}

		if size == 0 {
			// 文件结束或错误行
			// break
			continue
		} else if size < 2 {
			if !LoadNoFreq {
				// 无效行
				continue
			} else {
				freqText = "2"
			}
		} else if size == 2 {
			// 没有词性标注时设为空字符串
			pos = ""
		}

		// 解析词频
		var err error
		frequency, err = strconv.Atoi(freqText)
		if err != nil {
			continue
		}

		// 过滤频率太小的词
		if frequency < MinTokenFreq {
			continue
		}
		// 过滤, 降低词频
		if len([]rune(text)) < 2 {
			// continue
			frequency = 2
		}

		// 将分词添加到字典中
		words := splitTextToWords([]byte(text))
		token := Token{text: words, frequency: frequency, pos: pos}
		seg.dict.addToken(token)
	}

	return nil
}

// LoadBuiltinDict 加载内置词典
func LoadBuiltinDict(dictStr string) string {

	if dictStr == "zh" {
		return dict.ZH
	} else if dictStr == "jp" {
		return dict.JP
	}

	return ""
}

// IsJp is jp char return true
func IsJp(segText string) bool {
	for _, r := range segText {
		jp := unicode.Is(unicode.Scripts["Hiragana"], r) ||
			unicode.Is(unicode.Scripts["Katakana"], r)
		if jp {
			return true
		}
	}
	return false
}

// CalcToken calc the segmenter token
func (seg *Segmenter) CalcToken() {
	// 计算每个分词的路径值，路径值含义见 Token 结构体的注释
	logTotalFrequency := float32(math.Log2(float64(seg.dict.totalFrequency)))
	for i := range seg.dict.tokens {
		token := &seg.dict.tokens[i]
		token.distance = logTotalFrequency - float32(math.Log2(float64(token.frequency)))
	}

	// 对每个分词进行细致划分，用于搜索引擎模式，
	// 该模式用法见 Token 结构体的注释。
	for i := range seg.dict.tokens {
		token := &seg.dict.tokens[i]
		segments := seg.segmentWords(token.text, true)

		// 计算需要添加的子分词数目
		numTokensToAdd := 0
		for iToken := 0; iToken < len(segments); iToken++ {
			// if len(segments[iToken].token.text) > 1 {
			// 略去字元长度为一的分词
			// TODO: 这值得进一步推敲，特别是当字典中有英文复合词的时候
			if len(segments[iToken].token.text) > 0 {
				hasJp := false
				if len(segments[iToken].token.text) == 1 {
					segText := string(segments[iToken].token.text[0])
					hasJp = IsJp(segText)
				}

				if !hasJp {
					numTokensToAdd++
				}
			}
		}
		token.segments = make([]*Segment, numTokensToAdd)

		// 添加子分词
		iSegmentsToAdd := 0
		for iToken := 0; iToken < len(segments); iToken++ {
			// if len(segments[iToken].token.text) > 1 {
			if len(segments[iToken].token.text) > 0 {
				hasJp := false
				if len(segments[iToken].token.text) == 1 {
					segText := string(segments[iToken].token.text[0])
					hasJp = IsJp(segText)
				}

				if !hasJp {
					token.segments[iSegmentsToAdd] = &segments[iToken]
					iSegmentsToAdd++
				}
			}
		}
	}

}
