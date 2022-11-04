<template>
  <div id="expresslayout">
    <left></left>

    <div id="expresslayout-right">
      <right-top></right-top>

      <div class="expresslayout-right-main">
          <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {computed, defineComponent, nextTick, ref, watch, defineProps, defineEmits, defineExpose} from "vue";
import {useStore} from 'vuex';
import {useRoute} from 'vue-router';
import {StateType as GlobalStateType} from '@/store/global';
import {StateType as UserStateType} from "@/store/user";

import Left from './components/Left.vue';
import RightTop from '@/layouts/IndexLayout/components/RightTop.vue';

// const props = defineProps({
//   msg: String,
// });
//
// const emits = defineEmits(['handle']);
// const click = () => {
//   emits('handle', '')
// }
//
// defineExpose({
//   name,
//   btn
// })

const store = useStore<{ Global: GlobalStateType; User: UserStateType; }>();
const route = useRoute();

</script>

<style lang="less">
#expresslayout {
  display: flex;
  height: 100vh;
  overflow: hidden;
}

#expresslayout-right {
  position: relative;
  flex: 1;
  overflow: auto;
  background-color: #f0f3f4;

  &.fiexd-header {
    display: flex;
    flex-direction: column;

    .expresslayout-right-main {
      flex: 1;
      overflow: auto;
    }
  }
}

.expresslayout-main-content {
  margin: 24px;
  position: relative;
}
</style>