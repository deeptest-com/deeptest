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
    <a-menu-item v-for="menu in menu.children" :key="menu.key">
      <template #icon>
        <IconSvg :type="menu.icon" class="icon"/>
      </template>
      <span class="menu-text">
                {{ menu.title }}
              </span>
    </a-menu-item>
  </a-sub-menu>
</template>

<script setup lang="ts">
import {defineProps, createVNode} from "vue";
import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';
import {CaretRightOutlined} from "@ant-design/icons-vue";
import {DESIGN_MENU_CONFIG} from "./config";
import IconSvg from "@/components/IconSvg";
import TreeSubMenu from "./TreeSubMenu.vue";

const props = defineProps(['menu']);
</script>

<style lang="less" scoped>
.menu-text {
  display: inline-block;
  //margin-right: 8px;
  margin-left: 8px;
}

.expand-icon {
  position: relative;
  left: 16px;
  font-size: 12px;
}
</style>
