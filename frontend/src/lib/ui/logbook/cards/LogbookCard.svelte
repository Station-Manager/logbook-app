<script lang="ts">
    import {handleAsyncError} from "$lib/utils/error-handler";
    import {onMount} from "svelte";
    import {types} from "$lib/wailsjs/go/models";
    import {GetQsoSlice} from "$lib/wailsjs/go/facade/Service";
    import {formatDate, parseDatabaseFreqToDottedKhz} from "@station-manager/shared-utils";
    import {Checkbox} from "bits-ui";
    import LogbookCardHeader from "$lib/ui/logbook/cards/LogbookCardHeader.svelte";
    import Pagination from "$lib/ui/logbook/components/Pagination.svelte";
    import {configState} from "$lib/states/config-state.svelte";
    import {logbookListState} from "$lib/states/logbook-list-state.svelte";

    let pageNum = 1;
    let tableRows: types.Qso[] = $state<types.Qso[]>([]);
    let selections: number[] = $state([]);
    let allSelected = $state(false);
    let pageSize = $state(configState.pageSize);

    const fetchPagedQsos = async (pageNum: number): Promise<void> => {
        try {
            tableRows = await GetQsoSlice(1, pageNum, 24);
        } catch (e: unknown) {
            handleAsyncError(e, 'LogbookCard.svelte->fetchPagedQsos()');
        }
    }

    const select = (rowId: number): void => {
        if (!selections.includes(rowId)) {
            selections.push(rowId)
            allSelected = selections.length >= tableRows.length;
        }
//        menuClose();
    };

    const deselect = (rowId: number): void => {
        selections = selections.filter(id => id !== rowId);
        allSelected = selections.length >= tableRows.length;
        // menuClose();
    };

    const toggleAll = (v: boolean) => {
        allSelected = v;
        if (v) {
            for (let row of tableRows) {
                select(row.id);
            }
        } else {
            selections = [];
        }
    };

    const editQso = (qsoId: number): void => {
        try {
            console.log(`Edit QSO with ID: ${qsoId}`);
        } catch (e: unknown) {
            handleAsyncError(e, 'LogbookCard.svelte->editQso()');
        }
    }

    const emailSentStatusCss = (qso: types.Qso): string => {
        let css = "flex font-semibold w-14";
        if (qso.sm_fwrd_by_email_status !== "Y") {
            css += " text-red-700";
        } else {
            css += " text-green-700";
        }
        return css;
    }

    onMount(() => {
        fetchPagedQsos(pageNum);
    });
</script>

<LogbookCardHeader/>
<div class="h-153">
    <div role="row" class="flex flex-row items-center font-bold border-b mb-0.75 h-8 border-gray-400">
        <div class="w-8 pt-1">
            <Checkbox.Root bind:checked={allSelected} onCheckedChange={(v) => {toggleAll(v);}} class="ring w-4 h-4 mx-2 rounded ring-gray-400">
                {#snippet children({checked})}
                    {#if checked}
                        <svg fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-4 text-white bg-indigo-500">
                            <path stroke-linecap="round" stroke-linejoin="round" d="m4.5 12.75 6 6 9-13.5" />
                        </svg>
                    {/if}
                {/snippet}
            </Checkbox.Root>
        </div>
        <div class="w-28">Date</div>
        <div class="w-24">Call</div>
        <div class="w-44">Name</div>
        <div class="w-14">Band</div>
        <div class="w-28">Frequency</div>
        <div class="w-14">Mode</div>
        <div class="w-32">Country</div>
        <div class="w-42">Notes</div>
        <div>&nbsp;</div>
    </div>
    <div role="table">
        {#each tableRows as qso, index (qso.id)}
            <div id="row-{index}" role="row" class="flex flex-row odd:bg-white even:bg-gray-200">
                <Checkbox.Root
                        checked={selections.includes(qso.id)} onCheckedChange={(v) => {if (v) {select(qso.id);} else {deselect(qso.id);}}}
                        class="ring rounded mt-1 mx-2 ring-gray-400 w-4 h-4">
                    {#snippet children({checked})}
                        {#if checked}
                            <svg fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-4 text-white bg-indigo-500">
                                <path stroke-linecap="round" stroke-linejoin="round" d="m4.5 12.75 6 6 9-13.5" />
                            </svg>
                        {/if}
                    {/snippet}
                </Checkbox.Root>
                <div class="w-28">{formatDate(qso.qso_date)}</div>
                <div class="w-28">{@render status(qso)}</div>
                <div class="w-44 overflow-hidden text-nowrap text-ellipsis pr-2" title="{qso.name}">{qso.name}</div>
                <div class="w-14">{qso.band}</div>
                <div class="w-28">{parseDatabaseFreqToDottedKhz(qso.freq)}</div>
                <div class="w-14">{qso.mode}</div>
                <div class="w-32 overflow-hidden text-nowrap text-ellipsis" title="{qso.country}">{qso.country}</div>
                <div class="w-51">{qso.notes}</div>
                <div class="flex text-xs items-center">
                    <button onclick={() => editQso(qso.id)} class="text-gray-400 font-semibold hover:text-indigo-600 cursor-pointer">Edit</button>
                </div>
            </div>
        {/each}
    </div>
</div>
<Pagination callback={fetchPagedQsos} pageNum={pageNum} totalItems={logbookListState.list.length} pageSize={pageSize}/>

{#snippet status(qso: types.Qso)}
    <span class="flex items-center">
        <span class={emailSentStatusCss(qso)}>
            {qso.call}
        </span>
        <span>
            {#if qso.qrzcom_qso_upload_status !== "Y"}
            <svg fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="ml-2 size-4.25 text-yellow-600"><path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126ZM12 15.75h.007v.008H12v-.008Z" /></svg>
            {/if}
        </span>
    </span>
{/snippet}
