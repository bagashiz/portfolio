import { env } from '$env/dynamic/private';

/**
 * @type {string} query - GraphQL query to fetch pinned github repositories
 */
const query = `{
  user(login: "bagashiz") {
    pinnedItems(first: 6, types: REPOSITORY) {
      nodes {
        ... on Repository {
          name
          description
          url
          stargazers {
            totalCount
          }
          forks {
            totalCount
          }
        }
      }
    }
  }
}`;

/** @type {import('./$types').PageServerLoad} */
export async function load() {
	try {
		const token = env.GITHUB_ACCESS_TOKEN || '';

		const res = await fetch('https://api.github.com/graphql', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `bearer ${token}`
			},
			body: JSON.stringify({ query })
		});

		const { data } = await res.json();

		/**
		 * @type {Project[]} projects - Array of pinned repositories
		 */
		const projects = data.user.pinnedItems.nodes;

		return {
			projects
		};
	} catch (e) {
		console.log(`fetch error: ${e}`);
		throw e;
	}
}
