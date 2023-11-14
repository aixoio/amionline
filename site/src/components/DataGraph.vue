<template>
    <div>
        <div ref="$graph"></div>
    </div>
</template>


<script setup lang="ts">
import * as d3 from "d3"
import { onMounted, ref } from "vue";
import is_dark from "../assets/ts/is_dark"

const $graph = ref((null as unknown) as HTMLElement)

const width = 640;
const height = 400;
const marginTop = 20;
const marginRight = 20;
const marginBottom = 30;
const marginLeft = 40;

const svg = d3.create("svg")
    .attr("width", width)
    .attr("height", height);


const x = d3.scaleUtc().domain([new Date("2023-01-01"), new Date("2024-01-01")]).range([marginLeft, width - marginRight]) // ! PLACE DATE RANGE HERE

const y = d3.scaleLinear().domain([0, 100]).range([height - marginBottom, marginTop]); // ! PLACE HIGH PING HERE

let line_color = "#222"

if (is_dark()) {
    line_color = "#eee"
}

svg.append("g")
    .attr("transform", `translate(0,${height - marginBottom})`)
    .attr("color", line_color)
    .call(d3.axisBottom(x));

svg.append("g")
    .attr("transform", `translate(${marginLeft},0)`)
    .attr("color", line_color)
    .call(d3.axisLeft(y));

onMounted(() => {
    $graph.value.append(svg.node()!)
})


</script>
