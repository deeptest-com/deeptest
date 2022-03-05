<template>
  <div class="parameters-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">查询参数</a-col>
        <a-col flex="80px" class="dp-right">
          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>帮助</template>
            <QuestionCircleOutlined class="dp-icon-btn"/>
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>全部清除</template>
            <DeleteOutlined class="dp-icon-btn"/>
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>新增</template>
            <PlusOutlined class="dp-icon-btn"/>
          </a-tooltip>
        </a-col>
      </a-row>
    </div>
    <div class="params">
      <a-row v-for="(item, idx) in modelData.children" :key="idx" type="flex" class="param">
        <a-col flex="1">
          <a-input v-model:value="item.key" @change="onParamChange(idx)" class="dp-bg-input-transparent" />
        </a-col>
        <a-col flex="1">
          <a-input v-model:value="item.val" @change="onParamChange(idx)" class="dp-bg-input-transparent" />
        </a-col>
        <a-col flex="80px" class="dp-right dp-icon-btn-container">
          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>禁用</template>
            <CheckCircleOutlined class="dp-icon-btn" />
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>启用</template>
            <CloseCircleOutlined class="dp-icon-btn" />
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>移除</template>
            <DeleteOutlined class="dp-icon-btn"/>
          </a-tooltip>
        </a-col>
      </a-row>
    </div>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined, DeleteOutlined, PlusOutlined, CheckCircleOutlined, CloseCircleOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";

interface RequestParametersSetupData {
  modelData: ComputedRef;

  onParamChange: (idx) => void;
  doSomething: (e) => void;
}

export default defineComponent({
  name: 'RequestParameters',
  components: {
    QuestionCircleOutlined, DeleteOutlined, PlusOutlined, CheckCircleOutlined, CloseCircleOutlined,
  },
  setup(props): RequestParametersSetupData {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const modelData = computed<any>(() => store.state.Interface.modelResult);

    const onParamChange = (idx) => {
      console.log('onParamChange', idx)
      if (modelData.value.children.length <= idx + 1
            && (modelData.value.children[idx].key !== '' || modelData.value.children[idx].val !== '')) {
        modelData.value.children.push({})
      }
    };

    const doSomething = (e) => {
      console.log('doSomething', e)
    };

    return {
      modelData,
      onParamChange,
      doSomething,
    }
  }
})

</script>

<style lang="less" scoped>
.parameters-main {
  height: 100%;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .params {
    height: calc(100% - 28px);
    overflow-y: auto;
    .param {
      padding: 2px 3px;
      border-bottom: 1px solid #d9d9d9;

      .ant-col {
        border-right: 1px solid #d9d9d9;

        input {
          margin-top: 1px;
        }
      }
    }
  }

}

</style>