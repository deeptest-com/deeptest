<template>
  <div class="env-main">
    <div class="head no-padding">
      <div class="title dp-bg-white">
        <div class="label ">选择环境</div>

        <div class="content ">
          <a-dropdown-button trigger="click" placement="bottomLeft" class="dp-dropdown">
            <div class="name">
              <a class="more dp-color-text">
                <span v-if="environmentData.id">{{environmentData.name}}</span>
              </a>
            </div>

            <template #icon><DownOutlined /></template>

            <template #overlay>
              <a-menu @click="select" class="select-env-menu">
                <a-menu-item v-for="item in environmentsData" :key="item.id"
                             :class="[{'dp-bg-selected-light':item.id === environmentData.id}]">
                  <span class="menu-item-var">
                    <span class="title">{{item.name}}</span>

                    <span @click.stop="edit(item)" class="act"><EditOutlined /></span>
                    <span @click.stop="remove(item)" class="act"><DeleteOutlined /></span>
                    <span @click.stop="copy(item)" class="act"><CopyOutlined /></span>
                  </span>
                </a-menu-item>
              </a-menu>
            </template>
          </a-dropdown-button>
        </div>
      </div>
    </div>

    <div class="env-var">
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
        <div v-if="environmentData.id" class="envs">
          <div class="env header">
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
          <div v-for="(item, idx) in environmentData.vars" :key="idx" class="env">
            <div class="left">
              <div class="name">
                <a-tooltip class="name" overlayClassName="dp-tip-small">
                  <template #title>{{item.name}}</template>
                  {{item.name}}
                </a-tooltip>
              </div>

              <div class="val">
                <a-tooltip class="val" overlayClassName="dp-tip-small">
                  <template #title>{{item.value}}</template>
                  {{item.value}}
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

        <div v-if="environmentData.vars?.length == 0">
          <Empty></Empty>
        </div>

      </div>
    </div>

    <div class="env-var">
      <div class="body">
        <div v-if="environmentData.id" class="envs">
          <div class="env">
            <div class="left">
              共享变量
            </div>
            <div class="right">
              <span class="dp-link">
                <a-tooltip overlayClassName="dp-tip-small">
                  <template #title>清除</template>
                  <ClearOutlined @click="clearShareVar" class="dp-icon-btn dp-trans-80"/>
                </a-tooltip>
              </span>
              <span class="dp-link">
                <a-tooltip overlayClassName="dp-tip-small">
                  <template #title>帮助</template>
                  <QuestionCircleOutlined class="dp-icon-btn dp-trans-80"/>
                </a-tooltip>
              </span>
            </div>
          </div>

          <div v-for="(item, idx) in validExtractorVariablesData" :key="idx" class="env">
            <div class="left">
              <div class="name">
                <a-tooltip overlayClassName="dp-tip-small">
                  <template #title>{{item.name}}</template>
                  {{item.name}}
                </a-tooltip>
              </div>

              <div class="val">
                <a-tooltip class="val" overlayClassName="dp-tip-small">
                  <template #title>{{item.value}}</template>
                  {{item.value==='extractor_err'? t(item.value+'_short') : item.value}}
                </a-tooltip>
              </div>
            </div>

            <div class="right">
              <DeleteOutlined @click="removeShareVar(item)"  class="dp-icon-btn dp-trans-80" />
            </div>
          </div>
        </div>

        <div v-if="environmentData.vars?.length == 0">
          <Empty></Empty>
        </div>

      </div>
    </div>

    <EnvEdit
        v-if="envEditVisible"
        :modelId="modelId"
        :interfaceId="interfaceData.id"
        :onFinish="envEditFinish"
        :onCancel="envEditCancel"
    />

    <EnvVarEdit
        v-if="envVarEditVisible"
        :model="envVal"
        :environmentId="environmentData.id"
        :onFinish="envVarEditFinish"
        :onCancel="envVarEditCancel"
    />

  </div>
</template>

<script setup lang="ts">
import {computed, defineComponent, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined,ImportOutlined, MoreOutlined, ClearOutlined, PlusOutlined,
  DownOutlined, EditOutlined, DeleteOutlined, CopyOutlined } from '@ant-design/icons-vue';
