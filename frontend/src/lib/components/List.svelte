<script>
	import Directory from './Directory.svelte';
	import File from './File.svelte';

	export let items, type;
</script>

<div class="mt-2 grid grid-cols-1 gap-2 sm:grid-cols-3 lg:grid-cols-4">
	{#each items as item}
		{#if type === 'Directory'}
			<Directory name={item} />
		{:else if type === 'File'}
			<File name={item} />
		{/if}
	{/each}
</div>
