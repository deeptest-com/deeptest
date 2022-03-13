<template>
  <div class="request-body-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">
          <span>
            JavaScript代码
          </span>
        </a-col>

        <a-col flex="100px" class="dp-right">
          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>帮助</template>
            <QuestionCircleOutlined class="dp-icon-btn"/>
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>格式化</template>
            <ClearOutlined class="dp-icon-btn" />
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>清除</template>
            <DeleteOutlined class="dp-icon-btn"/>
          </a-tooltip>
        </a-col>
      </a-row>
    </div>

    <div class="body">
      <div class="codes">
        <MonacoEditor
            class="editor"
            :value="modelData.validationScript"
            language="javascript"
            theme="vs"
            :options="editorOptions"
        />
      </div>
      <div class="refer">
        <div class="desc">验证脚本使用JavaScript编写，并在收到响应后执行。</div>

        <div class="title">代码片段：</div>
        <div>
          <a-link to="">Environment: Set an environment variable</a-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined, DeleteOutlined, ClearOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import ALink from "@/components/ALink/index.vue";
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {MonacoOptions} from "@/utils/const";

interface RequestValidationScriptSetupData {
  modelData: ComputedRef;
  editorOptions: Ref
}

export default defineComponent({
  name: 'RequestValidationScript',
  components: {
    ALink,
    QuestionCircleOutlined, DeleteOutlined, ClearOutlined, MonacoEditor,
  },
  setup(props): RequestValidationScriptSetupData {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const modelData = computed<any>(() => store.state.Interface.modelResult);
    const editorOptions = ref(MonacoOptions)

    function onJsonChange (value) {
      console.log('value:', value)
    }

    return {
      modelData,
      editorOptions,
    }
  }
})

</script>

<style lang="less">
.request-body-main {
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
.request-body-main {
  height: 100%;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .body {
    display: flex;
    height: calc(100% - 43px);
    overflow-y: auto;
    &>div {
      height: 100%;
    }

    .codes {
      flex: 1;
    }
    .refer {
      padding: 10px;
      width: 500px;
      .title {
        margin-top: 12px;
      }
      .desc {

      }
    }
  }
}
</style>