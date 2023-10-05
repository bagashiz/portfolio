<script>
	import BlogCard from '$lib/components/BlogCard.svelte';

	/** @type {import('./$types').PageData} */
	export let data;
</script>

<section id="blog">
	<h1><strong>Blog Posts</strong></h1>

	{#each data.blogs as blog}
		<BlogCard {blog} />
	{/each}
</section>

<style>
	h1 {
		justify-content: center;
		text-align: center;
		margin-bottom: 3rem;
	}
</style>
