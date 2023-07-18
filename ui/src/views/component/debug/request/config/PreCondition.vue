<template>
  <div class="pre-condition-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">
          <a-select size="small" :style="{width:'116px'}" :bordered="true"
                    v-model:value="conditionType">
            <!--            <a-select-option key="" value="">
                          控制器类型
                        </a-select-option>-->

            <a-select-option v-for="item in conditionTypes" :key="item.value" :value="item.value">
              {{ t(item.label) }}
            </a-select-option>
          </a-select> &nbsp;

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
                 :list="preConditions"
                 handle=".handle"
                 @end="handleDrop">
        <template #item="{ element }">

          <div class="collapse-item">
            <div class="header">
              <div @click.stop="expand(element)" class="title dp-link">
                <MenuOutlined class="handle dp-drag" /> &nbsp;

                {{ t(element.entityType) }} &nbsp;
                {{ element.desc }} -
                {{activeItem || 0}},{{element.id}}
              </div>
              <div class="buttons">
                <ClearOutlined v-if="activeItem === +element.id && element.entityType === ConditionType.script"
                               @click.stop="format(element)"  class="dp-icon-btn dp-trans-80" />
                &nbsp;

                <CheckCircleOutlined v-if="!element.disabled" @click.stop="disable(element)"
                                     class="dp-icon-btn dp-trans-80 dp-color-pass" />
                <CloseCircleOutlined v-if="element.disabled" @click.stop="disable(element)"
                                     class="dp-icon-btn dp-trans-80" />

                <DeleteOutlined @click.stop="remove(element)"  class="dp-icon-btn dp-trans-80" />
                &nbsp;

                <RightOutlined v-if="activeItem !== element.id"
                               @click.stop="expand(element)"  class="dp-icon-btn dp-trans-80" />
                <DownOutlined v-if="activeItem === element.id"
                              @click.stop="expand(element)"  class="dp-icon-btn dp-trans-80" />
              </div>
            </div>

            <div class="content" v-if="activeItem === +element.id">
              <Script v-if="element.entityType === ConditionType.script"
                      :condition="element" />
            </div>

          </div>

        </template>
      </draggable>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, inject, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined, CheckCircleOutlined, DeleteOutlined,
  ClearOutlined, MenuOutlined, RightOutlined, DownOutlined, CloseCircleOutlined } from '@ant-design/icons-vue';
import draggable from 'vuedraggable'
import {ConditionType, PreConditionType, UsedBy} from "@/utils/enum";
import {EnvDataItem} from "@/views/project-settings/data";

import {StateType as Debug} from "@/views/component/debug/store";
import {getEnumSelectItems} from "@/views/scenario/service";
import Script from "./pre-conditions/Script.vue";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import {confirmToDelete} from "@/utils/confirm";

const store = useStore<{  Debug: Debug }>();
const debugData = computed<any>(() => store.state.Debug.debugData);
const debugInfo = computed<any>(() => store.state.Debug.debugInfo);
const preConditions = computed<any>(() => store.state.Debug.preConditions);

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

const activeItem = ref(0)

const conditionType = ref(PreConditionType.script)
const conditionTypes = ref(getEnumSelectItems(PreConditionType))

const expand = (item) => {
  console.log('expand', item.id)

  if (activeItem.value === item.id) {
    activeItem.value = 0
  } else {
    activeItem.value = item.id
  }
}

const list = () => {
  console.log('list')
  store.dispatch('Debug/listPreCondition')
}
list()

const create = () => {
  console.log('create', conditionType.value)
  store.dispatch('Debug/createPreCondition', {
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
  store.dispatch('Debug/disablePreCondition', item.id)
}
const remove = (item) => {
  console.log('remove', item)

  confirmToDelete(`确定删除该${t(item.entityType)}？`, '', () => {
    store.dispatch('Debug/removePreCondition', item.id)
  })
}

function handleDrop(_e: any) {
  const envIdList = preConditions.value.map((e: EnvDataItem) => {
    return e.id;
  })

  store.dispatch('Debug/movePreCondition', {
    data: envIdList,
    info: debugInfo.value,
  })
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
