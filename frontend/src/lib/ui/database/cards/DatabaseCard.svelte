<script lang="ts">
    import DatabaseCardHeader from "$lib/ui/database/cards/DatabaseCardHeader.svelte";
    import DatabaseCardContent from "$lib/ui/database/cards/DatabaseCardContent.svelte";

    interface Props {
        dbFileList: string[];
    }

    let {
        dbFileList = []
    }: Props = $props();
</script>

<DatabaseCardHeader/>
<DatabaseCardContent dbFileList={dbFileList}/>