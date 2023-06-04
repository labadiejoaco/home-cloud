<script>
	import { Alert, Button, Input, Label, Modal } from 'flowbite-svelte';
	import { enhance } from '$app/forms';
	import { page } from '$app/stores';

	let createDirModal = false;
</script>

<Button on:click={() => (createDirModal = true)} class="mb-4">
	<svg
		xmlns="http://www.w3.org/2000/svg"
		fill="none"
		viewBox="0 0 24 24"
		stroke-width="1.5"
		stroke="currentColor"
		class="mr-2 h-5 min-h-[1.25rem] w-5 min-w-[1.25rem]"
	>
		<path
			stroke-linecap="round"
			stroke-linejoin="round"
			d="M12 10.5v6m3-3H9m4.06-7.19l-2.12-2.12a1.5 1.5 0 00-1.061-.44H4.5A2.25 2.25 0 002.25 6v12a2.25 2.25 0 002.25 2.25h15A2.25 2.25 0 0021.75 18V9a2.25 2.25 0 00-2.25-2.25h-5.379a1.5 1.5 0 01-1.06-.44z"
		/>
	</svg>
	Create Directory
</Button>
<Modal title="Create Directory" bind:open={createDirModal}>
	{#if $page.form?.form === 'create-directory' && $page.form?.success}
		<Alert color="green">
			<span>{$page.form.message}</span>
		</Alert>
	{/if}
	{#if $page.form?.form === 'create-directory' && $page.form?.error}
		<Alert color="red">
			<span>{$page.form.error}</span>
		</Alert>
	{/if}
	<form id="create-directory" method="POST" action="?/create" use:enhance>
		<div>
			<Label for="name-input" class="mb-2">Name</Label>
			<Input id="name-input" name="name" placeholder="Name" required />
			<input type="hidden" name="path" value={$page.url.pathname} />
		</div>
	</form>
	<svelte:fragment slot="footer">
		<Button type="submit" form="create-directory">Create</Button>
		<Button on:click={() => (createDirModal = false)} color="alternative">Close</Button>
	</svelte:fragment>
</Modal>
