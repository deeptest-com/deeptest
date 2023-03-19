<template>
  <div class="response-image-main">
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
          :value="responseData.content"
          :language="responseData.contentLang"
          theme="vs"
          :options="editorOptions"
      />
    </div>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, inject, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { DownloadOutlined, CopyOutlined, ClearOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface1/store";
import {isInArray} from "@/utils/array";
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {MonacoOptions} from "@/utils/const";
import {Interface, Response} from "@/views/interface1/data";
import {UsedBy} from "@/utils/enum";
import {StateType as ScenarioStateType} from "@/views/scenario/store";

export default defineComponent({
  name: 'ResponseLensImage',
  components: {
    MonacoEditor,
    CopyOutlined, DownloadOutlined, ClearOutlined,
  },

  computed: {
  },

  setup(props) {
    const {t} = useI18n();
    const usedBy = inject('usedBy') as UsedBy
    const store = useStore<{ Interface1: StateType, Scenario: ScenarioStateType }>();
    const responseData = computed<Response>(() =>
        usedBy === UsedBy.interface ? store.state.Interface1.responseData : store.state.Scenario.responseData);

    const editorOptions = ref(Object.assign({usedWith: 'response'}, MonacoOptions) )

    return {
      responseData,
      editorOptions,
    }
  }
})

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
    height: calc(100% - 30px);
    overflow-y: hidden;
    &>div {
      height: 100%;
    }
  }
}
</style>