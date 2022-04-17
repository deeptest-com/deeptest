<template>
  <div class="history-main">
    <div class="head">
      <div class="title">
       历史
      </div>
    </div>

    <div class="body">
      <div class="histories">
        <div v-for="(item, idx) in requestsData" :key="idx" class="history">
          <div class="left">
            <span @click="loadHistory(item.id)" class="dp-link" title="点击加载历史请求数据">{{item.name}}</span>
          </div>
          <div class="right">
            <span @click="removeHistory(item.id)" class="dp-link"><DeleteOutlined /></span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { DeleteOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import Empty from "@/components/others/empty.vue";
import {Interface} from "@/views/interface/data";

export default defineComponent({
  name: 'RequestHistory',
  components: {
    DeleteOutlined,
  },

  computed: {
  },

  setup(props) {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);
    const requestsData = computed<any[]>(() => store.state.Interface.requestsData);

    const loadHistory = (id) => {
      console.log('loadHistory', id)
      store.dispatch('Interface/loadHistory', id)
    }

    const removeHistory = (id) => {
      console.log('removeHistory', id)
      store.dispatch('Interface/removeRequest', {id: id, interfaceId: interfaceData.value.id})
    }

    return {
      interfaceData,
      requestsData,

      loadHistory,
      removeHistory,
    }
  }
})

</script>

<style lang="less" scoped>
.history-main {
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
    height: calc(100% - 30px);
    overflow-y: hidden;

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
        }
      }
    }
  }
}
</style>