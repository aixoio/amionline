<template>
  <div>
    <h1 class="center p-2 m-4 title">Am I Online?</h1>
    <nav>
      <OptionsMenu></OptionsMenu>
    </nav>
    <DataGraph v-if="canshowgraph" class="m-8 center"></DataGraph>
  </div>
</template>

<script setup lang="ts">
import DataGraph from '@/components/DataGraph.vue';
import OptionsMenu from '@/components/OptionsMenu.vue';
import { storeToRefs } from 'pinia';
import { onMounted, ref, watch } from 'vue';
import { get_last_20_events } from '../assets/ts/api';
import { useDataStore } from '../stores/datastore';

const datastore = useDataStore()
const { loaded } = storeToRefs(datastore)

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
