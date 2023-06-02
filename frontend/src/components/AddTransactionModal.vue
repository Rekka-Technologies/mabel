<script lang="ts">

import { LedgerRecord } from '../models/LedgerRecord';

import Modal from './Modal.vue';
import InputField from './InputField.vue';
import ButtonInput, { btnType } from './ButtonInput.vue';

export default {
    name: 'AddTransactionModal',
    components: {
        Modal,
        InputField,
        ButtonInput
    },
    props: {
        open: { // Only opens modal when value is changed
            type: Boolean,
            default: false
        },
        SubmitTransaction: {
            type: Function,
            default: (_record: LedgerRecord, closeCb: Function) => { closeCb(); }
        },

    },
    data() {
        return {
            newTransaction: LedgerRecord.empty(),
            openModel: false,

            BTN_ADD: btnType.primary,
            BTN_CLOSE: btnType.danger_outline,
        }
    },
    watch: {
        open: function () {
            // Open modal on change
            this.modalOpen();
        },
    },
    methods: {
        cleanModal() {
            this.newTransaction = LedgerRecord.empty();
        },
        modalOpen() {
            this.openModel = true;
            this.cleanModal();
        },
        modalClose() {
            this.openModel = false;
            this.cleanModal();
        },
    },
}

</script>

<template>
    <Modal :class="{ 'hidden': !openModel }" :OnClose="() => modalClose()">
        <template #header>
            <h1>Add a New Transaction</h1>
        </template>
        <template #body>
            <InputField :value="newTransaction.name" label="Name" placeholder="eg. Network Switch"
                :OnChange="(value: string) => newTransaction.name = value" />
            <vue-tailwind-datepicker as-single v-model="newTransaction.date"
                :formatter="{ date: 'DD/MM/YYYY', month: 'MMM' }" />
            <InputField :value="newTransaction.amount" label="Amount" placeholder="eg. $36.30"
                :OnChange="(value: string) => newTransaction.amount = +value"
                :validate="(value: string) => !isNaN(parseFloat(value))" />
            <InputField :value="newTransaction.reference" label="Reference [Optional]" placeholder="eg. #000001"
                :OnChange="(value: string) => newTransaction.reference = value" />
        </template>
        <template #footer>
            <ButtonInput text="Close" :type="BTN_CLOSE" :OnClick="() => SubmitTransaction(newTransaction, modalClose)" />
            <ButtonInput text="Submit" :type="BTN_ADD" :OnClick="() => SubmitTransaction(newTransaction, modalClose)" />
        </template>

    </Modal>
</template>