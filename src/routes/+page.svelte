<script>
	import EduCard from '$lib/components/EduCard.svelte';
	import JobCard from '$lib/components/JobCard.svelte';
	import SkillsCard from '$lib/components/SkillsCard.svelte';
	import SocialIcon from '$lib/components/SocialIcon.svelte';
	import { educations, experiences, skills, socials, volunteering, workflows } from '$lib/store';
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

<hr />

<!-- Education -->
<section id="education">
	<h2><strong>Education</strong></h2>

	{#each educations as edu}
		<EduCard {...edu} />
	{/each}
</section>

<hr />

<!-- Skills -->
<section id="skills">
	<h2><strong>Skills</strong></h2>

	<SkillsCard {skills} />

	<h3>Workflows</h3>

	{#each workflows as workflow}
		<ul class="fa-ul">
			<li id="workflow">
				<span class="fa-li">
					<i class="fas fa-check" />
				</span>
				{workflow}
			</li>
		</ul>
	{/each}
</section>

<style>
	h3 {
		margin-bottom: 1rem;
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

	#socials {
		margin-top: 2.5rem;
	}

	#info {
		padding-left: 0rem;
	}

	#info li {
		list-style-type: none;
	}

	#workflow {
		list-style-type: none;
	}
</style>
