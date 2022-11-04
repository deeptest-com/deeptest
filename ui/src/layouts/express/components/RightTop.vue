<template>
  <div id="expresslayout-right-top">
    <div class="expresslayout-right-top-top">
      <div class="expresslayout-top-menu">
      </div>

      <div class="expresslayout-top-menu-right">
      </div>
    </div>

    <RightTopWebsocket/>
  </div>
</template>
<script setup lang="ts">
import {defineComponent, PropType, toRefs} from "vue";
import {useI18n} from "vue-i18n";
import RightTopWebsocket from './RightTopWebsocket.vue';

const {t} = useI18n();
// const {topNavEnable} = toRefs(props);

</script>

<style lang="less">
#expresslayout-right-top {
  width: 100%;
  height: 50px;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
  z-index: 9;

  .expresslayout-right-top-top {
    display: flex;
    width: 100%;
    height: 50px;
    background-color: #222834;
    color: #c0c4cc;
  }
}
</style>