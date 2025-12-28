<script lang="ts">
	import './layout.css';
	import favicon from '../../static/favicon.png';

	let { children } = $props();
</script>

<header>
</header>
<main>
	{@render children()}
</main>
