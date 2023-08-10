<template>
  <div class="processor_custom_code-edit dp-proccessors-container">
    <ProcessorHeader v-if="fullscreen"/>
    <div class="content">
      <div class="codes">
        <MonacoEditor theme="vs" language="typescript" class="editor"
                      :value="modelRef.content || ''"
                      :timestamp="timestamp"
                      :options="editorOptions"
                      @change="editorChange"/>
      </div>

      <div class="refer">
        <div class="desc">预请求脚本使用JavaScript编写，并在请求发送前执行。</div>

        <div class="title">代码片段：</div>
        <div>
          <!-- <div @click="addSnippet('environment_get')" class="dp-link-primary">Get an environment variable</div>
               <div @click="addSnippet('environment_set')" class="dp-link-primary">Set an environment variable</div>
               <div @click="addSnippet('environment_clear')" class="dp-link-primary">Clear an environment variable</div>-->

          <div @click="addSnippet('variables_get')" class="dp-link-primary">Get an variable</div>
          <div @click="addSnippet('variables_set')" class="dp-link-primary">Set an variable</div>
          <div @click="addSnippet('variables_clear')" class="dp-link-primary">Clear an variable</div>

          <div @click="addSnippet('datapool_get')" class="dp-link-primary">Get datapool variable</div>
        </div>
      </div>
    </div>

    <a-row>
      <a-col offset="2">
        <a-button type="primary" @click.prevent="save">保存</a-button>
      </a-col>
    </a-row>
  </div>
</template>

<script setup lang="ts">
import {computed, ref, watch, inject} from "vue";
import {useStore} from "vuex";
import {message, notification} from "ant-design-vue";
import {MonacoOptions, NotificationKeyCommon} from "@/utils/const";
import {StateType as ScenarioStateType} from "../../../../../store";
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import ProcessorHeader from '../../common/ProcessorHeader.vue';
const store = useStore<{ Scenario: ScenarioStateType; }>();
const modelRef: any = computed<boolean>(() => store.state.Scenario.nodeData);

const fullscreen = inject('fullscreen');

const editorOptions = ref(Object.assign({
      usedWith: 'request',
      initTsModules: true,

      allowNonTsExtensions: true,
      minimap: {
        enabled: false
      },
    }, MonacoOptions
))

const save = async () => {
  const res = await store.dispatch('Scenario/saveProcessor', {
    ...modelRef.value,
    content: modelRef.value.content,
  })

  if (res === true) {
    notification.success({
      key: NotificationKeyCommon,
      message: `保存成功`,
    });
  } else {
    notification.error({
      key: NotificationKeyCommon,
      message: `保存失败`,
    });
  }
}

const addSnippet = (snippetName) => {
  store.dispatch('Scenario/addSnippet', snippetName)
}
const editorChange = (newScriptCode) => {
  modelRef.value.content = newScriptCode;
}

const timestamp = ref('')
watch(modelRef, (newVal) => {
  timestamp.value = Date.now() + ''
}, {immediate: true, deep: true})

</script>

<style lang="less" scoped>
.processor_custom_code-edit {
  height: 100%;
  width: 100%;

  .content {
    height: calc(100% - 32px);
    display: flex;

    & > div {
      height: 100%;
    }

    .codes {
      flex: 1;
    }

    .refer {
      width: 260px;
      padding: 10px;
      overflow-y: auto;

      .title {
        margin-top: 12px;
      }

      .desc {

      }
    }
  }

  .codes {
    height: 100%;
    min-height: 160px;

    .editor {
      height: 100%;
      min-height: 160px;
    }
  }
}
</style>
