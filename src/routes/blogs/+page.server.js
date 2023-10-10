import { env } from '$env/dynamic/private';
import redis from '$lib/redis';

/** @type {import('./$types').PageServerLoad} */
export async function load() {
	try {
		const cacheKey = 'blogs';
		const cacheData = await redis.get(cacheKey);

		if (cacheData) {
			/**
			 * @type {Blog[]} blogs - Array of dev.to blogs
			 */
			const blogs = JSON.parse(cacheData);

			return {
				blogs
			};
		}

		const timestamp = new Date().toISOString();
		console.log(`[${timestamp}] Cache miss, getting data from dev.to API`);

		const url = env.BLOG_URL || '';
		const res = await fetch(url);

		/**
		 * @type {Blog[]} blogs - Array of dev.to blogs
		 */
		const blogs = await res.json();

		await redis.setex(cacheKey, 3600, JSON.stringify(blogs));

		return {
			blogs
		};
	} catch (e) {
		console.log(`error fetching blogs: ${e}`);
		throw e;
	}
}
