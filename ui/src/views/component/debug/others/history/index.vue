<template>
  <div class="history-main">
    <div class="head">
      <div class="title">
       执行历史
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

<script lang="ts">
import {computed, defineComponent, inject} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { DeleteOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface1/store";
import {Interface} from "@/views/interface1/data";
import {UsedBy} from "@/utils/enum";
import {StateType as ScenarioStateType} from "@/views/scenario/store";

export default defineComponent({
  name: 'RequestHistory',
  components: {
    DeleteOutlined,
  },

  computed: {
  },

  setup(props) {
    const usedBy = inject('usedBy') as UsedBy
    const {t} = useI18n();

    const store = useStore<{ Interface1: StateType, Scenario: ScenarioStateType }>();
    const interfaceData = computed<Interface>(
        () => usedBy === UsedBy.interface ? store.state.Interface1.interfaceData : store.state.Scenario.interfaceData);
    const invocationsData = computed<any[]>(() =>
        usedBy === UsedBy.interface ? store.state.Interface1.invocationsData : store.state.Scenario.invocationsData);

    const getRequestAsInterface = (id) => {
      console.log('getRequestAsInterface', id)
      usedBy === UsedBy.interface ? store.dispatch('Interface1/getInvocationAsInterface', id) :
          store.dispatch('Scenario/getInvocationAsInterface', id)
    }

    const removeHistory = (id) => {
      console.log('removeHistory', id)
      usedBy === UsedBy.interface ? store.dispatch('Interface1/removeInvocation', {id: id, interfaceId: interfaceData.value.id}) :
          store.dispatch('Scenario/removeInvocation', {id: id, interfaceId: interfaceData.value.id})
    }

    const mouseOver = (event) => {
      // console.log('mouseOver', event)
      event.currentTarget.querySelector(".link").style.display = 'block'
    }
     const mouseLeave = (event) => {
       // console.log('mouseLeave', event);
       event.currentTarget.querySelector(".link").style.display = 'none'
    }

    return {
      interfaceData,
      invocationsData,

      getRequestAsInterface,
      removeHistory,
      mouseOver,
      mouseLeave
    }
  }
})

</script>

<style lang="less" scoped>
.history-main {
  display: flex;
  flex-direction: column;

  height: 100%;
  .head {
    padding: 0 5px;
    border-bottom: 1px solid #d9d9d9;
    height: 32px;
    line-height: 32px;
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
    flex: 1;
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
          width: 20px;
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