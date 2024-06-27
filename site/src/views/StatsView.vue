<template>
    <div>
        <nav class="md:flex w-full min-w-full border-b border-b-black/20 pb-4 justify-between">
            <nav class="md:flex block justify-start mt-5 pt-4 pl-4 pr-4">
                <h1 class="title">Am I Online? - Statistics</h1>
                <div class="flex gap-4">
                    <OptionsMenu></OptionsMenu>
                    <ReloadButton></ReloadButton>
                    <span class="text flex items-center">{{ get_event_name }}</span>
                    <span v-show="dataset != null && !dataset.cached"
                        class="bg-red-500 flex items-center p-1 w-fit uppercase rounded-lg text-sm ml-2 mr-2 font-bold text-white">Uncached</span>
                </div>
            </nav>
            <nav class="flex justify-end m-3 gap-5">
                <DownloadButton class="flex items-center"></DownloadButton>
                <HomeLink class="flex items-center"></HomeLink>
            </nav>
        </nav>
        <div class="grid lg:grid-cols-2 grid-cols-1 lg:w-full w-[75%] h-full">
            <div class="grid md:grid-cols-2 grid-cols-1 w-full h-full">
                <SuccessSplitGraphVue v-if="canshowresults" class="m-8 center w-full h-full"></SuccessSplitGraphVue>
                <TimeingSplitGraph v-if="canshowresults" class="m-8 center w-full h-full"></TimeingSplitGraph>
            </div>
            <TimeMSData v-if="canshowresults" class="m-8"></TimeMSData>
        </div>
        <Loading v-if="!loaded" class="flex justify-center items-center h-full mt-56"></Loading>
    </div>
</template>

<script setup lang="ts">
import OptionsMenu from '@/components/OptionsMenu.vue';
import HomeLink from "@/components/HomeLink.vue"
import ReloadButton from '@/components/ReloadButton.vue';
import Loading from '@/components/Loading.vue';
import SuccessSplitGraphVue from '@/components/SuccessSplitGraph.vue';
import TimeingSplitGraph from '@/components/TimeingSplitGraph.vue';
import TimeMsData from '@/components/TimeMsData.vue';
import DownloadButton from '@/components/DownloadButton.vue';
import { storeToRefs } from "pinia"
import { useDataStore } from "../stores/datastore"
import { computed, onMounted, ref, watch } from 'vue';
import { get_last_20_events } from '../assets/ts/api';

const datastore = useDataStore()
const { loaded, dataset } = storeToRefs(datastore)

let canshowresults = ref(false)

onMounted(async () => {
    if (dataset.value != null) {
        canshowresults.value = true
        return
    }
    const data = await get_last_20_events()
    datastore.dataset = data
    datastore.loaded = true
    datastore.store_type = "20"
    canshowresults.value = true
})

const get_event_name = computed(() => {
    switch (datastore.store_type) {
        case "20":
            return "Last 20 events"
            break;
        case "un20":
            return "Last 20 events"
            break;
        case "all":
            return "All events"
            break;
        case "unall":
            return "All events"
            break;
    }
})

watch(loaded, (val) => {
  canshowresults.value = val
})

</script>
