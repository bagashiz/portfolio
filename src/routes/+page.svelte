<script>
	import JobCard from '$lib/components/JobCard.svelte';
	import SocialIcon from '$lib/components/SocialIcon.svelte';
	import { experiences, socials, volunteering } from '$lib/store';
</script>

<!-- Profile -->
<section id="profile">
	<hgroup>
		<h1><strong>Bagas Hizbullah</strong></h1>
		<ul id="info">
			<li>
				<i class="fa-solid fa-location-dot" />
				West Java, Indonesia
			</li>
			<li>
				<a href="mailto:bagash.office@simplelogin.com">
					<i class="fa-solid fa-envelope" />
					bagash.office@simplelogin.com
				</a>
			</li>
		</ul>
	</hgroup>

	<p>
		I am an Information Systems Student at
		<a href="https://ub.ac.id" target="_blank">Brawijaya University</a>, passionate about web
		development, Linux, and Free and Open Source Software (FOSS). I am eager to explore
		opportunities and contribute significantly to the IT industry.
	</p>

	<!-- Socials -->
	<div id="socials">
		{#each socials as social}
			<SocialIcon {...social} />
		{/each}
	</div>
</section>

<hr />

<!-- Experience -->
<section id="experience">
	<h2><strong>Experience</strong></h2>

	{#each experiences as exp}
		<JobCard {...exp} />
	{/each}
</section>

<hr />

<!-- Volunteering -->
<section id="volunteering">
	<h2><strong>Volunteering</strong></h2>

	{#each volunteering as vol}
		<JobCard {...vol} />
	{/each}
</section>

<style>
	#socials {
		margin-top: 2.5rem;
	}

	#info {
		padding-left: 0rem;
	}

	#info li {
		list-style-type: none;
	}

	hr {
		border: 0.05rem solid;
		opacity: 0.2;
		margin-top: 2rem;
		margin-bottom: 2rem;
	}

	i {
		color: var(--primary);
	}
</style>
