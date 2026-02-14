<script lang="ts">
//    import {getFocusContext} from "$lib/states/focus-context.svelte";
    import {onMount} from "svelte";
import {qsoEditState} from "$lib/states/qso-edit-state.svelte";

    interface Props {
        qsoId: number
    }

//    const focusContext = getFocusContext();

    let {qsoId}: Props = $props();

    const onCancel = () => {
        if (qsoEditState.panelOpen) {
            qsoEditState.panelOpen = false;
        }
    }

    onMount(() => {
       console.log("EditQsoPanel mounted:", {qsoId});
    });
</script>

<div>
    <div class="h-129">
        <label for="call">Callsign</label>
        <input id="call"/>
    </div>
    <div class="flex gap-2 justify-end">
        <button class="">Save</button>
        <button class="" onclick={onCancel}>Cancel</button>
    </div>
</div>
