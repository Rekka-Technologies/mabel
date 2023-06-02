import { defineStore } from 'pinia'

import { LedgerRecord } from '../models/LedgerRecord'

const BASE_URL = 'http://localhost:8080/api'

export const useTransactionStore = defineStore('transaction', {
    state: () => ({ 
        transactions: [] as LedgerRecord[],
    }),
    getters: {
        getTransactions(): LedgerRecord[] {
            return this.transactions
        }
    },
    actions: {
        fetchTransactions(successCb: (data : LedgerRecord[]) => void, errorCb: (err : string) => void) {
            fetch(`${BASE_URL}/transactions`, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': 'Bearer ' + localStorage.getItem('mable') || '',
                    },
                })
                .then((response) => {
                    response.json()
                        .then((data) => {
                            if (response.status !== 200) {
                                throw new Error(data.error)
                            }

                            // Set the transactions
                            this.transactions = LedgerRecord.fromJsonArray(data.data)
                            successCb(this.transactions)
                        })
                        .catch((error) => {
                            errorCb(error.message)
                        }
                    )
                })
                .catch((error) => {
                    errorCb(error.message)
                }
            )
        },
        addTransaction(transaction: LedgerRecord, successCb: () => void, errorCb: (err : string) => void) {
            // Validate the transaction
            if (!transaction.Validate()) {
                errorCb('Invalid transaction')
                return
            }

            // Send the transaction to the backend
            fetch(`${BASE_URL}/transactions`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': 'Bearer ' + localStorage.getItem('mable') || '',
                    },
                    body: transaction.toJson(),
                })
                .then((response) => {
                    response.json()
                        .then((data) => {
                            if (response.status !== 200) {
                                throw new Error(data.error)
                            }

                            // Update State with the new transaction
                            this.transactions.push(LedgerRecord.fromJson(data.data))
                            successCb()
                        })
                        .catch((error) => {
                            errorCb(error.message)
                        }
                    )
                })
                .catch((error) => {
                    errorCb(error.message)
                }
            )
        }
      }
  })