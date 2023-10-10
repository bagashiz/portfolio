# Portfolio

My personal portfolio website for showcasing my resume, projects, and blog posts. The site is built using SvelteKit and styled with Pico CSS.

The featured GitHub projects are dynamically retrieved through the power of the [GitHub GraphQL API](https://docs.github.com/en/graphql). The blog posts are seamlessly pulled in using the [Dev.to API](https://docs.dev.to/api). Additionally, Redis is used to cache the GitHub and Dev.to API responses for 1 hour to reduce the number of API calls. Icons are provided by [Font Awesome](https://fontawesome.com) through their kit from the CDN. I've also implemented the new [View Transition API](https://developer.mozilla.org/en-US/docs/Web/API/View_Transitions_API) feature to enhance the user experience.

## Demo

Check out the live demo at [bagashiz.me](https://bagashiz.me)!

## Dependencies

- [SvelteKit](https://kit.svelte.dev)
- [IORedis](https://github.com/redis/ioredis)
- [Pico CSS](https://picocss.com)
- [Font Awesome Icons](https://fontawesome.com)
