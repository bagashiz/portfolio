import { env } from '$env/dynamic/private';
import redis from '$lib/scripts/redis';

/** @type {import('./$types').PageServerLoad} */
export async function load() {
	const key = 'blogs';
	const url = env.BLOG_URL;
	return {
		streamed: {
			/**
			 * @type {Promise<Blog[]>} blogs - Array of dev.to blogs
			 */
			blogs: new Promise((resolve) => {
				redis.get(key).then((data) => {
					if (data) {
						const blogs = JSON.parse(data);
						resolve(blogs);
					} else {
						const timestamp = new Date().toISOString();
						console.log(`[${timestamp}] Cache miss, getting data from dev.to API`);

						fetch(url)
							.then((res) => res.json())
							.then((blogs) => {
								redis.setex(key, 3600, JSON.stringify(blogs));
								resolve(blogs);
							});
					}
				});
			})
		}
	};
}
