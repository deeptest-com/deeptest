<template>
  <div class="pre-script-main">
    <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
      <div class="content">
        <div class="codes">
          <MonacoEditor theme="vs" language="typescript" class="editor"
                        :value="model.content"
                        :options="editorOptions"
                        @change="editorChange" />
        </div>

        <div class="refer">
          <div class="desc">预请求脚本使用JavaScript编写，并在请求发送前执行。</div>

          <div class="title">代码片段：</div>
          <div>
            <!--          <div @click="addSnippet('environment_get')" class="dp-link-primary">Get an environment variable</div>
                      <div @click="addSnippet('environment_set')" class="dp-link-primary">Set an environment variable</div>
                      <div @click="addSnippet('environment_clear')" class="dp-link-primary">Clear an environment variable</div>-->

            <div @click="addSnippet('variables_get')" class="dp-link-primary">Get an variable</div>
            <div @click="addSnippet('variables_set')" class="dp-link-primary">Set an variable</div>
            <div @click="addSnippet('variables_clear')" class="dp-link-primary">Clear an variable</div>

            <div @click="addSnippet('datapool_get')" class="dp-link-primary">Get datapool variable</div>
          </div>
        </div>
      </div>

      <a-form-item :wrapper-col="{ span: wrapperCol.span, offset: labelCol.span }">
        <a-button type="primary" @click="save" class="dp-btn-gap">保存</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, inject, reactive, ref} from "vue";
import {message, Form} from 'ant-design-vue';
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined, DeleteOutlined, ClearOutlined } from '@ant-design/icons-vue';
import {UsedBy} from "@/utils/enum";

import {StateType as Debug} from "@/views/component/debug/store";
import {MonacoOptions} from "@/utils/const";
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";

const useForm = Form.useForm;
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

const store = useStore<{  Debug: Debug }>();

const debugInfo = computed<any>(() => store.state.Debug.debugInfo);
const debugData = computed<any>(() => store.state.Debug.debugData);
const model = computed<any>(() => store.state.Debug.scriptData);

const props = defineProps({
  condition: {
    type: Object,
    required: true,
  },
})

const load = () => {
  console.log('load', props.condition)
  store.dispatch('Debug/getScript', props.condition.entityId)
}
load()

const editorOptions = ref(Object.assign({
      usedWith: 'request',
      initTsModules: true,

      allowNonTsExtensions: true,
      minimap: {
        enabled: false
      },
    }, MonacoOptions
))

const addSnippet = (snippetName) => {
  store.dispatch('Debug/addSnippet', snippetName)
}
const editorChange = (newScriptCode) => {
  model.value.content = newScriptCode;
}

const rules = reactive({
  content: [
    { required: true, message: '请输入脚本内容', trigger: 'blur' },
  ]
} as any);

let { resetFields, validate, validateInfos } = useForm(model, rules);

const save = () => {
  console.log('save', model.value)
  validate().then(() => {
    model.value.debugInterfaceId = debugInfo.value.debugInterfaceId
    model.value.endpointInterfaceId = debugInfo.value.endpointInterfaceId
    model.value.projectId = debugData.value.projectId

    store.dispatch('Debug/saveScript', model.value)
  })
}

const labelCol = { span: 0 }
const wrapperCol = { span: 24 }

</script>

<style lang="less" scoped>
.pre-script-main {
  height: 100%;
  width: 100%;

  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .content {
    display: flex;
    height: calc(100% - 28px);
    &>div {
      height: 100%;
    }

    .codes {
      flex: 1;
    }
    .refer {
      width: 260px;
      padding: 10px;
      overflow-y: auto;

      .title {
        margin-top: 12px;
      }
      .desc {

      }
    }
  }

  .codes {
    height: 100%;
    min-height: 160px;

    .editor {
      height: 100%;
      min-height: 160px;
    }
  }
}
</style>