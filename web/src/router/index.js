import { createMemoryHistory, createRouter } from 'vue-router'

import homeComponent from '../views/homeComponent.vue'
import loginComponent from '../views/loginComponent.vue'

const routes = [
  {
    path: '/home',
    name: 'home',
    component: homeComponent
  },
  {
    path: '/',
    name: 'login',
    component: loginComponent
  }
]

const router = createRouter({
  history: createMemoryHistory(),
  routes,
})

export default router;