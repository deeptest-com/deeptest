<template>
  <div class="post-condition-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">
          <a-select size="small" :style="{width:'116px'}" :bordered="true"
                    v-model:value="conditionType">
            <template v-for="item in conditionTypes" :key="item.value">
              <a-select-option v-if="item.value !== ConditionType.checkpoint" :value="item.value">
                {{ t(item.label) }}
              </a-select-option>
            </template>
          </a-select> &nbsp;

          <a-button @click="create" type="primary" size="small">添加处理</a-button>
        </a-col>

        <a-col flex="100px" class="dp-right">
          <Tips section="i" title="帮助" />
        </a-col>
      </a-row>
    </div>

    <div class="content">
      <draggable tag="div" item-key="name" class="collapse-list"
                 :list="postConditions || []"
                 handle=".handle"
                 @end="move">
        <template #item="{ element }">

          <div :class="[activePostCondition.id === +element.id ? 'active' : '']" class="collapse-item">
            <div class="header">
              <div @click.stop="expand(element)" class="title dp-link dp-ellipsis">
                <icon-svg class="handle dp-drag icon" type="move" />

                <icon-svg v-if="element.entityType === ConditionType.extractor"
                          type="variable"
                          class="icon variable" />
                <icon-svg v-if="element.entityType === ConditionType.checkpoint"
                          type="checkpoint"
                          class="icon"  />
                <icon-svg v-if="element.entityType === ConditionType.script"
                          type="script"
                          class="icon"  />

                <span v-html="element.desc || t(element.entityType)"></span>
              </div>
              <div class="buttons">
                <icon-svg class="icon dp-link-primary dp-icon-large" type="save"
                          title="保存"
                          v-if="activePostCondition.id === element.id"
                          @click.stop="save(element)" />

                <ClearOutlined v-if="activePostCondition.id === +element.id && element.entityType === ConditionType.script"
                               @click.stop="format(element)"
                               class="dp-icon-btn dp-trans-80"
                               title="格式化"/>&nbsp;

                <CheckCircleOutlined v-if="!element.disabled" @click.stop="disable(element)"
                                     class="dp-icon-btn dp-trans-80 dp-color-pass" title="启用" />
                <CloseCircleOutlined v-if="element.disabled" @click.stop="disable(element)"
                                     class="dp-icon-btn dp-trans-80" title="禁用" />
                <DeleteOutlined @click.stop="remove(element)"
                                class="dp-icon-btn dp-trans-80" title="删除" />

                <FullscreenOutlined v-if="activePostCondition.id === element.id"
                                    @click.stop="openFullscreen(element)"
                                    class="dp-icon-btn dp-trans-80" title="全屏" />

                <RightOutlined v-if="activePostCondition.id !== element.id"
                               @click.stop="expand(element)"
                               class="dp-icon-btn dp-trans-80" />
                <DownOutlined v-if="activePostCondition.id === element.id"
                              @click.stop="expand(element)"
                              class="dp-icon-btn dp-trans-80" />
              </div>
            </div>

            <div class="content" v-if="activePostCondition.id === +element.id">
              <Extractor v-if="element.entityType === ConditionType.extractor"
                         :condition="activePostCondition"
                          :finish="list"/>

              <Checkpoint v-if="element.entityType === ConditionType.checkpoint"
                          :condition="activePostCondition"
                          :finish="list"/>

              <Script v-if="element.entityType === ConditionType.script"
                      :condition="activePostCondition"
                      :finish="list"/>
            </div>
          </div>

        </template>
      </draggable>
    </div>

    <FullScreenPopup v-if="fullscreen"
                     :visible="fullscreen"
                     :model="activePostCondition"
                     :onCancel="closeFullScreen" />
  </div>
</template>

<script setup lang="ts">
import {computed, inject, ref, watch, getCurrentInstance, ComponentInternalInstance} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { CheckCircleOutlined, DeleteOutlined,
  ClearOutlined, RightOutlined,
  DownOutlined, CloseCircleOutlined, FullscreenOutlined } from '@ant-design/icons-vue';
import draggable from 'vuedraggable'
import Tips from "@/components/Tips/index.vue";
import {ConditionType, UsedBy} from "@/utils/enum";
import {EnvDataItem} from "@/views/project-settings/data";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import {confirmToDelete} from "@/utils/confirm";
import {StateType as Debug} from "@/views/component/debug/store";
import {getEnumSelectItems} from "@/views/scenario/service";
import IconSvg from "@/components/IconSvg";

import Extractor from "./conditions-post/Extractor.vue";
import Checkpoint from "./conditions-post/Checkpoint.vue";
import Script from "./conditions-post/Script.vue";
import Cookie from "./conditions-post/Cookie.vue";
import FullScreenPopup from "./ConditionPopup.vue";

const store = useStore<{  Debug: Debug }>();
const debugData = computed<any>(() => store.state.Debug.debugData);
const debugInfo = computed<any>(() => store.state.Debug.debugInfo);
const postConditions = computed<any>(() => store.state.Debug.postConditions);
const activePostCondition = computed<any>(() => store.state.Debug.activePostCondition);

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

const fullscreen = ref(false)

const conditionType = ref(ConditionType.extractor)
const conditionTypes = ref(getEnumSelectItems(ConditionType))

const expand = (item) => {
  console.log('expand', item)
  store.commit('Debug/setActivePostCondition', item)
}

const list = () => {
  console.log('list')
  store.dispatch('Debug/listPostCondition')
}

watch(debugData, (newVal) => {
  console.log('watch debugData')
  list()
}, {immediate: true, deep: true});

const create = () => {
  console.log('create', conditionType.value)
  store.dispatch('Debug/createPostCondition', {
    entityType: conditionType.value,
    ...debugInfo.value,
  })
}

const format = (item) => {
  console.log('format', item)
  bus.emit(settings.eventEditorAction, {act: settings.eventTypeFormat})
}
const disable = (item) => {
  console.log('disable', item)
  store.dispatch('Debug/disablePostCondition', item)
}
const remove = (item) => {
  console.log('remove', item)

  confirmToDelete(`确定删除该${t(item.entityType)}？`, '', () => {
    store.dispatch('Debug/removePostCondition', item)
  })
}
function move(_e: any) {
  const envIdList = postConditions.value.map((e: EnvDataItem) => {
    return e.id;
  })

  store.dispatch('Debug/movePostCondition', {
    data: envIdList,
    info: debugInfo.value,
    entityType: '',
  })
}

const save = (item) => {
  console.log('save', item)
  bus.emit(settings.eventConditionSave, {});
}

const openFullscreen = (item) => {
  console.log('openFullscreen', item)
  fullscreen.value = true
}
const closeFullScreen = (item) => {
  console.log('closeFullScreen', item)
  fullscreen.value = false
}

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
  display: flex;
  flex-direction: column;

  .head {
    height: 30px;
    padding: 2px 3px;
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
      padding: 0;

      .collapse-item {
        margin: 4px;
        border-radius: 5px;
        border: 1px solid #d9d9d9;

        &.active {
          border: 1px solid #1890ff;
        }

        .header {
          height: 28px;
          padding: 3px;
          background-color: #fafafa;
          border-radius: 5px;

          display: flex;
          .title {
            flex: 1;

            .icon {
              margin-right: 3px;
              &.variable {
                font-size: 16px;
                vertical-align: -0.2em;
              }
            }
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
