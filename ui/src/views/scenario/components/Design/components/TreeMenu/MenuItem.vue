<template>
  <a-menu-item :key="menu.key">
    <template #icon>
      <IconSvg :type="menu.icon" class="icon"/>
    </template>
    <span class="menu-text">
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
  margin-left: 8px;
}
</style>
