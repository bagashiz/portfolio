<script>
	import { page } from '$app/stores';

	/** @type {Record<number, string>} map of status codes to emojis */
	const emojis = {
		400: '❌', // bad request
		401: '🔒', // unauthorized
		403: '🚫', // forbidden
		404: '🔍', // not found
		500: '🛠', // internal server error
		502: '🔌', // bad gateway
		503: '🔧', // service unavailable
		504: '⏳' // gateway timeout
	};
</script>

<section id="error" class="container">
	<h1><strong>{$page.status}</strong></h1>
	<h2>{$page.error?.message}</h2>
	<span id="emoji">
		{emojis[$page.status] ?? emojis[500]}
	</span>
</section>

<style>
	h1,
	h2 {
		margin-bottom: 1rem;
	}

	#error {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		text-align: center;
		margin: 0rem;
	}

	#emoji {
		font-size: 10rem;
	}
</style>
