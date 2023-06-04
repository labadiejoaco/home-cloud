<script>
	import '../app.css';
	import CreateDirectory from '$lib/components/CreateDirectory.svelte';
	import UploadFiles from '$lib/components/UploadFiles.svelte';
	import Navigate from '$lib/components/Navigate.svelte';
</script>

<main class="m-auto max-w-4xl space-y-12 overflow-hidden px-4 py-12">
	<h1 class="text-center text-4xl">Home Cloud</h1>
	<div class="flex flex-col">
		<UploadFiles />
		<CreateDirectory />
		<Navigate />
	</div>
	<slot />
</main>
