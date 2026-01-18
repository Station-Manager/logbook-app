<script lang="ts">
    import {onMount} from "svelte";

    interface Props {
        callback: (page: number) => void;
        pageNum: number;
        totalItems: number;
        pageSize: number;
    }

    let {callback = $bindable(), pageNum, totalItems, pageSize}: Props = $props();
    let startIndex: number = $state(0);
    let pageCount: number = $state(0); // Total number of pages
    let endIndex: number = $state(0); // Number of items on the last page

    const update = (pgNum: number): {startInx: number, endInx: number} => {
        if (pgNum < 1) {
            return {startInx: 0, endInx: 0};
        }
        const start = (pgNum * pageSize) - pageSize + 1;
        const end = pgNum * pageSize;
        if (end > totalItems) {
            return {startInx: start, endInx: totalItems};
        }
        return {startInx: start, endInx: end};
    }

    const onClickNext = () => {
        if (pageNum >= pageCount) {
            return;
        }
        pageNum++;
        const up = update(pageNum);
        startIndex = up.startInx;
        endIndex = up.endInx;
        callback(pageNum);
    }

    const onClickPrevious = () => {
        if (pageNum <= 1) {
            return;
        }
        pageNum--;
        const up = update(pageNum);
        startIndex = up.startInx;
        endIndex = up.endInx;
        callback(pageNum);
    }

    onMount(() => {
        const up = update(pageNum);
        startIndex = up.startInx;
        endIndex = up.endInx;
        pageCount = Math.ceil(totalItems / pageSize);
    });
</script>

<div class="flex flex-row h-11.75 bg-white items-center justify-between mr-2">
    <div></div>
    <div class="text-sm text-gray-700">
        <p>
            Showing
            <span class="font-medium">{startIndex}</span>
            to
            <span class="font-medium">{endIndex}</span>
            of
            <span class="font-medium">{totalItems}</span>
            entries
        </p>
    </div>
    <div class="flex flex-row space-x-2">
        <div>
            <button onclick={onClickPrevious} class="text-sm cursor-pointer w-20 border border-gray-300 rounded-md py-1 hover:border-gray-400 hover:bg-gray-100">Previous</button>
        </div>
        <div>
            <button onclick={onClickNext} class="text-sm cursor-pointer w-20 border border-gray-300 rounded-md py-1 hover:border-gray-400 hover:bg-gray-100">Next</button>
        </div>
    </div>
</div>