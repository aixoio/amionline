<template>
    <div>
        <div class="flex items-center" ref="$div" @mouseenter="onHover" @mouseleave="onHoverLeave">
            <ChartBarSquareIcon class="w-8 h-8 text-gray-700 dark:text-gray-200 cursor-pointer" @click="go_to_stats_page"></ChartBarSquareIcon>
        </div>
        <div ref="$tooltip" class="bg-white/20 dark:bg-black/20 shadow-lg backdrop-blur-lg p-1.5 rounded-md border border-white/70" :class="{ show: tooltipOpen, hide: !tooltipOpen }" :style="floatingStyles">
            <span class="text">Statistics</span>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ChartBarSquareIcon } from "@heroicons/vue/24/outline"
import { useRouter } from "vue-router";
import { useFloating } from '@floating-ui/vue';
import { flip, offset, shift } from "@floating-ui/dom";
import { onUpdated, ref } from "vue";

const router = useRouter()


function go_to_stats_page() {
    router.push({
        name: "stats"
    })
}

const $div = ref((null as unknown) as HTMLElement)
const $tooltip = ref((null as unknown) as HTMLElement)

const tooltipOpen = ref(false)

const { floatingStyles, update } = useFloating($div, $tooltip, {
    placement: "left",
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

</script>

<style scoped lang="scss">
.show {
    display: block;
}

.hide {
    display: none;
}
</style>
