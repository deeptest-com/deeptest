<!-- ::::参数设置器-->
<template>
  <div class="main">
    <a-input :value="fieldState.name"
             @change="handleChangeName"
             style="width: 300px"
             :disabled="hasRef"
             placeholder="输入字段名称">
      <template #addonAfter>
        <a-select
            :value="fieldState.type"
            placeholder="请选择类型"
            @change="handleTypeChange"
            :disabled="hasRef"
            :options="pathParamsDataTypesOpts"
            style="width: 100px"/>
      </template>
    </a-input>
    <a-input :value="fieldState.description"
             placeholder="输入描述信息"
             :disabled="hasRef"
             @change="handleDescChange"
             style="width: 300px">
      <template #addonAfter>
        <a-space :size="8">
          <!-- ::::关联组件 -->
          <a-popover v-if="!hideRef" v-model:visible="showRef" :title="null" trigger="click">
            <template #content>
              <div class="ref-content">
                <a-form :layout="'vertical'">
                  <a-form-item label="关联组件">
                    <a-select
                        :value="fieldState.ref"
                        style="width: 300px"
                        :allowClear="true"
                        placeholder="请选择需要关联的组件"
                        @change="handleRefChange"
                        :options="refsOptions?.[fieldState.type] || []"
                    ></a-select>
                  </a-form-item>
                </a-form>
              </div>
            </template>
            <a-tooltip v-if="!hideRef" placement="topLeft" arrow-point-at-center title="关联组件">
              <LinkOutlined @click="setRef" :class="{'ref-active':hasRef}"/>
            </a-tooltip>
          </a-popover>
          <!-- ::::是否必填 -->
          <a-tooltip v-if="!hideRequire" placement="top" arrow-point-at-center :title="hasRef ? '引用类型不支持修改' : '是否必填'">
            <InfoCircleOutlined :class="{'disabled':hasRef}" v-if="!fieldState.required" @click="setRequire"/>
            <InfoCircleTwoTone  :class="{'disabled':hasRef}" v-if="fieldState.required" @click="setRequire"/>
          </a-tooltip>
          <a-popover v-if="!hideOther" v-model:visible="showOther" placement="topLeft" :title="null" trigger="click">
            <template #content>
              <div class="other-props-content">
                <a-form :layout='"vertical"' v-if="otherProps?.value === fieldState.type">
                  <div class="card-title">Other {{ otherProps.props.label }}</div>
                  <a-row type="flex" :gutter="24">
                    <a-col :span="12" class="card-content" v-for="opt in otherProps.props.options" :key="opt.name">
                      <a-form-item
                          class="card-content-form-item"
                          :labelAlign="'right'"
                          :label="opt.label">
                        <a-select
                            v-if="opt.component === 'selectTag'"
                            :value="fieldState[opt.name] || []"
                            mode="tags"
                            :placeholder="opt.placeholder"
                            :disabled="hasRef"
                            @change="(val) => {
                            handleOthersPropsChange(opt.name, val);
                          }"/>
                        <a-select
                            v-if="opt.component === 'select'"
                            :value="fieldState[opt.name] || null"
                            :options="opt.options"
                            :disabled="hasRef"
                            :placeholder="opt.placeholder"
                            @change="(val) => {
                            handleOthersPropsChange(opt.name,val);
                          }"
                        />
                        <a-input
                            v-if="opt.component === 'input'"
                            :value="fieldState[opt.name] || null"
                            :disabled="hasRef"
                            @change="(e) => {
                            handleOthersPropsChange(opt.name, e.target.value);
                          }"
                            :placeholder="opt.placeholder"/>
                        <a-input-number
                            v-if="opt.component === 'inputNumber'"
                            :value="fieldState[opt.name] || null"
                            :disabled="hasRef"
                            class="input-number"
                            @change="(val) => {
                            handleOthersPropsChange(opt.name,val);
                          }"
                            :placeholder="opt.placeholder"/>
                        <a-switch
                            v-if="opt.component === 'switch'"
                            :checked="fieldState[opt.name] || false"
                            :disabled="hasRef"
                            @change="(val) => {
                            handleOthersPropsChange(opt.name,val);
                          }"/>
                      </a-form-item>
                    </a-col>
                  </a-row>
                </a-form>
              </div>
            </template>
            <a-tooltip v-if="!hideOther" placement="topLeft" arrow-point-at-center title="设置其他属性">
              <SettingOutlined @click="setOtherProps"/>
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
  watch, computed,
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
import {useStore} from "vuex";
import {Endpoint} from "@/views/endpoint/data";
import {message} from "ant-design-vue";

interface fieldStateType {
  name: string;
  required: boolean;
  type: string;
  description: string | null
  ref: string | null;
  format: string,
  default: string,
  example: string
  pattern: string,
  minLength: number
  maxLength: number
}
const props = defineProps(['fieldData', 'hideRequire', 'hideRef', 'hideOther', 'hideDel']);
const emit = defineEmits(['del', 'setRef', 'settingOther', 'change']);

const store = useStore<{ Endpoint, ProjectGlobal, ServeGlobal }>();
const refsOptions: any = computed<Endpoint>(() => store.state.Endpoint.refsOptions);
const currServe = computed<any>(() => store.state.ServeGlobal.currServe);

const fieldState = ref<fieldStateType | any>({
  name: '',
  required: false,
  type: 'string',
  description: null,
  ref: '',
  format: '',
  default: '',
  example: '',
  pattern: '',
});

const showRef = ref(false);
const otherProps: any = ref({});
const showOther = ref(false);

const hasRef = computed(() => {
  return !!fieldState?.value?.ref;
});

function handleChangeName(e: any) {
  const name = e.target.value;
  const reg = /^[\w-]*$/;
  if (!reg.test(name)) {
    return;
  }
  fieldState.value.name = name;
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
  fieldState.value.ref = val;
  const selectedRefDetail =  refsOptions.value?.[fieldState.value.type]?.find((item: any) => {
      return item.ref === val;
  });
  const content = JSON.parse(selectedRefDetail.content);
  fieldState.value.name = selectedRefDetail.name;
  fieldState.value.description = selectedRefDetail.description || '';
  fieldState.value = {
    ...fieldState.value,
    ...content
  }
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
  if(hasRef.value) {
    return;
  }
  fieldState.value.required = !fieldState.value.required;
  emit('change', fieldState.value);
}

function setOtherProps() {
  if(hasRef.value) return;
  showOther.value = true;
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
async function setRef(data: any) {
  showRef.value = true;
  await store.dispatch('Endpoint/getRefsOptions', {
    serveId: currServe.value.id,
    type: fieldState.value.type
  });
}

watch(() => {
  return props.fieldData
}, (newVal) => {
  if (newVal) { // && newVal.type) {
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
.main {
  margin: 16px auto;
}

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

.card-content-form-item {
  margin-bottom: 12px;
}

.input-number {
  width: 100%;
}


:deep(.ref-active svg){
  color: #0000cc;
}
.disabled{
  cursor: not-allowed;
}
</style>
