// @ts-check

(async () => {

    const jp = /^\?jp$/.test(location.search)

    /** @type {HTMLButtonElement} */
    const runBtn = document.querySelector("#runButton")
    runBtn.disabled = true

    const wasmURL = `https://cdn.staticaly.com/gh/Xmader/gse-wasm/master/dist/gse${jp ? "_full" : ""}.wasm`
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

    seg.LoadDict(jp ? "jp" : "zh")

    runBtn.disabled = false

})()
