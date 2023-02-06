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
            v-model:value="interfaceData.preRequestScript"
            language="typescript"
            theme="vs"
            :options="editorOptions"
            @change="editorChange"
        />
      </div>

      <div class="refer">
        <div class="desc">预请求脚本使用 JavaScript 编写，并在请求发送前执行。</div>

        <div class="title">代码片段：</div>
        <div>
          <div @click="addSnippet('environment_get')" class="dp-link-primary">Set an environment variable</div>
          <div @click="addSnippet('environment_set')" class="dp-link-primary">Get an environment variable</div>
          <div @click="addSnippet('environment_clear')" class="dp-link-primary">Clear an environment variable</div>

          <div @click="addSnippet('variables_get')" class="dp-link-primary">Set an variable</div>
          <div @click="addSnippet('variables_set')" class="dp-link-primary">Get an variable</div>
          <div @click="addSnippet('variables_clear')" class="dp-link-primary">Clear an variable</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, ComputedRef, defineComponent, inject, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined, DeleteOutlined, ClearOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import ALink from "@/components/ALink/index.vue";
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {MonacoOptions, NotificationKeyCommon} from "@/utils/const";
import {Interface} from "@/views/interface/data";
import {UsedBy} from "@/utils/enum";
import {StateType as ScenarioStateType} from "@/views/scenario/store";
import {inviteUser} from "@/views/user/info/service";
import {notification} from "ant-design-vue";
import {getSnippet} from "@/views/interface/service";

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();
const store = useStore<{ Interface: StateType, Scenario: ScenarioStateType }>();
const interfaceData = computed<Interface>(
    () => usedBy === UsedBy.interface ? store.state.Interface.interfaceData : store.state.Scenario.interfaceData);

const editorOptions = ref(Object.assign({
    usedWith: 'request',
    allowNonTsExtensions: true,
    minimap: {
      enabled: false
    },
  }, MonacoOptions
))

const addSnippet = (name) => {
  store.dispatch('Interface/addSnippet', name)
}

const editorChange = (newScriptCode) => {
  interfaceData.value.preRequestScript = newScriptCode;
}

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
}
</style>

<style lang="less">
</style>