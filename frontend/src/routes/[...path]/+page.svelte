<script>
	import List from '$lib/components/List.svelte';

	/** @type {import('./$types').PageData} */
	export let data;
</script>

<div class="space-y-6">
	<div>
		<h2 class="text-lg">Directories</h2>
		{#if data.directories.length > 0}
			<List items={data.directories} type="Directory" />
		{:else}
			<p class="text-gray-600">No directories found</p>
		{/if}
	</div>
	<div>
		<h2 class="text-lg">Files</h2>
		{#if data.files.length > 0}
			<List items={data.files} type="File" />
		{:else}
			<p class="text-gray-600">No files found</p>
		{/if}
	</div>
</div>
