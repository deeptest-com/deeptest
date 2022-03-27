<template>
  <div class="pre-body-main">
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
            <QuestionCircleOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>格式化</template>
            <ClearOutlined class="dp-icon-btn dp-trans-80" />
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>清除</template>
            <DeleteOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>
        </a-col>
      </a-row>
    </div>

    <div class="content">
      <div class="codes">
        <MonacoEditor
            class="editor"
            :value="requestData.preRequestScript"
            language="javascript"
            theme="vs"
            :options="editorOptions"
        />
      </div>

      <div class="refer">
        <div class="desc">预请求脚本使用 JavaScript 编写，并在请求发送前执行。</div>

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

interface RequestPreRequestScriptSetupData {
  requestData: ComputedRef;
  editorOptions: Ref
}

export default defineComponent({
  name: 'RequestPreRequestScript',
  components: {
    ALink,
    QuestionCircleOutlined, DeleteOutlined, ClearOutlined, MonacoEditor,
  },
  setup(props): RequestPreRequestScriptSetupData {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const requestData = computed<any>(() => store.state.Interface.requestData);
    const editorOptions = ref(MonacoOptions)

    function onJsonChange (value) {
      console.log('value:', value)
    }

    return {
      requestData,
      editorOptions,
    }
  }
})

</script>

<style lang="less" scoped>
.pre-body-main {
  height: 100%;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .content {
    display: flex;
    height: calc(100% - 28px);
    &>div {
      height: 100%;
    }

    .codes {
      flex: 1;
    }
    .refer {
      padding: 10px;
      width: 360px;

      .title {
        margin-top: 12px;
      }
      .desc {

      }
    }
  }
}
</style>

<style lang="less">
</style>