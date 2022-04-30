<template>
  <div class="response-checkpoint-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">
        </a-col>
      </a-row>
    </div>

    <div class="body">

    </div>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { CloseOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import {Interface, Response} from "@/views/interface/data";

export default defineComponent({
  name: 'ResponseCheckpoint',
  components: {
  },

  computed: {
  },

  setup(props) {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const responseData = computed<Response>(() => store.state.Interface.responseData);

    return {
      responseData,
    }
  }
})

</script>

<style lang="less">
.response-checkpoint-main {
}
</style>

<style lang="less" scoped>
.response-checkpoint-main {
  height: 100%;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .body {
    padding: 6px;
    height: calc(100% - 30px);

  }
}
</style>