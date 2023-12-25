<script>
	import { page } from '$app/stores';

	import '../app.css';
	import Header from '../lib/Header.svelte';
	import Footer from '../lib/Footer.svelte';

	$: path = $page.url.pathname;
</script>

<svelte:head>
	<title>SplitBill | Split The Expense</title>
</svelte:head>

{#if path == '/'}
	<Header />
	<slot />
	<Footer />
{:else if path == '/signup' || path == '/signin' || path == '/password'}
	<slot />
	<Footer />
{:else}
	<slot />
{/if}
