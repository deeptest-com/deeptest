<template>
  <div class="processor_interface_default-main">
    <DesignInterface />
  </div>
</template>

<script setup lang="ts">
import {computed, reactive, ref} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {useI18n} from "vue-i18n";
import DesignInterface from '@/views/interface/components/designer/Interface.vue';
import {StateType as ScenarioStateType} from "../../../../store";

const router = useRouter();
const {t} = useI18n();

const store = useStore<{ Scenario: ScenarioStateType; }>();
const modelRef = computed<boolean>(() => store.state.Scenario.nodeData);

</script>

<style lang="less" scoped>
.processor_interface_default-main {
  height: 100%;
}
</style>
