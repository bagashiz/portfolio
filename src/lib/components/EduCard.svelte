<script>
	/** @type {string} school - The school name. */
	export let school;
	/** @type {string} url - The school website link. */
	export let url;
	/** @type {string} major - The major of study. */
	export let major;
	/** @type {string} start - The start date of the education. */
	export let start;
	/** @type {string} end - The end date of the education. */
	export let end;
	/** @type {string} description - The education description. */
	export let description;
</script>

<article>
	<header>
		<hgroup>
			<h3>
				<a href={url} target="_blank">{school}</a>
			</h3>
			<h4>{major}</h4>
		</hgroup>
		<p class="date">{start} - {end}</p>
	</header>
	<p><strong>{description}</strong></p>
</article>

<style>
	article {
		padding-bottom: 0.5rem;
		margin-top: 2rem;
		margin-bottom: 2rem;
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

	a {
		color: var(--text-color);
		text-decoration: none;
	}

	.date {
		color: var(--primary);
	}
</style>
