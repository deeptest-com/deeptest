<template>
  <div class="processor_data_excel-main">
    <a-card :bordered="false">
      <div class="top-header-tip">
        <a-alert message="说明：数据迭代处理器将循环读取文件中的行内容，并将读取的内容赋值给指定的变量" type="info" show-icon/>
      </div>

      <a-form
          ref="formRef"
          :rules="rules"
          :model="formState"
          :label-col="{ span: 4 }"
          @submit.prevent
          :wrapper-col="{ span: 16 }">
        <a-form-item label="变量名称" name="variableName">
          <a-input v-model:value="formState.variableName"/>
          <div v-if="formState.variableName" class="dp-input-tip">
            可使用 {{ '${' + formState.variableName + '.列名' + '}' }} 访问数据变量
          </div>
        </a-form-item>

        <a-form-item label="上传文件" name="url">
          <div class="upload-file">
            <div class="upload-container">
              <a-upload :beforeUpload="upload"
                        :fileList="fileList"
                        :showUploadList="true"
                        :accept="'.xls, .xlsx, .csv, .txt'">
                <a-button>
                  <UploadOutlined/>上传文件
                </a-button>
              </a-upload>
            </div>
          </div>
          <span class="dp-input-tip">仅支持excel、csv、 text三种文件格式</span>
        </a-form-item>

        <a-form-item label="分隔符" name="separator">
          <a-input style="width: 200px;"
              v-model:value="formState.separator"/>
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
          <a-input v-model:value="formState.comments"/>
        </a-form-item>

        <a-form-item :wrapper-col="{ span: 16, offset: 4 }">
          <a-button type="primary" @click.prevent="submit">保存</a-button>
        </a-form-item>

      </a-form>

    </a-card>
  </div>
</template>

<script setup lang="ts">
import {computed, ref, watch} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';
import {StateType as ScenarioStateType} from "../../../../../store";
import {UploadOutlined} from "@ant-design/icons-vue";
import settings from "@/config/settings";
import {getUrls} from "@/utils/request";
import {getToken} from "@/utils/localToken";
import {uploadRequest} from "@/utils/upload";

const useForm = Form.useForm;

const router = useRouter();

const {t} = useI18n();

const formRef = ref();

const rules:any = ref({
  variableName: [
    {required: true, message: '请输入变量名称', trigger: 'blur'},
  ],
  url: [
    {required: true, message: '请选择文件', trigger: 'blur'},
  ],
});


const store = useStore<{ Scenario: ScenarioStateType; }>();
const nodeData: any = computed<boolean>(() => store.state.Scenario.nodeData);


const formState = ref({
  variableName: '',
  url: '',
  separator: '',
  repeatTimes: 1,
  isRand: false,
  comments: '',
});


const fileList = ref([]);
const upload = async (file, fileList) => {
  // debugger;
  const res = await uploadRequest(file)
  formState.value.url = res.path;
  // debugger;
  // fileList.value = [{
  //   ...file,
  //   status: 'done',
  //   url: res.path,
  // }];
  return false
}

watch(() => {
  return nodeData.value;
}, (val: any) => {
  if (!val) return;
  formState.value.variableName = val.variableName;
  formState.value.url = val.url;
  formState.value.separator = val.separator;
  formState.value.repeatTimes = val.repeatTimes;
  formState.value.isRand = val.isRand;
  formState.value.comments = val.comments;
});

watch(() => {
  return formState.value.url;
}, (val: any) => {
  if (!val) return;
  //  .txt 结尾
  if(val.endsWith('.txt')) {
    rules.value.separator = [
      {required: true, message: '请输入分隔符', trigger: 'blur'},
    ]
  }


}, {deep: true});

const submit = async () => {
  formRef.value
      .validate()
      .then(async () => {
        // 下面代码改成 await 的方式
        const res = await store.dispatch('Scenario/saveProcessor', {
          ...nodeData.value,
          ...formState.value,
        });
        if (res === true) {
          message.success('保存成功');
        } else {
          message.error('保存失败');
        }
      })
      .catch(error => {
        console.log('error', error);
      });
};


</script>

<style lang="less" scoped>
.top-header-tip {
  position: relative;
  margin: 6px auto 24px 60px;
}
</style>
