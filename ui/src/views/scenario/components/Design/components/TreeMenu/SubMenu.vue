<template>
  <a-sub-menu v-if="menu?.children?.length" :key="menu.key" :expandIcon="null" :icon="null">
    <template #icon>
      <IconSvg :type="menu.icon" class="icon"/>
    </template>
    <template #expandIcon>
      <CaretRightOutlined class="expand-icon icon"/>
    </template>
    <template #title>
      <span class="sub-menu-text menu-text">
         {{ menu.title }}
      </span>
    </template>
    <template v-for="menu in menu.children">
      <SubMenu :key="menu.key" :menu="menu" v-if="menu?.children?.length"/>
      <MenuItem v-else :key="menu.key" :menu="menu" :disabled="menu?.disabled"/>
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
  margin-left: 8px;
  width: 72px;
}
.expand-icon {
  position: relative;
  left: 16px;
  font-size: 12px;
}
</style>
