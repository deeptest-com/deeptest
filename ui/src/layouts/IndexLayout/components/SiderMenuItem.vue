<template>
  <template v-if="!item.hidden">
    <a-menu-item class="left-menu-item" v-if="(item.children && item.children.length === 0) || !item.children" :key="item.path">
      <a-link :to="item.path">
        <Icon v-if="item.icon" :type="item.icon" class="anticon" />
        <span class="left-menu-title" style="margin-left: 5px;">{{t(item.title)}}</span>
      </a-link>
    </a-menu-item>
    <a-sub-menu v-else :key="`sub_${item.path}`">
      <template #title>
        <Icon v-if="item.icon" :type="item.icon" class="anticon" />
        <span class="left-menu-title" style="margin-left: 5px;">{{t(item.title)}}</span>
      </template>
      <a-menu-item v-for="childrenItem in item.children" :key="childrenItem.path">
        <a-link :to="childrenItem.path">
          <span class="left-menu-title" style="margin-left: 5px;">{{t(childrenItem.title)}}</span>
        </a-link>
      </a-menu-item>
    </a-sub-menu>
  </template>
</template>
<script lang="ts">

import { defineComponent, PropType, toRefs, computed, ComputedRef, Ref } from 'vue';
import { useI18n } from "vue-i18n";
import { RoutesDataItem, getRouteBelongTopMenu, hasChildRoute } from '@/utils/routes';
import ALink from '@/components/ALink/index.vue';
import Icon from "./Icon.vue";

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
    }
  },
  components: {
    ALink,
    Icon
  },
  setup(props) {

    const { routeItem } = toRefs(props);
    const { t } = useI18n();

    const topMenuPath = computed<string>(()=> getRouteBelongTopMenu(routeItem.value as RoutesDataItem));

    return {
      item: routeItem,
      topMenuPath,
      hasChildRoute,
      t
    }

  }
})
</script>