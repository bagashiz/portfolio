# Portfolio

## Description

My personal portfolio website for showcasing my resume, projects, and blog posts. The site is built using Go, Templ, HTMX, and Pico CSS. The site also uses [KeyDB](https://keydb.dev) as a drop-in replacement for Redis.

The featured GitHub projects are dynamically retrieved through the power of the [GitHub GraphQL API](https://docs.github.com/en/graphql). The blog posts are seamlessly pulled in using the [Dev.to API](https://docs.dev.to/api). Additionally, KeyDB is used to cache the GitHub and Dev.to API responses for 1 hour to reduce the number of API calls. Icons are provided by [Font Awesome](https://fontawesome.com) through their kit from the CDN. I've also implemented the new [View Transition API](https://developer.mozilla.org/en-US/docs/Web/API/View_Transitions_API) feature to enhance the user experience.

There is an old stack version of the site in the `sveltekit` branch. The old stack version of the site was built using SveteKit.

## Demo

Check out the live demo at [bagashiz.xyz](https://bagashiz.xyz)!

## Dependencies

- [Go](https://golang.org) version 1.22 or higher
- [NPM](https://www.npmjs.com) (for installing [terser](https://terser.org))
- [Templ](https://templ.guide)
- [go-redis](https://github.com/redis/go-redis)
- [Taskfile](https://taskfile.dev)
