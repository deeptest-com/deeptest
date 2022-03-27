<template>
  <div class="headers-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">响应头</a-col>
        <a-col flex="80px" class="dp-right">
          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>复制</template>
            <CopyOutlined class="dp-icon-btn dp-trans-80" />
          </a-tooltip>
        </a-col>
      </a-row>
    </div>
    <div class="params">
      <a-row v-for="(item, idx) in requestData.children" :key="idx" type="flex" class="param">
        <a-col flex="1">
          <a-input v-model:value="item.key" @change="onParamChange(idx)" class="dp-bg-input-transparent" />
        </a-col>
        <a-col flex="1">
          <a-input v-model:value="item.val" @change="onParamChange(idx)" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>
    </div>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {
  QuestionCircleOutlined,
  DeleteOutlined,
  PlusOutlined,
  CheckCircleOutlined,
  CloseCircleOutlined,
  CopyOutlined
} from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";

interface ResponseHeadersSetupData {
  requestData: ComputedRef;

  onParamChange: (idx) => void;
  doSomething: (e) => void;
}

export default defineComponent({
  name: 'ResponseHeaders',
  components: {
    CopyOutlined,
  },
  setup(props): ResponseHeadersSetupData {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const requestData = computed<any>(() => store.state.Interface.requestData);

    const onParamChange = (idx) => {
      console.log('onParamChange', idx)
      if (requestData.value.children.length <= idx + 1
          && (requestData.value.children[idx].key !== '' || requestData.value.children[idx].val !== '')) {
        requestData.value.children.push({})
      }
    };

    const doSomething = (e) => {
      console.log('doSomething', e)
    };

    return {
      requestData,
      onParamChange,
      doSomething,
    }
  }
})

</script>

<style lang="less" scoped>
.headers-main {
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