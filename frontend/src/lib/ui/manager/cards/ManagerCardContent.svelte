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
    <div class="flex flex-col w-[320px] border-r border-r-gray-300">
        <table>
            <thead>
                <tr>
                    <th>Logbooks</th>
                </tr>
            </thead>
            <tbody>
        {#each logbookList as lb (lb.id)}
            <tr><td>{lb.name}</td></tr>
        {/each}
            </tbody>
        </table>
    </div>
    <div class="flex flex-col w-full">Col 2</div>
</div>
