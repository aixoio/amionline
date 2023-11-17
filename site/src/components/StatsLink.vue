<template>
    <div>
        <div class="flex items-center">
            <ChartBarSquareIcon class="w-8 h-8 text-gray-700 dark:text-gray-200 cursor-pointer" @click="go_to_stats_page"></ChartBarSquareIcon>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ChartBarSquareIcon } from "@heroicons/vue/24/outline"
import { useRouter } from "vue-router";

const router = useRouter()


function go_to_stats_page() {
    router.push({
        name: "stats"
    })
}

</script>
