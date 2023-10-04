<script>
	/** @type {string[]} skills - The list of programming languages and tools. */
	export let skills;
</script>

<article>
	<header>
		<h3>Programming Languages & Tools</h3>
	</header>
	<ul id="tools-list">
		{#each skills as skill}
			<li class="dev-logo">
				<i class="fab fa-{skill}" />
			</li>
		{/each}
	</ul>
</article>

<style>
	h3 {
		margin-bottom: 1rem;
	}

	article {
		padding-top: 1rem;
		padding-bottom: 0rem;
	}

	header {
		padding-top: 1rem;
		padding-bottom: 0rem;
		margin-bottom: 0rem;
	}

	#tools-list {
		font-size: 3rem;
		padding-left: 0;
		list-style: none;
		text-align: center;
	}

	.dev-logo {
		display: inline-block;
		margin-right: 0.25rem;
		margin-left: 0.25rem;
	}

	.dev-logo:first-child {
		margin-left: 0rem;
	}

	.dev-logo:last-child {
		margin-right: 0rem;
	}
</style>
