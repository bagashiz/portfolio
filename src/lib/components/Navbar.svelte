<script>
	import { ThemeSwitcher } from '$lib/scripts/minimal-theme-switcher';
	import { onMount } from 'svelte';

	onMount(() => {
		ThemeSwitcher.init();
	});
</script>

<nav id="navbar" class="container-fluid">
	<ul>
		<li>
			<a href="/" class="contrast"><strong>Resume</strong></a>
		</li>
	</ul>
	<ul>
		<li>
			<a href="/projects" class="secondary outline">Projects</a>
		</li>
		<li>
			<a href="/blogs" class="secondary outline">Blogs</a>
		</li>
		<li>
			<details role="list" dir="rtl">
				<summary aria-haspopup="listbox" class="secondary">Theme</summary>
				<ul role="listbox">
					<li><a href="./" data-theme-switcher="auto">Auto</a></li>
					<li><a href="./" data-theme-switcher="light">Light</a></li>
					<li><a href="./" data-theme-switcher="dark">Dark</a></li>
				</ul>
			</details>
		</li>
	</ul>
</nav>

<style>
	#navbar {
		view-transition-name: navbar;
	}
</style>
