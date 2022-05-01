<template>
  <div class="response-checkpoint-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="80px">编号</a-col>
        <a-col flex="100px">类型</a-col>
        <a-col flex="150px">变量名</a-col>
        <a-col flex="1">运算符</a-col>
        <a-col flex="100px">数值</a-col>

        <a-col flex="100px" class="dp-right">
          <PlusOutlined @click.stop="add" class="dp-icon-btn dp-trans-80" />
        </a-col>
      </a-row>
    </div>

    <div class="body">
      <a-row v-for="(item, idx) in interfaceData.checkpoints" :key="idx" type="flex">
        <a-col flex="80px">{{idx + 1}}</a-col>
        <a-col flex="100px">{{ item.src }}</a-col>
        <a-col flex="150px">{{ item.type }}</a-col>
        <a-col flex="1">{{ item.expression }}</a-col>
        <a-col flex="100px">{{ item.variable }}</a-col>

        <a-col flex="100px" class="dp-right">
          <a-tooltip v-if="!item.disabled" @click="disable(item)" overlayClassName="dp-tip-small">
            <template #title>禁用</template>
            <CheckCircleOutlined class="dp-icon-btn dp-trans-80" />
          </a-tooltip>

          <a-tooltip v-if="item.disabled" @click="disable(item)" overlayClassName="dp-tip-small">
            <template #title>启用</template>
            <CloseCircleOutlined class="dp-icon-btn dp-trans-80 dp-light" />
          </a-tooltip>

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
          <a-form-item label="类型" v-bind="validateInfos.type">
            <a-select v-model:value="model.type"
                      @blur="validate('type', { trigger: 'change' }).catch(() => {})">
              <a-select-option v-for="(item, idx) in types" :key="idx" :value="item.value">
                {{ t(item.label) }}
              </a-select-option>
            </a-select>
          </a-form-item>

          <a-form-item v-if="model.type === 'extractor'" label="变量名" v-bind="validateInfos.variable">
            <a-input v-model:value="model.variable"
                     @blur="validate('variable', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>

          <a-form-item label="运算符" v-bind="validateInfos.operator">
            <a-select v-model:value="model.operator"
                      @blur="validate('operator', { trigger: 'change' }).catch(() => {})">
              <a-select-option v-for="(item, idx) in operators" :key="idx" :value="item.value">
                {{ item.label }}
              </a-select-option>
            </a-select>
          </a-form-item>

          <a-form-item label="数值" v-bind="validateInfos.value">
            <a-input v-model:value="model.value"
                     @blur="validate('value', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>

          <a-form-item :wrapper-col="{ span: wrapperCol.span, offset: labelCol.span }">
            <a-button type="primary" @click="save" class="dp-btn-gap">保存</a-button> &nbsp;
            <a-button @click="cancel" class="dp-btn-gap">取消</a-button>
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
import { PlusOutlined, EditOutlined, DeleteOutlined, CloseCircleOutlined, CheckCircleOutlined} from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import {Interface, Response} from "@/views/interface/data";
import {getEnumSelectItems} from "@/views/interface/service";
import {CheckpointOperator, CheckpointType, ExtractorSrc, ExtractorType} from "@/views/interface/consts";

const useForm = Form.useForm;

export default defineComponent({
  name: 'ResponseCheckpoint',
  components: {
    PlusOutlined, EditOutlined, DeleteOutlined, CloseCircleOutlined, CheckCircleOutlined,
  },

  computed: {
  },

  setup(props) {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();

    const types = getEnumSelectItems(CheckpointType)
    const operators = getEnumSelectItems(CheckpointOperator)

    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);
    const responseData = computed<Response>(() => store.state.Interface.responseData);

    store.dispatch('Interface/listCheckpoint')

    const model = ref({})
    const editVisible = ref(false)

    const rules = reactive({
      type: [
        { required: true, message: '请选择类型', trigger: 'blur' },
      ],
      variable: [
        { required: true, message: '请输入变量名', trigger: 'blur' },
      ],
      operator: [
        { required: true, message: '请选择操作', trigger: 'change' },
      ],
      value: [
        { required: true, message: '请输入数值', trigger: 'blur' },
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

    const disable = (item) => {
      console.log('disabled')
      item.disabled = !item.disabled
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
      disable,
      save,
      cancel,

      rules,
      validate,
      validateInfos,
      resetFields,

      types,
      operators,
      labelCol: { span: 6 },
      wrapperCol: { span: 16 },
    }
  }
})

</script>

<style lang="less">
.response-checkpoint-main {
}
</style>

<style lang="less" scoped>
.response-checkpoint-main {
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