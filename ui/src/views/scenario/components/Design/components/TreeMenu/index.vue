<template>
  <a-dropdown>
    <slot name="button"/>
    <template #overlay>
      <a-menu @click="selectMenu">
        <template v-for="menu in menus">
          <SubMenu :key="menu.key" :menu="menu" v-if="menu?.children?.length" :disabled="menu.disabled"/>
          <a-menu-divider v-else-if="menu?.key === 'divider'" :key="menu.key"/>
          <MenuItem v-else :menu="menu" :key="menu.key" :disabled="menu.disabled"/>
        </template>
      </a-menu>
    </template>
  </a-dropdown>
</template>

<script setup lang="ts">
import {defineEmits, defineProps, computed, watch} from "vue";
import cloneDeep from "lodash/cloneDeep";

const props = defineProps(['treeNode']);
const emit = defineEmits(['selectMenu']);

function selectMenu(info) {
  emit('selectMenu', info, props.treeNode)
}

import {DESIGN_MENU_CONFIG} from "../../config";
import SubMenu from "./SubMenu.vue";
import MenuItem from "./MenuItem.vue";


/**
 * 根据当前的节点类型，过滤掉不需要的菜单
 * */
const menus = computed(() => {
  const nodeType = props?.treeNode?.entityType;
  const disable = props?.treeNode?.disable;
  if (!nodeType) {
    return [];
  }
  const src = cloneDeep(DESIGN_MENU_CONFIG);
  // 处理菜单的 disabled状态
  // 不再递归，因为最多三层，直接处理了
  src.forEach((menu: any) => {
    if (menu?.enableInNodeTypes) {
      menu.disabled = !menu?.enableInNodeTypes?.includes(nodeType);
    }
    if (menu?.children?.length) {
      menu.children.forEach((subMenu: any) => {
        if (subMenu?.enableInNodeTypes) {
          subMenu.disabled = !subMenu.enableInNodeTypes?.includes(nodeType);
        }
      })
    }
  })
  // 递归过滤
  function filterMenu(menu) {
    // 如果有 showInNodeTypes，且当前节点类型不在 showInNodeTypes 中，则不显示
    if (menu?.showInNodeTypes?.length && !menu?.showInNodeTypes?.includes(nodeType)) {
      return false;
    } else if (menu?.hideInNodeTypes?.includes(nodeType)) {
      return false;
    }
    // 如果有 disabled，且当前节点 disabled 为 true，则不显示
    if (menu.key === 'disable' && disable) {
      return false;
    }
    // 如果是启用状态
    if (menu.key === 'enable' && !disable) {
      return false;
    }

    if (menu?.children?.length) {
      menu.children = menu.children.filter((subMenu: any) => {
        return filterMenu(subMenu);
      })
    }
    return true;
  }

  return src.filter((menu) => {
    return filterMenu(menu);
  })
});

</script>

