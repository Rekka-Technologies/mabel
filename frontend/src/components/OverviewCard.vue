<script lang="ts">

export const enum OverviewCardType {
    Blue = 0,
    Green = 1,
    Red = 2,
}

export default {
    name: 'OverviewCard',
    props: {
        title: String,
        value: Number,
        type: Number,
    },
    computed: {
        cardStyle() {
            let style = ['bg-gradient-to-tr', 'px-5', 'py-2', 'rounded-lg', 'w-1/3', 'mx-5']

            switch (this.type) {
                case OverviewCardType.Blue:
                    style.push('from-cyan-500', 'to-blue-500')
                    break
                case OverviewCardType.Green:
                    style.push('from-teal-500', 'to-green-500')
                    break
                case OverviewCardType.Red:
                    style.push('from-yellow-500', 'to-red-500')
                    break
                default:
                    style.push('from-gray-500', 'to-gray-500')
                    break
            }

            return style
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