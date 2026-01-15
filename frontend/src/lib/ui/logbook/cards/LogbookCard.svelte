<script lang="ts">
    import {handleAsyncError} from "$lib/utils/error-handler";
    import {onMount} from "svelte";
    import {types} from "$lib/wailsjs/go/models";
    import {GetQsoSlice} from "$lib/wailsjs/go/facade/Service";

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

<div>
    Logbook Card {tableRows.length}
</div>
