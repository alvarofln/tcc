
<script lang="ts">
	import { browser } from "$app/environment";
	import { RadioGroup, RadioItem } from "@skeletonlabs/skeleton";
    import { searchType, searchRhymingWord, searchSimilarWord } from "$lib/stores/LyricStore";
	import { SearchType, type Word } from "$lib/models";
    
    let textArea: HTMLTextAreaElement;
    let similarWords: string = '';
    let rhymingWords: string = '';

    async function loadRhymingWords(word: Word){ 
        if (!word || !browser){
            return;
        } 
        const response = await fetch(`/api/words-rhyming?q=${word.name}`, {
            method: "GET"
        }
        )
        rhymingWords = '';
        textArea.scrollTop = 0;

        const words = await response.json() as Word[];

        rhymingWords = words.map(w => w.name).join('\n');
    }

    async function loadSimilarWords(word: Word) {
        if (!word || !browser){
            return;
        } 
        const response = await fetch(`/api/words/${word.id}/similars`, {
            method: "GET"
        })
        similarWords = '';
        textArea.scrollTop = 0;
        const words = await response.json() as Word[];

        similarWords = words.map(w => w.name).join('\n');
    }

 
    $: loadRhymingWords($searchRhymingWord);
    $: loadSimilarWords($searchSimilarWord);
    
 


    
</script>

<div class="flex">
    <!-- svelte-ignore a11y-label-has-associated-control -->
    <label class="label">
        <span>Modo busca</span>
        <RadioGroup class="rounded-container-token flex w-full">
            <RadioItem bind:group={$searchType} name="searchType" value={SearchType.Similar} class="flex-1">Relacionadas</RadioItem>
            <RadioItem bind:group={$searchType} name="searchType" value={SearchType.Rhyme} class="flex-1">Rimas</RadioItem>
        </RadioGroup>
    </label>
</div>
<label class="label mt-4">

    {#if $searchType === SearchType.Similar}

    <span class="block mb-2">
       <i>{$searchSimilarWord?.name || "Nenhuma palavra selecionada"}...</i>
    </span>
    
    <textarea class="textarea w-full" rows="20" bind:this={textArea} bind:value={similarWords} readonly></textarea>
    {:else}
    <span class="block mb-2">
       <i>{$searchRhymingWord?.name || "Nenhuma palavra selecionada"}...</i>
    </span>
    <textarea class="textarea w-full" rows="20" bind:this={textArea} bind:value={rhymingWords} readonly></textarea>
    {/if}
    
</label>