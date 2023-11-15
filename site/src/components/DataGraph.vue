<template>
    <div>
        <div class="$graph" ref="$graph"></div>
    </div>
</template>


<script setup lang="ts">
import * as d3 from "d3"
import { onMounted, ref } from "vue";
import is_dark from "../assets/ts/is_dark"
import { get_all_events, get_last_20_events, type Event } from "../assets/ts/api"

const $graph = ref((null as unknown) as HTMLElement)

onMounted(async () => {

    const dataset = await get_last_20_events()


    const width = 640;
    const height = 400;
    const marginTop = 20;
    const marginRight = 20;
    const marginBottom = 30;
    const marginLeft = 40;
    const padding = 5

    const svg = d3.create("svg")
        .attr("width", width)
        .attr("height", height);

    let best_color = "#222"

    if (is_dark()) {
        best_color = "#eee"
    }

    const newest_date = new Date((dataset.events as Event[]).reduce((maxDate, event) => {
        return event.time_of_request > maxDate ? event.time_of_request : maxDate;
    }, (dataset.events as Event[])[0].time_of_request));

    const oldest_date = new Date((dataset.events as Event[]).reduce((minDate, event) => {
        return event.time_of_request < minDate ? event.time_of_request : minDate;
    }, (dataset.events as Event[])[0].time_of_request));

    const x = d3.scaleUtc()
        .domain([newest_date, oldest_date])
        .range([0, width - 1])

    const sortedTimeValues = (dataset.events as Event[]).map(event => event.time_ms).sort((a, b) => a - b)

    const y = d3.scaleLinear()
        .domain([sortedTimeValues[sortedTimeValues.length - 1], 0])
        .range([0, height - marginBottom])

    svg.selectAll("rect")
        .data(dataset.events as Event[])
        .enter()
        .append("rect")
        .attr("x", (item: Event, i: number) => i * (width / (dataset.events as Event[]).length <= 0 ? 5 : width / (dataset.events as Event[]).length) + marginLeft + 1)
        .attr("y", (item) => height - y(item.time_ms) - marginBottom - 1)
        .attr("height", (item) => y(item.time_ms))
        .attr("width", (width / (dataset.events as Event[]).length <= 0 ? 5 : width / (dataset.events as Event[]).length) - padding)
        .style("fill", (i: Event) => {
            if (i.time_ms <= 100 && i.success) {
                return "#0f0"
            } else if (i.time_ms >= 1000 || !i.success) {
                return "#f00"
            }
            return "#fa5"
        });

    svg.append("g")
        .attr("transform", `translate(${marginLeft},${height - marginBottom})`)
        .call(d3.axisBottom(x))
        .style("color", best_color)

    svg.append("g")
        .attr("transform", `translate(${marginLeft},0)`)
        .style("color", best_color)
        .call(d3.axisLeft(y));

    $graph.value.append(svg.node()!)
})


</script>
