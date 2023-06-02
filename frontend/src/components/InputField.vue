<script lang="ts">

export default {
    name: 'InputField',
    props: {
        value: {
            type: String,
            required: true,
            default: ''
        },
        OnChange: {
            type: Function,
            required: true,
            default: (_value: string) => {}
        },
        OnEnter: {
            type: Function,
            required: false,
            default: () => {}
        },
        label: {
            type: String,
            required: false,
            default: ''
        },
        disabled: {
            type: Boolean,
            required: false,
            default: false
        },
        hidden: {
            type: Boolean,
            required: false,
            default: false
        },
        validate: {
            type: Function,
            required: false,
            default: (_value : string) => true
        },
        placeholder: {
            type: String,
            required: false,
            default: ''
        },
    },
    computed: {
        validValue() {
            return this.validate(this.value);
        },
        inputType() {
            return this.hidden ? 'password' : 'text';
        }
    }
}


</script>

<template>
    <div class="w-full">
        <!-- Optional Label for the Input Field -->
        <p v-if="label">{{ label }}</p>

        <input :type="inputType" v-if="!disabled"
            class="block w-full p-2 border rounded-lg bg-white border-gray-200 placeholder-gray-300 text-black focus:ring-gray-500 focus:border-gray-500"
            :class="{'border-red-400 focus:ring-red-400 focus:border-red-400': !validValue}"
            :placeholder="placeholder" :value="value"
            @input="(e) => OnChange((e.target as HTMLInputElement).value)"
            @keyup.enter="() => { if (validValue) OnEnter()} "
        >

        <!-- What to display if the input field is disabled -->
        <input :type="inputType" v-else disabled
            class="block w-full select-none p-2 rounded-lg border bg-gray-200 border-gray-200 placeholder-gray-10 text-white"
            :placeholder="placeholder" :value="value"
            @input="(e) => OnChange((e.target as HTMLInputElement).value)"
        >
    </div>

</template>