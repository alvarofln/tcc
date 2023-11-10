<script lang="ts">
	import { SearchType, type Word } from "$lib/models";
	import { processedLines, searchType, searchRhymingWord, searchSimilarWord } from "$lib/stores/LyricStore";

    function onWordClick(word: Word) {
        if ($searchType === SearchType.Similar) {
            $searchSimilarWord = word;
        }else{
            $searchRhymingWord = word;
        }
    };

</script>
<ol class="list">
    {#each $processedLines as line, i}
        <li>
            <span class="badge-icon p-4 variant-soft-primary" >{line.syllableCount}</span>
            {#each line?.words as word}
                <button class="chip variant-soft-tertiary" on:click={()=>onWordClick(word)} disabled={$searchType === SearchType.Similar && word.id < 0}>{word.name}</button>
            {/each}
        </li>
    {/each}
</ol>