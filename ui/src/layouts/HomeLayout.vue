<template>
    <div class="home-wrap">
        <div class="home-header">
            <div class="home-header-left" @click="handleRedirect">
            </div>
            <div class="home-header-right">
                <UserSetting :theme="'white-theme'" />
            </div>
        </div>
        <router-view></router-view>
    </div>
</template>
<script lang="ts">
import { defineComponent } from 'vue';
import { useRouter } from 'vue-router';
import UserSetting from './IndexLayout/components/RightTopSettings.vue';

export default defineComponent({
    name: 'HomeLayout',
    components: {
        UserSetting
    },
    setup() {
        const router = useRouter();

        function handleRedirect() {
            router.push('/');
        }
        return {
            handleRedirect
        }
    }
})
</script>
<style scoped lang="less">
.home-wrap {
    background: #F5F5F5;

    .home-header {
        width: 100%;
        height: 64px;
        background-color: #2A325A;
        display: flex;
        align-items: center;
        justify-content: space-between;
        padding-left: 24px;
        box-sizing: border-box;

        .home-header-left {
            width: 105px;
            height: 35px;
            background-image: url('../assets/images/logo-deep.png');
            background-repeat: no-repeat;
            background-size: 100% 100%;
            cursor: pointer;
        }
    }
}
</style>
