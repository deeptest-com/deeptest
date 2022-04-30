<template>
  <div class="response-extractor-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1"></a-col>
        <a-col flex="100px" class="dp-right">
          <PlusOutlined @click.stop="add" class="dp-icon-btn dp-trans-80" />
        </a-col>
      </a-row>
    </div>

    <div class="body">
      <a-row v-for="(item, idx) in interfaceData.extractors" :key="idx" type="flex">
        <a-col flex="1"></a-col>
        <a-col flex="100px" class="dp-right">
          <EditOutlined @click.stop="edit" class="dp-icon-btn dp-trans-80" />
          <DeleteOutlined @click.stop="remove" class="dp-icon-btn dp-trans-80" />=
        </a-col>
      </a-row>
    </div>

    <a-modal
        :title="model.id ? '编辑' : '创建' + '变量'"
        :destroy-on-close="true"
        :mask-closable="false"
        :visible="editVisible"
        :onCancel="cancel"
        :footer="null"
    >
      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-form-item label="来源" v-bind="validateInfos.src">
            <a-select v-model:value="model.src"
                      @blur="validate('src', { trigger: 'change' }).catch(() => {})">
              <a-select-option v-for="(item, idx) in srcOptions" :key="idx" value="headers">
                {{ t(item.label) }}
              </a-select-option>
            </a-select>
          </a-form-item>

          <a-form-item label="类型" v-bind="validateInfos.type">
            <a-select v-model:value="model.type"
                      @blur="validate('type', { trigger: 'change' }).catch(() => {})">
              <a-select-option v-for="(item, idx) in typeOptions" :key="idx" value="headers">
                {{ t(item.label) }}
              </a-select-option>
            </a-select>
          </a-form-item>

          <a-form-item label="表达式" v-bind="validateInfos.expression">
            <a-input v-model:value="model.expression"
                     @blur="validate('expression', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>

          <a-form-item :wrapper-col="{ span: wrapperCol.span, offset: labelCol.span }">
            <a-button type="primary" @click="save" class="dp-btn-gap">保存</a-button> &nbsp;
            <a-button @click="() => cancel" class="dp-btn-gap">取消</a-button>
          </a-form-item>

        </a-form>

      </div>

    </a-modal>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, PropType, reactive, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {message, Form} from 'ant-design-vue';
import { PlusOutlined, EditOutlined, DeleteOutlined, } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import {Interface, Response} from "@/views/interface/data";
import {getEnumSelectItems} from "@/views/interface/service";
import {ExtractorSrc, ExtractorType} from "@/views/interface/consts";

const useForm = Form.useForm;

export default defineComponent({
  name: 'ResponseExtractor',
  components: {
    PlusOutlined, EditOutlined, DeleteOutlined,
  },

  computed: {
  },

  setup(props) {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();

    const srcOptions = getEnumSelectItems(ExtractorSrc)
    const typeOptions = getEnumSelectItems(ExtractorType)

    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);
    const responseData = computed<Response>(() => store.state.Interface.responseData);

    const model = ref({})
    const editVisible = ref(false)

    const rules = reactive({
      name: [
        { required: true, message: '请输入变量名', trigger: 'blur' },
      ],
    });

    const { resetFields, validate, validateInfos } = useForm(model, rules);

    const add = () => {
      editVisible.value = true
      console.log('add', editVisible.value)
    }

    const edit = () => {
      console.log('edit')
    }

    const save = () => {
      console.log('save')
    }

    const cancel = () => {
      console.log('cancel')
      editVisible.value = false
    }

    const remove = () => {
      console.log('remove')
    }

    return {
      t,
      interfaceData,
      responseData,
      model,
      editVisible,

      add,
      edit,
      remove,
      save,
      cancel,

      rules,
      validate,
      validateInfos,
      resetFields,

      srcOptions,
      typeOptions,
      labelCol: { span: 6 },
      wrapperCol: { span: 16 },
    }
  }
})

</script>

<style lang="less">
.response-extractor-main {
}
</style>

<style lang="less" scoped>
.response-extractor-main {
  height: 100%;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .body {
    padding: 6px;
    height: calc(100% - 30px);

  }
}
</style>