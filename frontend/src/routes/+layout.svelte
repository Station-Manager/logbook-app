<script lang="ts">
	import './layout.css';
	import MainNav from "$lib/ui/MainNav.svelte";

	let { children } = $props();
</script>

<header>
	<MainNav/>
</header>
<main>
	{@render children()}
</main>
