<template>
  <div class="request-body-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">原始请求体</a-col>
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
      <Vue3JsonEditor
          v-model="modelData"
          mode="code"
          :show-btns="false"
          :expandedOnStart="false"
          @json-change="onJsonChange"
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
import { Vue3JsonEditor } from 'vue3-json-editor'

interface RequestBodySetupData {
  modelData: ComputedRef;

  onJsonChange: (v) => void;
  doSomething: (e) => void;
}

export default defineComponent({
  name: 'RequestBody',
  components: {
    Vue3JsonEditor,
    QuestionCircleOutlined, DeleteOutlined, ClearOutlined, ImportOutlined,
  },
  setup(props): RequestBodySetupData {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const modelData = computed<any>(() => store.state.Interface.modelResult);

    function onJsonChange (value) {
      console.log('value:', value)
    }
    const doSomething = (e) => {
      console.log('doSomething', e)
    };

    return {
      modelData,
      onJsonChange,
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