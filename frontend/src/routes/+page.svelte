<script lang="ts">
	import FormEditor from '$lib/components/FormEditorContent.svelte';
	import MeterContent from '$lib/components/MeterContent.svelte';
	import RightContent from '$lib/components/SearchContent.svelte';
	import WordsContent from '$lib/components/WordsContent.svelte';
import { RadioGroup, RadioItem } from '@skeletonlabs/skeleton';

    let viewModes = ['Palavras', 'Métrica']
    let selectedViewMode = viewModes[0]

	
</script>

<div class="container h-full mx-auto flex">
	<div class="min-w-[400px] text-token card p-4 space-y-8">
		<FormEditor/>
	</div>
	<div class="min-w-[600px] w-full text-token card p-4 space-y-8">
        <div class="flex">
        <!-- svelte-ignore a11y-label-has-associated-control -->
        <label class="label">
            <span>Modo de visualização</span>
        <RadioGroup rounded="rounded-container-token" display="flex">
            {#each viewModes as viewMode}
            <RadioItem bind:group={selectedViewMode} name={viewMode} value={viewMode}>{viewMode}</RadioItem>
            {/each}
       </RadioGroup>
        </label>
    </div>
    
    {#if selectedViewMode === 'Métrica'}
        <MeterContent/>
    {:else}
		<WordsContent/>
    {/if}
	</div>
    <div class="min-w-[300px] text-token card p-4 space-y-8">
    <RightContent/>
    </div>
</div>
