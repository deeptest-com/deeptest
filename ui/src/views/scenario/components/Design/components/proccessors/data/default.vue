<template>
  <div class="processor_data_excel-main dp-processors-container">
    <ProcessorHeader/>
    <a-card :bordered="false">
      <div class="top-header-tip">
        <a-alert message="说明：数据迭代处理器将循环读取文件中的行内容，并将读取的内容赋值给指定的变量" type="info" show-icon/>
      </div>

      <a-form-item :label-col="{ span: 4 }" :wrapper-col="{ span: 16 }" label="数据来源">
        <a-radio-group v-model:value="formState.src">
          <a-radio v-for="(item, idx) in srcOptions" :key="idx" :value="item.value">
            {{ t(item.label) }}
          </a-radio>
        </a-radio-group>
      </a-form-item>

      <a-form :label-col="{ span: 4 }" :wrapper-col="{ span: 16 }"
          @submit.prevent>
        <a-form-item label="变量名称" name="variableName" v-bind="validateInfos.variableName" required>
          <a-input v-model:value="formState.variableName"
                   @blur="validate('variableName', { trigger: 'blur' }).catch(() => {})" />

          <div v-if="formState.variableName" class="dp-input-tip">
            可使用 {{ '${' + formState.variableName + '.列名' + '}' }} 访问数据变量
          </div>
        </a-form-item>

        <a-form-item v-if="formState.src === DataSrc.fileUpload"
                     label="上传文件" name="url" v-bind="validateInfos.url" required>
          <div class="upload-file">
            <div class="upload-container">
              <a-upload
                        :customRequest="upload"
                        :showUploadList="false"
                        :accept="extStr">
                <a-button><UploadOutlined/>上传文件</a-button>
              </a-upload>
            </div>
            <div class="upload-path">
              <span class="dp-input-tip" :class="[isWrongFileFormat ? 'dp-red' : '']">
                仅支持csv格式文件
              </span>
            </div>
          </div>

          <div>{{formState.url}}</div>
        </a-form-item>

        <a-form-item v-if="formState.src === DataSrc.datapool"
                     label="数据池" v-bind="validateInfos.datapoolId" required>
          <a-select v-model:value="formState.datapoolId"
                    @blur="validate('datapoolId', { trigger: 'change' }).catch(() => {})">
            <a-select-option :key="0" value="">请选择</a-select-option>
            <a-select-option v-for="(item, idx) in datapools" :key="idx" :value="''+item.id">
              {{item.name}}
            </a-select-option>
          </a-select>
        </a-form-item>

        <a-form-item label="分隔符" name="separator"
                     v-if="formState.format === 'txt'"
                     v-bind="validateInfos.separator">
          <a-input style="width: 200px;"
                   v-model:value="formState.separator"
                   @blur="validate('separator', { trigger: 'blur' }).catch(() => {})"/>

          <div class="dp-input-tip">一行多列内容可以使用分隔符来分隔</div>
        </a-form-item>

        <a-form-item label="重复次数" name="repeatTimes">
          <a-input-number style="width: 200px;"
                          v-model:value="formState.repeatTimes"/>

          <div class="dp-input-tip">将按指定次数循环读取文件内容</div>
        </a-form-item>

        <a-form-item label="是否随机" name="isRand">
          <a-switch v-model:checked="formState.isRand"/>

          <div class="dp-input-tip">开关打开，将按照随机顺序读取文件行内容</div>
        </a-form-item>

        <a-form-item label="备注" name="comments">
          <a-textarea v-model:value="formState.comments" :rows="3"/>
        </a-form-item>

        <a-form-item class="processor-btn" :wrapper-col="{ span: 16, offset: 4 }">
          <a-button type="primary" @click.prevent="submit">保存</a-button>
        </a-form-item>

      </a-form>

    </a-card>
  </div>
</template>

