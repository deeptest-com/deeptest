<template>
  <template v-if="!item.hidden">
    <a-menu-item :key="item.path">
      <a-link :to="item.path">
        <Icon v-if="item.icon" :type="item.icon" class="anticon" />
        <span style="margin-left: 5px;">{{t(item.title)}}</span>
      </a-link>
    </a-menu-item>
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