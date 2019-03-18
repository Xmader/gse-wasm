// @ts-check

const Init = require("gse-wasm");

(async () => {

    const wasmURL = "https://unpkg.com/gse-wasm/dist/gse.wasm"  // wasm 文件 URL
    const { gse } = await Init(wasmURL)

    const seg = gse.Segmenter

    // 加载词典
    // 中文和日语词典分别可以使用 "zh" 和 "jp" 表示
    seg.LoadDict("zh")

    // 默认加载中文词典
    // seg.LoadDict()

    // 加载词典字符串
    // seg.LoadDict(`
    //     联邦共和国 32 nt
    //     新星共和国 32 ns
    //     山达尔星新星联邦共和国 64 ns
    //     山达尔星新星联邦共和国联邦政府 32 nt
    // `)

    // 合并加载
    // seg.LoadDict("jp", "迪拜 113 ns\n哈里法 3 n\n哈利法塔 3 nr")

    let text = "你好世界, Hello world."

    let hmm = seg.Cut(text, true)
    console.log("hmm cut: ", hmm)

    hmm = seg.CutSearch(text, true)
    console.log("hmm cut: ", hmm)

    hmm = seg.CutAll(text)
    console.log("cut all: ", hmm)

    // 分词文本
    text = "山达尔星联邦共和国联邦政府"

    // 处理分词结果
    // 支持普通模式和搜索模式两种分词
    // 搜索模式主要用于给搜索引擎提供尽可能多的关键字
    console.log("输出分词结果, 类型为字符串, 使用搜索模式:")
    console.log(seg.String(text, true))

    console.log("输出分词结果, 类型为字符串数组, 不使用搜索模式:")
    console.log(seg.Slice(text))


})()
