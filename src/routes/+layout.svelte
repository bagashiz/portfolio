<script>
	import Footer from '$lib/components/Footer.svelte';
	import Navbar from '$lib/components/Navbar.svelte';
	import '@picocss/pico';
	import '../app.css';
</script>

<Navbar />

<main class="container">
	<slot />
</main>

<Footer />
