<template>
  <a-menu-item :key="menu.key">
    <template v-if="menu.icon" #icon>
      <IconSvg :type="menu.icon" class="icon dp-icon-large"/>
    </template>
    <span :class="['menu-text', !menu.icon && 'menu-text-no-margin' ]">
      {{ menu.title }}
     </span>
  </a-menu-item>
</template>
<script setup lang="ts">
import {defineProps, createVNode} from "vue";
import IconSvg from "@/components/IconSvg";

const props = defineProps(['menu']);

</script>

<style lang="less" scoped>
.menu-text {
  margin-left: 4px;

  &.menu-text-no-margin {
    margin: 0;
  }
}
</style>
