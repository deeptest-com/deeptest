<template>
  <div id="expresslayout-right-top">
    <div class="expresslayout-right-top-top">
      <a-row type="flex" class="nav">
        <a-col flex="auto" class="form">
          <a-row type="flex">
            <a-col flex="90px">
              <a-upload
                  v-model:file-list="fileList"
                  name="file"
                  :multiple="true"
                  action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
                  :headers="headers"
                  @change="handleChange"
              >
                <a-button>
                  <upload-outlined></upload-outlined>
                  <span>上传文件</span>
                </a-button>
              </a-upload>
            </a-col>
            <a-col flex="auto">
              <a-input-search
                  v-model:value="value"
                  placeholder="输入地址"
                  enter-button="加载"
                  @search="onSearch"
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

const {t} = useI18n();
// const {topNavEnable} = toRefs(props);

const handleChange = (info) => {
  if (info.file.status !== 'uploading') {
    console.log(info.file, info.fileList);
  }
  if (info.file.status === 'done') {
    message.success(`${info.file.name} file uploaded successfully`);
  } else if (info.file.status === 'error') {
    message.error(`${info.file.name} file upload failed.`);
  }
};

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
      }
      .github {
        text-align: right;
      }
    }

  }
}
</style>