<template>
  <div class="env-main">
    <div class="head no-padding">
      <div class="title">
        <a-select @change="select" v-model:value="environmentData.id" style="width: 100%;">
          <a-select-option value="">选择环境</a-select-option>
          <a-select-option v-for="(item, idx) in environmentsData" :value="item.id" :key="idx">
            {{ item.name }}
          </a-select-option>
        </a-select>
      </div>
    </div>
    <div class="head">
      <div class="title">
        <span @click="create" class="dp-link">
          <PlusOutlined class="dp-icon-btn dp-trans-80" />
          <span>新建</span>
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
      <div class="envs">
        <div v-for="(item, idx) in environmentData.vars" :key="idx" class="env">
          <div class="left">
            <EnvironmentOutlined class="dp-icon-btn dp-trans-80" />
            <span class="name">{{item.name}}</span>
          </div>
          <div class="right">
            <a-dropdown>
              <template #overlay>
                <a-menu>
                  <a-menu-item @click="edit(item)" key="edit">编辑</a-menu-item>
                  <a-menu-item @click="remove(item)" key="edit">删除</a-menu-item>
                  <a-menu-item @click="copy(item)" key="copy">复制</a-menu-item>
                </a-menu>
              </template>
              <a class="more dp-color-text" @click.prevent>
                <MoreOutlined />
              </a>
            </a-dropdown>
          </div>
        </div>
      </div>

      <div v-if="true">
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

  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined,ImportOutlined, MoreOutlined, EnvironmentOutlined, PlusOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import Empty from "@/components/others/empty.vue";
import {Interface} from "@/views/interface/data";
import EnvEdit from "./edit.vue";

export default defineComponent({
  name: 'RequestEnv',
  components: {
    QuestionCircleOutlined, ImportOutlined, Empty, MoreOutlined, EnvironmentOutlined, PlusOutlined,
    EnvEdit,
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

    const select = (val) => {
      console.log('select', val)
      store.dispatch('Interface/changeEnvironment', val)
    }

    const create = () => {
      console.log('create')
      modelId.value = 0
      envEditVisible.value = true
    }

    const edit = (val) => {
      console.log('edit', val)
      modelId.value = val
      envEditVisible.value = true
    }

    const remove = (val) => {
      console.log('removeHistory', val)
    }

    const copy = (val) => {
      console.log('edit', val)
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

    return {
      interfaceData,
      environmentsData,
      environmentData,

      envEditVisible,
      modelId,
      envEditFinish,
      envEditCancel,

      select,
      create,
      edit,
      remove,
      copy,

      get,
    }
  }
})

</script>

<style lang="less" scoped>
.env-main {
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
    }
    .acts {
      width: 50px;
      text-align: right;
    }
  }
  .body {
    height: calc(100% - 30px);
    overflow-y: hidden;

    .btn-wrapper {
      text-align: center;
    }
    .envs {
      .env {
        display: flex;
        padding: 3px 2px 4px 4px;
        line-height: 16px;
        .left {
          flex: 1;
          .name {
            margin-left: 0;
          }
          cursor: pointer;
        }
        .right {
          width: 20px;
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