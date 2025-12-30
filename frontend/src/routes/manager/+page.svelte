<script lang="ts">
    import ManagerCard from "$lib/ui/manager/cards/ManagerCard.svelte";
    import { types } from '$lib/wailsjs/go/models';

    interface Props {
        list: types.Logbook[];
    }

    export let data: Props = {list: []};
    const list = data.list;
</script>

<ManagerCard logbookList={list}/>
