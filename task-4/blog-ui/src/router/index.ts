import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'Login',
      component: () => import('@/views/Login.vue'),
      meta: { guest: true },
    },
    {
      path: '/register',
      name: 'Register',
      component: () => import('@/views/Register.vue'),
      meta: { guest: true },
    },
    {
      path: '/',
      component: () => import('@/layouts/MainLayout.vue'),
      meta: { requiresAuth: true },
      children: [
        {
          path: '',
          name: 'PostList',
          component: () => import('@/views/PostList.vue'),
        },
        {
          path: 'posts/create',
          name: 'PostCreate',
          component: () => import('@/views/PostEditor.vue'),
        },
        {
          path: 'posts/:id',
          name: 'PostDetail',
          component: () => import('@/views/PostDetail.vue'),
          props: true,
        },
        {
          path: 'posts/:id/edit',
          name: 'PostEdit',
          component: () => import('@/views/PostEditor.vue'),
          props: true,
        },
      ],
    },
    {
      path: '/:pathMatch(.*)*',
      redirect: '/',
    },
  ],
})

router.beforeEach((to) => {
  const userStore = useUserStore()
  if (to.meta.requiresAuth && !userStore.isLoggedIn) {
    return { name: 'Login', query: { redirect: to.fullPath } }
  }
  if (to.meta.guest && userStore.isLoggedIn) {
    return { name: 'PostList' }
  }
  return true
})

export default router
