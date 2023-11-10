
<script lang="ts">
	import { browser } from "$app/environment";
	import type { LyricLine } from "$lib/models";
    import { title, body, processedLines } from "$lib/stores/LyricStore";




    async function processLyric(title: string, body: string) {
        if (!body && !title || !browser){
            $processedLines = [];
            return
        } 
        
        const response = await fetch(`/api/lyrics`, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({title, body})
        }
        )

      
        const data = await response.json();
        $processedLines = data.lines as LyricLine[];
    }

    $: processLyric($title, $body);



</script>
<label class="label">
    <span>Título</span>
    <input class="input" type="text" placeholder="Meu título..." bind:value={$title} />
</label>
<label class="label">
    <span>Letra</span>
    <textarea class="textarea" rows="20" placeholder="Minha letra..." bind:value={$body} />
</label>