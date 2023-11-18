<template>
    <div>
        <canvas class="$graph" ref="$graph"></canvas>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { Chart, type ChartItem } from "chart.js/auto"
import { useDataStore } from '../stores/datastore';
import { storeToRefs } from 'pinia';
import { type Event } from '@/assets/ts/api';
import { round } from "mathjs"

const $graph = ref((null as unknown) as HTMLCanvasElement)

const datastore = useDataStore()
const { dataset } = storeToRefs(datastore)

function parse_success_data(events: Event[]): number {
    let count = 0
    for (let i = 0; i < events.length; i++) {
        if (events[i].success) {
            count++
        }
    }
    return count
}

function parse_fail_data(events: Event[]): number {
    let count = 0
    for (let i = 0; i < events.length; i++) {
        if (!events[i].success) {
            count++
        }
    }
    return count
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
                            if (i.dataIndex == 0) {
                                const success_percent = round((parse_success_data(dataset.value.events as Event[]) / (dataset.value.events as Event[]).length) * 100, 4)
                                return `${success_percent}% - ${parse_success_data(dataset.value.events as Event[])}`
                            }
                            const fail_percent = round((parse_fail_data(dataset.value.events as Event[]) / (dataset.value.events as Event[]).length) * 100, 4)
                            return `${fail_percent}% - ${parse_fail_data(dataset.value.events as Event[])}`

                        }
                    },
                },
            },
        },
        data: {
            labels: ["Online", "Offline"],
            datasets: [
                {
                    backgroundColor: ["#5bff2e", "#ff2e2e"],
                    data: [parse_success_data(dataset.value.events as Event[]), parse_fail_data(dataset.value.events as Event[])],
                },
            ]
        }
    })
})

</script>
