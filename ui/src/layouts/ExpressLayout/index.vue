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

<script setup lang="ts">
import {computed, defineComponent, nextTick, ref, watch, defineProps, defineEmits, defineExpose} from "vue";
import {useStore} from 'vuex';
import {useRoute} from 'vue-router';
import {StateType as GlobalStateType} from '@/store/global';

import Left from './components/Left.vue';
import RightTop from './components/RightTop.vue';

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

const store = useStore<{ Global: GlobalStateType; }>();
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
  background-color: #FFFFFF;

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