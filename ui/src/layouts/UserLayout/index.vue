<template>
  <div class="user-layout">
    <div class="logo" :class="{'ly-logo':isLeyanEnv}"></div>
    <div class="right-main">
      <div class="right-content">
        <router-view></router-view>
      </div>
      <div class="right-footer">
        Copyright © 2018 @deeptest.com 
      </div>
    </div>
    <div class="lang">
    </div>
  </div>
</template>
<script lang="ts">
import { computed, defineComponent, onMounted, ref } from "vue";
import { useRoute } from 'vue-router';
import { getRouteItem, RoutesDataItem, vueRoutes } from '@/utils/routes';
import UserLayoutRoutes from './routes';
import useTitle from '@/composables/useTitle';

export default defineComponent({
  name: 'UserLayout',
  components: {
  },
  setup() {

    let isLeyanEnv = process.env.VUE_APP_DEPLOY_ENV === 'ly';

    const route = useRoute();

    // 所有菜单路由
    const menuData = ref<RoutesDataItem[]>(vueRoutes(UserLayoutRoutes, '/user'));

    // 当前路由 item
    const routeItem = computed<RoutesDataItem>(() => getRouteItem(route.path, menuData.value as RoutesDataItem[]));

    // 设置title
    useTitle(routeItem);

    onMounted(() => {
      const appLoadingEl = document.getElementsByClassName('app-loading');
      if (appLoadingEl[0]) {
        appLoadingEl[0].classList.add('hide');
        setTimeout(() => {
            document.body.removeChild(appLoadingEl[0]);
        }, 600);
      }
    });

    return {
      isLeyanEnv
    }

  }
})
</script>
<style lang="less" scoped>
.user-layout {
  display: flex;
  justify-content: flex-end;
  min-width: 1440px;
  min-height: 960px;
  width: 100vw;
  height: 100vh;
  background-image: url('../../assets/images/bg-smooth.png');
  background-position: center;
  background-size: cover;

  .logo {
    width: 200px;
    height: 66.66667px;
    position: fixed;
    left: 72px;
    top: 62px;
    background-image: url('../../assets/images/logo.png');
    background-position: center;
    background-size: cover;
   background-repeat: no-repeat;
    &.ly-logo{
      background-image: url("https://leyanapi.nancalcloud.com/upload/images/202306291016448.svg");
      background-size: contain;
    }
  }

  .right-main {
    min-width: 684px;
    width: 50%;
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;

    .right-content {
      padding-bottom: 86px;
    }

    .right-footer {
      width: 412px;
      font-family: 'PingFang SC';
      font-weight: 400;
      font-size: 12px;
      line-height: 20px;
      text-align: center;
      letter-spacing: -0.559811px;
      color: rgba(255, 255, 255, 0.6);
      position: absolute;
      bottom: 46px;
      left: 50%;
      transform: translateX(-50%);
    }
  }

  .lang {
    position: absolute;
    top: 20px;
    right: 50px;
    color: #ffffff;
    font-size: 16px;
  }
}
</style>
