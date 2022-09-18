<template>
  <div class="env-main">
    <div class="head no-padding">
      <div class="title">
        <a-dropdown-button trigger="click" placement="bottomLeft">
          <div style="width:138px;">
            <a class="more dp-color-text">
              <span v-if="environmentData.id">{{environmentData.name}}</span>
              <span v-if="!environmentData.id">选择环境</span>
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
    <div class="head">
      <div class="title">
        <span @click="create" class="dp-link">
          <PlusOutlined class="dp-icon-btn dp-trans-80" />
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
        <div class="env">
          <div class="left"></div>
          <div class="right" style="width: 48px;">
            <span @click="clearVar" class="dp-link">
              <ClearOutlined class="dp-icon-btn dp-trans-80" />
            </span>
            <span @click="createVar" class="dp-link">
              <PlusOutlined class="dp-icon-btn dp-trans-80" />
            </span>
          </div>
        </div>
        <div v-for="(item, idx) in environmentData.vars" :key="idx" class="env">
          <div class="left">
            <span class="name">{{item.name}}</span>
            <span class="val" :title="item.value">{{item.value}}</span>
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

<script lang="ts">
import {computed, defineComponent, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined,ImportOutlined, MoreOutlined, ClearOutlined, PlusOutlined,
  DownOutlined, EditOutlined, DeleteOutlined, CopyOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import Empty from "@/components/others/empty.vue";
import {Interface} from "@/views/interface/data";
import EnvEdit from "./edit.vue";
import EnvVarEdit from "./edit-var.vue"

export default defineComponent({
  name: 'RequestEnv',
  components: {
    EnvEdit, EnvVarEdit,
    QuestionCircleOutlined, ImportOutlined, Empty, MoreOutlined, PlusOutlined, ClearOutlined,
    DownOutlined, EditOutlined, DeleteOutlined, CopyOutlined,
  },

  computed: {
  },

  setup(props) {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);
    const environmentsData = computed<any[]>(() => store.state.Interface.environmentsData);
    const environmentData = computed<any[]>(() => store.state.Interface.environmentData);

    const envEditVisible = ref(false)
    const modelId = ref(0)

    const envVarEditVisible = ref(false)
    const envVal = ref({})

    const select = (val) => {
      console.log('select', val.key)
      store.dispatch('Interface/changeEnvironment', val.key)
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

    const remove = (val) => {
      console.log('remove', val)
      store.dispatch('Interface/removeEnvironment', val.id)
    }

    const copy = (val) => {
      console.log('copy', val)
      store.dispatch('Interface/copyEnvironment', val.id)
    }

    const get = (id) => {
      console.log('get', id)
      store.dispatch('Interface/getEnvironment', id)
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
      store.dispatch('Interface/removeEnvironmentVar', item.id)
    }
    const clearVar  = () => {
      console.log('clearVar')
      store.dispatch('Interface/clearEnvironmentVar')
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

    return {
      interfaceData,
      environmentsData,
      environmentData,

      envEditVisible,
      modelId,
      envEditFinish,
      envEditCancel,

      envVarEditVisible,
      envVal,
      envVarEditFinish,
      envVarEditCancel,

      get,
      select,
      create,
      edit,
      remove,
      copy,

      createVar,
      editVar,
      removeVar,
      clearVar,
    }
  }
})

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
    }
    .acts {
      width: 50px;
      text-align: right;
    }
  }
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
        line-height: 16px;
        .left {
          flex: 1;
          display: flex;
          .name {
            margin-left: 0;
            flex: 1;
            overflow: hidden;
            white-space:nowrap;
            text-overflow :ellipsis;
          }
          .val {
            flex: 2;
            padding-left: 8px;
            overflow: hidden;
            white-space:nowrap;
            text-overflow :ellipsis;
            width: 0;
          }
        }
        .right {
          width: 24px;
          text-align: center;
          .more {
            font-weight: bolder;
            font-size: 16px;
          }
        }
      }
    }
  }
}
</style>