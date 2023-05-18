<template>
  <div class="scenario-edit-main">
    <a-card :bordered="false">
      <template #title>
        <div>{{modelId > 0 ? '编辑场景' : '新建场景'}}</div>
      </template>

      <template #extra></template>

      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-form-item label="名称" v-bind="validateInfos.name">
            <a-input v-model:value="modelRef.name"
                     @blur="validate('name', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>

          <a-form-item label="描述" v-bind="validateInfos.desc">
            <a-input v-model:value="modelRef.desc"
                     @blur="validate('desc', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>

          <a-form-item v-if="modelId > 0" label="是否禁用">
            <a-switch v-model:checked="modelRef.disabled" />
          </a-form-item>

          <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
            <a-button type="primary" @click.prevent="submitForm">保存</a-button>
            <a-button style="margin-left: 10px" @click="resetFields">重置</a-button>
          </a-form-item>
        </a-form>
      </div>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import {defineComponent, computed, ref, reactive, ComputedRef, defineProps, PropType} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import { useI18n } from "vue-i18n";
import {Form} from 'ant-design-vue';
const useForm = Form.useForm;
import {StateType} from "../store";
import {get} from "@/views/scenario/service";

const router = useRouter();
const { t } = useI18n();

const props = defineProps({
  modelId: {
    type: Number,
    required: true
  },
  categoryId: {
    type: Number,
    required: true
  },
  onFinish: {
    type: Function as PropType<() => void>,
    required: true
  }
})

const rulesRef = reactive({
  name: [
    { required: true, message: '请输入名称', trigger: 'blur' },
  ],
});

const store = useStore<{ Scenario: StateType }>();
const modelRef = ref({} as any)
const { resetFields, validate, validateInfos } = useForm(modelRef, rulesRef);

const getData = (id: number) => {
  if (id === 0) {
    modelRef.value = {}
    return
  }

  get(id).then((json) => {
    if (json.code === 0) {
      modelRef.value = json.data
    }
  })
}
getData(props.modelId)

const submitForm = async() => {
  validate().then(() => {
    console.log(modelRef);
    modelRef.value.categoryId = props.categoryId

    store.dispatch('Scenario/saveScenario', modelRef.value).then((res) => {
      console.log('res', res)
      if (res === true) {
        props.onFinish()
      }
    })
  })
  .catch(err => {
    console.log('error', err);
  });
};

const labelCol = { span: 4 }
const wrapperCol = { span: 18 }

</script>

<style lang="less" scoped>
.scenario-edit-main {

}
</style>
