<template>
  <div class="history-main">
    <div class="head">
      <div class="title">
        执行历史
      </div>
      <div class="acts">
        <CloseOutlined @click="close" class="dp-icon-btn dp-trans-80"/>
      </div>
    </div>

    <div class="body">
      <div class="histories">
        <div v-for="(item, idx) in invocationsData" :key="idx"
             @mouseover="mouseOver"
             @mouseout="mouseLeave"
             class="history dp-link">
          <div class="left">
            <span @click="getRequestAsInterface(item.id)" title="点击加载历史请求数据">{{item.name}}</span>
          </div>
          <div class="right">
            <span @click="removeHistory(item.id)" class="link"><DeleteOutlined /></span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, inject} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { DeleteOutlined, CloseOutlined } from '@ant-design/icons-vue';
import {UsedBy} from "@/utils/enum";

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

import {StateType as Debug} from "@/views/component/debug/store";
import {getRightTabPanelPosition} from "@/utils/dom";
const store = useStore<{  Debug: Debug }>();

const debugData = computed<any>(() => store.state.Debug.debugData);

const invocationsData = computed<any[]>(() => store.state.Debug.invocationsData);

const props = defineProps({
  onClose: {
    type: Function,
    required: false,
  },
})

store.dispatch('Debug/listInvocation')
store.dispatch('Debug/getLastInvocationResp')

const getRequestAsInterface = (id) => {
  store.dispatch('Debug/getInvocationAsInterface', id)
}

const removeHistory = (id) => {
  console.log('removeHistory', id)
  store.dispatch('Debug/removeInvocation', id)
}

const close = () => {
  if (props.onClose) props.onClose()
}

const mouseOver = (event) => {
  // console.log('mouseOver', event)
  event.currentTarget.querySelector(".link").style.display = 'block'
}
const mouseLeave = (event) => {
  // console.log('mouseLeave', event);
  event.currentTarget.querySelector(".link").style.display = 'none'
}

</script>

<style lang="less" scoped>
.history-main {
  height: 100%;

  .head {
    padding: 0 5px 5px 5px;
    border-bottom: 1px solid #d9d9d9;

    line-height: 32px;
    display: flex;
    .title {
      flex: 1;
      font-weight: bolder;
    }
    .acts {
      width: 50px;
      text-align: right;
    }
  }

  .body {
    flex: 1;
    height: calc(100% - 42px);
    overflow-y: auto;

    .btn-wrapper {
      text-align: center;
    }
    .histories {
      padding: 3px;

      .history {
        display: flex;
        padding: 6px;
        line-height: 16px;
        border-bottom: 1px solid #eaeaee;

        .left {
          flex: 1;
          .name {
            margin-left: 0;
            cursor: pointer;
          }
        }
        .right {
          width: 16px;
          cursor: pointer;
          .link {
            display: none;
          }
        }
      }
    }
  }
}
</style>