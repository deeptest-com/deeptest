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
            :value="fieldState.type"
            placeholder="请选择类型"
            @change="handleTypeChange"
            :options="pathParamsDataTypesOpts"
            style="width: 100px"/>
      </template>
    </a-input>
    <a-input :value="fieldState.description"
             placeholder="输入描述信息"
             @change="handleDescChange"
             style="width: 300px">
      <template #addonAfter>
        <a-space :size="8">
          <a-popover v-if="!hideRef" v-model:visible="showRef" :title="null" trigger="click">
            <template #content>
              <div class="ref-content">
                <a-form-item label="关联组件">
                  <a-select
                      :value="fieldState.$ref"
                      style="width: 200px"
                      :allowClear="true"
                      placeholder="请选择需要关联的组件"
                      @change="handleRefChange"
                      :options="refsOptions"
                  ></a-select>
                </a-form-item>
              </div>
            </template>
            <a-tooltip v-if="!hideRef" placement="topLeft" arrow-point-at-center title="关联组件">
              <LinkOutlined @click="setRef"/>
            </a-tooltip>
          </a-popover>
          <a-tooltip v-if="!hideRequire" placement="topLeft" arrow-point-at-center title="是否必填">
            <InfoCircleOutlined v-if="!fieldState.required" @click="setRequire"/>
            <InfoCircleTwoTone v-if="fieldState.required" @click="setRequire"/>
          </a-tooltip>
          <a-popover v-if="!hideOther" v-model:visible="showOther" :title="null" trigger="click">
            <template #content>
              <div class="other-props-content">
                <a-form v-if="otherProps?.value === fieldState.type">
                  <div class="card-title">Other {{ otherProps.props.label }}</div>
                  <div class="card-content" v-for="opt in otherProps.props.options" :key="opt.name">
                    <a-form-item
                        class="card-content-form-item"
                        :labelAlign="'right'"
                        :labelCol="{ span: 6 }"
                        :wrapperCol="{ span: 16 }"
                        :label="opt.label">
                      <a-select
                          v-if="opt.component === 'selectTag'"
                          :value="fieldState[opt.name] || []"
                          mode="tags"
                          :placeholder="opt.placeholder"
                          @change="(val) => {
                            handleOthersPropsChange(opt.name, val);
                          }"/>
                      <a-select
                          v-if="opt.component === 'select'"
                          :value="fieldState[opt.name] || null"
                          :options="opt.options"
                          :placeholder="opt.placeholder"
                          @change="(val) => {
                            handleOthersPropsChange(opt.name,val);
                          }"
                      />
                      <a-input
                          v-if="opt.component === 'input'"
                          :value="fieldState[opt.name] || null"
                          @change="(e) => {
                            handleOthersPropsChange(opt.name, e.target.value);
                          }"
                          :placeholder="opt.placeholder"/>
                      <a-input-number
                          v-if="opt.component === 'inputNumber'"
                          :value="fieldState[opt.name] || null"
                          id="inputNumber"
                          @change="(val) => {
                            handleOthersPropsChange(opt.name,val);
                          }"
                          :placeholder="opt.placeholder"/>
                      <a-switch
                          v-if="opt.component === 'switch'"
                          :checked="fieldState[opt.name] || false"
                          @change="(val) => {
                            handleOthersPropsChange(opt.name,val);
                          }"/>
                    </a-form-item>
                  </div>
                </a-form>
              </div>
            </template>
            <a-tooltip v-if="!hideOther" placement="topLeft" arrow-point-at-center title="设置其他属性">
              <SettingOutlined @click="showOther = true"/>
            </a-tooltip>
          </a-popover>
          <a-tooltip v-if="!hideDel" placement="topLeft" arrow-point-at-center title="删除">
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
  onMounted,
  watch
} from 'vue';

import {pathParamsDataTypesOpts, paramsSchemaDataTypes} from '@/config/constant';

import {
  SettingOutlined,
  DeleteOutlined,
  InfoCircleOutlined,
  LinkOutlined,
  InfoCircleTwoTone
} from '@ant-design/icons-vue';
import {cloneByJSON} from "@/utils/object";

const props = defineProps(['fieldData', 'hideRequire', 'hideRef', 'hideOther', 'hideDel', 'refsOptions']);

const emit = defineEmits(['del', 'setRef', 'settingOther', 'change']);

interface fieldStateType {
  name: string;
  required: boolean;
  type: string;
  description: string | null
  $ref: string | null;
  format: string,
  default: string,
  example: string
  pattern: string,
  minLength: number
  maxLength: number
}

const fieldState = ref<fieldStateType | any>({
  name: '',
  required: false,
  type: 'string',
  description: null,
  $ref: null,
  format: '',
  default: '',
  example: '',
  pattern: '',
});

const showRef = ref(false);
const otherProps: any = ref({});

const showOther = ref(false);


function handleChangeName(e: any) {
  emit('change', fieldState.value);
}

function handleTypeChange(val: any) {
  fieldState.value.type = val;
  emit('change', fieldState.value);
}

function handleDescChange(e: any) {
  fieldState.value.description = e.target.value;
  emit('change', fieldState.value);
}


function handleRefChange(val: any) {
  fieldState.value.$ref = val;
  emit('change', fieldState.value);
}

function handleOthersPropsChange(key: string, val: any) {
  fieldState.value[key] = val;
  emit('change', fieldState.value);
}


/**
 * 设置该属性为必填项
 * */
function setRequire() {
  fieldState.value.required = !fieldState.value.required;
  emit('change', fieldState.value);
}

/**
 * 设置属性格式
 * */
function settingOther(data: any) {
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
function setRef(data: any) {
  showRef.value = true;
  // emit('setRef', data);
}

watch(() => {
  return props.fieldData
}, (newVal) => {
  if (newVal && newVal.type) {
    fieldState.value = newVal;
  }
}, {
  immediate: true
})


watch(() => {
  return fieldState?.value?.type
}, (newVal: string) => {
  if (newVal) {
    otherProps.value = cloneByJSON(paramsSchemaDataTypes)[newVal];
  }
}, {
  immediate: true
})

</script>

<style lang="less" scoped>
.requireActived {
  color: #0000cc;
}

.ref-content {
  margin: 8px;
}

.other-props-content {
  width: 600px;

  .card-title {
    font-weight: bold;
    margin-bottom: 16px;
  }
}

</style>
