<!-- ::::
  参数设置器
 -->
<template>
  <div style="margin-bottom: 16px;">
    <a-input v-model:value="fieldState.name"
             @change="handleChangeName"
             style="width: 300px"
             placeholder="输入字段名称">
      <template #addonAfter>
        <a-select
            v-model:value="fieldState.type"
            :options="pathParamsDataTypesOpts"
            style="width: 100px"></a-select>
      </template>
    </a-input>
    <a-input v-model:value="fieldState.desc"
             placeholder="输入描述信息"
             style="width: 300px">
      <template #addonAfter>
        <a-space :size="8">
          <a-tooltip v-if="showRequire" placement="topLeft" arrow-point-at-center title="是否必填">
            <InfoCircleOutlined v-if="fieldState.required" @click="setRequire"/>
            <InfoCircleTwoTone v-if="!fieldState.required" @click="setRequire"/>
          </a-tooltip>
          <a-tooltip placement="topLeft" arrow-point-at-center title="删除">
            <DeleteOutlined @click="del"/>
          </a-tooltip>
        </a-space>
      </template>
    </a-input>
  </div>
</template>

<script lang="ts" setup>
import {
  ref,
  defineProps,
  defineEmits,
  onMounted, watch
} from 'vue';

import {requestMethodOpts, pathParamsDataTypesOpts} from '@/config/constant';

import {
  SettingOutlined,
  DeleteOutlined,
  InfoCircleOutlined,
  LinkOutlined,
  InfoCircleTwoTone
} from '@ant-design/icons-vue';

const props = defineProps(['fieldData','showRequire'])

const emit = defineEmits(['del', 'setRef', 'settingOther', 'change']);

interface fieldStateType {
  name: string;
  required: boolean;
  type: string;
  description: string;
}

const fieldState = ref<fieldStateType | any>({});

watch(() => {
  return props.fieldData
}, (newVal) => {
  fieldState.value = newVal;
}, {
  immediate: true
})

function handleChangeName(e) {
  emit('change',fieldState.value);
}

/**
 * 设置该属性为必填项
 * */
function setRequire() {
  fieldState.value.required = !fieldState.value.required;
  emit('change',fieldState.value);
}

/**
 * 设置属性格式
 * */
function settingOther(data) {
  emit('settingOther', data);
}

/**
 * 删除
 */
function del() {
  emit('del', fieldState.value);
}

/**
 * ref
 * */
function setRef(data) {
  emit('setRef', data);
}

</script>

<style lang="less" scoped>
.requireActived {
  color: #0000cc;
}
</style>
