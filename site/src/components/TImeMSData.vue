<template>
    <div>
        <h2 class="text-xl text">Time/ms data</h2>
        <div class="grid grid-cols-1 gap-0.5">
            <span class="text block">Mean: <span class="font-bold">{{ get_mean }}</span></span>
            <span class="text block">Max: <span class="font-bold">{{ get_max }}</span></span>
            <span class="text block">Min: <span class="font-bold">{{ get_min }}</span></span>
            <span class="text block">Median: <span class="font-bold">{{ get_median }}</span></span>
            <span class="text block">Mode: <span class="font-bold">{{ get_formated_mean }}</span></span>
            <span class="text block">Range: <span class="font-bold">{{ get_range }}</span></span>
        </div>
    </div>
</template>

<script setup lang="ts">
import { max, mean, median, min, mode, number, round } from "mathjs";
import { computed } from "vue";
import { useDataStore } from "../stores/datastore";

const datastore = useDataStore()

const get_mean = computed(() => {
    return round(mean(...datastore.dataset.events!.map((x) => x.time_ms)), 4)
})

const get_max = computed(() => {
    return round(max(...datastore.dataset.events!.map((x) => x.time_ms)), 4)
})

const get_min = computed(() => {
    return round(min(...datastore.dataset.events!.map((x) => x.time_ms)), 4)
})

const get_median = computed(() => {
    return round(median(...datastore.dataset.events!.map((x) => x.time_ms)), 4)
})

const get_mode = computed(() => {
    return ((mode(...datastore.dataset.events!.map((x) => x.time_ms)) as unknown) as number[]).map((x) => round(x, 4))
})

const get_range = computed(() => {
    return round(get_max.value - get_min.value, 4)
})

const get_formated_mean = computed(() => {
    let str = ""
    const data = get_mode.value!
    for (let i = 0; i < data.length; i++) {
        if (i == data.length - 1) {
            str += `${data[i]}`
        } else {
            str += `${data[i]}, `
        }
    }
    return str
})

</script>
