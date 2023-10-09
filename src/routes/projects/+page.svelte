<script>
	import ProjectCard from '$lib/components/ProjectCard.svelte';

	export let data;
</script>

<svelte:head>
	<title>Bagas Hizbullah's Projects</title>
</svelte:head>

<section id="projects">
	<h1><strong>Pinned GitHub Projects</strong></h1>

	<div class="grid">
		{#each data.projects as project}
			<ProjectCard {project} />
		{/each}
	</div>
</section>

<style>
	h1 {
		justify-content: center;
		text-align: center;
		margin-bottom: 3rem;
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
