<script lang="ts">
    import {handleAsyncError} from "$lib/utils/error-handler";
    import {onMount} from "svelte";
    import {types} from "$lib/wailsjs/go/models";
    import {GetQsoSlice} from "$lib/wailsjs/go/facade/Service";
    import {formatDate, parseDatabaseFreqToDottedKhz} from "@station-manager/shared-utils";
    import {Checkbox} from "bits-ui";

    let pageNum = 1;
    let tableRows: types.Qso[] = $state<types.Qso[]>([]);
    let selections: number[] = $state([]);
    let allSelected = $state(false);

    const fetchPagedQsos = async (pageNum: number): Promise<void> => {
        try {
            tableRows = await GetQsoSlice(1, pageNum, 20);
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

    onMount(() => {
        fetchPagedQsos(pageNum);
    });
</script>

<div class="">
    Logbook Card {tableRows.length}
    <div role="row" class="flex flex-row font-bold border-b mb-[3px] border-gray-400">
        <div class="w-8">
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
        <div class="w-32">Call</div>
        <div class="w-14">Band</div>
        <div class="w-28">Frequency</div>
        <div class="w-12">Mode</div>
        <div>Notes</div>
    </div>
    <div role="table">
        {#each tableRows as qso, index (qso.id)}
            <div id="row-{index}" role="row" class="flex flex-row">
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
                <div class="w-32">{qso.call}</div>
                <div class="w-14">{qso.band}</div>
                <div class="w-28">{parseDatabaseFreqToDottedKhz(qso.freq)}</div>
                <div class="w-12">{qso.mode}</div>
                <div>{qso.notes}</div>
            </div>
        {/each}
    </div>
</div>
