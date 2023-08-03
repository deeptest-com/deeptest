<template>
  <a-dropdown>
    <slot name="button"/>
    <template #overlay>
      <a-menu>
        <template v-for="menu in DESIGN_MENU_CONFIG">
          <TreeSubMenu :key="menu.key" :menu="menu" v-if="menu?.children?.length"/>
          <a-menu-item v-else :key="menu.key">
            <template #icon>
              <IconSvg :type="menu.icon" class="icon"/>
            </template>
            <span class="menu-text">
                {{ menu.title }}
            </span>
          </a-menu-item>
        </template>
      </a-menu>
    </template>
  </a-dropdown>
</template>

<script setup lang="ts">
import {defineProps, createVNode} from "vue";
import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';
import {CaretRightOutlined} from "@ant-design/icons-vue";
import {DESIGN_MENU_CONFIG} from "./config";
import IconSvg from "@/components/IconSvg";
import TreeSubMenu from "./TreeSubMenu.vue";


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
