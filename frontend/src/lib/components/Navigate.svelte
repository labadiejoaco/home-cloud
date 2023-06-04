<script>
	import { Button } from 'flowbite-svelte';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';

	const navigateUp = () => {
		const pathSegments = $page.url.pathname.split('/').filter(Boolean);

		const newPath = '/' + pathSegments.slice(0, -1).join('/');

		goto(newPath);
	};
</script>

<Button on:click={navigateUp}>
	<svg
		xmlns="http://www.w3.org/2000/svg"
		fill="none"
		viewBox="0 0 24 24"
		stroke-width="1.5"
		stroke="currentColor"
		class="mr-2 h-5 min-h-[1.25rem] w-5 min-w-[1.25rem] scale-x-[-1]"
	>
		<path
			stroke-linecap="round"
			stroke-linejoin="round"
			d="M9 9l6-6m0 0l6 6m-6-6v12a6 6 0 01-12 0v-3"
		/>
	</svg>
	Navigate Up
</Button>
