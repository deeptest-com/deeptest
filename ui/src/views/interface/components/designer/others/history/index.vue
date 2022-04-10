<template>
  <div class="history-main">
    <div class="head">
      <div class="title">
       历史
      </div>
    </div>

    <div class="body">
    </div>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined,ImportOutlined, MoreOutlined, PlusOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import Empty from "@/components/others/empty.vue";
import {Interface} from "@/views/interface/data";

export default defineComponent({
  name: 'RequestHistory',
  components: {

  },

  computed: {
  },

  setup(props) {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

    return {
      interfaceData,
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
    .historys {
      .history {
        display: flex;
        padding: 3px 2px 2px 4px;
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