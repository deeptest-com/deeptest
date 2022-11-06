<template>
  <div id="expresslayout-right-top">
    <div class="expresslayout-right-top-top">
      <a-row type="flex" class="nav">
        <a-col flex="auto" class="form">
          <a-row type="flex">
            <a-col flex="90px" class="from-file">
              <a-input-search
                  v-model:value="modelRef.file"
                  @search="load('file')">
                <template #enterButton>
                  <a-button>上传文件</a-button>
                </template>
              </a-input-search>
            </a-col>

            <a-col flex="auto" class="from-url">
              <a-input-search
                  v-model:value="modelRef.url"
                  placeholder="输入地址"
                  enter-button="加载"
                  @search="load('url')"
              />
            </a-col>
          </a-row>
        </a-col>

        <a-col flex="145px" class="github">
          <iframe src="https://ghbtns.com/github-btn.html?user=aaronchen2k&amp;repo=deeptest&amp;type=star&amp;count=true&amp;size=large"
                  frameborder="0" scrolling="0" width="136px" height="30px"></iframe>
        </a-col>
      </a-row>
    </div>

<!--    <RightTopWebsocket/>-->
  </div>

</template>
<script setup lang="ts">
import {defineComponent, PropType, toRefs, ref} from "vue";
import { message } from 'ant-design-vue';
import {useI18n} from "vue-i18n";
import { UploadOutlined } from '@ant-design/icons-vue';

import RightTopWebsocket from './RightTopWebsocket.vue';
import settings from "@/config/settings";
import {submitSpec} from "@/views/interface/service";

const {t} = useI18n();
// const {topNavEnable} = toRefs(props);

const isElectron = ref(!!window.require)
const modelRef = ref({
  type: 'openapi3',
  url: 'https://gitee.com/deeptest-com/deeptest/raw/main/xdoc/openapi/openapi3/callbacks.yml'
  //  url: 'https://gitee.com/deeptest-com/deeptest/raw/main/xdoc/openapi/swagger/swagger.json'
  // url: 'https://gitee.com/deeptest-com/deeptest/raw/main/xdoc/openapi/postman/v21/PostmantoOpenAPI.json'
} as any)

let ipcRenderer = undefined as any

if (isElectron.value && !ipcRenderer) {
  ipcRenderer = window.require('electron').ipcRenderer

  ipcRenderer.on(settings.electronMsgReplay, (event, data) => {
    console.log('from electron: ', data)

    submitSpec(data).then((json) => {
      if (json.code === 0) {
        console.log('submitSpec', json.data)
      } else {
        message.error(json.msg)
      }
    })
  })
}

const load = (src) => {
  console.log('load')
  if (!isElectron.value) return

  const data = {act: 'loadSpec', type: modelRef.value.type, src: src} as any
  if (src === 'url') {
    data.url = modelRef.value.url
    data.type = undefined
  }

  console.log(data)
  ipcRenderer.send(settings.electronMsg, data)
}

const fileList = ref([]);

</script>

<style lang="less">
#expresslayout-right-top {
  width: 100%;
  height: 50px;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
  z-index: 9;

  .expresslayout-right-top-top {
    display: flex;
    padding: 10px 5px 10px 10px;
    width: 100%;
    height: 50px;
    background-color: #FAFAFA;
    color: #343333;

    .nav {
      width: 100%;
      .form {
        .from-file {
          .ant-input-search {

            input {
              display: none;
            }
          }
        }
      }
      .github {
        text-align: right;
      }
    }

  }
}
</style>