import { SearchType, type LyricLine, type Word } from '$lib/models';
import { derived, writable } from 'svelte/store';

export const title = writable('');
export const body = writable('');
export const lyric = derived([title, body], ([$title, $body]) => {title; body; return $title + "\n" + $body;});

export const processedLines = writable<LyricLine[]>([]);

export const searchType = writable<SearchType>(SearchType.Similar);

export const searchSimilarWord = writable<Word>();
export const searchRhymingWord = writable<Word>();