<script setup lang="ts">
import {computed, reactive, ref, watch} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {useI18n} from "vue-i18n";
import {Form, notification} from 'ant-design-vue';
import debounce from "lodash.debounce";
import {StateType as ScenarioStateType} from "../../../../../store";
import {UploadOutlined} from "@ant-design/icons-vue";
import {uploadRequest} from "@/utils/upload";
import ProcessorHeader from '../../common/ProcessorHeader.vue';
import {getEnumArr, getEnumSelectItems} from "@/utils/comm";
import {DataFileExt, DataSrc, ExtractorSrc} from "@/utils/enum";
import {isInArray} from "@/utils/array";
import {listDatapool} from "@/views/project-settings/service";
import {notifyError, notifySuccess} from "@/utils/notify";
const useForm = Form.useForm;

const router = useRouter();
const {t} = useI18n();

const srcOptions = getEnumSelectItems(DataSrc)

const datapools = ref([] as any[])
const loadDatapools = async () => {
  console.log('loadDatapools')
  const resp = await listDatapool({})
  if (resp.code === 0) {
    datapools.value = resp.data.result
  }
}
loadDatapools()

const store = useStore<{ Scenario: ScenarioStateType; }>();
const nodeData: any = computed<boolean>(() => store.state.Scenario.nodeData);

const extArr = getEnumArr(DataFileExt)
const extStr = extArr.join(',')

const formState = ref({
  src: DataSrc.fileUpload,
  variableName: '',
  url: '',
  datapoolId: '',
  format: '',
  separator: '',
  repeatTimes: 1,
  isRand: false,
  comments: '',
});
const rulesRef:any = computed(() => { return {
  variableName: [
    {required: true, message: '请输入变量名称', trigger: 'blur'},
  ],
  url: formState.value.src === DataSrc.fileUpload ? [{required: true, message: '请上传文件', trigger: 'blur'}] : [],
  datapoolId: formState.value.src === DataSrc.datapool ? [{required: true, message: '请选择数据池', trigger: 'change'}] : [],
  separator: formState.value.format === 'txt' ? [{required: true, message: '请输入分隔符', trigger: 'blur'}] : [],
}})

const {resetFields, validate, validateInfos} = useForm(formState, rulesRef);

const isWrongFileFormat = ref(false)

const upload = async (e) => {
  const file = e.file;
  const ext = file.name.substr(file.name.lastIndexOf('.'));
  console.log('upload', file, ext)

  if (!isInArray(ext, extArr)) {
    isWrongFileFormat.value = true
    return false
  } else {
    isWrongFileFormat.value = false
  }

  const res = await uploadRequest(file)
  formState.value.url = res.path;
  formState.value.format = res.format;

  if(formState?.value?.format === 'txt') {
    formState.value.separator = ',';
  }

  return false
}

watch(nodeData, (val: any) => {
  console.log('watch nodeData')
  if (!val) return;

  formState.value.variableName = val.variableName;
  formState.value.src = val.src || DataSrc.fileUpload;
  formState.value.url = val.url;
  formState.value.datapoolId = val.datapoolId ? ''+val.datapoolId : '';
  formState.value.separator = val.separator;
  formState.value.repeatTimes = val.repeatTimes || 1;
  formState.value.isRand = val.isRand;
  formState.value.comments = val.comments;
}, {deep: true, immediate: true});

const submit = debounce(async () => {
  console.log('rulesRef', rulesRef.value, formState.value.datapoolId)

  validate().then(async () => {
        // 下面代码改成 await 的方式
    const data = {
      ...nodeData.value,
      ...formState.value,
    }
    data.datapoolId = +data.datapoolId

        const res = await store.dispatch('Scenario/saveProcessor', data);
        if (res === true) {
          notifySuccess(`保存成功`);
        } else {
          notifyError(`保存失败`);
        }
      })
      .catch(error => {
        console.log('error', error);
      });
}, 300);


</script>

<style lang="less" scoped>
.top-header-tip {
  position: relative;
  margin: 6px auto 24px 60px;
}
</style>
