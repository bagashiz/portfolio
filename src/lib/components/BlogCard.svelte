<script>
	/** @type {Blog} */
	export let blog;

	/**
	 * @type {string} publishDate - The date the blog was published in "Month Day, Year" format
	 */
	const publishDate = new Date(blog.published_at).toLocaleDateString('en-GB', {
		year: 'numeric',
		month: 'long',
		day: 'numeric'
	});
</script>

<article class="blog-card">
	<header>
		<img src={blog.cover_image} alt={blog.title} />
	</header>

	<div class="blog-content">
		<hgroup>
			<a href={blog.url} target="_blank" rel="noopener noreferrer">
				<h2><strong>{blog.title}</strong></h2>
			</a>
			<small>Published on {publishDate}</small>
		</hgroup>

		<div class="badges">
			<a href={blog.url} class="badge outline" role="button" target="_blank">
				<i class="fa-solid fa-heart" />
				<span>{blog.public_reactions_count}</span>
			</a>
			<a href={blog.url} class="badge outline" role="button" target="_blank">
				<i class="fa-solid fa-comment" />
				<span>{blog.comments_count}</span>
			</a>
			<a href={blog.url} class="badge outline" role="button" target="_blank">
				<i class="fa-solid fa-clock" />
				<span>{blog.reading_time_minutes} min.</span>
			</a>
		</div>

		<p>
			{blog.description}<a href={blog.url} target="_blank" rel="noopener noreferrer">Read More</a>
		</p>
	</div>
</article>

<style>
	header,
	h2 {
		margin-bottom: 0rem;
	}

	hgroup {
		margin-bottom: 0.5rem;
	}

	small {
		color: var(--primary);
	}

	p {
		margin-top: 0.5rem;
		margin-bottom: 0.5rem;
	}

	a {
		text-decoration: none;
	}

	.blog-card {
		padding-top: 1rem;
		padding-bottom: 1rem;
		border-radius: 0.75rem;
		overflow: hidden;
	}

	.blog-card header {
		padding: 0rem;
	}

	.blog-card header img {
		width: 100%;
		height: 100%;
		object-fit: cover;
	}

	.blog-content {
		padding: 0.5rem;
	}

	.badges {
		display: flex;
		padding-top: 0.25rem;
		padding-bottom: 0.25rem;
	}

	.badge {
		display: flex;
		align-items: center;
		font-size: medium;
		font-weight: bold;
		margin-right: 0.5rem;
		padding: 0.25rem;
	}

	.badge i {
		margin-right: 0.25rem;
	}
</style>
