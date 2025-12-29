<script lang="ts">
    interface Props {
        dbFileList: string[];
    }

    let {
        dbFileList = [],
    }: Props = $props();

</script>

<div class="flex flex-row gap-x-2 p-8">
    <div class="flex flex-col w-[320px] border-r border-r-gray-300">
        <table>
            <tbody>
        {#each dbFileList as file (file)}
            <tr><td>{file}</td></tr>
        {/each}
            </tbody>
        </table>
    </div>
    <div class="flex flex-col w-full">Col 2</div>
</div>
