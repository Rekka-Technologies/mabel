import { createRouter, createWebHistory } from 'vue-router'

import { useAuthStore } from './store/auth'

export const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            name: 'home',
            path: '/',
            redirect: '/ledger',
        },
        {
            name: 'login',
            path: '/login',
            component: () => import('./views/Login.vue'),
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

router.beforeEach(async (to) => {
    // redirect to login page if not logged in and trying to access a restricted page
    const publicPages = ['/login'];
    const authRequired = !publicPages.includes(to.path);
    const auth = useAuthStore();

    if (authRequired && !auth.token) {
        auth.returnUrl = to.fullPath;
        return '/login';
    }
});