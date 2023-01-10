<template>
  <div class="processor_interface_default-main">
    <DesignInterface :usedBy="designByScenario" />
  </div>
</template>

<script setup lang="ts">
import {useRouter} from "vue-router";
import {useI18n} from "vue-i18n";
import DesignInterface from '@/views/interface/components/designer/Interface.vue';
import {UsedBy} from "@/utils/enum";

const router = useRouter();
const {t} = useI18n();

const designByScenario = UsedBy.scenario

</script>

<style lang="less" scoped>
.processor_interface_default-main {
  height: 100%;
}
</style>
