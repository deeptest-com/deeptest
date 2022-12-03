<template>
  <div class="formdata-main">
    <div class="dp-param-grid">
      <div class="params">
        <a-row v-for="(item, idx) in interfaceData.bodyFormData" :key="idx" type="flex" class="param">
          <a-col flex="1">
            <a-input v-model:value="item.name" @change="onParamChange(idx)" class="dp-bg-input-transparent" />
          </a-col>
          <a-col width="60px">
            <a-select
                v-model:value="item.type"
                @change="onParamChange(idx)"
                :bordered="false"
            >
              <a-select-option value="text">Text</a-select-option>
              <a-select-option value="file">File</a-select-option>
            </a-select>
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
  </div>
</template>

<script setup lang="ts">
import {computed, ComputedRef, defineComponent, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined, DeleteOutlined, PlusOutlined, CheckCircleOutlined, CloseCircleOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import {Param, Interface, BodyFormDataItem} from "@/views/interface/data";

const {t} = useI18n();
const store = useStore<{ Interface: StateType }>();
const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

const onFormDataChange = (idx) => {
  console.log('onFormDataChange', idx)
  if (interfaceData.value.bodyFormData.length <= idx + 1
      && (interfaceData.value.bodyFormData[idx].name !== '' || interfaceData.value.bodyFormData[idx].value !== '')) {
    interfaceData.value.bodyFormData.push({type: 'text'} as BodyFormDataItem)
  }
}

const add = () => {
  console.log('add')
  interfaceData.value.bodyFormData.push({type: 'text'} as BodyFormDataItem)
}
const removeAll = () => {
  console.log('removeAll', interfaceData.value.bodyFormData)
  interfaceData.value.bodyFormData = [{type: 'text'} as BodyFormDataItem]
}

const disable = (idx) => {
  console.log('enable', idx)
  interfaceData.value.bodyFormData[idx].disabled = !interfaceData.value.bodyFormData[idx].disabled
}
const remove = (idx) => {
  console.log('remove')
  interfaceData.value.bodyFormData.splice(idx, 1)
  add()
}
const insert = (idx) => {
  console.log('insert')
  interfaceData.value.bodyFormData.splice(idx+1, 0, {} as BodyFormDataItem)
}

</script>

<style lang="less" scoped>
.formdata-main {
  height: 100%;
}

</style>