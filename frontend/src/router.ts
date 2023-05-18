import { App } from 'vue';
import { createRouter as createVueRouter, createWebHistory, Router } from 'vue-router'


export function createRouter(_app: App): Router {
    return createVueRouter({
        history: createWebHistory(),
        routes: [
            {
                name: 'home',
                path: '/',
                redirect: '/ledger',
            },
            {
                name: 'ledger',
                path: '/ledger',
                component: () => import('./views/Ledger.vue'),
            },

            {
                path: '/:pathMatch(.*)*',
                component: () => import('./views/NotFound.vue'),
            }

        ],
    });
}