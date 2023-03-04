<template>
  <div class="plan-edit-main">
    <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
      <a-form-item label="名称" v-bind="validateInfos.name">
        <a-input v-if="editField==='name'"
                 v-model:value="modelRef.name"
                 @focusout="saveName"
                 @pressEnter="saveName"/>

        <span v-else>
              {{ modelRef.name }}
              <edit-outlined class="editable-cell-icon" @click="edit('name')"/>
            </span>
      </a-form-item>

      <a-form-item label="描述" v-bind="validateInfos.desc">
        <a-input v-if="editField==='desc'"
                 v-model:value="modelRef.desc"
                 @focusout="saveDesc"
                 @pressEnter="saveDesc"/>

        <span v-else>
              {{ modelRef.desc }}
              <edit-outlined class="editable-cell-icon" @click="edit('desc')"/>
            </span>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup lang="ts">
import {defineProps, PropType, reactive, ref} from "vue";
import {EditOutlined} from '@ant-design/icons-vue';
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';
import {StateType} from "../store";
import {get} from "@/views/plan/service";

const useForm = Form.useForm;

const router = useRouter();
const {t} = useI18n();

const props = defineProps({
  modelId: {
    type: Number,
    required: true
  },
  categoryId: {
    type: Number,
    required: true
  },
  onFieldSaved: {
    type: Function as PropType<() => void>,
    required: true
  }
})

const rulesRef = reactive({
  name: [
    {required: true, message: '请输入名称', trigger: 'blur'},
  ],
});

const store = useStore<{ Plan: StateType }>();
const modelRef = ref({} as any)
const {resetFields, validate, validateInfos} = useForm(modelRef, rulesRef);

const editField = ref('')

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

const edit = (field) => {
  console.log('edit')
  editField.value = field
}

const saveName = () => {
  console.log('saveName')
  if (!modelRef.value.name) return
  saveModel()
}
const saveDesc = () => {
  console.log('saveDesc')
  if (!modelRef.value.desc) return
  saveModel()
}

const saveModel = async () => {
  console.log('saveModel');
  store.dispatch('Plan/savePlan', modelRef.value).then((res) => {
    console.log('res', res)
    editField.value = ''
    if (res === true) {
      props.onFieldSaved()
    }
  })
};

const labelCol = {span: 2}
const wrapperCol = {span: 15}

</script>

<style lang="less" scoped>
.plan-edit-main {

}
</style>