import {StateType as InterfaceStateType} from "@/views/interface/store";
import {StateType as EnvironmentStateType} from "@/store/environment";
import Empty from "@/components/others/empty.vue";
import {Interface} from "@/views/interface/data";
import EnvEdit from "./edit.vue";
import EnvVarEdit from "./edit-var.vue"
import {StateType as ProjectStateType} from "@/store/project";

    const {t} = useI18n();
    const store = useStore<{ Interface: InterfaceStateType, ProjectGlobal: ProjectStateType, Environment: EnvironmentStateType }>();
    const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);
    const environmentsData = computed<any[]>(() => store.state.Environment.environmentsData);
    const environmentData = computed<any>(() => store.state.Environment.environmentData);
    const validExtractorVariablesData = computed(() => store.state.Interface.validExtractorVariablesData);

    store.dispatch('Environment/listEnvironment')
    if (currProject.value.id)
      store.dispatch('Environment/getEnvironment', {id: 0, projectId: currProject.value.id})

    watch(currProject, () => {
      console.log('watch currProject', currProject.value.id)
      store.dispatch('Environment/getEnvironment', {id: 0, projectId: currProject.value.id})
    }, {deep: false})

    const envEditVisible = ref(false)
    const modelId = ref(0)

    const envVarEditVisible = ref(false)
    const envVal = ref({})

    const select = (val) => {
      console.log('select', val.key)
      store.dispatch('Environment/changeEnvironment', {id: val.key, projectId: currProject.value.id})
    }

    const create = () => {
      console.log('create')
      modelId.value = 0
      envEditVisible.value = true
    }

    const edit = (val) => {
      console.log('edit', val)
      modelId.value = val.id
      envEditVisible.value = true
    }

    const remove = (item) => {
      console.log('remove', item)
      store.dispatch('Environment/removeEnvironment', item)
    }

    const copy = (item) => {
      console.log('copy', item)
      store.dispatch('Environment/copyEnvironment', item)
    }

    const envEditFinish = () => {
      console.log('envEditFinish')
      envEditVisible.value = false
    }
    const envEditCancel = () => {
      console.log('envEditCancel')
      envEditVisible.value = false
    }

    const createVar  = () => {
      console.log('createVar')
      envVarEditVisible.value = true
    }
    const editVar = (item) => {
      console.log('editVar', item)
      envVal.value = item
      envVarEditVisible.value = true
    }
    const removeVar = (item) => {
      console.log('removeVar', item)
      store.dispatch('Environment/removeEnvironmentVar', item.id)
    }
    const clearVar  = () => {
      console.log('clearVar')
      store.dispatch('Environment/clearEnvironmentVar', environmentData.value.id)
    }
    const envVarEditFinish = () => {
      console.log('envVarEditFinish')
      envVal.value = {}
      envVarEditVisible.value = false
    }
    const envVarEditCancel = () => {
      console.log('envVarEditCancel')
      envVal.value = {}
      envVarEditVisible.value = false
    }

    const clearShareVar  = () => {
      console.log('clearShareVar')
      store.dispatch('Interface/clearShareVar', interfaceData.value.id)
    }
    const removeShareVar = (item) => {
      console.log('removeShareVar', item)
      store.dispatch('Interface/removeShareVar', item.id)
    }

</script>

<style lang="less">
.select-env-menu {
  .menu-item-var {
    display: flex;
    width: 200px;
    .title {
      flex: 1;
    }
    .act {
      width: 18px;
    }
  }
}
</style>

<style lang="less" scoped>
.env-main {
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

  .env-var {
    flex: 1;
    display: flex;
    flex-direction: column;

    .body {
      flex: 1;
      overflow-y: auto;

      .btn-wrapper {
        text-align: center;
      }
      .envs {
        padding: 3px 2px;
        .env {
          display: flex;
          padding: 3px 2px 4px 4px;
          line-height: 20px;

          &:first-child {
            border-bottom: 1px solid #eaeaee;
            .right {
              width: 48px;
              .dp-link {
                display: inline-block;
                width: 24px;
              }
            }
          }

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
            .val {
              flex: 3;
              padding-left: 8px;
              overflow: hidden;
              white-space:nowrap;
              text-overflow :ellipsis;
              width: 0;
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
  }

}
</style>