<script>
	import Navbar from '$lib/components/Navbar.svelte';
	import '@picocss/pico';
	import '../app.css';
</script>

<Navbar />

<main class="container">
	<slot />
</main>
