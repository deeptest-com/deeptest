<template>
  <div class="response-image-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">
          <span style="margin-left:5px;">IMAGE</span>
        </a-col>

        <a-col flex="100px" class="dp-right">
<!--          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>格式化</template>
            <ClearOutlined class="dp-icon-btn dp-trans-80" />
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>复制</template>
            <CopyOutlined class="dp-icon-btn dp-trans-80" />
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>下载</template>
            <DownloadOutlined class="dp-icon-btn dp-trans-80" />
          </a-tooltip>-->
        </a-col>
      </a-row>
    </div>

    <div class="body">
      <img class="image" :src="'data:' + responseData.contentType + ';base64,' + responseData.content" />
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, inject, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { DownloadOutlined, CopyOutlined, ClearOutlined } from '@ant-design/icons-vue';
import {StateType as Debug} from "@/views/component/debug/store";
import {MonacoOptions} from "@/utils/const";
import {UsedBy} from "@/utils/enum";

const {t} = useI18n();
const usedBy = inject('usedBy') as UsedBy

const store = useStore<{  Debug: Debug }>();

const debugData = computed<any>(() => store.state.Debug.debugData);
const responseData = computed<any>(() => store.state.Debug.responseData);

const editorOptions = ref(Object.assign({usedWith: 'response',readOnly:false}, MonacoOptions) )

</script>

<style lang="less">
.response-image-main {
  .imageeditor-vue {
    height: 100%;
    .imageeditor-menu {
      display: none;
    }
    .imageeditor-outer {
      margin: 0;
      padding: 0;
      height: 100%;
      .ace-imageeditor {
        height: 100%;
      }
    }
  }
}
</style>

<style lang="less" scoped>
.response-image-main {
  height: 100%;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .body {
    height: calc(100% - 28px);
    width: 100%;
    overflow-y: auto;
    word-wrap: break-word;

    .image {
      max-width: 100%;
      width: auto;
    }
  }
}
</style>
