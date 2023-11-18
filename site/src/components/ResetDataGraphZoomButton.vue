<template>
    <div>
        <div class="flex items-center">
            <MagnifyingGlassMinusIcon class="w-8 h-8 text-gray-700 dark:text-gray-200 cursor-pointer" @click="reset_data_graph_zoom_btn"></MagnifyingGlassMinusIcon>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useDataGraphStore } from "@/stores/datagraphstore";
import { MagnifyingGlassMinusIcon } from "@heroicons/vue/24/outline"

const datagraphstore = useDataGraphStore()

function reset_data_graph_zoom_btn() {
    datagraphstore.tiggerZoomReset()
}

</script>
