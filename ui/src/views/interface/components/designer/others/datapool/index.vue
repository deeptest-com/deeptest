<template>
  <div class="datapool-main">
    <div class="datapool-var">
      <div class="head">
        <div class="title">
        <span @click="create" class="dp-link">
          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>新建环境</template>
            <PlusOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>
        </span>
        </div>
        <div class="acts">
          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>帮助</template>
            <QuestionCircleOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>导入/导出</template>
            <ImportOutlined class="dp-icon-btn dp-trans-60" />
          </a-tooltip>
        </div>
      </div>
      <div class="body">
        <div v-if="datapoolData.id" class="datapools">
          <div class="datapool header">
            <div class="left">
              环境变量
            </div>
            <div class="right" style="width: 48px;">
            <span @click="clearVar" class="dp-link">
              <a-tooltip overlayClassName="dp-tip-small">
                <template #title>清除变量</template>
                <ClearOutlined class="dp-icon-btn dp-trans-80"/>
              </a-tooltip>
            </span>
              <span @click="createVar" class="dp-link">
              <a-tooltip overlayClassName="dp-tip-small">
                <template #title>添加变量</template>
                <PlusOutlined class="dp-icon-btn dp-trans-80"/>
              </a-tooltip>
            </span>
            </div>
          </div>

          <div v-for="(item, idx) in datapoolData.vars" :key="idx" class="datapool">
            <div class="left">
              <div class="name">
                <a-tooltip class="name" overlayClassName="dp-tip-small">
                  <template #title>{{item.name}}</template>
                  {{item.name}}
                </a-tooltip>
              </div>

              <div class="val">
                <a-tooltip class="val" overlayClassName="dp-tip-small">
                  <template #title>{{item.rightValue}}</template>
                  {{item.rightValue}}
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
                    <a-menu-item @click="editVar(item)" key="edit">编辑</a-menu-item>
                    <a-menu-item @click="removeVar(item)" key="edit">删除</a-menu-item>
                  </a-menu>
                </template>
              </a-dropdown>
            </div>
          </div>
        </div>

        <div v-if="datapoolData.vars?.length == 0">
          <Empty></Empty>
        </div>

      </div>
    </div>

    <EnvEdit
        v-if="datapoolEditVisible"
        :modelId="modelId"
        :interfaceId="interfaceData.id"
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
import EnvEdit from "./edit.vue";

import {StateType as ProjectStateType} from "@/store/project";
import {UsedBy} from "@/utils/enum";
const usedBy = inject('usedBy') as UsedBy

const {t} = useI18n();
const store = useStore<{ ProjectGlobal: ProjectStateType, Datapool: DatapoolStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);

const datapoolsData = computed<any[]>(() => store.state.Datapool.datapoolData);

store.dispatch('Datapool/listDatapool')
if (currProject.value.id)
  store.dispatch('Datapool/getDatapool', {id: 0, projectId: currProject.value.id})

watch(currProject, () => {
  console.log('watch currProject', currProject.value.id)
  store.dispatch('Datapool/getDatapool', {id: 0, projectId: currProject.value.id})
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

}
</style>