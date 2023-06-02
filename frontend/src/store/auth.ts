import { defineStore } from 'pinia'

import { router } from '../router'

const BASE_URL = 'http://localhost:8080/auth'

export const useAuthStore = defineStore('auth', {
    state: () => ({ 
        token: localStorage.getItem('mable'),
        returnUrl: '',
    }),
    getters: {
        getUserName(): string {
            if (!this.token) return ''

            const parsedToken = JSON.parse(atob(this.token.split('.')[1]))
            return !parsedToken.user_name ? '' : parsedToken.user_name.toString()
        }
    },
    actions: {
        register(username: string, password: string, successCb: () => void, errorCb: (err : string) => void) {
            fetch(`${BASE_URL}/register`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ username, password }),
                })
                .then((response) => {
                    response.json()
                        .then((data) => {
                            if (response.status !== 200) {
                                throw new Error(data.error)
                            }

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
        },
        login(username: string, password: string, errorCb: (err : string) => void) {
            fetch(`${BASE_URL}/login`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ username, password }),
                })
                .then((response) => {
                    response.json()
                        .then((data) => {
                            if (response.status !== 200) {
                                throw new Error(data.error)
                            }

                            // Store the token in local storage (for now - should be in a cookie)
                            // TODO: Store the token in a cookie
                            localStorage.setItem('mable', data.token)
                            this.token = data.token

                            // redirect to previous url or default to home page
                            if (this.returnUrl !== '') {
                                router.push(this.returnUrl)
                            } else {
                                router.push('/')
                            }
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
        logout() {
            // remove user from local storage to log user out
            localStorage.removeItem('mable')
            this.token = ''
            router.push('/login')
        }
    },
  })