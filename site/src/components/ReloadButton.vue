<template>
    <div>
        <ArrowPathIcon class="w-8 h-8 text-gray-700 dark:text-gray-200 cursor-pointer" @click="reload"></ArrowPathIcon>
    </div>
</template>

<script setup lang="ts">
import { ArrowPathIcon } from "@heroicons/vue/20/solid"
import { useDataStore } from "../stores/datastore";
import { load_20, load_all, load_uncached_20, load_uncached_all } from "@/assets/ts/loader";

const datastore = useDataStore()

function reload() {
    switch (datastore.store_type) {
        case "20":
            load_20()
            break;
        case "all":
            load_all()
            break;
        case "un20":
            load_uncached_20()
            break;
        case "unall":
            load_uncached_all()
            break;
    }
}

</script>
