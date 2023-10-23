<script>
	import EduCard from '$lib/components/EduCard.svelte';
	import JobCard from '$lib/components/JobCard.svelte';
	import SkillsCard from '$lib/components/SkillsCard.svelte';
	import SocialIcon from '$lib/components/SocialIcon.svelte';
	import {
		awards,
		certifications,
		educations,
		experiences,
		interests,
		skills,
		socials,
		volunteering,
		workflows
	} from '$lib/store';
</script>

<svelte:head>
	<title>Bagas Hizbullah's Resume</title>
</svelte:head>

<!-- Profile -->
<section id="profile">
	<hgroup>
		<h1><strong>Bagas Hizbullah</strong></h1>
		<ul id="info">
			<li class="icon-li">
				<i class="fa-solid fa-location-dot" />
				West Java, Indonesia
			</li>
			<li class="icon-li">
				<a href="mailto:bagash.office@simplelogin.com">
					<i class="fa-solid fa-envelope" />
					bagash.office@simplelogin.com
				</a>
			</li>
		</ul>
	</hgroup>

	<p id="bio">
		I am an Information Systems Student at
		<a href="https://ub.ac.id" target="_blank">Brawijaya University</a>, passionate about web
		development, Linux, and Free and Open Source Software (FOSS). I have practical experience in web
		development, with a focus on backend development using the Go programming language. In addition,
		I also have valuable insight into Linux system administration and cloud computing through
		practical experience with AWS and DigitalOcean.
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

	<h3>Programming Languages & Tools</h3>

	<SkillsCard {skills} />

	<h3>Workflows</h3>

	<div id="workflows">
		{#each workflows as workflow}
			<ul class="fa-ul">
				<li class="icon-li">
					<span class="fa-li">
						<i class="fas fa-check" />
					</span>
					{workflow}
				</li>
			</ul>
		{/each}
	</div>
</section>

<hr />

<!-- Interests -->
<section id="interests">
	<h2><strong>Interests</strong></h2>

	<div id="interests-desc">
		{#each interests as interest}
			<p>{interest}</p>
		{/each}
	</div>
</section>

<hr />

<!-- Awards & Certifications -->
<section id="awards-certifications">
	<h2><strong>Awards & Certifications</strong></h2>

	<div id="awards">
		<h3>Awards</h3>

		<ul class="fa-ul">
			{#each awards as award}
				<li class="award-cert-li icon-li">
					<p class="award-cert-p">
						<span class="fa-li"><i class="fas fa-trophy" /></span>
						<strong>
							{award.place}<sup>{award.suffix}</sup> Place &bull; {award.host} &bull; {award.competition}
						</strong>
					</p>
					<p>
						<em>({award.translation})</em>
					</p>
				</li>
			{/each}
		</ul>
	</div>

	<div id="certificates">
		<h3>Certificates</h3>

		<ul class="fa-ul">
			{#each certifications as cert}
				<li class="award-cert-li icon-li">
					<p class="award-cert-p">
						<span class="fa-li"><i class="fas fa-certificate" /></span>
						<strong>{cert.title}</strong>
					</p>
					<em>
						Credential ID: <a target="_blank" href={cert.credential_url}>{cert.credential_id}</a>
					</em>
				</li>
			{/each}
		</ul>
	</div>
</section>

<style>
	h2,
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

	#bio,
	#interests-desc {
		text-align: justify;
	}

	#socials {
		margin-top: 2.5rem;
	}

	#info {
		padding-left: 0rem;
	}

	.award-cert-p {
		margin-bottom: 0.5rem;
	}

	.award-cert-li {
		margin-bottom: 1rem;
	}

	.icon-li {
		list-style-type: none;
	}
</style>
