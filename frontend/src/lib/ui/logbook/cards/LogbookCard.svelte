<script lang="ts">
    import {handleAsyncError} from "$lib/utils/error-handler";
    import {onMount} from "svelte";
    import {types} from "$lib/wailsjs/go/models";
    import {GetQsoCount, GetQsoSlice} from "$lib/wailsjs/go/facade/Service";
    import {formatDate, parseDatabaseFreqToDottedKhz} from "@station-manager/shared-utils";
    import {Checkbox} from "bits-ui";
    import LogbookCardHeader from "$lib/ui/logbook/cards/LogbookCardHeader.svelte";
    import Pagination from "$lib/ui/logbook/components/Pagination.svelte";
    import {configState} from "$lib/states/config-state.svelte";
    import {sqlite} from "$lib/wailsjs/go/models";

    let pageNum = 1;

    let totalItems = $state(0);
    let tableRows: types.Qso[] = $state<types.Qso[]>([]);
    let selections: number[] = $state([]);
    let allSelected = $state(false);
    let pageSize = $state(configState.pageSize);
    let showEditPanel = $state(false);

    const fetchPagedQsos = async (pageNum: number): Promise<void> => {
        try {
            tableRows = await GetQsoSlice(configState.logbook.id, pageNum, pageSize, sqlite.Ordering.DESC);
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
            showEditPanel = true;
        } catch (e: unknown) {
            handleAsyncError(e, 'LogbookCard.svelte->editQso()');
        }
    }

    onMount(async (): Promise<void> => {
        try {
            totalItems = await GetQsoCount(configState.logbook.id);
            await fetchPagedQsos(pageNum);
        } catch (e: unknown) {
            handleAsyncError(e, 'LogbookCard.svelte->onMount()');
        }
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
        <div class="w-28">Call</div>
        <div class="w-44">Name</div>
        <div class="w-14">Band</div>
        <div class="w-28">Frequency</div>
        <div class="w-14">Mode</div>
        <div class="w-36">Country</div>
        <div class="w-8">
            <svg fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 16.5V9.75m0 0 3 3m-3-3-3 3M6.75 19.5a4.5 4.5 0 0 1-1.41-8.775 5.25 5.25 0 0 1 10.233-2.33 3 3 0 0 1 3.758 3.848A3.752 3.752 0 0 1 18 19.5H6.75Z" />
            </svg>
        </div>
        <div>
            <svg fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="mb-1 size-5 -rotate-30">
                <path stroke-linecap="round" stroke-linejoin="round" d="M6 12 3.269 3.125A59.769 59.769 0 0 1 21.485 12 59.768 59.768 0 0 1 3.27 20.875L5.999 12Zm0 0h7.5" />
            </svg>
        </div>
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
                <div class="w-28">{qso.call}</div>
                <div class="w-44 overflow-hidden text-nowrap text-ellipsis pr-2" title="{qso.name}">{qso.name}</div>
                <div class="w-14">{qso.band}</div>
                <div class="w-28">{parseDatabaseFreqToDottedKhz(qso.freq)}</div>
                <div class="w-14">{qso.mode}</div>
                <div class="w-36 overflow-hidden text-nowrap text-ellipsis" title="{qso.country}">{qso.country}</div>
                <div class="w-8">{@render uploadStatus(qso)}</div>
                <div class="w-8">{@render emailStatus(qso)}</div>
                <div class="flex text-xs items-center">
                    <button onclick={() => editQso(qso.id)} class="text-gray-400 font-semibold hover:text-indigo-600 cursor-pointer">Edit</button>
                </div>
            </div>
        {/each}
    </div>
</div>
{#if showEditPanel}
<div class="absolute top-25 left-0 w-full h-164.25 bg-gray-200 opacity-90 z-10">
    <div class="bg-white rounded-lg m-8 h-150">
Test
    </div>
</div>

{/if}
<Pagination callback={fetchPagedQsos} pageNum={pageNum} totalItems={totalItems} pageSize={pageSize}/>

{#snippet uploadStatus(qso: types.Qso)}
    <span class="flex items-center">
        {#if qso.qrzcom_qso_upload_status !== "Y"}
        <svg fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-4.25 text-yellow-600">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126ZM12 15.75h.007v.008H12v-.008Z" />
        </svg>
        {:else}
        <svg fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-5 text-green-600">
            <path stroke-linecap="round" stroke-linejoin="round" d="m4.5 12.75 6 6 9-13.5" />
        </svg>
        {/if}
    </span>
{/snippet}

{#snippet emailStatus(qso: types.Qso)}
    <span class="flex items-center">
        {#if qso.sm_fwrd_by_email_status !== "Y"}
        <svg fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-5 text-red-600">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18 18 6M6 6l12 12" />
        </svg>
        {:else}
        <svg fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-5 text-green-600">
            <path stroke-linecap="round" stroke-linejoin="round" d="m4.5 12.75 6 6 9-13.5" />
        </svg>
        {/if}
    </span>
{/snippet}
