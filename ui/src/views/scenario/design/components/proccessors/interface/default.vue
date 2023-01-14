<template>
  <div class="processor_interface_default-main">
    <DesignInterface />
  </div>
</template>

<script setup lang="ts">
import {useRouter} from "vue-router";
import {useI18n} from "vue-i18n";
import DesignInterface from '@/views/interface/components/designer/Interface.vue';

const router = useRouter();
const {t} = useI18n();

</script>

<style lang="less" scoped>
.processor_interface_default-main {
  height: 100%;
}
</style>
