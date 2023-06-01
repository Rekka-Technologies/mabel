<template>
    <div v-if="disabled"
        class="select-none text-center py-2 flex rounded-lg border bg-gray-50/25 border-gray-400 text-gray-300 "
        :class="additionalClasses">
        <p class="text-center mx-auto my-auto">{{text}}</p>
    </div>
    <div @click="clickEvent" v-else
        :class="btnClass.concat(' ', additionalClasses)"
        class="select-none text-center py-2 group flex cursor-pointer rounded-lg ">
        
        <p class="text-center mx-auto my-auto">
            {{text}}
        </p>
    </div>
</template>

<script lang="ts">

export const btnType = {
    // Blue Buttons
    primary: 0,
    primary_outline: 1,

    // Green Buttons
    secondary: 2,
    secondary_outline: 3,
    
    // Red Buttons
    danger: 4,
    danger_outline: 5, 

    // Red Buttons
    grey_outline: 6,
}

export default {
    name: 'ButtonInput',
    props: {
        text: {
            type: String,
            required: true
        },
        type: {
            type: Number,
            required: false,
            default: btnType.primary
        },
        fit: {
            type: Boolean,
            required: false,
            default: false
        },
        class: {
            type: String,
            required: false,
            default: ''
        },
        OnClick: {
            type: Function,
            required: true
        },
        loader: {
            type: Boolean,
            required: false,
            default: false
        },
        disabled: {
            type: Boolean,
            required: false,
            default: false
        },
    },
    data() {
        return {
            loading: false,
        }
    },
    methods: {
        clickEvent() {
            if (this.loader) {
                this.loading = true;
                this.OnClick(this.stopLoading);
                return;
            }
            this.OnClick();
        },
        stopLoading() {
            this.loading = false;
        }
    },
    computed: {
        additionalClasses() {
            return this.class + (this.fit ? ' w-fit h-fit py-1 px-5' : 'w-full');
        },
        btnClass() {
            switch (this.type) {
                case btnType.primary_outline:
                    return 'border-2 border-blue-400 text-blue-400 hover:bg-blue-400 hover:text-white';
                case btnType.secondary:
                    return 'border-2 border-green-400 bg-green-400 text-white hover:bg-green-500 hover:border-green-500';
                case btnType.secondary_outline:
                    return 'border-2 border-green-400 text-green-400 hover:bg-green-400 hover:text-white';
                case btnType.danger:
                    return 'border-2 border-red-400 bg-red-400 text-white hover:bg-red-500 hover:border-red-500';
                case btnType.danger_outline:
                    return 'border-2 border-red-400 text-red-400 hover:bg-red-400 hover:text-white';
                case btnType.grey_outline:
                    return 'border text-gray-300 border-gray-500 hover:text-white hover:bg-gray-500';
                default:
                    return 'border-2 border-blue-400 bg-blue-400 text-white hover:bg-blue-500 hover:border-blue-500';
            }
        },
        isLoading() {
            return this.loading && this.loader;
        },
        loadingColor() {
            switch (this.type) {
                case btnType.primary_outline:
                    return 'bg-blue-400 group-hover:bg-white';
                case btnType.secondary:
                    return 'bg-green-400';
                case btnType.secondary_outline:
                    return 'bg-green-400 group-hover:bg-white';
                case btnType.danger:
                    return 'bg-red-400';
                case btnType.danger_outline:
                    return 'bg-red-400 group-hover:bg-white';
                case btnType.grey_outline:
                    return 'bg-gray-300 group-hover:bg-white';
                default:
                    return 'bg-white';
            }
        }
    }
}

</script>
