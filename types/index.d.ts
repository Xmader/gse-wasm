// Type definitions for gse-wasm
// Project: https://github.com/Xmader/gse-wasm/
// Definitions by: Xmader <https://github.com/Xmader>

type AddTokenFunction = (text: string, frequency: number, po?: string) => null;
type CutFunction = (str: string, hmm?: boolean) => string[]

export interface Segmenter {
    LoadDict: (...dictStrList: string[]) => null;
    SetDict: (dictData: Uint8Array, length: number) => null;
    AddToken: AddTokenFunction;
    AddTokenForce: AddTokenFunction;
    CalcToken: () => null;
    Cut: CutFunction;
    CutSearch: CutFunction;
    CutAll: (str: string) => string[];
    String: (str: string, searchMode?: boolean) => string;
    Slice: (str: string, searchMode?: boolean) => string[];
    HMMCut: (str: string) => string[];
}

export interface Gse {
    Segmenter: Segmenter;
}

/**
 * @param {string} wasmFile 使用的 wasm 文件的 URL
 */
export type InitFunction = (wasmFile: string) => Promise<{ gse: Gse }>

declare const Init: InitFunction
export default Init
