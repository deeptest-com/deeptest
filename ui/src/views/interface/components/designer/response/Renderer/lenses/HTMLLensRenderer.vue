<template>
  <div class="response-json-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">
          <span>响应体</span>
        </a-col>

        <a-col flex="100px" class="dp-right">
          <a-tooltip overlayClassName="dp-tip-small">
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
          </a-tooltip>
        </a-col>
      </a-row>
    </div>

    <div class="body">
      <MonacoEditor
          class="editor"
          :value="modelData.body"
          :language="codeLang"
          theme="vs"
          :options="editorOptions"
      />
    </div>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { DownloadOutlined, CopyOutlined, ClearOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import {isInArray} from "@/utils/array";
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {MonacoOptions} from "@/utils/const";

interface ResponseLensHtmlSetupData {
  modelData: ComputedRef;
  editorOptions: Ref
  codeLang: ComputedRef<boolean>;
}

export default defineComponent({
  name: 'ResponseLensHtml',
  components: {
    MonacoEditor,
    CopyOutlined, DownloadOutlined, ClearOutlined,
  },

  computed: {
  },

  setup(props): ResponseLensHtmlSetupData {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const modelData = computed<any>(() => store.state.Interface.modelResult);
    const codeLang = computed(() => {
      return getCodeLang()
    })
    const editorOptions = ref(MonacoOptions)

    const getCodeLang = () => {
      if (isInArray(modelData.value.bodyType, ['json', 'xml', 'html', 'text'])) {
        return modelData.value.bodyType
      } else {
        return 'plaintext'
      }
    }

    return {
      modelData,
      editorOptions,
      codeLang,
    }
  }
})

</script>

<style lang="less">
.response-json-main {
  .jsoneditor-vue {
    height: 100%;
    .jsoneditor-menu {
      display: none;
    }
    .jsoneditor-outer {
      margin: 0;
      padding: 0;
      height: 100%;
      .ace-jsoneditor {
        height: 100%;
      }
    }
  }
}
</style>

<style lang="less" scoped>
.response-json-main {
  height: 100%;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .body {
    height: calc(100% - 30px);
    overflow-y: hidden;
    &>div {
      height: 100%;
    }
  }
}
</style>