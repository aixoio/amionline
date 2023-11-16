import { type Events_Responce } from "@/assets/ts/api"
import { defineStore } from "pinia"
import { ref } from "vue"

export const useDataStore = defineStore("data-store", () => {
    const dataset = ref((null as unknown) as Events_Responce)

    const loaded = ref(false)

    const store_type = ref("unloaded")

    return {
        dataset,
        loaded,
        store_type,
    }
})
