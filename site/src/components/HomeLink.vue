<template>
    <div>
        <div class="flex items-center">
            <HomeIcon class="w-8 h-8 text-gray-700 dark:text-gray-200 cursor-pointer" @click="go_to_home_page"></HomeIcon>
        </div>
    </div>
</template>

<script setup lang="ts">
import { HomeIcon } from "@heroicons/vue/24/outline"
import { useRouter } from "vue-router";

const router = useRouter()


function go_to_home_page() {
    router.push({
        name: "home"
    })
}

</script>
