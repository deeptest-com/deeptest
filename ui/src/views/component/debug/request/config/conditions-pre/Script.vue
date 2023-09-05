<template>
  <div class="pre-script-main">
      <div class="content">
        <div class="codes">
          <MonacoEditor 
            theme="vs" language="typescript" class="editor"
            customId="pre-script-main-codes"
            :value="scriptData.content"
            :timestamp="timestamp"
            :options="editorOptions"
            @change="editorChange" />
        </div>

        <div class="refer">
          <div class="desc">预请求脚本使用JavaScript编写，并在请求发送前执行。</div>

          <div class="title">代码片段：</div>
          <div>
            <!--      <div @click="addSnippet('environment_get')" class="dp-link-primary">Get an environment variable</div>
                      <div @click="addSnippet('environment_set')" class="dp-link-primary">Set an environment variable</div>
                      <div @click="addSnippet('environment_clear')" class="dp-link-primary">Clear an environment variable</div>-->

            <div @click="addSnippet('variables_get')" class="dp-link-primary">Get an variable</div>
            <div @click="addSnippet('variables_set')" class="dp-link-primary">Set an variable</div>
            <div @click="addSnippet('variables_clear')" class="dp-link-primary">Clear an variable</div>

            <div @click="addSnippet('datapool_get')" class="dp-link-primary">Get datapool variable</div>
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
const scriptData = computed<any>(() => store.state.Debug.scriptData);

const timestamp = ref('')
watch(scriptData, (newVal) => {
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
  scriptData.value.content = newScriptCode;
}

const save = async () => {
  console.log('save', scriptData.value)

  scriptData.value.debugInterfaceId = debugInfo.value.debugInterfaceId
  scriptData.value.endpointInterfaceId = debugInfo.value.endpointInterfaceId
  scriptData.value.projectId = debugData.value.projectId

  const result = await store.dispatch('Debug/saveScript', scriptData.value)
  if (result) {
    notifySuccess(`保存成功`)
  } else {
    notifyError(`保存失败`);
  }
}

onMounted(() => {
  console.log('onMounted')
  bus.on(settings.eventConditionSave, save);
  bus.on(settings.paneResizeTop, () => {
      bus.emit(settings.eventEditorAction, {
        act: 'heightChanged',
        container: 'codes',
        id: 'pre-script-main-codes'
      })
    })
})
onBeforeUnmount( () => {
  console.log('onBeforeUnmount')
  bus.off(settings.eventConditionSave, save);
})

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
    height: 100%;
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