import { purgeCss } from 'vite-plugin-tailwind-purgecss';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	server: {
        proxy: {
            '/api': 'http://127.0.0.1:8080',
        },
    },
	plugins: [sveltekit(), purgeCss()]
	
});
