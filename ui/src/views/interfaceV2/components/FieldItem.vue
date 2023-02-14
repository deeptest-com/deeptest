<template>
  <div style="margin-bottom: 16px;">
    <a-input v-model:value="fieldState.name"
             @change="handleChangeName"
             style="width: 300px"
             placeholder="path params">
      <template #addonAfter>
        <a-select
            v-model:value="fieldState.type"
            :options="pathParamsDataTypesOpts"
            style="width: 100px"></a-select>
      </template>
    </a-input>

    <a-input v-model:value="fieldState.desc"
             placeholder="path params description"
             style="width: 300px">
      <template #addonAfter>
        <a-space :size="8">
          <a-popconfirm placement="topLeft" ok-text="Yes" cancel-text="No" @confirm="confirm">
            <template #title>
              设置 ref
            </template>
            <a-tooltip placement="topLeft" arrow-point-at-center title="Reference">
              <LinkOutlined @click="setRef"/>
            </a-tooltip>
          </a-popconfirm>
          <a-tooltip placement="topLeft" arrow-point-at-center title="其他属性设置">
            <setting-outlined @click="settingOther"/>
          </a-tooltip>
          <a-tooltip placement="topLeft" arrow-point-at-center title="是否必填">
            <InfoCircleOutlined v-if="fieldState.required" @click="setRequire"/>
            <InfoCircleTwoTone v-if="!fieldState.required" @click="setRequire"/>
          </a-tooltip>
          <a-tooltip placement="topLeft" arrow-point-at-center title="删除">
            <DeleteOutlined @click="del(fieldState)"/>
          </a-tooltip>
        </a-space>
      </template>
    </a-input>

  </div>
</template>

<script lang="ts" setup>
import {ValidateErrorEntity} from 'ant-design-vue/es/form/interface';
import {
  defineComponent,
  reactive,
  ref,
  toRaw,
  UnwrapRef,
  defineProps,
  defineEmits,
  watch,
  onUnmounted,
  onMounted
} from 'vue';
import {requestMethodOpts, pathParamsDataTypesOpts} from '@/config/constant';
import {PlusOutlined, EditOutlined} from '@ant-design/icons-vue';
import contenteditable from 'vue-contenteditable'

import {
  SettingOutlined,
  DeleteOutlined,
  InfoCircleOutlined,
  LinkOutlined,
  InfoCircleTwoTone
} from '@ant-design/icons-vue';


const props = defineProps(['fieldData'])


const emit = defineEmits(['del', 'setRequire', 'setRef', 'settingOther', 'paramsNameChange']);

const fieldState = ref({});

onMounted(() => {
  fieldState.value = props.fieldData;
})


function handleChangeName(e) {
  emit('paramsNameChange',e.target.value);
}

/**
 * 设置该属性为必填项
 * */
function setRequire() {
  fieldState.value.required = !fieldState.value.required;
  emit('setRequire');
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
function del(data) {
  emit('del', data);
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
