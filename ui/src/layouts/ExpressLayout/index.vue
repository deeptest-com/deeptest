<template>
  <div class="expresslayout">
    <div class="expresslayout-left">
      <left></left>
    </div>

    <div class="expresslayout-right">
      <div class="expresslayout-right-top">
        <right-top></right-top>
      </div>

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
.expresslayout {
  display: flex;
  height: 100vh;
  overflow: hidden;

  .expresslayout-left {
    width: 280px;
  }

  .expresslayout-right {
    flex: 1;
    height: 100%;
    background-color: #FFFFFF;

    .expresslayout-right-top {
      height: 50px;
    }

    .expresslayout-right-main {
      height: calc(100% - 50px);
      .express-main {
        padding: 0 12px;
        position: relative;
        overflow-y: auto;
      }
    }
  }
}

</style>