<template>
    <div>
        <div ref="$div" @mouseenter="onHover" @mouseleave="onHoverLeave">
            <ArrowPathIcon class="w-8 h-8 text-gray-700 dark:text-gray-200 cursor-pointer" @click="reload"></ArrowPathIcon>
        </div>
        <div ref="$tooltip" class="bg-white/20 dark:bg-black/20 shadow-lg backdrop-blur-lg p-1.5 rounded-md border border-white/70" :class="{ show: tooltipOpen, hide: !tooltipOpen }" :style="floatingStyles">
            <span class="text">Refresh</span>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ArrowPathIcon } from "@heroicons/vue/20/solid"
import { useDataStore } from "../stores/datastore";
import { load_20, load_all, load_uncached_20, load_uncached_all } from "@/assets/ts/loader";
import { onUpdated, ref } from "vue";
import { useFloating } from '@floating-ui/vue';
import { flip, offset, shift } from "@floating-ui/dom";

const datastore = useDataStore()

const $div = ref((null as unknown) as HTMLElement)
const $tooltip = ref((null as unknown) as HTMLElement)

const tooltipOpen = ref(false)

const { floatingStyles, update } = useFloating($div, $tooltip, {
    placement: "bottom",
    open: tooltipOpen,
    middleware: [
        shift(),
        offset(10),
        flip(),
    ]
})

onUpdated(() => {
    update();
})


async function onHover() {
    tooltipOpen.value = true;
}

async function onHoverLeave() {
    tooltipOpen.value = false;
}

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

<style scoped lang="scss">
.show {
    display: block;
}

.hide {
    display: none;
}
</style>
