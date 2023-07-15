<template>
  <div class="post-condition-main">
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
      <draggable tag="div" item-key="name" class="collapse-list"
                 :list="postConditions"
                 handle=".handle"
                 @end="handleDrop">
        <template #item="{ element }">

          <div class="collapse-item">
            <div class="header">
              <div class="title">
                <MenuOutlined class="handle" /> &nbsp;

                {{ t(element.entityType) }} &nbsp;
                {{ element.desc }} -
                {{activeItem || 0}},{{element.id}}
              </div>
              <div class="buttons">
                <CheckCircleOutlined @click.stop="disable(element)"  class="dp-icon-btn dp-trans-80" />
                <DeleteOutlined @click.stop="remove(element)"  class="dp-icon-btn dp-trans-80" />
                &nbsp;&nbsp;
                <RightOutlined v-if="activeItem !== element.id"
                               @click.stop="expand(element)"  class="dp-icon-btn dp-trans-80" />
                <DownOutlined v-if="activeItem === element.id"
                              @click.stop="expand(element)"  class="dp-icon-btn dp-trans-80" />
              </div>
            </div>

            <div class="content" v-if="activeItem === +element.id">
                <Checkpoint v-if="element.entityType === ConditionType.checkpoint"
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
  ClearOutlined, MenuOutlined, RightOutlined, DownOutlined } from '@ant-design/icons-vue';
import draggable from 'vuedraggable'
import {ConditionType, UsedBy} from "@/utils/enum";

import {StateType as Debug} from "@/views/component/debug/store";
import {getEnumSelectItems} from "@/views/scenario/service";
import Checkpoint from "./post-conditions/Checkpoint.vue";
import {EnvDataItem} from "@/views/project-settings/data";

const store = useStore<{  Debug: Debug }>();
const debugData = computed<any>(() => store.state.Debug.debugData);
const debugInfo = computed<any>(() => store.state.Debug.debugInfo);
const postConditions = computed<any>(() => store.state.Debug.postConditions);

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

const activeItem = ref(0)

const conditionType = ref(ConditionType.checkpoint)
const conditionTypes = ref(getEnumSelectItems(ConditionType))

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
  store.dispatch('Debug/listPostCondition')
}
list()

const create = () => {
  console.log('create', conditionType.value)
  store.dispatch('Debug/createPostCondition', {
    entityType: conditionType.value,
    ...debugInfo.value,
  })
}

const disable = (item) => {
  console.log('disable', item)
}
const remove = (item) => {
  console.log('remove', item)
}

function handleDrop(_e: any) {
  const envIdList = postConditions.value.map((e: EnvDataItem) => {
    return e.id;
  })

  store.dispatch('Debug/movePostCondition', {
    data: envIdList,
    info: debugInfo.value,
  })
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
