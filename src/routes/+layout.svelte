<script>
	import Footer from '$lib/components/Footer.svelte';
	import Navbar from '$lib/components/Navbar.svelte';
	import '@picocss/pico';
	import '../app.css';
</script>

<svelte:head>
	<title>Bagas Hizbullah's Resume</title>
</svelte:head>

<Navbar />

<main class="container">
	<slot />
</main>

<Footer />
