<!-- ::::场景执行器详情：header 头部信息  -->
<template>
  <div class="header-con">
    <div class="left">
      <span class="title"><IconSvg :type="icon" class="prefix-icon-svg"/> {{ scenarioType }}</span>
      <div class="name">
        <EditAndShow placeholder="修改名称"
                     :value="nodeData?.name || ''"
                     @update="updateTitle"/>
      </div>
    </div>
    <div class="right" v-if="showRight">
      <IconSvg :type="'arrange-link'" class="prefix-icon-svg"/>
      绑定接口：<a href="javascript:void (0)">{{ linkedInterfaceName }}</a>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, watch} from "vue";
import {useStore} from "vuex";
import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Scenario} from "@/views/scenario/store";
import EditAndShow from "@/components/EditAndShow/index.vue";
import IconSvg from "@/components/IconSvg";
import {DESIGN_TYPE_ICON_MAP, scenarioTypeMapToText} from "../../config";
import {message} from "ant-design-vue";

const store = useStore<{ Debug: Debug, Scenario: Scenario }>();
const nodeData: any = computed<boolean>(() => store.state.Scenario.nodeData);


const linkedInterfaceName = computed(() => {
  return nodeData?.value?.srcName || nodeData?.value?.name;
})


const showRight = computed(() => {
  return nodeData.value?.processorType === 'processor_interface_default';
})

const icon = computed(() => {
  const processorInterfaceSrc = nodeData.value?.processorInterfaceSrc;
  if (processorInterfaceSrc) {
    return DESIGN_TYPE_ICON_MAP[processorInterfaceSrc] || 'interface';
  }
  return DESIGN_TYPE_ICON_MAP[nodeData?.value?.processorType] || 'interface';
});

const scenarioType = computed(() => {
  const processorInterfaceSrc = nodeData.value?.processorInterfaceSrc;
  if (processorInterfaceSrc) {
    return scenarioTypeMapToText[processorInterfaceSrc] || '接口定义';
  }
  return scenarioTypeMapToText[nodeData.value?.processorType] || '接口定义';
});

// 更新标题
async function updateTitle(title) {
  store.dispatch('Scenario/saveProcessor', {
    ...nodeData.value,
    name: title,
  }).then((res) => {
    if (res === true) {
      message.success('修改场景名称成功');
    } else {
      message.error('修改场景名称失败');
    }
  })
}


</script>

<style lang="less" scoped>
.header-con {
  display: flex;
  height: 40px;
  align-items: center;
  justify-content: space-between;
  border-bottom: 1px solid rgba(0, 0, 0, 0.06);
  //margin-bottom: 16px;
  padding-left: 16px;

  .title {
    font-size: 16px;
    font-weight: bold;
    margin-right: 16px;
  }

  .left{
    width: 400px;
  }

  .name {
    width: 200px;
    display: inline-block;
  }

  .right {
    text-align: right;
    margin-right: 16px;
    flex:1;
    max-width: 400px;
  //  超出省略号
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  border-bottom: 1px solid rgba(0, 0, 0, 0.06);

  .prefix-icon-svg {
    margin-right: 3px;
  }
}

</style>

