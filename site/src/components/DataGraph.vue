<template>
    <div>
        <div class="$graph" ref="$graph"></div>
    </div>
</template>


<script setup lang="ts">
import { onMounted, ref } from "vue";
import is_dark from "../assets/ts/is_dark"
import { get_last_20_events, type Event } from "../assets/ts/api"

const $graph = ref((null as unknown) as HTMLElement)

onMounted(async () => {

    

})


</script>
