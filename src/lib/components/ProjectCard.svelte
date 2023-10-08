<script>
	/** @type {Project} */
	export let project;
</script>

<article>
	<h2>
		<a href={project.url} target="_blank">
			<strong>{project.name}</strong>
		</a>
	</h2>

	<div class="badges">
		<a href={project.url} class="badge outline" role="button" target="_blank">
			<i class="fa-solid fa-star" />
			<span>{project.stargazers.totalCount}</span>
		</a>
		<a href={project.url} class="badge outline" role="button" target="_blank">
			<i class="fa-solid fa-code-fork" />
			<span>{project.forks.totalCount}</span>
		</a>
	</div>

	<p>{project.description}</p>
</article>

<style>
	article {
		padding: 0.5rem;
		margin-top: 1rem;
		margin-bottom: 1rem;
		border-radius: 0.75rem;
		overflow: hidden;
	}

	h2 {
		padding: 0.5rem;
		margin: 0rem;
	}

	h2 a {
		text-decoration: none;
		color: var(--text-color);
	}

	h2 a:hover {
		color: var(--primary);
	}

	p {
		padding: 0.5rem;
		margin: 0rem;
	}

	.badges {
		display: flex;
		align-items: center;
		list-style-type: none;
		color: var(--primary);
		margin: 0rem;
		padding-top: 0.25rem;
		padding-bottom: 0.25rem;
		padding-left: 0.5rem;
		padding-right: 0.5rem;
	}

	.badge {
		display: flex;
		align-items: center;
		font-size: medium;
		border: 1px solid;
		border-radius: 0.25rem;
		margin-right: 0.5rem;
		padding: 0.1rem;
		padding-left: 0.25rem;
		padding-right: 0.25rem;
	}

	.badge i {
		margin-right: 0.25rem;
	}
</style>
