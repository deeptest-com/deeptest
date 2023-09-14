<template>
  <a-tabs v-if="(tabs || []).length > 0" v-model:activeKey="activeRoute" @change="handleChange">
    <a-tab-pane 
      v-for="item in tabs" 
      :key="item?.path" 
      :tab="t(item.title)" />
  </a-tabs>
</template>
<script lang="ts" setup>
import { computed, ref, watch } from 'vue';
import { useStore } from 'vuex';  
import { useI18n } from "vue-i18n";
import { RoutesDataItem, vueRoutes } from '@/utils/routes';
import IndexLayoutRoutes from '@/layouts/IndexLayout/routes';
import { StateType as GlobalStateType } from "@/store/global";
import { RouteMenuType } from '@/types/permission';
import { useRouter } from 'vue-router';

const { t } = useI18n();
const store = useStore<{ Global: GlobalStateType }>();
const menuData: RoutesDataItem[] = vueRoutes(IndexLayoutRoutes);
const permissionRouteMenuMap = computed(() => store.state.Global.permissionMenuMap);
const tabs = computed(() => {
  const projectSettingMenu = menuData.find(e => e.path.includes('project-setting'));
  const routeList = projectSettingMenu?.children?.slice(1);
  return routeList?.filter(routeItem => !routeItem.hidden && permissionRouteMenuMap.value && permissionRouteMenuMap.value[RouteMenuType[`${routeItem.meta?.code}`]]);
});
const router = useRouter();

const activeRoute = ref<any>('');

const handleChange = (e) => {
  if (e === '/project-setting/enviroment') {
    router.push('/project-setting/enviroment/var');
  } else {
    router.push(e);
  }
}

watch(() => {
  return [router.currentRoute.value.path, tabs.value];
}, (v) => {
  const [path, list = []] = v;
  if (path && Array.isArray(list) && list.length > 0) {
    const find = list?.find((route: any) => path.includes(route.path));
    activeRoute.value = find && find.path;
    console.log(activeRoute.value);
  }
})
</script>