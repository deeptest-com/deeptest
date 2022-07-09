<template>
  <div class="processor_assertion_default-main">
    <a-card :bordered="false">
      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">

          <a-form-item label="备注" v-bind="validateInfos.comments">
            <a-input v-model:value="modelRef.comments"/>
          </a-form-item>

          <a-form-item label="变量" v-bind="validateInfos.variable">
            <a-input v-model:value="model.variable"
                     @blur="validate('variable', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>

          <a-form-item label="操作">
            <a-select v-model:value="model.opt">
              <a-select-option v-for="(item, idx) in optOptions" :key="idx" :value="item.value">
                {{ t(item.label) }}
              </a-select-option>
            </a-select>
          </a-form-item>

          <a-form-item label="取值" v-bind="validateInfos.value">
            <a-input v-model:value="modelRef.value"/>
            <div class="dp-tip-small">常量或用${name}表示的变量</div>
          </a-form-item>

          <a-form-item :wrapper-col="{ span: 16, offset: 4 }">
            <a-button type="primary" @click.prevent="submitForm">保存</a-button>
            <a-button style="margin-left: 10px" @click="resetFields">重置</a-button>
          </a-form-item>
        </a-form>
      </div>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import {computed, reactive, ref} from "vue";
import {Form} from "ant-design-vue";
import {useRouter} from "vue-router";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {StateType as ScenarioStateType} from "@/views/scenario/store";

const useForm = Form.useForm;

const router = useRouter();

const {t} = useI18n();

const formRef = ref();

const rulesRef = reactive({
  variable: [
    {required: true, message: '请输入取值', trigger: 'blur'},
  ],
  value: [
    {required: true, message: '请输入取值', trigger: 'blur'},
  ],
});

const store = useStore<{ Scenario: ScenarioStateType; }>();
const modelRef = computed<boolean>(() => store.state.Scenario.nodeData);
const {resetFields, validate, validateInfos} = useForm(modelRef, rulesRef);

const submitForm = async () => {
  validate()
      .then(() => {
        console.log(modelRef);

        // store.dispatch('Project/saveProject', modelRef.value).then((res) => {
        //   console.log('res', res)
        //   if (res === true) {
        //     message.success(`保存项目成功`);
        //     router.replace('/project/list')
        //   } else {
        //     message.error(`保存项目失败`);
        //   }
        // })
      })
      .catch(err => {
        console.log('error', err);
      });
};

const optOptions = [
  {label: '等于', value: 'equal'},
  {label: '不等于', value: 'not_equal'},
  {label: '包含', value: 'contain'},
  {label: '不包含', value: 'not_contain'},
]

const labelCol = { span: 4 }
const wrapperCol = { span: 16 }

</script>

<style lang="less" scoped>
.processor_assertion_default-main {
  .icons {
    text-align: right;
    line-height: 32px;
  }
}
</style>
