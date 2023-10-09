<script>
	import { onNavigate } from '$app/navigation';
	import Footer from '$lib/components/Footer.svelte';
	import Navbar from '$lib/components/Navbar.svelte';
	import '@picocss/pico';
	import '../app.css';

	// view transition between routes
	onNavigate((navigation) => {
		// @ts-ignore
		if (!document.startViewTransition) return;

		return new Promise((resolve) => {
			// @ts-ignore
			document.startViewTransition(async () => {
				resolve();
				await navigation.complete;
			});
		});
	});
</script>

<svelte:head>
	<title>Bagas Hizbullah's Resume</title>
</svelte:head>

<Navbar />

<main class="container">
	<slot />
</main>

<Footer />
