import { useDataStore } from "@/stores/datastore"
import { get_all_events, get_all_uncached_events, get_last_20_events, get_last_20_uncached_events } from "./api"

const datastore = useDataStore()

export async function load_20() {
    datastore.loaded = false
    const data = await get_last_20_events()
    datastore.loaded = true
    datastore.store_type = "20"
    datastore.dataset = data
}

export async function load_all() {
    datastore.loaded = false
    const data = await get_all_events()
    datastore.loaded = true
    datastore.store_type = "all"
    datastore.dataset = data
}

export async function load_uncached_20() {
    datastore.loaded = false
    const data = await get_last_20_uncached_events()
    datastore.loaded = true
    datastore.store_type = "un20"
    datastore.dataset = data
}

export async function load_uncached_all() {
    datastore.loaded = false
    const data = await get_all_uncached_events()
    datastore.loaded = true
    datastore.store_type = "unall"
    datastore.dataset = data
}
