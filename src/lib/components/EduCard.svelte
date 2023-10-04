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
				<a href={url}>{school}</a>
			</h3>
			<h4>{major}</h4>
		</hgroup>
		<p>{start} - {end}</p>
	</header>
	<p>{description}</p>
</article>

<style>
	article {
		padding-top: 1rem;
		padding-bottom: 0.25rem;
	}

	header {
		padding-top: 1rem;
		padding-bottom: 0rem;
		margin-bottom: 0.5rem;
	}

	a {
		color: var(--text-color);
		text-decoration: none;
	}

	p {
		margin-bottom: 0.5rem;
	}
</style>
