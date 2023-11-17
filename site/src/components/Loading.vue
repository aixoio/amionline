<template>
    <div>
        <div class="flex gap-5 bg-gray-50/30 dark:bg-gray-950/30 w-fit p-2 m-2 rounded-lg shadow-lg">
            <ArrowPathIcon class="fill-gray-950 dark:fill-gray-100 w-16 h-16 spin"></ArrowPathIcon>
            <span class="text flex items-center font-semibold">Loading...</span>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ArrowPathIcon } from "@heroicons/vue/20/solid"

</script>

<style scoped lang="scss">

.spin {
    animation: spin ease-in-out infinite 1s;
}

@keyframes spin {
    0% {
        transform: rotate(0deg);
    }

    50% {
        transform: rotate(180deg);
    }

    100% {
        transform: rotate(360deg);
    }
}

</style>
