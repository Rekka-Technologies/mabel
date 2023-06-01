import { defineStore } from 'pinia'

import { LedgerRecord } from '../models/LedgerRecord'

export const useAuthStore = defineStore('transaction', {
    state: () => ({ 
        transactions: [] as LedgerRecord[],
    }),
    actions: {
      actions: {
        async getTransactions() {
            const response = await fetch('http://localhost:8000/api/transactions/', {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + localStorage.getItem('mable') || '',
                },
            })
            if (response.status !== 200) {
                throw new Error('Failed to fetch transactions')
            }

            const data = await response.json()
            this.transactions = LedgerRecord.fromJsonArray(data.data)
        },
        async addTransaction(transaction: LedgerRecord) {
            const response = await fetch('http://localhost:8000/api/transactions/', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + localStorage.getItem('mable') || '',
                },
                body: JSON.stringify(transaction),
            })

            if (response.status !== 200) {
                throw new Error('Failed to add transactions')
            }

            const data = await response.json()
            this.transactions.push(LedgerRecord.fromJson(data.data))
        }
      }
    },
  })