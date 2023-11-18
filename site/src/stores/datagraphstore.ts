import { defineStore } from "pinia";
import { ref } from "vue";

export const useDataGraphStore = defineStore("data-graph-store", () => {

    const tiggerZoomResetValue = ref(true)

    const tiggerZoomReset = () => {
        tiggerZoomResetValue.value = !tiggerZoomResetValue.value
    }

    return {
        tiggerZoomResetValue,
        tiggerZoomReset,
    }

})
