import { defineStore } from "pinia";
import { ref } from "vue";

export const useDownloadButtonStore = defineStore("download-button-store", () => {

    const download_as = ref("JSON")

    return {
        download_as
    }

})
