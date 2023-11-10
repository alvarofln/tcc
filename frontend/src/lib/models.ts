export interface Word {
    id: number;
    name: string;
    syllables: string[]
    stressType: number;
}

export interface LyricLine {
    words: Word[];
    syllableCount: number;
}
export enum SearchType { Similar, Rhyme }
