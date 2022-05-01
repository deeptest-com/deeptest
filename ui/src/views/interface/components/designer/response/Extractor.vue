<template>
  <div class="response-extractor-main">
    <div class="head">
      <a-row type="flex" class="extractor">
        <a-col flex="60px">编号</a-col>
        <a-col flex="80px">来源</a-col>
        <a-col flex="100px">提取类型</a-col>
        <a-col flex="1">表达式</a-col>
        <a-col flex="150px">环境变量</a-col>
        <a-col flex="100px">提取结果</a-col>

        <a-col flex="100px" class="dp-right">
          <PlusOutlined @click.stop="add" class="dp-icon-btn dp-trans-80" />
        </a-col>
      </a-row>
    </div>

    <div class="items">
      <a-row v-for="(item, idx) in checkpointsData" :key="idx" type="flex" class="item">
        <a-col flex="60px">{{idx + 1}}</a-col>
        <a-col flex="80px">{{ item.src }}</a-col>
        <a-col flex="100px">{{ item.type }}</a-col>
        <a-col flex="1">{{ item.expression }}</a-col>
        <a-col flex="150px">{{ item.variable }}</a-col>
        <a-col flex="100px">{{item.result}}</a-col>

        <a-col flex="100px" class="dp-right">
          <a-tooltip v-if="!item.disabled" @click="disable(item)" overlayClassName="dp-tip-small">
            <template #title>禁用</template>
            <CheckCircleOutlined class="dp-icon-btn dp-trans-80" />
          </a-tooltip>

          <a-tooltip v-if="item.disabled" @click="disable(item)" overlayClassName="dp-tip-small">
            <template #title>启用</template>
            <CloseCircleOutlined class="dp-icon-btn dp-trans-80 dp-light" />
          </a-tooltip>

          <EditOutlined @click.stop="edit(item)" class="dp-icon-btn dp-trans-80" />
          <DeleteOutlined @click.stop="remove(item)" class="dp-icon-btn dp-trans-80" />
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
              <a-select-option v-for="(item, idx) in srcOptions" :key="idx" :value="item.value">
                {{ t(item.label) }}
              </a-select-option>
            </a-select>
          </a-form-item>

          <a-form-item v-if="model.src === 'body'" label="提取方法" v-bind="validateInfos.type">
            <a-select v-model:value="model.type"
                      @blur="validate('type', { trigger: 'change' }).catch(() => {})">
              <a-select-option v-for="(item, idx) in typeOptions" :key="idx" :value="item.value">
                {{ t(item.label) }}
              </a-select-option>
            </a-select>
          </a-form-item>

          <a-form-item :label="model.src === 'body' ? '表达式' : '键值'" v-bind="validateInfos.expression">
            <a-input v-model:value="model.expression"
                     @blur="validate('expression', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>

          <a-form-item label="环境变量" v-bind="validateInfos.variable">
            <a-input v-model:value="model.variable"
                     @blur="validate('variable', { trigger: 'blur' }).catch(() => {})" />
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
import {Extractor, Interface, Response} from "@/views/interface/data";
import {getEnumSelectItems} from "@/views/interface/service";
import {ExtractorSrc, ExtractorType} from "@/views/interface/consts";

const useForm = Form.useForm;

export default defineComponent({
  name: 'ResponseExtractor',
  components: {
    PlusOutlined, EditOutlined, DeleteOutlined, CloseCircleOutlined, CheckCircleOutlined,
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
    const checkpointsData = computed(() => store.state.Interface.extractorsData);

    store.dispatch('Interface/listExtractor')

    const model = ref({src: '', type: '', expression: '', variable: ''} as Extractor)
    const results = ref({})
    const editVisible = ref(false)

    const rules = reactive({
      src: [
        { required: true, message: '请选择来源', trigger: 'change' },
      ],
      type: [
        { required: true, message: '请选择类型', trigger: 'change' },
      ],
      expression: [
        { required: true, message: '请输入表达式', trigger: 'blur' },
      ],
      variable: [
        { required: true, message: '请输入变量名', trigger: 'blur' },
      ],
    });

    const { resetFields, validate, validateInfos } = useForm(model, rules);

    const add = () => {
      editVisible.value = true
      console.log('add', editVisible.value)
      model.value = {src: '', type: '', expression: '', variable: ''} as Extractor
    }

    const edit = (item) => {
      console.log('edit')
      model.value = item
      editVisible.value = true
    }

    const save = () => {
      console.log('save')
      model.value.interfaceId = interfaceData.value.id
      store.dispatch('Interface/saveExtractor', model.value).then((result) => {
        if (result) {
          editVisible.value = false
        }
      })
    }

    const cancel = () => {
      console.log('cancel')
      editVisible.value = false
    }

    const remove = (item) => {
      console.log('remove')
      store.dispatch('Interface/removeExtractor', item.id)
    }

    const disable = (item) => {
      console.log('disabled')
      item.disabled = !item.disabled
      store.dispatch('Interface/saveExtractor', item)
    }

    return {
      t,
      interfaceData,
      responseData,
      checkpointsData,
      model,
      results,
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

      srcOptions,
      typeOptions,
      labelCol: { span: 6 },
      wrapperCol: { span: 16 },
    }
  }
})

</script>

<style lang="less">
</style>

<style lang="less" scoped>
.response-extractor-main {
  height: 100%;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .items {
    padding: 6px;
    height: calc(100% - 30px);
  }
  .item {

  }
}
</style>