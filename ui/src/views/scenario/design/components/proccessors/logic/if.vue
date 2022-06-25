<template>
  <div class="processor_simple-main">
    <a-card :bordered="false">
      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-form-item :wrapper-col="{ span: 16, offset: 2 }">
            <a-row v-if="!editMap.name" type="flex">
              <a-col flex="1">
                <span class="icons">{{modelRef.name}}</span>
              </a-col>

              <a-col flex="16px" />

              <a-col flex="36px" class="icons">
                <EditOutlined @click="editName()" />
              </a-col>
            </a-row>

            <a-row v-if="editMap.name" type="flex">
              <a-col flex="1">
                <a-input v-model:value="modelRef.name" />
              </a-col>

              <a-col flex="16px" />

              <a-col flex="36px" class="icons">
                <CheckOutlined @click="saveName()" />&nbsp;
                <CloseOutlined @click="cancelName()" />
              </a-col>
            </a-row>
          </a-form-item>

          <a-form-item label="备注" v-bind="validateInfos.comments">
            <a-input v-model:value="modelRef.comments"/>
          </a-form-item>

          <a-form-item label="表达式" v-bind="validateInfos.expression">
            <a-textarea v-model:value="modelRef.expression"
                        :auto-size="{ minRows: 3, maxRows: 6 }" />
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
import {Form, message} from 'ant-design-vue';
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

const editMap = ref({} as any)
const editName = () => {
  editMap.value.name = !editMap.value.name
}
const saveName = () => {
  store.dispatch('Scenario/saveProcessorName', modelRef.value).then((res) => {
    if (res === true) {
      editMap.value.name = false
    }
  })
}
const cancelName = () => {
  editMap.value.name = false
}

const submitForm = async () => {
  validate()
      .then(() => {
        store.dispatch('Scenario/saveProcessor', modelRef.value).then((res) => {
          if (res === true) {
            message.success(`保存成功`);
          } else {
            message.error(`保存失败`);
          }
        })
      })
};


const labelCol = { span: 2 }
const wrapperCol = { span: 16 }

</script>

<style lang="less" scoped>
.processor_simple-main {
  .icons {
    text-align: right;
    line-height: 32px;
  }
}
</style>
