<script lang="ts">
import InputField from '../components/InputField.vue';
import ButtonInput from '../components/ButtonInput.vue';

import { mapActions } from 'pinia'
import { useAuthStore } from '../store/auth';

export default {
    name: 'Login',
    components: {
        InputField,
        ButtonInput
    },
    data() {
        return {
            show_register: false,
            Username: '',
            Password: '',
            errorMsg: '',
        }
    },
    computed: {
        showRegister() {
            return this.show_register;
        },
        btnText() {
            return this.show_register ? 'Register' : 'Login';
        },
    },
    methods: {
        ...mapActions(useAuthStore, ['login', 'register']),
        submit() {
            if (this.showRegister) {
                // Register
                this.register(this.Username, this.Password,
                    () => {
                        console.log("Registered")
                    },
                    (err : string) => {
                        // Failure
                        this.errorMsg = "Failed to register: " + err;
                    }    
                );
            } else {
                // Login
                this.login(this.Username, this.Password,
                    (err : string) => {
                        // Failure
                        this.errorMsg = "Failed to login: " + err;
                    }    
                );
            }
        }
    }
}

</script>

<template>
    <div class="mx-auto max-w-sm p-4 h-screen flex">
        <div class="bg-gray-100 rounded-lg my-auto w-full p-2 flex flex-col space-y-3">
            <h1 class="text-gray-900 text-2xl text-center font-sans ">Welcome!</h1>

            <p v-if="showRegister" class="text-gray-900">Already have an account? Click <span class="text-blue-500 cursor-pointer" @click="() => show_register = !show_register">here</span> to login.</p>
            <p v-else class="text-gray-900">Please login to continue to your Ledger. If you are a new user, click <span class="text-blue-500 cursor-pointer" @click="() => show_register = !show_register">here</span> to register.</p>
            
            
            <input-field placeholder="Username" :value="Username" :OnChange="(val : string) => { Username = val}" />
            <input-field placeholder="Password" :value="Password" :OnChange="(val : string) => { Password = val}" />
            <button-input :text="btnText" :OnClick="submit" />

            <p v-if="errorMsg" class="text-red-500">{{ errorMsg }}</p>
        </div>
    </div>
</template>