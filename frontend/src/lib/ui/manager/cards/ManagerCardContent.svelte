<script lang="ts">
    import {types} from "$lib/wailsjs/go/models";

    interface Props {
        logbookList: types.Logbook[];
    }

    let {
        logbookList = [],
    }: Props = $props();

    const newLogbookAction = (): void => {
        console.log("New logbook action");
    }

    const loadLogbookAction = (id: number): void => {
        console.log("Load logbook action", id);
    }
</script>

<div class="flex flex-row gap-x-2 p-8">
    <div class="flex flex-col w-[460px] h-[580px] border-r border-r-gray-300 pr-4 gap-y-4">
        <button onclick={newLogbookAction} class="flex flex-col border border-gray-300 p-4 rounded-md hover:border-gray-400 hover:cursor-pointer">
            <span class="font-semibold">New Log Book</span>
        </button>
        {#each logbookList as lb (lb.id)}
        <button onclick={() => loadLogbookAction(lb.id)} class="flex flex-col border border-gray-300 p-4 rounded-md text-left hover:border-gray-400 hover:cursor-pointer">
            <span class="font-semibold">{lb.name}</span>
            <span class="text-sm text-gray-400">{lb.description}</span>
        </button>
        {/each}
    </div>
    <div class="flex flex-col w-full">

    </div>
</div>
