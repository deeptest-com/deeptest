<template>
  <div class="pre-condition-main">
    <div class="head">
      <a-row type="flex" class="row">
        <a-col flex="1" class="left">
          <icon-svg type="script" class="icon"  />
          <span>JavaScript代码</span>
        </a-col>

        <a-col flex="100px" class="dp-right">
          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>保存</template>
            <icon-svg type="save" class="dp-icon dp-link-primary dp-icon-large"
                      @click.stop="save" />
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>帮助</template>
            <QuestionCircleOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>清除</template>
            <DeleteOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>全屏</template>
            <FullscreenOutlined @click.stop="openFullscreen()"  class="dp-icon-btn dp-trans-80" />
          </a-tooltip>

        </a-col>
      </a-row>
    </div>

    <div class="content">
      <Script v-if="preConditions.length > 0" :condition="preConditions[0]" />
    </div>

    <FullScreenPopup v-if="preConditions.length > 0 && fullscreen"
                     :visible="fullscreen"
                     :model="preConditions[0]"
                     :onCancel="closeFullScreen"/>
  </div>
</template>

<script setup lang="ts">
import {computed, inject, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined, DeleteOutlined, FullscreenOutlined } from '@ant-design/icons-vue';
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import {UsedBy} from "@/utils/enum";
import IconSvg from "@/components/IconSvg";

import {StateType as Debug} from "@/views/component/debug/store";
import Script from "./conditions-pre/Script.vue";
import FullScreenPopup from "./ConditionPopup.vue";

const store = useStore<{  Debug: Debug }>();
const debugData = computed<any>(() => store.state.Debug.debugData);
const debugInfo = computed<any>(() => store.state.Debug.debugInfo);
const preConditions = computed<any>(() => store.state.Debug.preConditions);

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

const fullscreen = ref(false)

const list = () => {
  console.log('list')
  store.dispatch('Debug/listPreCondition')
}

watch(debugData, (newVal) => {
  console.log('watch debugData')
  list()
}, {immediate: true, deep: true});

const format = (item) => {
  console.log('format', item)
  bus.emit(settings.eventEditorAction, {act: settings.eventTypeFormat})
}

const save = () => {
  console.log('save')
  bus.emit(settings.eventConditionSave, {});
}

const openFullscreen = () => {
  console.log('openFullscreen')
  fullscreen.value = true
}
const closeFullScreen = () => {
  console.log('closeFullScreen')
  fullscreen.value = false
}

</script>

<style lang="less">
.pre-condition-main {
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
.pre-condition-main {
  height: 100%;
  display: flex;
  flex-direction: column;

  .head {
    height: 30px;
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;

    .row {
      .left {
        .icon {
          margin-right: 5px;
        }
      }
    }
  }
  .content {
    flex: 1;
    height: calc(100% - 30px);
    overflow-y: auto;

    display: flex;
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

    .collapse-list {
      height: 100%;
      width: 100%;
      padding: 3px 0;

      .collapse-item {
        width: 100%;
        border: 1px solid #d9d9d9;
        border-bottom: 0;
        border-radius: 2px;

        &:last-child {
          border-radius: 0 0 2px 2px;
          border-bottom: 1px solid #d9d9d9;
        }

        .header {
          height: 38px;
          line-height: 22px;
          padding: 10px;
          background-color: #fafafa;

          display: flex;
          .title {
            flex: 1;
            font-weight: bolder;
          }
          .buttons {
            width: 160px;
            text-align: right;
          }
        }
        .content {
          padding: 16px 10px;
          width: 100%;
        }
      }
    }
  }
}
</style>

<style lang="less" scoped>

</style>
