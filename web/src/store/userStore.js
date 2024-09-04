// src/stores/userStore.js
import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', () => {
    let userName = ''
    let isLogin = false
    let token = ''
    return { userName, isLogin, token }
  });
