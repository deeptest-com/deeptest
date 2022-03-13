<template>
  <div class="request-body-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">
          <span>
            原始请求体
          </span>

          <a-select
              ref="bodyType"
              :options="bodyTypes"
              v-model:value="modelData.bodyType"
              size="small"
              :dropdownMatchSelectWidth="false"
              :bordered="false"
          >
          </a-select>
        </a-col>

        <a-col flex="100px" class="dp-right">
          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>帮助</template>
            <QuestionCircleOutlined class="dp-icon-btn"/>
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>全部清除</template>
            <DeleteOutlined class="dp-icon-btn"/>
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>格式化</template>
            <ClearOutlined class="dp-icon-btn" />
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>导入</template>
            <ImportOutlined class="dp-icon-btn" />
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
import { QuestionCircleOutlined, DeleteOutlined, ClearOutlined, ImportOutlined, CheckCircleOutlined, CloseCircleOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {isInArray} from "@/utils/array";

interface RequestBodySetupData {
  modelData: ComputedRef;
  bodyTypes: Ref<any[]>
  editorOptions: Ref

  codeLang: ComputedRef<boolean>;
  doSomething: (e) => void;
}

export default defineComponent({
  name: 'RequestBody',
  components: {
    MonacoEditor,
    QuestionCircleOutlined, DeleteOutlined, ClearOutlined, ImportOutlined,
  },

  computed: {
  },

  setup(props): RequestBodySetupData {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const modelData = computed<any>(() => store.state.Interface.modelResult);
    const codeLang = computed(() => {
      return getCodeLang()
    })

    const editorOptions = ref({
      colorDecorators: true,
      lineHeight: 24,
      tabSize: 2,
      autoIndent: true,
      formatOnPaste: true,
      formatOnType: true
    })

    const bodyTypes = ref([
      {value: 'json', label: 'application/json'},
      {value: 'xml', label: 'application/xml'},
      {value: 'formUrlencoded', label: 'application/x-www-form-urlencoded'},
      {value: 'formData', label: 'application/form-data'},
      {value: 'html', label: 'text/html'},
      {value: 'text', label: 'text/text'},
    ])

    const getCodeLang = () => {
      if (isInArray(modelData.value.bodyType, ['json', 'xml', 'html', 'text'])) {
        return modelData.value.bodyType
      } else {
        return 'plaintext'
      }
    }

    const doSomething = (e) => {
      console.log('doSomething', e)
    };

    return {
      modelData,
      editorOptions,
      bodyTypes,
      codeLang,
      doSomething,
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
    height: calc(100% - 43px);
    overflow-y: auto;
    &>div {
      height: 100%;
    }
  }
}
</style>