<template>
  <a-dropdown>
    <slot name="button"/>
    <template #overlay>
      <a-menu>
        <template v-for="menu in DESIGN_MENU_CONFIG">
          <SubMenu :key="menu.key" :menu="menu" v-if="menu?.children?.length"/>
          <a-menu-divider v-else-if="menu?.key === 'divider'" :key="menu.key"/>
          <MenuItem v-else :menu="menu" :key="menu.key"/>
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
import SubMenu from "./SubMenu.vue";
import MenuItem from "./MenuItem.vue";
</script>

