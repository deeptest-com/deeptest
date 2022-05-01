<template>
  <div class="headers-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">请求头列表</a-col>
        <a-col flex="80px" class="dp-right">
          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>帮助</template>
            <QuestionCircleOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>

          <a-tooltip @click="removeAll" overlayClassName="dp-tip-small">
            <template #title>全部清除</template>
            <DeleteOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>

          <a-tooltip @click="add" overlayClassName="dp-tip-small">
            <template #title>新增</template>
            <PlusOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>
        </a-col>
      </a-row>
    </div>
    <div class="headers">
      <a-row v-for="(item, idx) in interfaceData.headers" :key="idx" type="flex" class="header">
        <a-col flex="1">
          <a-input v-model:value="item.name" @change="onParamChange(idx)" class="dp-bg-input-transparent" />
        </a-col>
        <a-col flex="1">
          <a-input v-model:value="item.value" @change="onParamChange(idx)" class="dp-bg-input-transparent" />
        </a-col>
        <a-col flex="80px" class="dp-right dp-icon-btn-container">
          <a-tooltip v-if="!item.disabled" @click="disable(idx)" overlayClassName="dp-tip-small">
            <template #title>禁用</template>
            <CheckCircleOutlined class="dp-icon-btn dp-trans-80" />
          </a-tooltip>

          <a-tooltip v-if="item.disabled" @click="disable(idx)" overlayClassName="dp-tip-small">
            <template #title>启用</template>
            <CloseCircleOutlined class="dp-icon-btn dp-trans-80 dp-light" />
          </a-tooltip>

          <a-tooltip @click="remove(idx)" overlayClassName="dp-tip-small">
            <template #title>移除</template>
            <DeleteOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>

          <a-tooltip @click="insert(idx)" overlayClassName="dp-tip-small">
            <template #title>插入</template>
            <PlusOutlined class="dp-icon-btn dp-trans-80"/>
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
import {Header, Interface} from "@/views/interface/data";

export default defineComponent({
  name: 'RequestHeaders',
  components: {
    QuestionCircleOutlined, DeleteOutlined, PlusOutlined, CheckCircleOutlined, CloseCircleOutlined,
  },
  setup(props) {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

    const onParamChange = (idx) => {
      console.log('onParamChange', idx)
      if (interfaceData.value.headers.length <= idx + 1
          && (interfaceData.value.headers[idx].name !== '' || interfaceData.value.headers[idx].value !== '')) {
        interfaceData.value.headers.push({} as Header)
      }
    };

    const add = () => {
      console.log('add')
      interfaceData.value.params.push({} as Header)
    }
    const removeAll = () => {
      console.log('removeAll', interfaceData.value.headers)
      interfaceData.value.headers = [{} as Header]
    }

    const disable = (idx) => {
      console.log('enable', idx)
      interfaceData.value.headers[idx].disabled = !interfaceData.value.headers[idx].disabled
    }
    const remove = (idx) => {
      console.log('remove')
      interfaceData.value.headers.splice(idx, 1)
      add()
    }
    const insert = (idx) => {
      console.log('insert')
      interfaceData.value.headers.splice(idx+1, 0, {} as Header)
    }

    return {
      interfaceData,
      onParamChange,
      add,
      removeAll,
      disable,
      remove,
      insert,
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
  .headers {
    height: calc(100% - 28px);
    overflow-y: auto;
    .header {
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