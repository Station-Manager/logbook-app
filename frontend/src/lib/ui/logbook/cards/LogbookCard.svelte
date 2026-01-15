<script lang="ts">
    import {handleAsyncError} from "$lib/utils/error-handler";
    import {onMount} from "svelte";
    import {types} from "$lib/wailsjs/go/models";
    import {GetQsoSlice} from "$lib/wailsjs/go/facade/Service";
    import {formatDate} from "../../../../../../../shared-utils/src";

    let pageNum = 1;
    let tableRows: types.Qso[] = $state<types.Qso[]>([]);

    const fetchPagedQsos = async (pageNum: number): Promise<void> => {
        try {
            tableRows = await GetQsoSlice(1, pageNum, 20);
        } catch (e: unknown) {
            handleAsyncError(e, 'LogbookCard.svelte->fetchPagedQsos()');
        }
    }

    onMount(() => {
        fetchPagedQsos(pageNum);
    });
</script>

<div class="">
    Logbook Card {tableRows.length}
    <div role="table">
        {#each tableRows as qso, index (qso.id)}
            <div id="row-{index}" role="row" class="flex flex-row">
                <div>{formatDate(qso.qso_date)}</div>
                <div>{qso.freq}</div>
            </div>
        {/each}
    </div>
</div>
