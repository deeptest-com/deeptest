<template>
  <div class="response-extractor-main">
    <div class="head">
      <a-row type="flex" class="extractor">
        <a-col flex="60px">编号</a-col>
        <a-col flex="80px">来源</a-col>
        <a-col flex="100px">提取类型</a-col>
        <a-col flex="1">表达式 / 键值</a-col>
        <a-col flex="150px">环境变量</a-col>
        <a-col flex="1">提取结果</a-col>

        <a-col flex="100px" class="dp-right">
          <PlusOutlined @click.stop="add" class="dp-icon-btn dp-trans-80" />
        </a-col>
      </a-row>
    </div>

    <div class="items">
      <a-row v-for="(item, idx) in extractorsData" :key="idx" type="flex" class="item">
        <a-col flex="60px">{{idx + 1}}</a-col>
        <a-col flex="80px">{{ t(item.src) }}</a-col>
        <a-col flex="100px">{{ item.type ? t(item.type) : '' }}</a-col>
        <a-col flex="1">
          {{ item.src === ExtractorSrc.body ? item.expression : item.key }}
        </a-col>
        <a-col flex="150px">{{ item.variable }}</a-col>
        <a-col flex="1">{{item.result}}</a-col>

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
          <a-form-item label="数据来源" v-bind="validateInfos.src">
            <a-radio-group name="srcGroup" @change="selectSrc" v-model:value="model.src"
                           @blur="validate('src', { trigger: 'change' }).catch(() => {})">
              <a-radio v-for="(item, idx) in srcOptions" :key="idx" :value="item.value">
                {{ t(item.label) }}
              </a-radio>
            </a-radio-group>

          </a-form-item>

          <a-form-item v-if="model.src === 'body'" label="提取方法" v-bind="validateInfos.type">
            <a-select v-model:value="model.type"
                      @blur="validate('type', { trigger: 'change' }).catch(() => {})">
              <a-select-option v-for="(item, idx) in typeOptions" :key="idx" :value="item.value">
                {{ t(item.label) }}
              </a-select-option>
            </a-select>
          </a-form-item>

          <a-form-item v-if="model.src !== 'body'"  label="键值" v-bind="validateInfos.key">
            <a-input v-model:value="model.key"
                     @blur="validate('key', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>

          <a-form-item v-if="model.src === 'body'" label="元素路径" v-bind="validateInfos.expression">
            <a-input v-model:value="model.expression"
                     @blur="validate('expression', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>

          <a-form-item label="变量名称" v-bind="validateInfos.variable">
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
import {computed, defineComponent, reactive, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {Form} from 'ant-design-vue';
import {
  CheckCircleOutlined,
  CloseCircleOutlined,
  DeleteOutlined,
  EditOutlined,
  PlusOutlined
} from '@ant-design/icons-vue';
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
    const extractorsData = computed(() => store.state.Interface.extractorsData);

    store.dispatch('Interface/listExtractor')

    const model = ref({src: ExtractorSrc.header, type: ExtractorType.jsonquery, expression: '', variable: ''} as Extractor)
    const results = ref({})
    const editVisible = ref(false)

    const typeRequired = { required: true, message: '请选择类型', trigger: 'change' }
    const expressionRequired = { required: true, message: '请输入元素路径', trigger: 'blur' }
    const keyRequired = { required: true, message: '请输入键值', trigger: 'blur' }

    const rules = reactive({
      src: [
        { required: true, message: '请选择来源', trigger: 'change' },
      ],
      type: [
        typeRequired,
      ],
      expression: [
        expressionRequired,
      ],
      key: [
        keyRequired,
      ],
      variable: [
        { required: true, message: '请输入变量名', trigger: 'blur' },
      ],
    });

    const { resetFields, validate, validateInfos } = useForm(model, rules);

    const add = () => {
      editVisible.value = true
      model.value = {src: ExtractorSrc.header, type: ExtractorType.jsonquery, expression: '', variable: ''} as Extractor

      selectSrc()
    }

    const edit = (item) => {
      console.log('edit')
      model.value = item
      editVisible.value = true

      selectSrc()
    }

    const save = () => {
      console.log('save')
      validate().then(() => {
        model.value.interfaceId = interfaceData.value.id
        store.dispatch('Interface/saveExtractor', model.value).then((result) => {
          if (result) {
            editVisible.value = false
          }
        })
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

    const selectSrc = () => {
      console.log('selectSrc', model.value.src)

      if (model.value.src !== ExtractorSrc.body) {
        rules.key = [keyRequired]
        rules.expression = []
        rules.type = []
      } else {
        rules.key = []
        rules.expression = [expressionRequired]
        rules.type = [typeRequired]
      }
    }

    return {
      t,
      interfaceData,
      responseData,
      extractorsData,
      model,
      results,
      editVisible,
      ExtractorSrc,

      add,
      edit,
      remove,
      disable,
      save,
      cancel,
      selectSrc,

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