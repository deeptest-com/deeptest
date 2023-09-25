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
import {routes} from '@/config/routes';
import { StateType as GlobalStateType } from "@/store/global";
import { RouteMenuType } from '@/types/permission';
import { useRouter } from 'vue-router';

const { t } = useI18n();
const store = useStore<{ Global: GlobalStateType }>();

const r = routes[3] as any
const menuData: RoutesDataItem[] = vueRoutes(r.children);
const permissionRouteMenuMap = computed(() => store.state.Global.permissionMenuMap);

const tabs = computed(() => {
  console.log('999', tabs, menuData)

  const sysSettingMenu = menuData.find(e => e.path.includes('sys-setting'));
  const routeList = sysSettingMenu?.children;

  const ret = routeList?.filter(routeItem => {
    const notHidden = !routeItem.hidden

    const code = routeItem.meta?.code as string
    const key = RouteMenuType[code]
    const found = permissionRouteMenuMap.value[key]

    return notHidden && found
  })
  return ret
});

const router = useRouter();

const activeRoute = ref<any>('');

watch(() => {return [router.currentRoute.value.path, tabs.value];}, (v) => {
  const [path, list = []]: any = v;
  if (path && Array.isArray(list) && list.length > 0) {
    const find = list?.find((route: any) => path.includes(route.path));
    activeRoute.value = find && find.path;
    console.log(activeRoute.value);
  }
}, {immediate: true})

const handleChange = (item) => {
  console.log(item);
  router.push(item);
}
</script>