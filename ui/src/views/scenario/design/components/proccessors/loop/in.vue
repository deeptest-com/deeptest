<template>
  <div class="processor_loop_in-main">
    <a-card :bordered="false">
      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">

          <a-form-item label="备注" v-bind="validateInfos.comments">
            <a-input v-model:value="modelRef.comments"/>
          </a-form-item>

          <a-form-item :wrapper-col="{ span: 16, offset: 2 }">
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
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';
import {StateType as ScenarioStateType} from "../../../../store";
import {EditOutlined, CheckOutlined, CloseOutlined} from "@ant-design/icons-vue";

const useForm = Form.useForm;

const router = useRouter();

const {t} = useI18n();

const formRef = ref();

const rulesRef = reactive({
  name: [
    {required: true, message: '请输入名称', trigger: 'blur'},
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


const labelCol = { span: 2 }
const wrapperCol = { span: 16 }

</script>

<style lang="less" scoped>
.processor_loop_in-main {
  .icons {
    text-align: right;
    line-height: 32px;
  }
}
</style>
