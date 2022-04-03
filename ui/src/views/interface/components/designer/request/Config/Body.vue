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
              v-model:value="interfaceData.apiKey.transferMode"
              size="small"
              :dropdownMatchSelectWidth="false"
              :bordered="false"
          >
          </a-select>
        </a-col>

        <a-col flex="100px" class="dp-right">
          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>帮助</template>
            <QuestionCircleOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>全部清除</template>
            <DeleteOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>格式化</template>
            <ClearOutlined class="dp-icon-btn dp-trans-80" />
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>导入</template>
            <ImportOutlined class="dp-icon-btn dp-trans-80" />
          </a-tooltip>
        </a-col>
      </a-row>
    </div>

    <div class="body">
      <MonacoEditor
          class="editor"
          :value="interfaceData.body"
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
import {isInArray} from "@/utils/array";
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {MonacoOptions} from "@/utils/const";
import {Interface} from "@/views/interface/data";

export default defineComponent({
  name: 'RequestBody',
  components: {
    MonacoEditor,
    QuestionCircleOutlined, DeleteOutlined, ClearOutlined, ImportOutlined,
  },

  computed: {
  },

  setup(props) {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);
    const codeLang = computed(() => {
      return getCodeLang()
    })
    const editorOptions = ref(MonacoOptions)

    const bodyTypes = ref([
      {value: 'json', label: 'application/json'},
      {value: 'xml', label: 'application/xml'},
      {value: 'formUrlencoded', label: 'application/x-www-form-urlencoded'},
      {value: 'formData', label: 'application/form-data'},
      {value: 'html', label: 'text/html'},
      {value: 'text', label: 'text/text'},
    ])

    const getCodeLang = () => {
      if (isInArray(interfaceData.value.bodyType, ['json', 'xml', 'html', 'text'])) {
        return interfaceData.value.bodyType
      } else {
        return 'plaintext'
      }
    }

    return {
      interfaceData,
      editorOptions,
      bodyTypes,
      codeLang,
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
    height: calc(100% - 30px);
    overflow-y: hidden;
    &>div {
      height: 100%;
    }
  }
}
</style>