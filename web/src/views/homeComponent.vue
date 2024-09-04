<template>
    <div>Home</div>
</template>
  
<script setup>
import { useUserStore } from '@/store/userStore';
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';

const userStore = useUserStore();
const router = useRouter();
//如果未登录，导航到\login进行登录
onMounted(() => {
    if(!userStore.isLogin){
        router.replace({ path: '/login' });
    }
})

</script>
  
<style>

</style>
  