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
              <div class="item">
                <div>// 获取自动Mock响应数据</div>
                <div>var responseJson = fox.mockResponse.json();</div>
                <br />
                <div>// 修改数据里的分页信息</div>
                <div>// 修改page当前页数字段</div>
                <div>responseJson.page = fox.mockRequest.getParam('page');</div>
                <br />
                <div>// 修改total总页数字段</div>
                <div>responseJson.total = 120;</div>
                <br />
                <div>// 将修改后的结果写入fox.mockResponse</div>
                <div>fox.mockResponse.setBody(responseJson);</div>
              </div>
              <br />
              <div class="item">
                <div>请求对象: fox.mockRequest</div>
                <br />
                <div>fox.mockRequest.getParam(key: string) 获取请求参数，包括Path参数、Body参数、Query参数。</div>
                <div>fox.mockRequest.headers 请求的HTTP头</div>
                <div>fox.mockRequest.cookies 请求带的Cookies</div>
                <br />
                <div>响应对象: fox.mockResponse</div>
                <br />
                <div>fox.mockResponse.json() 系统自动生成的JSON格式响应数据</div>
                <div>fox.mockResponse.setBody(body: any) 设置接口返回Body, 参数支持JSON或字符串。</div>
                <div>fox.mockResponse.setCode(code: number) 设置接口返回的HTTP状态码</div>
                <div>fox.mockResponse.setDelay(duration: number) 设置Mock响应延时，单位为毫秒。</div>
                <div>fox.mockResponse.headers 响应的HTTP头</div>
              </div>
            </a-card>
          </div>

          <div class="title">代码片段：</div>
          <div>
            <div @click="addSnippet('get_mock_request')" class="dp-link-primary">
              获取请求
            </div>
            <div @click="addSnippet('set_mock_response')" class="dp-link-primary">
              设置响应
            </div>
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
  store.dispatch('Debug/addSnippet', snippetName)
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

const labelCol = { span: 0 }
const wrapperCol = { span: 24 }

</script>

<style lang="less">
.mock-script-main {
  .ant-card.sample-content .ant-card-body {
    padding: 10px !important;
  }
}
</style>

<style lang="less" scoped>
.mock-script-main {
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