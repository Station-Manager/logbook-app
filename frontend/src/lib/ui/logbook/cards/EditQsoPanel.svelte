<script lang="ts">
//    import {getFocusContext} from "$lib/states/focus-context.svelte";
    import {onMount} from "svelte";
    import {qsoEditState} from "$lib/states/qso-edit-state.svelte";
    import FormControls from "$lib/ui/logbook/components/FormControls.svelte";

    interface Props {
        qsoId: number
    }

//    const focusContext = getFocusContext();

    let {qsoId}: Props = $props();

    const onCancel = (): void => {
        if (qsoEditState.panelOpen) {
            qsoEditState.panelOpen = false;
        }
    }

    const onSubmit = (): void => {}

    onMount(() => {
       console.log("EditQsoPanel mounted:", {qsoId});
    });
</script>

<div>
    <div class="h-129">
        <label for="call">Callsign</label>
        <input id="call"/>
    </div>
    <div class="flex justify-end">
        <FormControls submitFunc={onSubmit} cancelFunc={onCancel}/>
    </div>
</div>
