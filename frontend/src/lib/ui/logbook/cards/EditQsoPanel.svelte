<script lang="ts">
    import {getFocusContext} from "$lib/states/focus-context.svelte";
    import {qsoEditState} from "$lib/states/qso-edit-state.svelte";
    import {onMount} from "svelte";

    interface Props {
        qsoId: number
    }

    const focusContext = getFocusContext();

    let {qsoId}: Props = $props();


    onMount(() => {
       console.log("EditQsoPanel mounted:", {qsoId});
    });
</script>

<div>
    <div>
        <label for="call">Callsign</label>
        <input id="call"/>
    </div>
</div>
