<template>
  <div class="env-main">
    <div class="env-var">
      <div class="body">
        <div class="envs">
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

          <div v-for="(item, idx) in debugData.shareVars" :key="idx" class="env">
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
                  {{item.rightValue==='extractor_err'? t(item.rightValue+'_short') : item.value}}
                </a-tooltip>
              </div>
            </div>

            <div class="right">
              <DeleteOutlined @click="removeShareVar(item)"  class="dp-icon-btn dp-trans-80" />
            </div>
          </div>
          <div v-if="!debugData.shareVars || debugData.shareVars.length===0" class="env">空</div>
        </div>

      </div>
    </div>

    <div class="env-var">
      <div class="body">
        <div class="envs">
          <div class="env header">
            <div class="left">
              环境变量
            </div>
            <div class="right">
              <span class="dp-link">
                <a-tooltip overlayClassName="dp-tip-small">
                  <template #title>帮助</template>
                  <QuestionCircleOutlined class="dp-icon-btn dp-trans-80"/>
                </a-tooltip>
              </span>
            </div>
          </div>

          <div v-for="(item, idx) in debugData.envVars" :key="idx" class="env">
            <div class="left">
              <div class="name">
                <a-tooltip class="name" overlayClassName="dp-tip-small">
                  <template #title>{{item.name}}</template>
                  {{item.name}}
                </a-tooltip>
              </div>

              <div class="val">
                <a-tooltip class="val" overlayClassName="dp-tip-small">
                  <template #title>{{item.localValue}}</template>
                  {{item.localValue}}
                </a-tooltip>
              </div>

              <div class="right"></div>

            </div>
          </div>
          <div v-if="!debugData.envVars || debugData.envVars.length===0" class="env">空</div>
        </div>
      </div>
    </div>

    <div class="env-var">
      <div class="body">
        <div class="envs">
          <div class="env">
            <div class="left">
              全局变量
            </div>
            <div class="right">
              <span class="dp-link">
                <a-tooltip overlayClassName="dp-tip-small">
                  <template #title>帮助</template>
                  <QuestionCircleOutlined class="dp-icon-btn dp-trans-80"/>
                </a-tooltip>
              </span>
            </div>
          </div>

          <div v-for="(item, idx) in debugData.globalEnvVars" :key="idx" class="env">
            <div class="left">
              <div class="name">
                <a-tooltip overlayClassName="dp-tip-small">
                  <template #title>{{item.name}}</template>
                  {{item.name}}
                </a-tooltip>
              </div>

              <div class="val">
                <a-tooltip class="val" overlayClassName="dp-tip-small">
                  <template #title>{{item.localValue}}</template>
                  {{item.localValue}}
                </a-tooltip>
              </div>
            </div>

            <div class="right"></div>
          </div>
          <div v-if="!debugData.globalEnvVars || debugData.globalEnvVars.length===0" class="env">空</div>
        </div>

      </div>
    </div>

    <div class="env-var">
      <div class="body">
        <div class="envs">
          <div class="env">
            <div class="left">
              全局参数
            </div>
            <div class="right">
              <span class="dp-link">
                <a-tooltip overlayClassName="dp-tip-small">
                  <template #title>帮助</template>
                  <QuestionCircleOutlined class="dp-icon-btn dp-trans-80"/>
                </a-tooltip>
              </span>
            </div>
          </div>

          <div v-for="(item, idx) in debugData.globalParamVars" :key="idx" class="env">
            <div class="left">
              <div class="name">
                <a-tooltip overlayClassName="dp-tip-small">
                  <template #title>{{item.name}}</template>
                  {{item.name}}
                </a-tooltip>
              </div>

              <div class="val">
                <a-tooltip class="val" overlayClassName="dp-tip-small">
                  <template #title>{{item.defaultValue}}</template>
                  {{item.defaultValue}}
                </a-tooltip>
              </div>
            </div>

            <div class="right"></div>
          </div>
          <div v-if="!debugData.globalParamVars || debugData.globalParamVars.length===0" class="env">空</div>
        </div>

      </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import {computed, defineComponent, inject, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined,ImportOutlined, MoreOutlined, ClearOutlined, PlusOutlined,
  DownOutlined, EditOutlined, DeleteOutlined, CopyOutlined } from '@ant-design/icons-vue';
import {StateType as EnvironmentStateType} from "@/store/environment";
import Empty from "@/components/others/empty.vue";
import {StateType as ProjectStateType} from "@/store/project";
import {UsedBy} from "@/utils/enum";
const usedBy = inject('usedBy') as UsedBy

const {t} = useI18n();

import {StateType as Debug} from "@/views/component/debug/store";
const store = useStore<{  Debug: Debug, ProjectGlobal: ProjectStateType, Environment: EnvironmentStateType }>();

const debugData = computed<any>(() => store.state.Debug.debugData);
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);

const clearShareVar  = () => {
  console.log('clearShareVar')
  store.dispatch('Debug/clearShareVar', {usedBy: usedBy})
}
const removeShareVar = (item) => {
  console.log('removeShareVar', item)
  store.dispatch('Debug/removeShareVar', {id: item.id})
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
    height: 0;

    .body {
      flex: 1;
      height: 100%;
      overflow-y: auto;

      .btn-wrapper {
        text-align: center;
      }
      .envs {
        padding: 3px 2px 10px 2px;
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
              flex: 3;
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
            text-align: right;
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