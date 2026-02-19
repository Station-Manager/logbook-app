<script lang="ts">
    import {sendEmailFormState, sendEmailState} from "$lib/states/send-email-state.svelte";
    import {inputBase} from "@station-manager/shared-utils";

    interface Props {
        selections: number[]
    }
    let {selections}: Props = $props();

    const cancelClickHandler = ():void => {
        sendEmailState.dialogOpen = false;
    }
</script>

<div class="flex absolute top-23.25 w-full h-155 bg-gray-300/50 z-30 pt-32 justify-center">
    <div class="flex flex-col bg-white rounded-lg shadow-2xl w-100 h-42.5 z-40 p-6">
        <div>
            <div>
                <label class="flex flex-row text-sm/5 font-medium text-gray-900" for="email">Send selected ({selections.length}) QSOs to:</label>
                <div class="mt-2 w-[320px]">
                    <input
                            bind:value={sendEmailFormState.email}
                            id="email"
                            type="email"
                            placeholder="Enter email address"
                            class={inputBase}>
                </div>
            </div>
        </div>
        <div class="mt-6 flex items-center justify-end gap-x-6">
            <button
                    onclick={cancelClickHandler}
                    type="button"
                    class="cursor-pointer rounded-md bg-white px-2.5 py-1.5 text-sm font-semibold text-gray-900 shadow-xs inset-ring inset-ring-gray-300 hover:bg-gray-50 dark:bg-white/10 dark:text-white dark:shadow-none dark:inset-ring-white/5 dark:hover:bg-white/20">
                Cancel</button>
            <button type="submit" class="cursor-pointer rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Send</button>
        </div>
    </div>
</div>
