import { BLOG_URL } from '$env/static/private';

/** @type {import('./$types').PageServerLoad} */
export async function load() {
	try {
		const url = BLOG_URL;

		const res = await fetch(url);

		const blogs = await res.json();

		return {
			blogs
		};
	} catch (e) {
		console.log(`fetch error: ${e}`);
		throw e;
	}
}
