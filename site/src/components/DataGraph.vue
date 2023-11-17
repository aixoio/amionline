<template>
    <div>
        <canvas class="$graph" ref="$graph"></canvas>
    </div>
</template>


<script setup lang="ts">
import { onMounted, ref } from "vue";
import { type Event } from "../assets/ts/api"
import { Chart, type ChartItem } from "chart.js/auto"
import { useDataStore } from "../stores/datastore"
import { storeToRefs } from "pinia";
import zoomPlugin from "chartjs-plugin-zoom"

Chart.register(zoomPlugin)

const $graph = ref((null as unknown) as HTMLCanvasElement)

function parse_date(events: Event[]): string[] {
    const e = events.sort((a, b) => {
        if (a.time_of_request > b.time_of_request) {
            return -1;
        }
        if (a.time_of_request < b.time_of_request) {
            return 1;
        }
        return 0;
    })

    const out: string[] = []

    for (let i = 0; i < e.length; i++) {
        out.push(new Date(e[i].time_of_request).toLocaleString())
    }

    return out;

}

function parse_time_ms(events: Event[]): number[] {
    const e = events.sort((a, b) => {
        if (a.time_of_request > b.time_of_request) {
            return -1;
        }
        if (a.time_of_request < b.time_of_request) {
            return 1;
        }
        return 0;
    })

    const out: number[] = []

    for (let i = 0; i < e.length; i++) {
        out.push(e[i].time_ms)
    }

    return out;

}

function parse_color(events: Event[]): string[] {
    const e = events.sort((a, b) => {
        if (a.time_of_request > b.time_of_request) {
            return -1;
        }
        if (a.time_of_request < b.time_of_request) {
            return 1;
        }
        return 0;
    })
    const out: string[] = []

    for (let i = 0; i < e.length; i++) {
        if (!e[i].success) {
            out.push("#6b0000") // dark red
        }
        else if (e[i].time_ms <= 75) {
            out.push("#36ff17") // green
        } else if (e[i].time_ms > 175) {
            out.push("#ff1717") // red
        } else {
            out.push("#ffae17") // orange
        }
    }

    return out

}

function longest_request_time(events: Event[]): number {
    let bigest = 0;
    for (let i = 0; i < events.length; i++) {
        bigest = Math.max(bigest, events[i].time_ms)
    }
    return bigest
}

const datastore = useDataStore()
const { dataset } = storeToRefs(datastore)

onMounted(() => {

    new Chart($graph.value! as ChartItem, {
        type: 'bar',
        options: {
            plugins: {
                legend: {
                    display: false,
                },
                tooltip: {
                    callbacks: {
                        label: (i) => {
                            const item = (dataset.value.events as Event[])[i.dataIndex]
                            
                            const time = new Date(item.time_of_request).toLocaleString()
                            const success = item.success ? "online" : "offline"
                            const time_ms = item.time_ms
                            const target_ip = item.target_ip

                            return `${time_ms}/ms to ${target_ip} at ${time} and we are ${success}`

                        }
                    },
                },
                zoom: {
                    zoom: {
                        wheel: {
                            enabled: true,
                        },
                        pinch: {
                            enabled: true,
                        },
                        mode: "xy",
                        drag: {
                            enabled: true,
                        },
                    },
                    limits: {
                        y: {min: 0, max: longest_request_time(dataset.value.events as Event[]) + 10}
                    }
                },
            },
            responsive: true,
        },
        data: {
            labels: parse_date(dataset.value.events as Event[]),
            datasets: [
                {
                    label: "",
                    backgroundColor: parse_color(dataset.value.events as Event[]),
                    data: parse_time_ms(dataset.value.events as Event[]),
                }
            ]
        }
    })

})


</script>
