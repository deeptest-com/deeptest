<template>
  <div class="response-checkpoint-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="50px">编号</a-col>
        <a-col flex="100px">类型</a-col>
        <a-col flex="120px">变量 / 键值</a-col>
        <a-col flex="60px">运算符</a-col>
        <a-col flex="100px">数值</a-col>
        <a-col flex="1">实际结果</a-col>
        <a-col flex="100px">状态</a-col>

        <a-col flex="100px" class="dp-right">
          <PlusOutlined @click.stop="add" class="dp-icon-btn dp-trans-80" />
        </a-col>
      </a-row>
    </div>

    <div class="body">
      <a-row v-for="(item, idx) in checkpointsData" :key="idx" type="flex">
        <a-col flex="50px">{{idx + 1}}</a-col>
        <a-col flex="100px">{{ t(item.type) }}</a-col>
        <a-col flex="120px">{{ item.type === CheckpointType.extractor ? item.extractorVariable : item.expression }} </a-col>
        <a-col flex="60px">{{ t(item.operator) }}</a-col>
        <a-col flex="100px">{{ item.value }}</a-col>
        <a-col flex="1" style="width: 0; word-break: break-word;">
          {{ item.actualResult }}
        </a-col>
        <a-col flex="100px" :class="getResultCls(item.resultStatus)">
          {{ item.resultStatus ? t(item.resultStatus)  : ''}}
        </a-col>

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
          <a-form-item label="类型" v-bind="validateInfos.type">
            <a-select v-model:value="model.type"
                      @change="selectType"
                      @blur="validate('type', { trigger: 'change' }).catch(() => {})">
              <a-select-option v-for="(item, idx) in types" :key="idx" :value="item.value">
                {{ t(item.label) }}
              </a-select-option>
            </a-select>
          </a-form-item>

          <a-form-item v-if="model.type === 'responseHeader'"
                       :label="model.type === 'responseHeader' ? '键值' : ''"
                       v-bind="validateInfos.expression">
            <a-input v-model:value="model.expression"
                     @blur="validate('expression', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>

          <a-form-item v-if="model.type === 'extractor'" label="变量名称" v-bind="validateInfos.extractorVariable">
            <a-select v-model:value="model.extractorVariable"
                      @blur="validate('extractorVariable', { trigger: 'blur' }).catch(() => {})">
              <a-select-option key="" value=""></a-select-option>
              <a-select-option v-for="(item, idx) in variables" :key="idx" :value="item.name">
                {{ item.name }}
              </a-select-option>
            </a-select>
          </a-form-item>

          <a-form-item label="运算符" v-bind="validateInfos.operator">
            {{ void (options = model.type === 'responseStatus' ? operatorsForCode :
              isInArray(model.type, ['responseHeader', 'responseBody']) ? operatorsForString : operators) }}

            <a-select v-model:value="model.operator"
                      @blur="validate('operator', { trigger: 'change' }).catch(() => {})">

              <a-select-option v-for="(item, idx) in options" :key="idx" :value="item.value">
                {{ t(item.label) }}
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
import {computed, ComputedRef, defineComponent, PropType, reactive, Ref, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {message, Form} from 'ant-design-vue';
import { PlusOutlined, EditOutlined, DeleteOutlined, CloseCircleOutlined, CheckCircleOutlined} from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import {Checkpoint, Extractor, Interface, Response} from "@/views/interface/data";
import {
  getEnumSelectItems,
  listExtractorVariable
} from "@/views/interface/service";
import {ComparisonOperator, CheckpointType} from "@/utils/enum";
import {isInArray} from "@/utils/array";
import {getResultCls} from "@/utils/dom"
import {getCompareOptsForRespCode, getCompareOptsForString} from "@/utils/compare";

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
    const operators = getEnumSelectItems(ComparisonOperator)
    const operatorsForString = getCompareOptsForString()
    const operatorsForCode = getCompareOptsForRespCode()

    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);
    const checkpointsData = computed(() => store.state.Interface.checkpointsData);

    watch(interfaceData, () => {
      console.log('watch interfaceData')
      listCheckPoint()
    }, {deep: true})

    const listCheckPoint = () => {
      store.dispatch('Interface/listCheckpoint')
    }
    listCheckPoint()

    const model = ref({
      type: CheckpointType.responseStatus,
      expression: '',
      extractorVariable: '',
      operator: ComparisonOperator.equal,
      value: ''} as Checkpoint)

    const variables = ref([])
    const editVisible = ref(false)

    const extractorVariableRequired = { required: true, message: '请选择变量', trigger: 'change' }
    const expressionRequired = { required: true, message: '请输入键值', trigger: 'blur' }
    const rules = reactive({
      type: [
        { required: true, message: '请选择类型', trigger: 'blur' },
      ],
      extractorVariable: [
        extractorVariableRequired
      ],
      expression: [
        expressionRequired
      ],
      operator: [
        { required: true, message: '请选择操作', trigger: 'change' },
      ],
      value: [
        { required: true, message: '请输入数值', trigger: 'blur' },
      ],
    } as any);

    const { resetFields, validate, validateInfos } = useForm(model, rules);

    const add = () => {
      console.log('add')
      editVisible.value = true
      model.value = {
        type: CheckpointType.responseStatus,
        expression: '',
        extractorVariable: '',
        operator: ComparisonOperator.equal,
        value: ''} as Checkpoint
      loadExtractorVariable()
    }

    const edit = (item) => {
      console.log('edit')
      model.value = item
      editVisible.value = true

      loadExtractorVariable()
    }

    const save = () => {
      console.log('save')
      validate().then(() => {
        model.value.interfaceId = interfaceData.value.id
        store.dispatch('Interface/saveCheckpoint', model.value).then((result) => {
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
      store.dispatch('Interface/removeCheckpoint', item.id)
    }

    const disable = (item) => {
      console.log('disabled')
      item.disabled = !item.disabled
      store.dispatch('Interface/saveCheckpoint', item)
    }

    const selectType = () => {
      console.log('selectType')

      if (model.value.type === CheckpointType.responseBody) {
        model.value.operator = ComparisonOperator.contain
      } else {
        model.value.operator = ComparisonOperator.equal
      }

      loadExtractorVariable()
    }

    const loadExtractorVariable = () => {
      if (model.value.type === CheckpointType.responseHeader) {
        rules.expression = [expressionRequired]
      } else {
        rules.expression = []
      }

      if (model.value.type === CheckpointType.extractor) {
        rules.extractorVariable = [extractorVariableRequired]

        listExtractorVariable(interfaceData.value.id).then((jsn) => {
          variables.value = jsn.data
        })
      } else {
        rules.extractorVariable = []
      }
    }

    return {
      t,
      interfaceData,
      checkpointsData,
      model,
      variables,
      editVisible,
      CheckpointType,
      getResultCls,
      add,
      edit,
      remove,
      disable,
      save,
      cancel,
      selectType,
      isInArray,

      rules,
      validate,
      validateInfos,
      resetFields,

      types,
      operators,
      operatorsForString,
      operatorsForCode,
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