// @ts-check

(async () => {

    /** @type {HTMLButtonElement} */
    const runBtn = document.querySelector("#runButton")
    runBtn.disabled = true

    const wasm = await Init("https://cdn.staticaly.com/gh/Xmader/gse-wasm/master/dist/gse.wasm")

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

    seg.LoadDict("zh")

    runBtn.disabled = false

})()
