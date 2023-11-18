<template>
    <div>
        <div class="flex items-center">
            <ArrowDownOnSquareIcon class="w-8 h-8 text-gray-700 dark:text-gray-200 cursor-pointer" @click="show_popup">
            </ArrowDownOnSquareIcon>
        </div>
        <TransitionRoot appear :show="can_show" as="template">
            <Dialog as="div" @close="close_popup" class="relative z-10">
                <TransitionChild as="template" enter="duration-300 ease-out" enter-from="opacity-0" enter-to="opacity-100"
                    leave="duration-200 ease-in" leave-from="opacity-100" leave-to="opacity-0">
                    <div class="fixed inset-0 bg-black/25" />
                </TransitionChild>

                <div class="fixed inset-0 overflow-y-auto">
                    <div class="flex min-h-full items-center justify-center p-4 text-center">
                        <TransitionChild as="template" enter="duration-300 ease-out" enter-from="opacity-0 scale-95"
                            enter-to="opacity-100 scale-100" leave="duration-200 ease-in" leave-from="opacity-100 scale-100"
                            leave-to="opacity-0 scale-95">
                            <DialogPanel
                                class="w-full max-w-md transform rounded-2xl bg-white/20 backdrop-blur-lg dark:bg-black/20 p-6 text-left align-middle shadow-xl transition-all border border-black/70 dark:border-white/70 z-20 overflow-visible">
                                <DialogTitle as="h3" class="text-lg font-medium leading-6 text">
                                    Download
                                </DialogTitle>
                                <div class="mt-2">
                                    <p class="text-sm text">
                                        This will download {{ get_event_name }}
                                    </p>
                                    <span class="text text-sm">The data will be download as:</span>

                                    <Menu as="div" class="relative inline-block text-left m-2">
                                        <MenuButton
                                            class="inline-flex w-full justify-center rounded-md bg-black/20 px-4 py-2 text-sm font-medium text-white hover:bg-black/30 focus:outline-none ring-2 ring-white/75">
                                            {{ download_as }}
                                            <ChevronDownIcon
                                                class="mr-1 ml-2 h-5 w-5 dark:text-indigo-200 text-black hover:text-violet-100"
                                                aria-hidden="true" />
                                        </MenuButton>
                                        <transition enter-active-class="transition duration-100 ease-out"
                                            enter-from-class="transform scale-95 opacity-0"
                                            enter-to-class="transform scale-100 opacity-100"
                                            leave-active-class="transition duration-75 ease-in"
                                            leave-from-class="transform scale-100 opacity-100"
                                            leave-to-class="transform scale-95 opacity-0">
                                            <MenuItems
                                                class="z-50 backdrop-blur-lg absolute right-0 mt-2 w-fit p-2 origin-top-right rounded-md bg-white/20 dark:bg-black/20 shadow-lg ring-1 ring-black/5 dark:ring-white/5 focus:outline-none">
                                                <menu-item v-slot="{ active }">
                                                    <button :class="[
                                                        active ? 'bg-indigo-500 text-white' : 'text-gray-900 dark:text-gray-100',
                                                        'group flex w-full items-center rounded-md px-2 py-2 text-sm',
                                                        download_as == 'JSON' ? 'bg-violet-500 text-white' : ''
                                                    ]" class="cursor-pointer mt-3" @click="download_as = 'JSON'">JSON</button>
                                                </menu-item>
                                                <menu-item v-slot="{ active }">
                                                    <button :class="[
                                                        active ? 'bg-indigo-500 text-white' : 'text-gray-900 dark:text-gray-100',
                                                        'group flex w-full items-center rounded-md px-2 py-2 text-sm',
                                                        download_as == 'CSV' ? 'bg-violet-500 text-white' : ''
                                                    ]" class="cursor-pointer mt-3" @click="download_as = 'CSV'">CSV</button>
                                                </menu-item>
                                            </MenuItems>
                                        </transition>
                                    </Menu>


                                </div>

                                <div class="mt-4 flex justify-start gap-4">
                                    <button type="button"
                                        class="inline-flex justify-center rounded-md border border-transparent bg-emerald-300 px-4 py-2 text-sm font-medium text-emerald-900 hover:bg-emerald-500 focus:outline-none focus-visible:ring-2 focus-visible:ring-emerald-500 focus-visible:ring-offset-2"
                                        @click="download">
                                        Download
                                    </button>
                                    <button type="button"
                                        class="inline-flex justify-center rounded-md border border-transparent bg-red-300 px-4 py-2 text-sm font-medium text-red-900 hover:bg-red-500 focus:outline-none focus-visible:ring-2 focus-visible:ring-red-500 focus-visible:ring-offset-2"
                                        @click="close_popup">
                                        Cancel
                                    </button>
                                </div>
                            </DialogPanel>
                        </TransitionChild>
                    </div>
                </div>
            </Dialog>
        </TransitionRoot>
    </div>
</template>

<script setup lang="ts">
import { ArrowDownOnSquareIcon, ChevronDownIcon } from '@heroicons/vue/24/outline';
import { computed, ref } from 'vue';
import { TransitionRoot, TransitionChild, Dialog, DialogPanel, DialogTitle, Menu, MenuButton, MenuItems, MenuItem } from '@headlessui/vue'
import { useDataStore } from '../stores/datastore';

const datastore = useDataStore()

const can_show = ref(false)
const download_as = ref("JSON")

function show_popup() {
    can_show.value = true
}

function close_popup() {
    can_show.value = false
}

function download() {
    close_popup()
}

const get_event_name = computed(() => {
    switch (datastore.store_type) {
        case "20":
            return "the last 20 events"
            break;
        case "un20":
            return "the last 20 events (Uncached)"
            break;
        case "all":
            return "all events"
            break;
        case "unall":
            return "all events (Uncached)"
            break;
    }
})


</script>
