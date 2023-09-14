<template>
  <template v-if="!item.hidden">
    <a-menu-item 
      class="left-menu-item"
      v-if="!item.children || item.children.length === 0"
      @click="handleRedirect(item.path)"
      :key="item.path">
      <Icon :type="isActive ? `${item.icon}-active` : item.icon" class="anticon" />
      <span class="left-menu-title">{{t(item.title)}}</span>
    </a-menu-item>

    <a-sub-menu v-else :key="`sub_${item.path}`">
      <template #title>
        <Icon v-if="item.icon" :type="item.icon" class="anticon" />
        <span class="left-menu-title">{{t(item.title)}}</span>
      </template>

      <a-menu-item v-for="childrenItem in item.children" :key="childrenItem.path">
        <span class="left-menu-title">{{t(childrenItem.title)}}</span>
      </a-menu-item>
    </a-sub-menu>
  </template>
</template>
<script lang="ts">

import { defineComponent, PropType, toRefs, computed, ComputedRef, Ref, watch } from 'vue';
import { useI18n } from "vue-i18n";
import { RoutesDataItem, getRouteBelongTopMenu, hasChildRoute } from '@/utils/routes';
import {DownOutlined, RightOutlined,} from '@ant-design/icons-vue';
import Icon from "./Icon.vue";
import { useRouter } from 'vue-router';

export default defineComponent({
  name: 'SiderMenuItem',
  props: {
    routeItem: {
      type: Object as PropType<RoutesDataItem>,
      required: true
    },
    topNavEnable: {
      type: Boolean,
      default: true
    },
    belongTopMenu: {
      type: String,
      default: ''
    },
    selectedKeys: {
      type: Array as PropType<string[]>,
      default: () => {
        return [];
      }
    }
  },
  components: {
    Icon
  },
  setup(props) {

    const { routeItem }: any = toRefs(props);
    const { t } = useI18n();
    const router = useRouter();

    const topMenuPath = computed<string>(()=> getRouteBelongTopMenu(routeItem.value as RoutesDataItem));

    const handleRedirect = (path: string) => {
      router.push(path);
    }

    const isActive = computed(() => {
      return (props.selectedKeys || []).includes(routeItem.value.path);
    });
  
    return {
      item: routeItem,
      topMenuPath,
      hasChildRoute,
      t,
      handleRedirect,
      isActive,
    }

  }
})
</script>