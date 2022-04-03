<template>
  <div class="validation-script">
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
            :value="interfaceData.validationScript"
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
import {Interface} from "@/views/interface/data";

export default defineComponent({
  name: 'RequestValidationScript',
  components: {
    ALink,
    QuestionCircleOutlined, DeleteOutlined, ClearOutlined, MonacoEditor,
  },
  setup(props) {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);
    const editorOptions = ref(MonacoOptions)

    function onJsonChange (value) {
      console.log('value:', value)
    }

    return {
      interfaceData,
      editorOptions,
    }
  }
})

</script>

<style lang="less">
</style>

<style lang="less" scoped>
.validation-script {
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
      width: 300px;
      padding: 10px;
      .title {
        margin-top: 12px;
      }
      .desc {

      }
    }
  }
}
</style>