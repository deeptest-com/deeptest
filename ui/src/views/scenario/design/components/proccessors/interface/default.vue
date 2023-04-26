<template>
  <div class="processor_interface_default-main">
<!--    <DebugComp />-->
  </div>
</template>

<script setup lang="ts">
import {useRouter} from "vue-router";
import {useI18n} from "vue-i18n";

import DebugComp from '@/views/component/debug/index.vue';

import {provide} from "vue";
import {UsedBy} from "@/utils/enum";

provide('usedBy', UsedBy.InterfaceDebug)

const router = useRouter();
const {t} = useI18n();

</script>

<style lang="less" scoped>
.processor_interface_default-main {
  height: 100%;
}
</style>
