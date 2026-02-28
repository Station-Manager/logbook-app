<script lang="ts">
//    import {getFocusContext} from "$lib/states/focus-context.svelte";
    import {onMount} from "svelte";
    import {qsoEditState} from "$lib/states/qso-edit-state.svelte";
    import FormControls from "$lib/ui/logbook/components/FormControls.svelte";
    import {inputBaseUppercase} from "@station-manager/shared-utils";
    import {types} from "$lib/wailsjs/go/models";
    import {GetQsoById} from "$lib/wailsjs/go/facade/Service";
    import {handleAsyncError} from "$lib/utils/error-handler";

    interface Props {
        qsoId: number
    }

//    const focusContext = getFocusContext();

    let {qsoId}: Props = $props();
    let qso: types.Qso = $state(new types.Qso());

    const onCancel = (): void => {
        if (qsoEditState.panelOpen) {
            qsoEditState.panelOpen = false;
        }
    }

    const onSubmit = (): void => {}

    onMount(async (): Promise<void> => {
       console.log("EditQsoPanel mounted:", {qsoId});
       try {
           qso = await GetQsoById(qsoId);
       } catch (e: unknown) {
           handleAsyncError(e, 'EditQsoPanel.svelte:onMount')
       }
    });
</script>

<div>
    <div class="flex h-126">
    <div class="h-32 w-full border">
        <div class="w-37.5">
            <label for="call">Callsign</label>
            <input
                    bind:value={qso.call}
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
