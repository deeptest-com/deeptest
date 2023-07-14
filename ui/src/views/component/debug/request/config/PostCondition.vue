<template>
  <div class="post-condition-main">
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
      <a-collapse v-model:activeKey="collActiveKey">
        <a-collapse-panel key="1" header="This is panel header 1">
          <p>1</p>
        </a-collapse-panel>

        <a-collapse-panel key="2" header="This is panel header 2" :disabled="false">
          <p>2</p>
        </a-collapse-panel>
        <a-collapse-panel key="2" header="This is panel header 3" :disabled="false">
          <p>2</p>
        </a-collapse-panel>
        <a-collapse-panel key="2" header="This is panel header 4" :disabled="false">
          <p>2</p>
        </a-collapse-panel>
        <a-collapse-panel key="2" header="This is panel header 5" :disabled="false">
          <p>2</p>
        </a-collapse-panel>
        <a-collapse-panel key="2" header="This is panel header 6" :disabled="false">
          <p>2</p>
        </a-collapse-panel>
        <a-collapse-panel key="2" header="This is panel header 7" :disabled="false">
          <p>2</p>
        </a-collapse-panel>
        <a-collapse-panel key="2" header="This is panel header 8" :disabled="false">
          <p>2</p>
        </a-collapse-panel>
      </a-collapse>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, inject, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined, DeleteOutlined, ClearOutlined } from '@ant-design/icons-vue';
import {UsedBy} from "@/utils/enum";

import {StateType as Debug} from "@/views/component/debug/store";
const store = useStore<{  Debug: Debug }>();
const debugData = computed<any>(() => store.state.Debug.debugData);

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

const collActiveKey = ref('')

</script>

<style lang="less">
.post-condition-main {
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

<style lang="less" scoped>
.post-condition-main {
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

<style lang="less" scoped>

</style>
