<script>
	export const ssr = false;
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';

	let time = 5;

	onMount(() => {
		const interval = setInterval(() => {
			time--;
			if (time === 0) {
				clearInterval(interval); // Clear the interval when time reaches 0
				goto('/');
			}
		}, 1000);

		// Make sure to clear the interval when the component is unmounted
		return () => clearInterval(interval);
	});
</script>

<h1 class="text-center" style="margin-top: 2rem;">
	{$page.status} - {$page.error.message} <br />
</h1>
<p class="text-center mt-10">Redirecting to home page in {time} seconds</p>
