<template>
  <div>
    <nav class="md:flex w-full min-w-full border-b border-b-black/20 pb-4 justify-between">
      <nav class="md:flex block justify-start mt-5 pt-4 pl-4 pr-4">
        <h1 class="title">Am I Online?</h1>
        <div class="flex gap-4">
          <OptionsMenu></OptionsMenu>
          <ReloadButton></ReloadButton>
          <span v-show="dataset != null && !dataset.cached"
            class="bg-red-500 flex items-center p-1 w-fit uppercase rounded-lg text-sm ml-2 mr-2 font-bold text-white">Uncached</span>
        </div>
      </nav>
      <nav class="flex justify-end m-3 gap-5">
        <DownloadButton class="flex items-center"></DownloadButton>
        <StatsLink class="flex items-center"></StatsLink>
      </nav>
    </nav>
    <DataGraph v-if="canshowgraph" class="m-8 center"></DataGraph>
    <Loading v-if="!loaded" class="flex justify-center items-center h-full mt-56"></Loading>
  </div>
</template>

<script setup lang="ts">
import DataGraph from '@/components/DataGraph.vue';
import OptionsMenu from '@/components/OptionsMenu.vue';
import ReloadButton from '@/components/ReloadButton.vue';
import Loading from '@/components/Loading.vue';
import StatsLink from '@/components/StatsLink.vue';
import DownloadButton from '@/components/DownloadButton.vue';
import { storeToRefs } from 'pinia';
import { onMounted, ref, watch } from 'vue';
import { get_last_20_events } from '../assets/ts/api';
import { useDataStore } from '../stores/datastore';

const datastore = useDataStore()
const { loaded, dataset } = storeToRefs(datastore)

onMounted(async () => {
  if (dataset.value != null) {
    canshowgraph.value = true
    return
  }
  const data = await get_last_20_events()
  datastore.dataset = data
  datastore.loaded = true
  datastore.store_type = "20"
  canshowgraph.value = true
})

let canshowgraph = ref(false)

watch(loaded, (val) => {
  canshowgraph.value = val
})

</script>
