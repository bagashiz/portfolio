<script>
	import JobCard from '$lib/components/JobCard.svelte';
	import { experiences, volunteering } from '$lib/store';
</script>

<!-- Profile -->
<section id="profile">
	<h1><strong>Bagas Hizbullah</strong></h1>

	<summary>
		I am an Information Systems Student at <a href="https://ub.ac.id">Brawijaya University</a>,
		passionate about web development, Linux, and Free and Open Source Software (FOSS). I am eager to
		explore opportunities and contribute significantly to the IT industry.
	</summary>
</section>

<!-- Experience -->
<section id="experience">
	<h2><strong>Experience</strong></h2>

	{#each experiences as exp}
		<JobCard {...exp} />
	{/each}
</section>

<!-- Volunteering -->
<section id="volunteering">
	<h2><strong>Volunteering</strong></h2>

	{#each volunteering as vol}
		<JobCard {...vol} />
	{/each}
</section>
