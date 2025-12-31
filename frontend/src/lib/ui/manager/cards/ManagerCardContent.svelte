<script lang="ts">
    import {types} from "$lib/wailsjs/go/models";

    interface Props {
        logbookList: types.Logbook[];
    }

    let {
        logbookList = [],
    }: Props = $props();
    let inputDisabled: boolean = $state(true);

    const newLogbookAction = (): void => {
        console.log("New logbook action");
    }

    const loadLogbookAction = (id: number): void => {
        console.log("Load logbook action", id);
    }
</script>

<div class="flex flex-row gap-x-2 p-8">
    <div class="flex flex-col w-[580px] h-[580px] border-r border-r-gray-300 pr-10 gap-y-4">
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
    <div class="flex flex-col w-full pl-10">
        <div class="w-[360px]">
            <div class="flex flex-col gap-y-4">
                <div>
                    <label for="name" class="block text-sm/6 font-medium text-gray-900 dark:text-white">Name</label>
                    <div class="mt-2">
                        <div class="flex items-center rounded-md bg-white pl-3 outline-1 -outline-offset-1 outline-gray-300 focus-within:outline-2 focus-within:-outline-offset-2 focus-within:outline-indigo-600 dark:bg-white/5 dark:outline-white/10 dark:focus-within:outline-indigo-500">
                            <input disabled={inputDisabled} id="name" type="text" placeholder="Log book name" class="block bg-white py-1.5 pr-3 pl-1 text-base text-gray-900 placeholder:text-gray-400 focus:outline-none"/>
                        </div>
                    </div>
                </div>
                <div>
                    <label for="about" class="block text-sm/6 font-medium text-gray-900">Description</label>
                    <div class="mt-2">
                        <textarea disabled={inputDisabled} id="about" name="about" rows="3" class="resize-none block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600"></textarea>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
