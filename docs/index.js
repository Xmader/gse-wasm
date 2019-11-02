// @ts-check

(async () => {

    const jp = /^\?jp$/.test(location.search)
    const lang = jp ? "jp" : "zh"

    /** @type {HTMLButtonElement} */
    const runBtn = document.querySelector("#runButton")
    runBtn.disabled = true

    const wasmURL = "https://cdn.staticaly.com/gh/Xmader/gse-wasm/master/dist/gse_lite.wasm"
    const wasm = await Init(wasmURL)

    /** @type {import("..").Gse} */
    const gse = wasm.gse
    const seg = gse.Segmenter

    window["gse"] = gse
    window["seg"] = seg

    /** @type {HTMLTextAreaElement} */
    const textarea = document.querySelector("#textarea")
    const results = document.querySelector("#results")
    const run = () => {
        results.textContent = seg.String(textarea.value)
    }

    runBtn.onclick = run

    /** @type {import("localforage")} */
    const store = localforage.createInstance({
        name: "gse-wasm"
    })
    const dictDataFile = `dict_data_${lang}.bin`

    const savedData = await store.getItem(dictDataFile)
    if (savedData) {
        seg.SetDict(savedData, savedData.length)
    } else {
        const dictDataURL = `https://raw.githubusercontent.com/Xmader/gse-wasm/master/dist/${dictDataFile}`
        const dictDataURLr = await fetch(dictDataURL)
        if (dictDataURLr.ok) {
            const dictData = new Uint8Array(await dictDataURLr.arrayBuffer())
            store.setItem(dictDataFile, dictData)
            seg.SetDict(dictData, dictData.length)
        } else {
            throw new Error(`${dictDataURLr.status} ${dictDataURLr.statusText}`)
        }
    }

    runBtn.disabled = false

})()
