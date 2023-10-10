<script>
	/** @type {string} href - The link to the social media. */
	export let href;
	/** @type {string} rel - The rel attribute of the link. */
	export let rel;
	/** @type {string} fa_class - The Font Awesome class name. */
	export let fa_class;
</script>

<a class="social-icon" role="button" target="_blank" {rel} {href}>
	<i class="fab fa-{fa_class}" />
</a>

<style>
	.social-icon {
		display: inline-flex;
		align-items: center;
		justify-content: center;
		height: 3.5rem;
		width: 3.5rem;
		background-color: var(--primary);
		color: var(--primary-inverse);
		border-radius: 100%;
		font-size: 1.5rem;
		margin-right: 0.5rem;
		margin-left: 0.5rem;
	}

	.social-icon:first-child {
		margin-left: 0;
	}

	.social-icon:last-child {
		margin-right: 0;
	}

	.social-icon:hover {
		background-color: var(--primary-hover);
	}
</style>
