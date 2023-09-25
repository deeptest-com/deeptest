<template>
  <div class="mock-script-main">
      <div class="content">
        <div class="codes">
          <MonacoEditor theme="vs" language="typescript" class="editor"
                        :value="mockScript.content"
                        :timestamp="timestamp"
                        :options="editorOptions"
                        @change="editorChange" />
        </div>

        <div class="refer">
          <div class="desc">可通过脚本获取用户请求参数，修改返回响应内容。</div>

          <div class="title">示例代码：</div>
          <div class="sample">
            <a-card class="sample-content">
              <pre>{{ doc }}</pre>
            </a-card>
          </div>

          <div class="title">代码片段：</div>
          <div>
            <div @click="addSnippet('get_param')" class="dp-link-primary">获取指定参数值</div>
            <div @click="addSnippet('get_header')" class="dp-link-primary">获取指定Header值</div>
            <div @click="addSnippet('get_cookie')" class="dp-link-primary">获取指定Cookie对象</div>

            <div @click="addSnippet('set_mock_resp_code')" class="dp-link-primary">设置响应码</div>
            <div @click="addSnippet('set_mock_resp_field')" class="dp-link-primary">修改JSON响应字段</div>
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

const store = useStore<{ Endpoint }>();
const endpoint = computed<any>(() => store.state.Endpoint.endpointDetail);
const mockScript = computed<any>(() => store.state.Endpoint.mockScript);

const timestamp = ref('')
watch(mockScript, (newVal) => {
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
  store.dispatch('Endpoint/addSnippet', snippetName)
}
const editorChange = (newScriptCode) => {
  mockScript.value.content = newScriptCode;
}

onMounted(() => {
  console.log('onMounted')
  bus.on(settings.paneResizeTop, () => {
      bus.emit(settings.eventEditorAction, {
        act: 'heightChanged',
        container: 'codes'
      })
    })
})
onBeforeUnmount( () => {
  console.log('onBeforeUnmount')
  bus.off(settings.paneResizeTop, () => {
      bus.emit(settings.eventEditorAction, {
        act: 'heightChanged',
        container: 'codes'
      })
    })
})

const doc = `
/** 获取请求和响应对象 */
var req = dt.request;
var resp = dt.response;

/** 获取请求参数、Header和Cookie */
var strVal = dt.getParam('name');
var strVal = dt.getHeader('name');
// cookie is an object with name, value, path, domain properties
var objVal = dt.getCookie('name');

/** 判断请求方法 */
if (dt.request.method.toLowerCase() == 'get') {
  // do something
}

/** 判断响应内容类型 */
if (dt.response.contentType == 'text/html') {
  // do something
}

/** 修改响应码 */
dt.response.statusCode = 404;

/** 修改JSON响应字段 */
dt.response.data.field1 = 'newValue';

/** 修改HTML响应内容 */
if (dt.response.contentType == 'text/html') {
  dt.response.data =
  dt.response.data.replace('old', 'new');
}
`.trim()

const labelCol = { span: 0 }
const wrapperCol = { span: 24 }

</script>

<style lang="less">
.mock-script-main {

}
</style>

<style lang="less" scoped>
.mock-script-main {
  height: 100%;
  width: 100%;

  .ant-card.sample-content .ant-card-body {
    padding: 10px !important;
  }

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
      height: 100%;
      width: 360px;
      padding: 10px;
      overflow-y: auto;

      .title {
        margin-top: 12px;
      }
      .desc {
      }
      .sample {
        .ant-card {
          background-color: rgb(246, 245, 245);
          border: 1px solid #d9d9d9;
          .ant-card-body {
            padding: 10px !important;
            .item {
              margin: 6px;
              div {
                word-break: break-all;
              }
            }
          }
        }
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