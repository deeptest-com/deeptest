<!-- ::::场景执行器详情：header 头部信息  -->
<template>
  <div class="header-con">
    <span class="title"><IconSvg :type="icon" class="prefix-icon-svg"/> {{ scenarioType }}</span>
    <div class="name">
      <EditAndShow placeholder="修改标题"
                   :value="nodeData?.name || ''"
                   @update="updateTitle"/>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, inject, ref} from "vue";
import {useStore} from "vuex";
import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Scenario} from "@/views/scenario/store";
import EditAndShow from "@/components/EditAndShow/index.vue";
import IconSvg from "@/components/IconSvg";
import {DESIGN_TYPE_ICON_MAP,scenarioTypeMapToText} from "../../config";
import {message} from "ant-design-vue";

const store = useStore<{ Debug: Debug, Scenario: Scenario }>();
const nodeData: any = computed<boolean>(() => store.state.Scenario.nodeData);

const icon = computed(() => {
  return DESIGN_TYPE_ICON_MAP[nodeData.value?.processorType] || '未定义';
});
const scenarioType = computed(() => {
  return scenarioTypeMapToText[nodeData.value?.processorType] || '未定义';
});

// 更新标题
async function updateTitle(title) {
  debugger;
  store.dispatch('Scenario/saveProcessor', {
    ...nodeData.value,
    name:title,
  }).then((res) => {
    if (res === true) {
      message.success('保存成功');
    } else {
      message.error('保存失败');
    }
  })
}


</script>

<style lang="less" scoped>
.header-con {
  display: flex;
  height: 40px;
  align-items: center;
  justify-content: flex-start;
  border-bottom: 1px solid rgba(0, 0, 0, 0.06);
  //margin-bottom: 16px;
  padding-left: 16px;

  .title {
    font-size: 16px;
    font-weight: bold;
    margin-right: 16px;
  }

  .name {
    width: 300px;
    display: inline-block;
  }

  border-bottom: 1px solid rgba(0, 0, 0, 0.06);
}

</style>

