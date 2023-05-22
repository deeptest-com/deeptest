<template>
  <div class="user-layout">
    <div class="logo"></div>
    <div class="right-main">
      <div class="right-content">
        <router-view></router-view>
      </div>
      <div class="right-footer">
        京ICP备11017824号-7 京ICP证130164号 北京市公安局朝阳分局备案编号:11010502034683
        Copyright © 2006-2023 ZCOOL中文 English
      </div>
    </div>
    <div class="lang">
    </div>
  </div>
</template>
<script lang="ts">
import { computed, defineComponent, ref } from "vue";
import { useRoute } from 'vue-router';
import { getRouteItem, RoutesDataItem, vueRoutes } from '@/utils/routes';
import UserLayoutRoutes from './routes';
import useTitle from '@/composables/useTitle';

export default defineComponent({
  name: 'UserLayout',
  components: {
  },
  setup() {
    const route = useRoute();

    // 所有菜单路由
    const menuData = ref<RoutesDataItem[]>(vueRoutes(UserLayoutRoutes, '/user'));

    // 当前路由 item
    const routeItem = computed<RoutesDataItem>(() => getRouteItem(route.path, menuData.value as RoutesDataItem[]));

    // 设置title
    useTitle(routeItem);

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