<template>
  <div class="home-wrap">
    <div class="home-header">
      <div class="home-header-left" :class="{'leyan-logo':isLeyanEnv}" @click="handleRedirect">
      </div>
      <div class="home-header-right">
        <UserSetting :theme="'white-theme'"/>
      </div>
    </div>
    <!-- <router-view></router-view> -->
    <slot />
    <RightTopUpdate />
  </div>
</template>
<script lang="ts">
import {defineComponent} from 'vue';
import {useRouter} from 'vue-router';
import UserSetting from './IndexLayout/components/RightTopSettings.vue';
import RightTopUpdate from './IndexLayout/components/RightTopUpdate.vue';

export default defineComponent({
  name: 'HomeLayout',
  components: {
    UserSetting,
    RightTopUpdate
  },
  setup() {
    const router = useRouter();
    let isLeyanEnv = process.env.VUE_APP_DEPLOY_ENV === 'ly';

    function handleRedirect() {
      router.push('/');
    }

    return {
      handleRedirect,
      isLeyanEnv
    }
  }
})
</script>
<style scoped lang="less">
.home-wrap {
  background: #F5F5F5;
  min-width: 1440px;

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
      background-image: url('../assets/images/logo.png');
      background-repeat: no-repeat;
      background-size: 100% 100%;
      cursor: pointer;
      &.leyan-logo{
        transform: scale(1.1) translateX(5px);

        background-image: url("https://od-1310531898.cos.ap-beijing.myqcloud.com/202306291016448.svg");
      }
    }
  }
}
</style>
