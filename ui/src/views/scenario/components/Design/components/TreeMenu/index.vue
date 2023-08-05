<template>
  <a-dropdown>
    <slot name="button"/>
    <template #overlay>
      <a-menu @click="selectMenu">
        <template v-for="menu in menus">
          <SubMenu :key="menu.key" :menu="menu" v-if="menu?.children?.length"/>
          <a-menu-divider v-else-if="menu?.key === 'divider'" :key="menu.key"/>
          <MenuItem v-else :menu="menu" :key="menu.key"/>
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
  console.log('832 menus', props?.treeNode);
  const nodeType = props?.treeNode?.entityType;

  if(!nodeType) {
    return [];
  }
  const src = cloneDeep(DESIGN_MENU_CONFIG);
  // 递归过滤
  function filterMenu(menu) {
    if(menu?.hideInNodeTypes?.includes(nodeType)) {
      return false;
    }
    if (menu?.children?.length) {
      menu.children = menu.children.filter((subMenu:any) => {
        return filterMenu(subMenu);
      })
    }
    return true;
  }
  return src.filter((menu) => {
    return filterMenu(menu);
  })
});


watch(() => {
  return menus.value;
}, (newVal) => {
  console.log('83222newVal', newVal);
},{
  deep: true,
  immediate: true,
})
</script>

