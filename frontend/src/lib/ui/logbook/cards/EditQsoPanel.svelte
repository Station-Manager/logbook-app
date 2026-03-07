<script lang="ts">
    import {onMount} from "svelte";
    import {type QsoEditState, qsoEditState} from "$lib/states/qso-edit-state.svelte";
    import FormControls from "$lib/ui/logbook/components/FormControls.svelte";
    import {
        inputBase,
        inputBaseUppercase,
    } from "@station-manager/shared-utils";
    import {GetQsoById} from "$lib/wailsjs/go/facade/Service";
    import {handleAsyncError} from "$lib/utils/error-handler";
    import {getFocusContext} from "@station-manager/shared-utils/svelte";

    interface Props {
        qsoId: number
    }

    const focusContext = getFocusContext();

    let {qsoId}: Props = $props();
    let qso: QsoEditState = $state(qsoEditState);

    const onCancel = (): void => {
        if (qsoEditState.panelOpen) {
            qsoEditState.panelOpen = false;
        }
    }

    const onSubmit = (): void => {}

    onMount(async (): Promise<void> => {
       try {
           await getFocusContext().focus('callsignInput');
           qso.fromQso(await GetQsoById(qsoId));
       } catch (e: unknown) {
           handleAsyncError(e, 'EditQsoPanel.svelte:onMount')
       }
    });
</script>

<div>
    <div class="flex flex-col h-126">
        <div class="flex flex-row gap-x-4 h-20 w-full">
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
                        bind:value={qso.freq}
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
                        bind:value={qso.freq_rx}
                        class={inputBase}
                        autocomplete="off"
                        spellcheck="false"
                        disabled/>
            </div>
        </div>
        <div class="flex flex-row gap-x-4 h-32 w-full">
            <div class="w-37.5">
                <label for="date_on">Date On</label>
                <input
                        id="date_on"
                        type="date"
                        bind:value={qso.qso_date}
                        class={inputBase}
                        autocomplete="off"
                        spellcheck="false"/>
            </div>
            <div class="w-37.5">
                <label for="date_off">Date Off</label>
                <input
                        id="date_off"
                        type="date"
                        bind:value={qso.qso_date_off}
                        class={inputBase}
                        autocomplete="off"
                        spellcheck="false"/>
            </div>
            <div class="w-37.5">
                <label for="time_on">Time On</label>
                <input
                        id="time_on"
                        type="time"
                        bind:value={qso.time_on}
                        class={inputBase}
                        autocomplete="off"
                        spellcheck="false"/>
            </div>
            <div class="w-37.5">
                <label for="time_off">Time Off</label>
                <input
                        id="time_off"
                        type="time"
                        bind:value={qso.time_off}
                        class={inputBase}
                        autocomplete="off"
                        spellcheck="false"/>
            </div>
        </div>
    </div>
    <div class="flex justify-end mt-2">
        <FormControls submitFunc={onSubmit} cancelFunc={onCancel}/>
    </div>
</div>
