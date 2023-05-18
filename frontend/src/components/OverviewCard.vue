<script lang="ts">

export default {
    name: 'OverviewCard',
    props: {
        title: String,
        value: Number,
        color1: String,
        color2: String,
    },
    computed: {
        colorFrom() {
            if (this.color1 == undefined)
                return 'from-cyan-500';
            return `from-${this.color1}-500`
        },
        colorTo() {
            if (this.color2 == undefined)
                return 'to-cyan-500';
            return `to-${this.color2}-500`
        },
        cardStyle() {
            return `bg-gradient-to-tr ${this.colorFrom} ${this.colorTo} px-5 py-2 rounded-lg w-1/3 mx-5`
        }
    },
}

</script>

<template>
    <div :class="cardStyle">
        <p class="text-sm text-white font-semibold uppercase">{{ title }}</p>
        <p class="text-2xl text-white font-semibold">{{ value }}</p>
    </div>
</template>