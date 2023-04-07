<template>
  <div class="datapool-main">
    <div class="datapool-var">
      <div class="head">
        <div class="title">
          <span class="dp-link">
            <span>数据池</span>
          </span>
        </div>
        <div class="acts">
          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>新建</template>
            <PlusOutlined @click="create" class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>帮助</template>
            <QuestionCircleOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>
        </div>
      </div>
      <div class="body">
        <div class="datapools">
          <div v-for="(item, idx) in datapoolsData" :key="idx" class="datapool">
            <div class="left">
              <div class="name">
                <a-tooltip class="name" overlayClassName="dp-tip-small">
                  <template #title>{{item.name}}</template>
                  {{item.name}}
                </a-tooltip>
              </div>
            </div>

            <div class="right">
              <a-dropdown>
                <a class="more dp-color-text" @click.prevent>
                  <MoreOutlined />
                </a>
                <template #overlay>
                  <a-menu>
                    <a-menu-item @click="edit(item)" key="edit">编辑</a-menu-item>
                    <a-menu-item @click="remove(item)" key="edit">删除</a-menu-item>
                  </a-menu>
                </template>
              </a-dropdown>
            </div>
          </div>
        </div>

        <div v-if="datapoolsData?.length == 0">
          <Empty></Empty>
        </div>

      </div>
    </div>

    <DatapoolEdit
        v-if="datapoolEditVisible"
        :modelId="modelId"
        :onFinish="datapoolEditFinish"
        :onCancel="datapoolEditCancel"
    />

  </div>
</template>

<script setup lang="ts">
import {computed, defineComponent, inject, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined,ImportOutlined, MoreOutlined, ClearOutlined, PlusOutlined} from '@ant-design/icons-vue';

import {StateType as DatapoolStateType} from "@/store/datapool";
import Empty from "@/components/others/empty.vue";
import DatapoolEdit from "./edit.vue";

import {StateType as ProjectStateType} from "@/store/project";
import {UsedBy} from "@/utils/enum";
const usedBy = inject('usedBy') as UsedBy

const {t} = useI18n();

const store = useStore<{ ProjectGlobal: ProjectStateType, Datapool: DatapoolStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const datapoolsData = computed<any[]>(() => store.state.Datapool.datapoolsData);

if (currProject.value.id)
  store.dispatch('Datapool/listDatapool')

watch(currProject, () => {
  console.log('watch currProject', currProject.value.id)
  store.dispatch('Datapool/listDatapool')
}, {deep: false})

const datapoolEditVisible = ref(false)
const modelId = ref(0)

const create = () => {
  console.log('create')
  modelId.value = 0
  datapoolEditVisible.value = true
}

const edit = (val) => {
  console.log('edit', val)
  modelId.value = val.id
  datapoolEditVisible.value = true
}

const remove = (item) => {
  console.log('remove', item)
  store.dispatch('Datapool/removeDatapool', item)
}

const datapoolEditFinish = () => {
  console.log('datapoolEditFinish')
  datapoolEditVisible.value = false
}
const datapoolEditCancel = () => {
  console.log('datapoolEditCancel')
  datapoolEditVisible.value = false
}

</script>

<style lang="less" scoped>
.datapool-main {
  display: flex;
  flex-direction: column;
  height: 100%;

  .head {
    padding: 0 3px;
    height: 32px;
    line-height: 32px;
    border-bottom: 1px solid #d9d9d9;
    display: flex;
    &.no-padding {
      padding: 0;
    }
    .title {
      flex: 1;
      display: flex;

      .label {
        padding: 0 5px;
        width: 68px;
      }
      .content {
        flex: 1;
      }

    }
    .acts {
      width: 50px;
      text-align: right;
    }
  }

  .datapools {
    .datapool {
      display: flex;
      padding: 3px 2px 4px 4px;
      line-height: 20px;
      border-bottom: 1px solid #eaeaee;

      .left {
        flex: 1;
        display: flex;

        .name {
          margin-left: 0;
          flex: 2;
          overflow: hidden;
          white-space:nowrap;
          text-overflow :ellipsis;
        }
      }
      .right {
        text-align: center;
        width: 24px;
        .more {
          display: inline-block;
          font-weight: bolder;
          font-size: 16px;
        }
      }
    }
  }
}
</style>