<script lang="ts">

// Store imports
import { mapActions, mapState } from 'pinia'
import { useTransactionStore } from '../store/transaction';
import { useAuthStore } from '../store/auth';

// Model and Component imports
import { LedgerRecord } from '../models/LedgerRecord';
import OverviewCard, {OverviewCardType} from '../components/OverviewCard.vue';
import InputField from '../components/InputField.vue';

// Search imports
import Fuse from 'fuse.js'

export default {
    name: 'Ledger',
    components: {
        OverviewCard,
        InputField,
    },
    data() {
        return {
            CARD_BLUE: OverviewCardType.Blue,
            CARD_GREEN: OverviewCardType.Green,
            CARD_RED: OverviewCardType.Red,

            searchQuery: '',
            searchedData: [] as LedgerRecord[],
            fuseSearch: new Fuse<LedgerRecord>([]),
            searchParams: ['date', 'name', 'reference', 'amount'],
        }
    },
    methods: {
        ...mapActions(useTransactionStore, ['fetchTransactions', 'addTransaction']),
        ...mapActions(useAuthStore, ['logout']),

        // Convert date from YYYY-MM-DD to DD/MM/YYYY
        convertDate(date: string): string {
            const dateParts = date.split('-');
            return `${dateParts[2]}/${dateParts[1]}/${dateParts[0]}`
        },
    },
    computed: {
        ...mapState(useAuthStore, ['getUserName']),
        ...mapState(useTransactionStore, ['getTransactions']),

        // Transactions to display
        displayTransactions(): LedgerRecord[] {
            if (this.searchQuery === '') {
                return this.getTransactions;
            }

            // Search for the query
            const results = this.fuseSearch.search(this.searchQuery);
            return results.map((r) => r.item);
        },

        credit(): string {
            let total = 0.0;
            this.getTransactions.map((t) => {
                if (t.amount > 0) {
                    total += t.amount;
                }
            });
            return `$${total.toFixed(2)}`;
        },
        debit(): string {
            let total = 0.0;
            this.getTransactions.map((t) => {
                if (t.amount < 0) {
                    total += (t.amount * -1.0);
                }
            });
            return `$${total.toFixed(2)}`;
        },
        numTransactions(): number {
            return this.getTransactions.length;
        },

        // Display the 
        lastUpdated(): string {
            let curDate = new Date(0);

            // Find the newest transaction
            this.getTransactions.map((t) => {
                const transactionDate = new Date(t.date);
                if (transactionDate > curDate) {
                    curDate = transactionDate;
                }
            });

            // There is a possibility of no transactions, so don't display 
            // a date.
            if (curDate.getTime() === 0) {
                return "No transactions";
            }

            // Convert to the DD/MM/YYYY format
            return this.convertDate(curDate.toISOString().split('T')[0]);
        },
    },
    watch: {
        // Update the searchable data when the transactions list is either 
        // ammended to or changed in general
        getTransactions: function(newVal) {
            this.fuseSearch = new Fuse<LedgerRecord>(newVal, {
                keys: this.searchParams,
                threshold: 0.3,
            });
        }
    },
    mounted() {
        this.fetchTransactions(() => {}, () => {console.log("Failed to get transactions")});
    }
}

</script>

<template>
    <div class="mx-auto max-w-screen-md p-4">
        <!-- Intro to page -->
        <div>
            <div class="flex flex-row items-center justify-between text-right">
                <h1 class="text-3xl font-mono font-medium">#Ledger</h1>
                <div class="my-1">
                    <p class="font-semibold text-gray-900 leading-none">{{ getUserName }}</p>
                    <p class="text-sm text-gray-500 cursor-pointer" @click="logout">Logout</p>
                </div>    
            </div>
            <p class="w-full bg-gray-100 border-l-8 border-gray-600 pl-2 mb-2">
                Last Updated: {{ lastUpdated }}
            </p>
            <p>
                Welcome to the ledger page. This is where you can view all the 
                transactions that have been logged. This also includes a quick 
                total overview.
            </p>
        </div>

        <!-- Overview Cards -->
        <div class="flex flex-row my-5 items-center flex-wrap justify-center sm:flex-nowrap sm:flex-row sm:justify-between">
            <OverviewCard title="Credit" :value="credit" :type="CARD_BLUE" />
            <OverviewCard title="Debit" :value="debit" :type="CARD_RED" />
            <OverviewCard class="mt-2 sm:mt-0" title="Transactions" :value="numTransactions" :type="CARD_GREEN" />
        </div>

        <!-- Options Bar -->
        <div class="flex flex-row text-lg mb-5 space-x-3">
            <!-- Search Bar -->
            <input-field :value="searchQuery" placeholder="Search..." :OnChange="(val : string) => searchQuery = val" />

            <!-- Export and Add Buttons -->
            <div class="bg-gray-600 hover:bg-gray-800 rounded flex px-2 cursor-pointer">
                <p class="m-auto text-white">Export</p>
            </div>
            <div class="bg-gray-600 hover:bg-gray-800 rounded flex px-2 cursor-pointer">
                <p class="m-auto text-white">Add</p>
            </div>
        </div>

        <!-- Table -->
        <div class="bg-gray-100 w-full rounded-lg px-4 py-2">
            <table class="table-auto w-full">
                <thead>
                    <tr>
                        <th class="text-left">Date</th>
                        <th class="text-left">Name</th>
                        <th>Reference</th>
                        <th>Amount</th>
                    </tr>
                </thead>
                <tbody>

                    <tr class="rounded cursor-pointer" v-for="t in displayTransactions" :key="t.id">
                        <td>{{ convertDate(t.date) }}</td>
                        <td>{{ t.name }}</td>
                        <td class="text-center">{{ t.reference }}</td>

                        <!-- Colour Ammount accordingly -->
                        <td v-if="t.amount > 0" class="text-green-500 text-right font-semibold">{{ t.amount.toFixed(2) }}</td>
                        <td v-else class="text-red-500 text-right font-semibold">{{ (t.amount * -1.0).toFixed(2) }}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>