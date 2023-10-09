<script>
	/** @type {string} job - The job title. */
	export let job;
	/** @type {string} url - The company website link. */
	export let url;
	/** @type {string} company - The company name. */
	export let company;
	/** @type {string} start - The start date of the job. */
	export let start;
	/** @type {string} end - The end date of the job. */
	export let end;
	/** @type {string[]} description - The job description. */
	export let description;
</script>

<article>
	<header>
		<hgroup>
			<h3>{job}</h3>
			<h4>
				<a href={url} target="_blank" class="secondary">{company}</a>
			</h4>
		</hgroup>
		<p class="date">{start} - {end}</p>
	</header>
	<details>
		<summary>
			<strong>Job Description</strong>
		</summary>
		<ul>
			{#each description as desc}
				<li>{desc}</li>
			{/each}
		</ul>
	</details>
</article>

<style>
	article {
		padding-top: 1rem;
		padding-bottom: 0rem;
		border-radius: 0.75rem;
	}

	header {
		padding-top: 1rem;
		padding-bottom: 0rem;
		margin-bottom: 1rem;
	}

	hgroup {
		margin-bottom: 0.75rem;
	}

	p {
		margin-bottom: 0.75rem;
	}

	.date {
		color: var(--primary);
	}
</style>
