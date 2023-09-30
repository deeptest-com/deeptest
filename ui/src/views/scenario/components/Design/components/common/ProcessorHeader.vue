<!-- ::::场景执行器详情：header 头部信息  -->
<template>
  <div class="header-con">
    <div class="left">
      <span class="title">
        <IconSvg :type="icon" class="prefix-icon-svg"/>
        {{ scenarioType }}
      </span>

      <div class="name">
        <EditAndShow placeholder="修改名称"
                     :autoFocus="false"
                     :key="nodeData?.id"
                     :value="nodeData?.name || ''"
                     :canEmpty="true"
                     :emptyValue="'未命名'"
                     :customClass="nodeData?.name? '':'defaultValue'"
                     @update="updateTitle"/>
      </div>
    </div>

     <div class="right" v-if="showRight">
       <span v-if="scenarioTypeBindText">
          <IconSvg :type="'arrange-link'" class="prefix-icon-svg"/>
          {{scenarioTypeBindText}}：<a href="javascript:void (0)">{{ linkedInterfaceName }}</a>
       </span>
       <Tips v-else :section="nodeData.processorType" :title="nodeData.processorType" />
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, watch,onMounted,ref} from "vue";
import {useStore} from "vuex";
import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Scenario} from "@/views/scenario/store";
import EditAndShow from "@/components/EditAndShow/index.vue";
import Tips from "@/components/Tips/index.vue";
import IconSvg from "@/components/IconSvg";
import {DESIGN_TYPE_ICON_MAP, scenarioTypeMapToText,scenarioTypeMapToBindText} from "../../config";
import {notifyError, notifySuccess} from "@/utils/notify";

const store = useStore<{ Debug: Debug, Scenario: Scenario }>();
const nodeData: any = computed<boolean>(() => store.state.Scenario.nodeData);

// 获取源接口名称
const linkedInterfaceName = computed(() => {
  return nodeData?.value?.srcName || nodeData?.value?.name;
})

const showRight = computed(() => {
  return true
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
const scenarioTypeBindText = computed(() => {
  const processorInterfaceSrc = nodeData.value?.processorInterfaceSrc;
  return scenarioTypeMapToBindText[processorInterfaceSrc]
});
// 更新标题
async function updateTitle(title) {
  store.dispatch('Scenario/saveProcessor', {
    ...nodeData.value,
    name: title,
  }).then((res) => {
    if (res === true) {
      notifySuccess('修改场景步骤名称成功');
    } else {
      notifyError('修改场景步骤名称失败');
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
    .defaultValue {
      color: #c1baba;
    }
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

