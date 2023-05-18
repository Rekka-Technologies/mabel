<script lang="ts">

import OverviewCard from '../components/OverviewCard.vue';

export default {
    name: 'Ledger',
    components: {
        OverviewCard,
    },
}

</script>

<template>
    <div class="mx-auto max-w-screen-md p-4">
        <!-- Intro to page -->
        <div>
            <h1 class="text-3xl font-mono font-medium">#Ledger</h1>    
            <p class="w-full bg-gray-100 border-l-8 border-gray-600 pl-2 mb-2">
                Last Updated: 16/06/2023
            </p>
            <p>
                Welcome to the ledger page. This is where you can view all the 
                transactions that have been logged. This also includes a quick 
                total overview.
            </p>
        </div>

        <!-- Overview Cards -->
        <div class="flex flex-row justify-between my-5">
            <OverviewCard title="Credit" value="$9,000" color1="cyan" color2="blue" />
            <OverviewCard title="Debit" value="$9,000" color1="yellow" color2="red" />
            <OverviewCard title="Transactions" value="9000" color1="teal" color2="green" />
        </div>

        <!-- Options Bar -->
        <div class="flex flex-row text-lg mb-5">
            <!-- Search Bar -->
            <input type="text" class="w-full border rounded px-2" placeholder="Search...">

            <!-- Export and Add Buttons -->
            <div class="bg-gray-600 hover:bg-gray-800 rounded flex px-2 ml-2 cursor-pointer">
                <p class="m-auto text-white">Export</p>
            </div>
            <div class="bg-gray-600 hover:bg-gray-800 rounded flex px-2 ml-2 cursor-pointer">
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
                        <th>Tags</th>
                    </tr>
                </thead>
                <tbody>
                    <tr class="rounded hover:bg-gray-500 hover:text-white cursor-pointer">
                        <td>16/06/2023</td>
                        <td>Tillered Invoice</td>
                        <td class="text-center">A000013</td>
                        <td class="text-green-500 text-right font-semibold">7000.00</td>
                        <td class="flex flex-row justify-center">
                            <p class="px-1 py-0.5 mx-0.5 bg-green-500 rounded text-white text-sm uppercase font-semibold">Wage</p>
                        </td>
                    </tr>
                    <tr class="rounded hover:bg-gray-500 hover:text-white cursor-pointer">
                        <td>16/06/2023</td>
                        <td>TP-Link Deco x20</td>
                        <td></td>
                        <td class="text-red-500 text-right font-semibold">329.00</td>
                        <td class="flex flex-row justify-center">
                            <p class="px-1 py-0.5 mx-0.5 bg-blue-500 rounded text-white text-sm uppercase font-semibold">Inventory</p>
                        </td>
                    </tr>
                    <tr class="rounded hover:bg-gray-500 hover:text-white cursor-pointer">
                        <td>16/06/2023</td>
                        <td>Rent & Utilities</td>
                        <td></td>
                        <td class="text-red-500 text-right font-semibold">2000.00</td>
                        <td class="flex flex-row justify-center">
                            <p class="px-1 py-0.5 mx-0.5 bg-purple-500 rounded text-white text-sm uppercase font-semibold">Utilities</p>
                        </td>
                    </tr>
                    <tr class="rounded hover:bg-gray-500 hover:text-white cursor-pointer">
                        <td>16/06/2023</td>
                        <td>VPS Server</td>
                        <td></td>
                        <td class="text-red-500 text-right font-semibold">6.32</td>
                        <td class="flex flex-row justify-center">
                            <p class="px-1 py-0.5 mx-0.5 bg-purple-500 rounded text-white text-sm uppercase font-semibold">Utilities</p>
                        </td>
                    </tr>
                    <tr class="rounded hover:bg-gray-500 hover:text-white cursor-pointer">
                        <td>16/06/2023</td>
                        <td>Super Contributions</td>
                        <td></td>
                        <td class="text-red-500 text-right font-semibold">700.00</td>
                        <td class="flex flex-row justify-center">
                            <p class="px-1 py-0.5 mx-0.5 bg-amber-500 rounded text-white text-sm uppercase font-semibold">Super</p>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>