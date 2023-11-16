<template>
    <div class="fixed top-16 w-56 text-right">
        <Menu as="div" class="relative inline-block text-left">
            <MenuButton
                class="inline-flex w-full justify-center rounded-md bg-black/20 px-4 py-2 text-sm font-medium text-white hover:bg-black/30 focus:outline-none focus-visible:ring-2 focus-visible:ring-white/75">
                Options
                <ChevronDownIcon class="mr-1 ml-2 h-5 w-5 dark:text-indigo-200 text-black hover:text-violet-100" aria-hidden="true" />
            </MenuButton>
            <transition enter-active-class="transition duration-100 ease-out"
                enter-from-class="transform scale-95 opacity-0" enter-to-class="transform scale-100 opacity-100"
                leave-active-class="transition duration-75 ease-in" leave-from-class="transform scale-100 opacity-100"
                leave-to-class="transform scale-95 opacity-0">
                <MenuItems
                    class="absolute right-0 mt-2 w-fit p-2 origin-top-right divide-y divide-gray-100 rounded-md bg-white/20 dark:bg-black/20 shadow-lg ring-1 ring-black/5 dark:ring-white/5 focus:outline-none">
                    <div class="px-1 py-1">
                        <menu-item v-slot="{ active }">
                            <button :class="[
                                active ? 'bg-indigo-500 text-white' : 'text-gray-900 dark:text-gray-100',
                                'group flex w-full items-center rounded-md px-2 py-2 text-sm',
                                store_type == '20' ? 'bg-violet-500 text-white' : ''
                            ]" class="cursor-pointer mt-3" @click="load_20">
                                Last 20 events
                            </button>
                        </menu-item>
                        <menu-item v-slot="{ active }">
                            <button :class="[
                                active ? 'bg-indigo-500 text-white' : 'text-gray-900 dark:text-gray-100',
                                'group flex w-full items-center rounded-md px-2 py-2 text-sm',
                                store_type == 'all' ? 'bg-violet-500 text-white' : ''
                            ]" class="cursor-pointer mt-3" @click="load_all">
                                All events
                            </button>
                        </menu-item>
                    </div>
                    <div class="px-1 py-1">
                        <menu-item v-slot="{ active }">
                            <button :class="[
                                active ? 'bg-indigo-500 text-white' : 'text-gray-900 dark:text-gray-100',
                                'group flex w-full items-center rounded-md px-2 py-2 text-sm',
                                store_type == 'un20' ? 'bg-violet-500 text-white' : ''
                            ]" class="cursor-pointer mt-3" @click="load_uncached_20">
                                <div>
                                    <span
                                        class="bg-red-500 m-1.5 p-1 w-fit uppercase rounded-lg text-sm inline-block font-bold text-white">Uncached</span>
                                    Last 20 events
                                </div>
                            </button>
                        </menu-item>
                        <menu-item v-slot="{ active }">
                            <button :class="[
                                active ? 'bg-indigo-500 text-white' : 'text-gray-900 dark:text-gray-100',
                                'group flex w-full items-center rounded-md px-2 py-2 text-sm',
                                store_type == 'unall' ? 'bg-violet-500 text-white' : ''
                            ]" class="cursor-pointer mt-3" @click="load_uncached_all">
                                <div>
                                    <span
                                        class="bg-red-500 m-1.5 p-1 w-fit uppercase rounded-lg text-sm inline-block font-bold text-white">Uncached</span>
                                    All events
                                </div>
                            </button>
                        </menu-item>
                    </div>
                </MenuItems>
            </transition>
        </Menu>
    </div>
</template>

<script setup lang="ts">
import { Menu, MenuButton, MenuItems, MenuItem } from "@headlessui/vue"
import { ChevronDownIcon } from "@heroicons/vue/20/solid"
import { storeToRefs } from "pinia"
import { load_20, load_all, load_uncached_20, load_uncached_all } from "@/assets/ts/loader"
import { useDataStore } from "../stores/datastore"

const datastore = useDataStore()
const { store_type } = storeToRefs(datastore)


</script>
