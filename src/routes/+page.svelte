<script>
	import EduCard from '$lib/components/EduCard.svelte';
	import ExpCard from '$lib/components/ExpCard.svelte';
	import SkillsCard from '$lib/components/SkillsCard.svelte';
	import SocialIcon from '$lib/components/SocialIcon.svelte';
	import {
		awards,
		certifications,
		educations,
		interests,
		skills,
		socials,
		volunteer,
		workflows,
		works
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
		I am an Information Systems student at
		<a href="https://ub.ac.id" target="_blank">Brawijaya University</a>. I specialize in web
		development, with a specific focus on back-end development. I have practical experience working
		on back-end projects using technologies such as Go, Node.js, and PHP with the Laravel web
		framework. Additionally, I have acquired knowledge in Linux system administration and cloud
		computing through hands-on experience with AWS, GCP, and other cloud providers.
	</p>

	<!-- Socials -->
	<div id="socials">
		{#each socials as social}
			<SocialIcon {...social} />
		{/each}
	</div>
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

<!-- Work Experience -->
<section id="work">
	<h2><strong>Work Experience</strong></h2>

	{#each works as work}
		<ExpCard {...work} />
	{/each}
</section>

<hr />

<!-- Volunteer Experience -->
<section id="volunteer">
	<h2><strong>Volunteer Work</strong></h2>

	{#each volunteer as vol}
		<ExpCard {...vol} />
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
