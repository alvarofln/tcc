<script lang="ts">
	import type { Word } from '$lib/models';
	import { processedLines } from '$lib/stores/LyricStore';

	$: headerSize = $processedLines.reduce((acc, line) => Math.max(acc, line.syllableCount), 0);
	$: headers = Array.from(Array(headerSize).keys()).map((i) => (i + 1).toString());

	function isStressed(word: Word, syllabeIndex: number) {
		return word.stressType > -1 && word.syllables?.length - word.stressType - 1 === syllabeIndex;
	}
</script>

<div class="overflow-x-auto">
	<table class="table">
		<thead class="table-head">
			<tr>
				{#each headers as header}
					<th>{header}</th>
				{/each}
			</tr>
		</thead>
		<tbody class="table-body">
			{#each $processedLines as line}
				<tr>
					{#each line.words as word}
						{#each word.syllables as syllable, i}
							{#if isStressed(word, i)}
								<td class="underline font-bold">
									{syllable}
								</td>
							{:else}
								<td>
									{syllable}
								</td>
							{/if}
						{/each}
					{/each}
				</tr>
			{/each}
		</tbody>
	</table>
</div>
