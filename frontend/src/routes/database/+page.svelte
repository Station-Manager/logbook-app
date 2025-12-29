<script lang="ts">
    import DatabaseCard from "$lib/ui/database/cards/DatabaseCard.svelte";

    interface DatabasePageData {
        list: string[];
    }

    export let data: DatabasePageData = {list: []};
    const dbFileList = data.list;
</script>

<DatabaseCard dbFileList={dbFileList}/>
