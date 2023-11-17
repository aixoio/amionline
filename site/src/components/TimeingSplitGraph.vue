<template>
    <div>
        <canvas class="$graph" ref="$graph"></canvas>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { Chart, type ChartItem } from "chart.js/auto"
import { round } from "mathjs"
import { type Event } from '@/assets/ts/api';
import { storeToRefs } from "pinia"
import { useDataStore } from "../stores/datastore"

const $graph = ref((null as unknown) as HTMLCanvasElement)

const datastore = useDataStore()
const { dataset } = storeToRefs(datastore)

function parse_data(e: Event[]): number[] {
    let fast = 0
    let ok = 0
    let slow = 0
    let unsuccessful = 0
    for (let i = 0; i < e.length; i++) {
        if (!e[i].success) {
            unsuccessful++
        }
        else if (e[i].time_ms <= 75) {
            fast++
        } else if (e[i].time_ms > 175) {
            slow++
        } else {
            ok++
        }
    }
    return [fast, ok, slow, unsuccessful]

}

onMounted(() => {
    new Chart($graph.value! as ChartItem, {
        type: "pie",
        options: {
            responsive: true,
            plugins: {
                tooltip: {
                    callbacks: {
                        label: (i) => {
                            const data = parse_data(dataset.value.events!)
                            const total = dataset.value.events!.length
                            switch (i.dataIndex) {
                                case 0:
                                    const fast_percent = round((data[0] / total) * 100, 4)
                                    return `${fast_percent}% - ${data[0]}`
                                    break;
                                case 1:
                                    const ok_percent = round((data[1] / total) * 100, 4)
                                    return `${ok_percent}% - ${data[1]}`
                                    break;
                                case 2:
                                    const slow_percent = round((data[2] / total) * 100, 4)
                                    return `${slow_percent}% - ${data[2]}`
                                    break;
                                case 3:
                                    const unsuccessful_percent = round((data[3] / total) * 100, 4)
                                    return `${unsuccessful_percent}% - ${data[3]}`
                                    break;
                            }
                        }
                    },
                },
            },
        },
        data: {
            labels: ["Fast", "Ok", "Slow", "Unsuccessful"],
            datasets: [
                {
                    backgroundColor: ["#36ff17", "#ffae17", "#ff1717", "#6b0000"],
                    data: parse_data(dataset.value.events!)
                },
            ]
        },
    })
})

</script>
