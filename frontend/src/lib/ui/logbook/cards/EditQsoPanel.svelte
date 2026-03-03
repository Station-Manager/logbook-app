<script lang="ts">
    import {onMount} from "svelte";
    import {qsoEditState} from "$lib/states/qso-edit-state.svelte";
    import FormControls from "$lib/ui/logbook/components/FormControls.svelte";
    import {inputBase, inputBaseUppercase, parseDatabaseFreqToDottedKhz} from "@station-manager/shared-utils";
    import {types} from "$lib/wailsjs/go/models";
    import {GetQsoById} from "$lib/wailsjs/go/facade/Service";
    import {handleAsyncError} from "$lib/utils/error-handler";
    import {getFocusContext} from "@station-manager/shared-utils/svelte";

    interface Props {
        qsoId: number
    }

    const focusContext = getFocusContext();

    let {qsoId}: Props = $props();
    let qso: types.Qso = $state(new types.Qso());
    let parsedFreq: string = $state('');
    let parsedFreqRx: string = $state('');

    const onCancel = (): void => {
        if (qsoEditState.panelOpen) {
            qsoEditState.panelOpen = false;
        }
    }

    const onSubmit = (): void => {}

    $effect(() => {
        parsedFreq = parseDatabaseFreqToDottedKhz(qso.freq);
        parsedFreqRx = parseDatabaseFreqToDottedKhz(qso.freq_rx);
    });

    onMount(async (): Promise<void> => {
       console.log("EditQsoPanel mounted:", {qsoId});

       try {
           await getFocusContext().focus('callsignInput');
           qso = await GetQsoById(qsoId);
       } catch (e: unknown) {
           handleAsyncError(e, 'EditQsoPanel.svelte:onMount')
       }
    });
</script>

<div>
    <div class="flex h-126">
    <div class="flex flex-row gap-x-4 h-32 w-full">
        <div class="w-37.5">
            <label for="call">Callsign</label>
            <input
                    bind:value={qso.call}
                    bind:this={focusContext.refs.callsignInput}
                    id="call"
                    autocomplete="off"
                    spellcheck="false"
                    class={inputBaseUppercase}/>
        </div>
        <div class="w-37.5">
            <label for="band">Band</label>
            <input
                    id="band"
                type="text"
                bind:value={qso.band}
                class={inputBase}
                autocomplete="off"
                spellcheck="false"
                disabled/>
        </div>
        <div class="w-37.5">
            <label for="freq">Frequency</label>
            <input
                    id="freq"
                    type="text"
                    bind:value={parsedFreq}
                    class={inputBase}
                    autocomplete="off"
                    spellcheck="false"
                    disabled/>
        </div>
        <div class="w-37.5">
            <label for="freq_rx">Frequency RX</label>
            <input
                    id="freq_rx"
                    type="text"
                    bind:value={parsedFreqRx}
                    class={inputBase}
                    autocomplete="off"
                    spellcheck="false"
                    disabled/>
        </div>
    </div>
    </div>
    <div class="flex justify-end mt-2">
        <FormControls submitFunc={onSubmit} cancelFunc={onCancel}/>
    </div>
</div>
