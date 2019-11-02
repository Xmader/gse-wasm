package main

import (
	"bytes"
	"compress/lzw"
	"encoding/gob"
	"io"
	"log"
	"os"

	gse "github.com/Xmader/gse-wasm/src"
	"github.com/Xmader/gse-wasm/src/cedar"
)

// ExportDictData lang: "zh" | "jp"
func ExportDictData(lang, outputDir string) {
	file, err := os.Create(outputDir + "/dict_data_" + lang + ".bin")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	compressor := lzw.NewWriter(file, lzw.MSB, 8)
	defer compressor.Close()

	var seg gse.Segmenter
	seg.LoadDict(lang)

	dict := *seg.Dictionary()

	var buf bytes.Buffer

	gob.Register(cedar.Cedar{})
	enc := gob.NewEncoder(&buf)
	enc.Encode(dict)

	io.Copy(compressor, &buf)
}

func main() {
	ExportDictData("zh", "../dist")
	ExportDictData("jp", "../dist")
}
