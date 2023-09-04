<template>
  <ContentPane>
    <template #left>
      <DiagnoseInterfaceTree/>
    </template>
    <template #right>
      <DiagnoseInterfaceDesign/>
    </template>
  </ContentPane>
</template>

<script setup lang="ts">
import DiagnoseInterfaceTree from './components/tree.vue';
import DiagnoseInterfaceDesign from './design/index.vue';
import ContentPane from '@/views/component/ContentPane/index.vue';
</script>

