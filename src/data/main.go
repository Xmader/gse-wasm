package main

import (
	"flag"
	"fmt"

	gse "github.com/go-ego/gse/src"
	"github.com/go-ego/gse/src/dict"
)

var (
	text = flag.String("text", "《复仇者联盟3：无限战争》是全片使用IMAX摄影机拍摄", "要分词的文本")
)

func main() {
	flag.Parse()

	var seg gse.Segmenter
	seg.LoadDict(dict.Dictionary)

	segments := seg.Segment([]byte(*text))
	fmt.Println(gse.ToString(segments, true))
}
