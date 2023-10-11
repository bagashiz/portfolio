import { env } from '$env/dynamic/private';
import redis from '$lib/scripts/redis';

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
          primaryLanguage {
            name
			color
          }
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
		const cacheKey = 'projects';
		const cacheData = await redis.get(cacheKey);

		if (cacheData) {
			/**
			 * @type {Project[]} projects - Array of pinned repositories
			 */
			const projects = JSON.parse(cacheData);

			return {
				projects
			};
		}

		const timestamp = new Date().toISOString();
		console.log(`[${timestamp}] cache miss, getting data from GitHub API`);

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

		await redis.setex(cacheKey, 3600, JSON.stringify(projects));

		return {
			projects
		};
	} catch (e) {
		console.log(`error fetching projects: ${e}`);
		throw e;
	}
}
