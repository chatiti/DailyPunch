<template>
    <div class="login-container">
        <p class="title">Daily<font color="#28a7cb"><i>Punch</i></font></p>
        <van-form>
            <van-field
                    v-model="username"
                    name="用户名"
                    label="用户名"
                    placeholder="用户名"
                    :rules="[{ required: true, message: '请填写用户名' }]"
            />
            <van-field
                    v-model="password"
                    type="password"
                    name="密码"
                    label="密码"
                    placeholder="密码"
                    :rules="[{ required: true, message: '请填写密码' }]"
            />
            <div class="button-container">
                <van-button round block type="primary" native-type="submit" @click="onSubmit">登录</van-button>
            </div>
            <div class="reg">
                <div @click="toRegister">没有账号？立即注册</div>
            </div>
        </van-form>
    </div>
</template>
  
<script setup>
import { useUserStore } from '@/store/userStore';
import { Axios } from 'axios';
import { Notify } from 'vant';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

let userStore = useUserStore();
let curRouter = useRouter();
let username = ref('');
let password = ref('');
async function onSubmit() {
    try {
        const res = await Axios.post('https://localhost:8080/Login', {
            username: username.value, 
            password: password.value,
        })
        userStore.userName = username.value
        userStore.isLogin = true
        userStore.token = res.data.Token
        Notify({ type: 'success', message: '登录成功' })
        curRouter.replace({ path: '/home' })
        
    } catch (error) {
        Notify({ type: 'danger', message: '登录失败，请检查用户名和密码' });
        console.error(error);
    }
}
</script>
  
<style scoped>
.login-container {
    top: 70%;
    width: 70%;
    margin: 0 auto;
    padding: 20px;
    background-color: #f8f9fa;
    border-radius: 8px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.title {
    text-align: center;
    font-size: 36px;
    font-weight: 900;
    color: #343a40;
    margin-bottom: 20px;
}

.button-container {
    margin: 20px 0;
}

.reg {
    text-align: center;
    color: #007bff;
    cursor: pointer;
    margin-top: 10px;
}

.reg:hover {
    text-decoration: underline;
}
</style>
  