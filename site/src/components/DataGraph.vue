<template>
    <div>
        <canvas class="$graph" ref="$graph"></canvas>
    </div>
</template>


<script setup lang="ts">
import { onMounted, ref } from "vue";
import { get_last_20_events, type Event } from "../assets/ts/api"
import { Chart, type ChartItem } from "chart.js/auto"

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
        if (e[i].time_ms <= 75 && e[i].success) {
            out.push("#36ff17") // green
        } else if (e[i].time_ms > 100 || !e[i].success) {
            out.push("#ff1717") // red
        } else {
            out.push("#ffae17") // orange
        }
    }

    return out

}

function parse_legends(events: Event[]): string[] {
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
        out.push(`Time/ms`)
    }

    return out

}

onMounted(async () => {

    const dataset = await get_last_20_events()

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
                            const item = (dataset.events as Event[])[i.dataIndex]
                            
                            const time = new Date(item.time_of_request).toLocaleString()
                            const success = item.success ? "online" : "offline"
                            const time_ms = item.time_ms
                            const target_ip = item.target_ip

                            return `${time_ms}/ms to ${target_ip} at ${time} and we are ${success}`

                        }
                    },
                },
            },
            responsive: true,
        },
        data: {
            labels: parse_date(dataset.events as Event[]),
            datasets: [
                {
                    label: "",
                    backgroundColor: parse_color(dataset.events as Event[]),
                    data: parse_time_ms(dataset.events as Event[]),
                }
            ]
        }
    })

})


</script>
