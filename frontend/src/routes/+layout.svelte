<script lang="ts">
	import './layout.css';
	import MainNav from "$lib/ui/MainNav.svelte";
	import {SvelteToast} from "@zerodevx/svelte-toast";

	let { children } = $props();
</script>

<header>
	<SvelteToast/>
	<MainNav/>
</header>
<main>
	{@render children()}
</main>
