
type AddTokenFunction = (text: string, frequency: number, po?: string) => null;
type CutFunction = (str: string, hmm?: boolean) => string[]

interface Segmenter {
    LoadDict: (...dictStrList: string[]) => null;
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

interface Gse {
    Segmenter: Segmenter;
}

export type InitFunction = () => Promise<{ gse: Gse }>

declare const Init: InitFunction

export default Init
