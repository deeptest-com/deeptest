<template>
  <div class="processor_custom_code-main dp-processors-container">
    <ProcessorHeader/>
    <CustomCodeHeader mode="smart" @update-screen="updateScreen" />
    <div class="content">
      <ProcessorCustomCodeEdit ref="processCodeEdit" />
    </div>

    <ProcessorPopup 
      v-if="fullscreen"
      :visible="fullscreen"
      @update-screen="updateScreen"
      :model="modelRef" />
  </div>
</template>

<script setup lang="ts">
import {computed, ref, provide} from "vue";
import {useStore} from "vuex";
import {StateType as ScenarioStateType} from "../../../../../store";
import ProcessorCustomCodeEdit from "./edit.vue";
import ProcessorPopup from "../../common/ProcessorPopup.vue";
import ProcessorHeader from '../../common/ProcessorHeader.vue';
import CustomCodeHeader from './header.vue';
const store = useStore<{ Scenario: ScenarioStateType; }>();
const modelRef: any = computed<boolean>(() => store.state.Scenario.nodeData);

const fullscreen = ref(false);
const processCodeEdit = ref();

const updateScreen = (e) => {
  console.log(e);
  fullscreen.value = e;
};

provide('fullscreen', computed(() => fullscreen.value));

</script>

<style lang="less" scoped>
.processor_custom_code-main {
  height: 100%;
  display: flex;
  flex-direction: column;

  .header {
    height: 32px;
    padding: 3px 8px;
    border: 1px solid #d9d9d9;
    background-color: #fafafa;
    border-radius: 3px;
  }

  .content {
    flex: 1;
  }
}
</style>
