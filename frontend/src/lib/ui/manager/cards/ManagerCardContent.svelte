<script lang="ts">
    import {types} from "$lib/wailsjs/go/models";

    interface Props {
        logbookList: types.Logbook[];
    }

    let {
        logbookList = [],
    }: Props = $props();
</script>

<div class="flex flex-row gap-x-2 p-8">
    <div class="flex flex-col w-[460px] border-r border-r-gray-300 pr-4">
        {#each logbookList as lb (lb.id)}
        <div class="flex flex-col border border-gray-300 p-4 rounded-md hover:border-gray-400 hover:cursor-pointer">
            <div class="font-semibold">{lb.name}</div>
            <div class="text-sm text-gray-400">{lb.description}</div>
        </div>
        {/each}
    </div>
    <div class="flex flex-col w-full">Col 2</div>
</div>
