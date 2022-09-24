<template>
  <div class="response-extractor-main">
    <div class="head">
      <a-row type="flex" class="extractor">
        <a-col flex="50px">编号</a-col>
        <a-col flex="70px">来源</a-col>
        <a-col flex="90px">提取类型</a-col>
        <a-col flex="1">表达式</a-col>
        <a-col flex="150px" style="padding-left: 10px;">变量</a-col>
        <a-col flex="1">提取结果</a-col>

        <a-col flex="100px" class="dp-right">
          <PlusOutlined @click.stop="add" class="dp-icon-btn dp-trans-80" />
        </a-col>
      </a-row>
    </div>

    <div class="items">
      <a-row v-for="(item, idx) in extractorsData" :key="idx" type="flex" class="item">
        <a-col flex="50px">{{idx + 1}}</a-col>
        <a-col flex="70px">{{ t(item.src) }}</a-col>
        <a-col flex="90px">{{ item.type ? t(item.type) : '' }}</a-col>
        <a-col flex="1"  style="width: 0; word-break: break-word;">
          <span v-if="item.src === ExtractorSrc.header">
            {{ item.key }}
          </span>
          <span v-if="item.src === ExtractorSrc.body">
            {{ item.type === ExtractorType.boundary ?
                `${item.boundaryStart}-${item.boundaryEnd}[${item.boundaryIndex}] ${item.boundaryIncluded}` :
                item.expression }}
          </span>
        </a-col>
        <a-col flex="150px" style="padding-left: 10px;">{{ item.variable }}</a-col>
        <a-col flex="1"  style="width: 0; word-break: break-word;">{{item.result}}</a-col>

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

          <!-- for body -->
          <a-form-item v-if="model.src === 'body'" label="提取方法" v-bind="validateInfos.type">
            <a-select v-model:value="model.type" @change="selectType"
                      @blur="validate('type', { trigger: 'change' }).catch(() => {})">
              <a-select-option v-for="(item, idx) in typeOptions" :key="idx" :value="item.value">
                {{ t(item.label) }}
              </a-select-option>
            </a-select>
          </a-form-item>

          <!-- for header -->
          <a-form-item v-if="model.src === 'header'"  label="键值" v-bind="validateInfos.key">
            <a-input v-model:value="model.key"
                     @blur="validate('key', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>

          <template v-if="model.src === 'body' && model.type === 'boundary'">
            <a-form-item label="边界开始" v-bind="validateInfos.boundaryStart">
              <a-input v-model:value="model.boundaryStart"
                       @blur="validate('boundaryStart', { trigger: 'blur' }).catch(() => {})" />
            </a-form-item>
            <a-form-item  label="边界结束" v-bind="validateInfos.boundaryEnd">
              <a-input v-model:value="model.boundaryEnd"
                       @blur="validate('boundaryEnd', { trigger: 'blur' }).catch(() => {})" />
            </a-form-item>
            <a-form-item  label="索引值">
              <a-input-number v-model:value="model.boundaryIndex" />
            </a-form-item>
            <a-form-item  label="是否包含边界">
              <a-switch v-model:checked="model.boundaryIncluded" />
            </a-form-item>
          </template>

          <a-form-item v-if="model.src === 'body' && model.type !== 'boundary'" label="元素路径" v-bind="validateInfos.expression">
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
import {computed, defineComponent, reactive, ref, watch} from "vue";
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
import {ExtractorSrc, ExtractorType} from "@/utils/enum";

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
    const extractorsData = computed(() => store.state.Interface.extractorsData);

    watch(interfaceData, () => {
      console.log('watch interfaceData')
      listExtractor()
    }, {deep: true})

    const listExtractor = () => {
      store.dispatch('Interface/listExtractor')
    }
    listExtractor()

    const model = ref({} as Extractor)
    const results = ref({})
    const editVisible = ref(false)

    const typeRequired = { required: true, message: '请选择类型', trigger: 'change' }
    const expressionRequired = { required: true, message: '请输入元素路径', trigger: 'blur' }
    const keyRequired = { required: true, message: '请输入键值', trigger: 'blur' }
    const boundaryStartRequired = { required: true, message: '请输入边界开始字符串', trigger: 'blur' }
    const boundaryEndRequired = { required: true, message: '请输入边界结束字符串', trigger: 'blur' }

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
      boundaryStart: [
        boundaryStartRequired,
      ],
      boundaryEnd: [
        boundaryEndRequired,
      ],
      variable: [
        { required: true, message: '请输入变量名', trigger: 'blur' },
      ],
    });

    const { resetFields, validate, validateInfos } = useForm(model, rules);

    const add = () => {
      editVisible.value = true
      model.value = {src: ExtractorSrc.header, type: ExtractorType.boundary, expression: '', variable: ''} as Extractor

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

      if (model.value.src === ExtractorSrc.header) {
        rules.key = [keyRequired]
        rules.expression = []
        rules.type = []
      } else {
        rules.key = []
        rules.expression = [expressionRequired]
        rules.type = [typeRequired]
      }

      selectType()
    }
    const selectType = () => {
      console.log('selectType', model.value.type)

      if (model.value.type === ExtractorType.boundary) {
        rules.boundaryStart = [boundaryStartRequired]
        rules.boundaryEnd = [boundaryEndRequired]
        rules.expression = []
      } else {
        rules.boundaryStart = []
        rules.boundaryEnd = []
        rules.expression = [expressionRequired]
      }
    }

    return {
      t,
      interfaceData,
      extractorsData,
      model,
      results,
      editVisible,
      ExtractorSrc,
      ExtractorType,

      add,
      edit,
      remove,
      disable,
      save,
      cancel,
      selectSrc,
      selectType,

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
    .ant-col {
      word-break: break-all;
    }
  }
}
</style>