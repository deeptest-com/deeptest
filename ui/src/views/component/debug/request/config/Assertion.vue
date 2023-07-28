<template>
  <div class="post-condition-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">
          <a-button @click="create" size="small">新建</a-button>
        </a-col>

        <a-col flex="100px" class="dp-right">
          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>帮助</template>
            <QuestionCircleOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>清除</template>
            <DeleteOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>
        </a-col>
      </a-row>
    </div>

    <div class="content">
      <draggable tag="div" item-key="name" class="collapse-list"
                 :list="assertionConditions || []"
                 handle=".handle"
                 @end="move">
        <template #item="{ element }">

          <div class="collapse-item">
            <div class="header">
              <div @click.stop="expand(element)" class="title dp-link">
                <icon-svg class="handle dp-drag icon" type="move"  />

                <icon-svg v-if="element.entityType === ConditionType.extractor"
                          type="variable"
                          class="icon variable" />
                <icon-svg v-if="element.entityType === ConditionType.checkpoint"
                          type="checkpoint"
                          class="icon"  />
                <icon-svg v-if="element.entityType === ConditionType.script"
                          type="script"
                          class="icon"  />

                {{ element.desc || t(element.entityType) }}
              </div>
              <div class="buttons">
                <ClearOutlined v-if="activeItem.id === +element.id && element.entityType === ConditionType.script"
                               @click.stop="format(element)"  class="dp-icon-btn dp-trans-80" />&nbsp;

                <CheckCircleOutlined v-if="!element.disabled" @click.stop="disable(element)"
                                     class="dp-icon-btn dp-trans-80 dp-color-pass" />
                <CloseCircleOutlined v-if="element.disabled" @click.stop="disable(element)"
                                     class="dp-icon-btn dp-trans-80" />
                <DeleteOutlined @click.stop="remove(element)"  class="dp-icon-btn dp-trans-80" />

                <SaveOutlined class="dp-icon-btn dp-trans-80" title="保存"
                              v-if="activeItem.id === element.id"
                              @click.stop="save(element)" />
                <FullscreenOutlined class="dp-icon-btn dp-trans-80"
                                    v-if="activeItem.id === element.id"
                                    @click.stop="openFullscreen(element)" />

                <RightOutlined v-if="activeItem.id !== element.id"
                               @click.stop="expand(element)"  class="dp-icon-btn dp-trans-80" />
                <DownOutlined v-if="activeItem.id === element.id"
                              @click.stop="expand(element)"  class="dp-icon-btn dp-trans-80" />
              </div>
            </div>

            <div class="content" v-if="activeItem.id === +element.id">
              <Checkpoint v-if="element.entityType === ConditionType.checkpoint"
                          :condition="element" />
            </div>
          </div>

        </template>
      </draggable>
    </div>

    <FullScreenPopup v-if="fullscreen"
                     :visible="fullscreen"
                     :model="activeItem"
                     :onCancel="closeFullScreen"/>
  </div>
</template>

<script setup lang="ts">
import {computed, inject, ref, watch, getCurrentInstance, ComponentInternalInstance} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined, CheckCircleOutlined, DeleteOutlined,
  ClearOutlined, MenuOutlined, RightOutlined,
  DownOutlined, CloseCircleOutlined, FullscreenOutlined, SaveOutlined } from '@ant-design/icons-vue';
import draggable from 'vuedraggable'
import {ConditionType, UsedBy} from "@/utils/enum";
import {EnvDataItem} from "@/views/project-settings/data";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import {confirmToDelete} from "@/utils/confirm";
import {StateType as Debug} from "@/views/component/debug/store";
import {getEnumSelectItems} from "@/views/scenario/service";
import IconSvg from "@/components/IconSvg";

import Checkpoint from "./conditions-post/Checkpoint.vue";
import FullScreenPopup from "./ConditionPopup.vue";

const store = useStore<{  Debug: Debug }>();
const debugData = computed<any>(() => store.state.Debug.debugData);
const debugInfo = computed<any>(() => store.state.Debug.debugInfo);
const assertionConditions = computed<any>(() => store.state.Debug.assertionConditions);

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

const fullscreen = ref(false)
const activeItem = ref({} as any)

const expand = (item) => {
  console.log('expand', item)

  if (activeItem.value.id === item.id) {
    activeItem.value = {}
  } else {
    activeItem.value = item
  }
}

const list = () => {
  console.log('list')
  store.dispatch('Debug/listAssertionCondition')
}

watch(debugData, (newVal) => {
  console.log('watch debugData')
  list()
}, {immediate: true, deep: true});

const create = () => {
  console.log('create', ConditionType.checkpoint)
  store.dispatch('Debug/createPostCondition', {
    entityType: ConditionType.checkpoint,
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
  const envIdList = assertionConditions.value.map((e: EnvDataItem) => {
    return e.id;
  })

  store.dispatch('Debug/movePostCondition', {
    data: envIdList,
    info: debugInfo.value,
    entityType: ConditionType.checkpoint,
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
    border-bottom: 1px solid #d9d9d9;
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
        border: 1px solid #d9d9d9;
        border-radius: 5px;

        .header {
          height: 28px;
          padding: 3px;
          background-color: #fafafa;
          border-radius: 5px;

          display: flex;
          .title {
            flex: 1;
            font-weight: bolder;

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

<style lang="less" scoped>

</style>
