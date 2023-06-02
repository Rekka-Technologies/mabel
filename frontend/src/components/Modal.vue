<script lang="ts">

export default {
    name: 'Modal',
    props: {
        OnClose: {
            type: Function,
            default: () => { }
        },
    },
}

</script>

<template>
    <div tabindex="-1"
        class="fixed flex top-0 left-0 right-0 z-50 w-full p-4 overflow-x-hidden overflow-y-auto md:inset-0 h-full bg-gray-900/75"
        @click.self="() => OnClose()">
        <div class="relative w-full max-w-2xl m-auto flex">
            <!-- Modal content -->
            <div class="relative my-auto rounded-lg shadow bg-white w-full">
                <!-- Modal header -->
                <div
                    class="flex items-start justify-between p-4 border-b rounded-t border-gray-600 text-xl font-semibold text-gray-600">
                    <slot name="header"></slot>
                    <button type="button"
                        class="text-gray-700 font-mono bg-transparent hover:bg-red-500 hover:text-white rounded-lg text-sm p-1.5 ml-auto inline-flex items-center"
                        @click="() => OnClose()">
                        <font-awesome-icon icon="xmark" class="text-xl" />
                        <span class="sr-only">Close modal</span>
                    </button>
                </div>
                <!-- Modal body -->
                <div class="p-6 space-y-6 text-base leading-relaxed text-gray-800">
                    <slot name="body"></slot>
                </div>
                <!-- Modal footer -->
                <div class="flex items-center p-6 space-x-2 border-t rounded-b border-gray-600">
                    <slot name="footer"></slot>
                </div>
            </div>
        </div>
    </div>
</template>


