<template>
  <a-sub-menu v-if="menu?.children?.length" :key="menu.key" :expandIcon="null" :icon="null">
    <template v-if="menu.icon" #icon>
      <IconSvg :type="menu.icon" class="icon"/>
    </template>
    <template #expandIcon>
      <CaretRightOutlined class="expand-icon icon"/>
    </template>
    <template #title>
      <span :class="['sub-menu-text', 'menu-text', !menu.icon && 'menu-text-no-margin']">
         {{ menu.title }}
      </span>
    </template>
    <template v-for="menu in menu.children" :key="menu.key">
      <div class="scenario-tree-menu" >
        <SubMenu  :menu="menu" v-if="menu?.children?.length"/>
        <MenuItem v-else  :menu="menu" :disabled="menu?.disabled"/>
      </div>
    </template>
  </a-sub-menu>
</template>
<script setup lang="ts">
import {defineProps} from "vue";
import {CaretRightOutlined} from "@ant-design/icons-vue";
import IconSvg from "@/components/IconSvg";
import MenuItem from "./MenuItem.vue";
import SubMenu from "./SubMenu.vue";
const props = defineProps(['menu']);
</script>

<style lang="less" scoped>
.menu-text {
  display: inline-block;
  margin-left: 3px;
  width: 72px;

  &.menu-text-no-margin {
    margin-left: 0;
    margin-right: 22px;
  }
}
.expand-icon {
  position: relative;
  left: 16px;
  font-size: 12px;
}

.scenario-tree-menu {

  :deep(.ant-dropdown-menu-item) {
    .svg-icon {
      font-size: 18px !important;
      vertical-align: -0.26em;
    }
  }
}
</style>
