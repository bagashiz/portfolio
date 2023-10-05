<script>
	/** @type {string[]} skills - The list of programming languages and tools. */
	export let skills;
</script>

<article>
	<ul id="tools-list">
		{#each skills as skill}
			<li class="dev-logo">
				<i class="fab fa-{skill}" />
			</li>
		{/each}
	</ul>
</article>

<style>
	article {
		padding: 0rem;
		margin-top: 1rem;
		margin-bottom: 1rem;
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
