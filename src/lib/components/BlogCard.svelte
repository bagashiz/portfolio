<script>
	/** @type any */
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
		<a href={blog.url} target="_blank" rel="noopener noreferrer">
			<h2><strong>{blog.title}</strong></h2>
		</a>

		<small>Published on {publishDate}</small>

		<p>
			{blog.description}<a href={blog.url} target="_blank" rel="noopener noreferrer">Read More</a>
		</p>
	</div>
</article>

<style>
	header,
	h2,
	p {
		margin-bottom: 0.5rem;
	}

	small {
		color: var(--primary);
	}

	p {
		margin-top: 1rem;
	}

	a {
		text-decoration: none;
	}

	.blog-card {
		padding-top: 1rem;
		padding-bottom: 1rem;
		margin-bottom: 4.5rem;
		border-radius: 0.75rem;
	}

	.blog-card header {
		padding: 0rem;
		overflow: hidden;
	}

	.blog-card header img {
		width: 100%;
		height: 100%;
		object-fit: cover;
	}

	.blog-content {
		padding: 0.5rem;
	}
</style>
