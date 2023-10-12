import { env } from '$env/dynamic/private';
import redis from '$lib/scripts/redis';

/**
 * @type {string} url - Github GraphQL API endpoint
 */
const url = 'https://api.github.com/graphql';

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
	const key = 'projects';
	const token = env.GITHUB_ACCESS_TOKEN;
	const reqInfo = {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `bearer ${token}`
		},
		body: JSON.stringify({ query })
	};

	return {
		streamed: {
			/**
			 * @type {Promise<Project[]>} projects - Array of GitHub projects
			 */
			projects: new Promise((resolve) => {
				redis.get(key).then((data) => {
					if (data) {
						const projects = JSON.parse(data);
						resolve(projects);
					} else {
						const timestamp = new Date().toISOString();
						console.log(`[${timestamp}] Cache miss, getting data from GitHub API`);

						fetch(url, reqInfo)
							.then((res) => res.json())
							.then(({ data }) => {
								const projects = data.user.pinnedItems.nodes;
								redis.setex(key, 3600, JSON.stringify(projects));
								resolve(projects);
							});
					}
				});
			})
		}
	};
}
