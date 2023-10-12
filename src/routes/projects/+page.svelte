<script>
	import ProjectCard from '$lib/components/ProjectCard.svelte';

	export let data;
</script>

<svelte:head>
	<title>Bagas Hizbullah's Projects</title>
</svelte:head>

<section id="projects">
	<h1><strong>Pinned GitHub Projects</strong></h1>

	{#await data.streamed.projects}
		<article aria-busy="true" id="skeleton" class="outline">
			<strong>Fetching data...</strong>
		</article>
	{:then projects}
		<div class="grid">
			{#each projects as project}
				<ProjectCard {project} />
			{/each}
		</div>
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

	.grid {
		display: grid;
		grid-template-columns: repeat(2, 1fr);
	}

	@media (max-width: 600px) {
		.grid {
			grid-template-columns: 1fr; /* Single column for smaller screens */
		}
	}
</style>
