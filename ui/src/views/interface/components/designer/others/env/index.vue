<template>
  <div class="env-main">
    <div class="head">
      <div class="title">
        <a-select :value="''" size="small" style="width: 100%;">
          <a-select-option value="">选择环境</a-select-option>
        </a-select>
      </div>
      <div class="acts">
        <a-tooltip overlayClassName="dp-tip-small">
          <template #title>帮助</template>
          <QuestionCircleOutlined class="dp-icon-btn"/>
        </a-tooltip>

        <a-tooltip overlayClassName="dp-tip-small">
          <template #title>导入/导出</template>
          <ImportOutlined class="dp-icon-btn" />
        </a-tooltip>
      </div>
    </div>

    <div class="body">
      <div v-if="true" class="envs">
        <div class="env">
          <div class="left">
            <EnvironmentOutlined />
            <span class="name">全局</span>
          </div>
          <div class="right">
            <a-dropdown>
              <template #overlay>
                <a-menu>
                  <a-menu-item key="edit">编辑</a-menu-item>
                  <a-menu-item key="edit">删除</a-menu-item>
                  <a-menu-item key="copy">复制</a-menu-item>
                </a-menu>
              </template>
              <a class="ant-dropdown-link more" @click.prevent>
                <MoreOutlined />
              </a>
            </a-dropdown>
          </div>
        </div>
        <div class="env">
          <div class="left">
            <EnvironmentOutlined />
            <span class="name">测试环境</span>
          </div>
          <div class="right">
            <a-dropdown>
              <template #overlay>
                <a-menu>
                  <a-menu-item key="edit">编辑</a-menu-item>
                  <a-menu-item key="edit">删除</a-menu-item>
                  <a-menu-item key="copy">复制</a-menu-item>
                </a-menu>
              </template>
              <a class="ant-dropdown-link more" @click.prevent>
                <MoreOutlined />
              </a>
            </a-dropdown>
          </div>
        </div>
      </div>

      <div v-if="true">
        <Empty></Empty>
        <div class="btn-wrapper">
          <a-button type="link">新建环境</a-button>
        </div>
      </div>

    </div>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined,ImportOutlined, MoreOutlined, EnvironmentOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import Empty from "@/components/others/empty.vue";

interface RequestEnvSetupData {
  modelData: ComputedRef;
}

export default defineComponent({
  name: 'RequestEnv',
  components: {
    QuestionCircleOutlined, ImportOutlined, Empty, MoreOutlined, EnvironmentOutlined
  },

  computed: {
  },

  setup(props): RequestEnvSetupData {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const modelData = computed<any>(() => store.state.Interface.modelResult);

    return {
      modelData,
    }
  }
})

</script>

<style lang="less" scoped>
.env-main {
  height: 100%;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
    display: flex;
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
        padding: 3px 2px 2px 8px;
        line-height: 16px;
        .left {
          flex: 1;
          .name {
            margin-left: 8px;
          }
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