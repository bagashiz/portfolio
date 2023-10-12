<script>
	import BlogCard from '$lib/components/BlogCard.svelte';

	/** @type {import('./$types').PageData} */
	export let data;
</script>

<svelte:head>
	<title>Bagas Hizbullah's Blogs</title>
</svelte:head>

<section id="blog">
	<h1><strong>Blog Posts</strong></h1>

	{#await data.streamed.blogs}
		<article aria-busy="true" id="skeleton" class="outline">
			<strong>Fetching data...</strong>
		</article>
	{:then blogs}
		{#each blogs as blog}
			<BlogCard {blog} />
		{/each}
	{:catch}
		<article id="skeleton" class="outline">
			<strong>Uh oh! Failed to fetch data.</strong>
		</article>
	{/await}
</section>

<style>
	h1 {
		justify-content: center;
		text-align: center;
		margin-bottom: 3rem;
	}

	#skeleton {
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 2rem;
	}
</style>
