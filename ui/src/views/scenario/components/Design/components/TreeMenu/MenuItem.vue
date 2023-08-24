<template>
  <a-menu-item :key="menu.key">
    <template #icon>
      <IconSvg v-if="menu.icon" :type="menu.icon" class="icon"/>
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

console.log(props.menu);

</script>

<style lang="less" scoped>
.menu-text {
  margin-left: 8px;

  &.menu-text-no-margin {
    margin: 0;
  }
}
</style>
