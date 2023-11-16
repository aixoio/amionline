<template>
  <div>
    <nav class="flex justify-start mt-5 border-b border-b-black/20 p-4">
      <h1 class="title">Am I Online?</h1>
      <div class="flex gap-4">
        <OptionsMenu></OptionsMenu>
        <ReloadButton></ReloadButton>
        <span v-show="dataset != null && !dataset.cached"
          class="bg-red-500 m-1.5 p-1 w-fit uppercase rounded-lg text-sm inline-block font-bold text-white">Uncached</span>
      </div>
    </nav>
    <DataGraph v-if="canshowgraph" class="m-8 center"></DataGraph>
  </div>
</template>

<script setup lang="ts">
import DataGraph from '@/components/DataGraph.vue';
import OptionsMenu from '@/components/OptionsMenu.vue';
import ReloadButton from '@/components/ReloadButton.vue';
import { storeToRefs } from 'pinia';
import { onMounted, ref, watch } from 'vue';
import { get_last_20_events } from '../assets/ts/api';
import { useDataStore } from '../stores/datastore';

const datastore = useDataStore()
const { loaded, dataset } = storeToRefs(datastore)

onMounted(async () => {
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
