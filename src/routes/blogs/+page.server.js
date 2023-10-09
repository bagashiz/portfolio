import { env } from '$env/dynamic/private';

/** @type {import('./$types').PageServerLoad} */
export async function load() {
	try {
		const url = env.BLOG_URL || '';

		const res = await fetch(url);

		/**
		 * @type {Blog[]} blogs - Array of dev.to blogs
		 */
		const blogs = await res.json();

		return {
			blogs
		};
	} catch (e) {
		console.log(`fetch error: ${e}`);
		throw e;
	}
}
