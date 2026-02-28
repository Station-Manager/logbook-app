<script lang="ts">
//    import {getFocusContext} from "$lib/states/focus-context.svelte";
    import {onMount} from "svelte";
    import {qsoEditState} from "$lib/states/qso-edit-state.svelte";
    import FormControls from "$lib/ui/logbook/components/FormControls.svelte";
    import {inputBaseUppercase} from "@station-manager/shared-utils";

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
    <div class="flex h-126">
    <div class="h-32 w-full border">
        <div class="w-37.5">
            <label for="call">Callsign</label>
            <input
                    id="call"
                    autocomplete="off"
                    spellcheck="false"
                    class={inputBaseUppercase}/>
        </div>
    </div>
    </div>
    <div class="flex justify-end mt-2">
        <FormControls submitFunc={onSubmit} cancelFunc={onCancel}/>
    </div>
</div>
