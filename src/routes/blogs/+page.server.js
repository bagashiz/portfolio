import { BLOG_URL } from '$env/static/private';

/** @type {import('./$types').PageServerLoad} */
export async function load() {
	const url = BLOG_URL;

	const res = await fetch(url);

	const blogs = await res.json();

	return {
		blogs
	};
}
