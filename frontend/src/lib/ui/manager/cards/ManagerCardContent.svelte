<script lang="ts">
    import {types} from "$lib/wailsjs/go/models";
    import {configState} from "$lib/states/config-state.svelte";
    import {onMount} from "svelte";
    import {logbookState} from "$lib/states/logbook-state.svelte";

    interface Props {
        logbookList: types.Logbook[];
    }

    let nameInput: HTMLInputElement;
    let callsignInput: HTMLInputElement;
    let descriptionInput: HTMLTextAreaElement;

    let {
        logbookList = [],
    }: Props = $props();
    let inputDisabled: boolean = $state(true);

    const newLogbookAction = (): void => {
        logbookState.clear();
        inputDisabled = false;
        nameInput.focus();
    }

    const loadLogbookAction = (id: number): void => {
        console.log("Load logbook action", id);
    }

    onMount((): void => {
        logbookState.fromLogbook(configState.logbook);
    });
</script>

<div class="flex flex-row gap-x-2 p-8">
    <div class="flex flex-col w-[580px] h-[580px] border-r border-r-gray-300 pr-10 gap-y-4">
        <button onclick={newLogbookAction} class="flex flex-col border border-gray-300 p-4 rounded-md hover:border-gray-400 hover:cursor-pointer">
            <span class="font-semibold">New Log Book</span>
        </button>
        {#each logbookList as lb (lb.id)}
        <button disabled={lb.id === configState.logbook.id} onclick={() => loadLogbookAction(lb.id)} class="flex flex-col border border-gray-300 p-4 rounded-md text-left hover:border-gray-400 hover:cursor-pointer disabled:hover:border-gray-300 disabled:cursor-not-allowed">
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
                        <div class="flex items-center rounded-md bg-white pl-3 outline-1 -outline-offset-1 outline-gray-300 focus-within:outline-2 focus-within:-outline-offset-2 focus-within:outline-indigo-600">
                            <input
                                    bind:this={nameInput}
                                    bind:value={logbookState.name}
                                    disabled={inputDisabled}
                                    id="name"
                                    type="text"
                                    placeholder="Log book name"
                                    class="block bg-white py-1.5 pr-3 pl-1 text-base text-gray-900 placeholder:text-gray-400 focus:outline-none"/>
                        </div>
                    </div>
                </div>
                <div>
                    <label for="callsign" class="block text-sm/6 font-medium text-gray-900 dark:text-white">Callsign</label>
                    <div class="mt-2">
                        <div class="flex items-center rounded-md bg-white pl-3 outline-1 -outline-offset-1 outline-gray-300 focus-within:outline-2 focus-within:-outline-offset-2 focus-within:outline-indigo-600 dark:bg-white/5 dark:outline-white/10 dark:focus-within:outline-indigo-500">
                            <input
                                    bind:this={callsignInput}
                                    bind:value={logbookState.callsign}
                                    disabled={inputDisabled}
                                    id="callsign"
                                    type="text"
                                    placeholder="Callsign"
                                    class="block bg-white py-1.5 pr-3 pl-1 text-base text-gray-900 placeholder:text-gray-400 focus:outline-none"/>
                        </div>
                    </div>
                </div>
                <div>
                    <label for="description" class="block text-sm/6 font-medium text-gray-900">Description</label>
                    <div class="mt-2">
                        <textarea
                                bind:this={descriptionInput}
                                bind:value={logbookState.description}
                                disabled={inputDisabled}
                                id="description"
                                rows="3"
                                class="resize-none block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600"></textarea>
                    </div>
                </div>
            </div>
        </div>
        <div class="flex justify-between mt-4">
            <div>
            <button class="rounded-md bg-red-600 px-2.5 py-1.5 text-sm font-semibold text-white shadow-xs hover:bg-red-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-red-600">Delete</button>
            </div>
            <div class="flex gap-x-4">
            <button class="rounded-md bg-indigo-600 px-2.5 py-1.5 text-sm font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Select</button>
            <button class="rounded-md bg-indigo-600 px-2.5 py-1.5 text-sm font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Update</button>
            <button class="rounded-md bg-white px-2.5 py-1.5 text-sm font-semibold text-gray-900 shadow-xs inset-ring inset-ring-gray-300 hover:bg-gray-50">Cancel</button>
            </div>
        </div>
    </div>
</div>
