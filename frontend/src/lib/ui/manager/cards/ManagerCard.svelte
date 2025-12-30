<script lang="ts">
    import ManagerCardHeader from "$lib/ui/manager/cards/ManagerCardHeader.svelte";
    import ManagerCardContent from "$lib/ui/manager/cards/ManagerCardContent.svelte";
    import {types} from "$lib/wailsjs/go/models";

    interface Props {
        logbookList: types.Logbook[];
    }

    let {
        logbookList = []
    }: Props = $props();
</script>

<ManagerCardHeader/>
<ManagerCardContent logbookList={logbookList}/>