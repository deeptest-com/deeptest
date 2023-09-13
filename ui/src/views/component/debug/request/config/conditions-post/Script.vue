<template>
  <div class="post-script-main">
    <div class="content">
      <div class="codes">
        <MonacoEditor theme="vs" language="typescript" class="editor"
                      :value="model.content"
                      :options="editorOptions"
                      :timestamp="timestamp"
                      @change="editorChange" />
      </div>

      <div class="refer">
        <div class="desc">后置处理脚本使用JavaScript编写，并在收到请求响应后执行。</div>

        <div class="title">代码片段：</div>
        <div>
          <!--      <div @click="addSnippet('environment_get')" class="dp-link-primary">Get an environment variable</div>
                    <div @click="addSnippet('environment_set')" class="dp-link-primary">Set an environment variable</div>
                    <div @click="addSnippet('environment_clear')" class="dp-link-primary">Clear an environment variable</div>-->
          <div @click="addSnippet('variables_get')" class="dp-link-primary">获取变量</div>
          <div @click="addSnippet('variables_set')" class="dp-link-primary">设置变量</div>
          <div @click="addSnippet('variables_clear')" class="dp-link-primary">清除变量</div>

          <div @click="addSnippet('datapool_get')" class="dp-link-primary">获取数据池变量</div>

          <div @click="addSnippet('log')" class="dp-link-primary">打印日志</div>
          <div @click="addSnippet('set_mock_resp_code')" class="dp-link-primary">设置响应码</div>
          <div @click="addSnippet('set_mock_resp_field')" class="dp-link-primary">修改JSON响应对象</div>
          <div @click="addSnippet('set_mock_resp_text')" class="dp-link-primary">修改字符串响应内容</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, inject, onBeforeUnmount, onMounted, reactive, ref, watch} from "vue";
import {message, Form, notification} from 'ant-design-vue';
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined, DeleteOutlined, ClearOutlined } from '@ant-design/icons-vue';
import {UsedBy} from "@/utils/enum";

import {StateType as Debug} from "@/views/component/debug/store";
import {MonacoOptions, NotificationKeyCommon} from "@/utils/const";
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import {notifyError, notifySuccess} from "@/utils/notify";

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
  finish: {
    type: Function,
    required: false,
  },
})

const load = () => {
  console.log('load script ...', props.condition)
  store.dispatch('Debug/getScript', props.condition.entityId)
}
load()

const timestamp = ref('')
watch(model, (newVal) => {
  timestamp.value = Date.now() + ''
}, {immediate: true, deep: true})

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
  console.log('editorChange', newScriptCode)
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

    store.dispatch('Debug/saveScript', model.value).then((result) => {
      if (result) {
        notifySuccess(`保存成功`);
        if (props.finish) {
          props.finish()
        }
      } else {
        notifyError(`保存失败`);
      }
    })
  })
}
const cancel = () => {
  console.log('cancel')
  if (props.finish) {
    props.finish()
  }
}

onMounted(() => {
  console.log('onMounted')
  bus.on(settings.eventConditionSave, save);
})
onBeforeUnmount( () => {
  console.log('onBeforeUnmount')
  bus.off(settings.eventConditionSave, save);
})

const labelCol = { span: 0 }
const wrapperCol = { span: 24 }

</script>

<style lang="less" scoped>
.post-script-main {
  height: 100%;
  width: 100%;

  .content {
    display: flex;
    height: calc(100% - 32px);

    &>div {
      height: 100%;
    }
    .codes {
      flex: 1;
      height: 100%;
      min-height: 160px;

      .editor {
        height: 100%;
        min-height: 160px;
      }
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
  .footer {

  }


}
</style>