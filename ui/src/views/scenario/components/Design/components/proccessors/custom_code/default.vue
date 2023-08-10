<template>
  <div class="processor_custom_code-main dp-proccessors-container">
    <ProcessorHeader/>
    <div class="header">
      <a-row type="flex" class="row">
        <a-col flex="1" class="left">
          <icon-svg type="script" class="icon" />&nbsp;
          <span>JavaScript代码</span>
        </a-col>

        <a-col flex="100px" class="dp-right">
          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>帮助</template>
            <QuestionCircleOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>全屏</template>
            <FullscreenOutlined @click.stop="openFullscreen()"  class="dp-icon-btn dp-trans-80" />
          </a-tooltip>

        </a-col>
      </a-row>
    </div>

    <div class="content">
      <ProcessorCustomCodeEdit />
    </div>

    <ProcessorPopup v-if="fullscreen"
                    :visible="fullscreen"
                    :model="modelRef"
                    :onCancel="closeFullScreen" />
  </div>
</template>

<script setup lang="ts">
import {computed, ref, watch, provide} from "vue";
import {useStore} from "vuex";
import {message} from "ant-design-vue";
import { QuestionCircleOutlined, FullscreenOutlined } from '@ant-design/icons-vue';
import IconSvg from "@/components/IconSvg";
import {StateType as ScenarioStateType} from "../../../../../store";
import ProcessorCustomCodeEdit from "./edit.vue";
import ProcessorPopup from "../../common/ProcessorPopup.vue";
import ProcessorHeader from '../../common/ProcessorHeader.vue';
const store = useStore<{ Scenario: ScenarioStateType; }>();
const modelRef: any = computed<boolean>(() => store.state.Scenario.nodeData);

const fullscreen = ref(false)

const openFullscreen = () => {
  console.log('openFullscreen')
  fullscreen.value = true
}
const closeFullScreen = () => {
  console.log('closeFullScreen')
  fullscreen.value = false
}

const save = async () => {
  const res = await store.dispatch('Scenario/saveProcessor', {
    ...modelRef.value,
    content: modelRef.value.content,
  })

  if (res === true) {
    message.success('保存成功');
  } else {
    message.error('保存失败');
  }
}

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
