<template>
  <div id="expresslayout-left">
    <div class="expresslayout-left-logo">
      <router-link to="/" class="logo-url">
        <div class="logo-title">DeepTest</div>
      </router-link>
    </div>

    <div class="expresslayout-left-menu">

    </div>
  </div>
</template>

<script lang="ts">

import { defineComponent, PropType} from "vue";

export default defineComponent({
  name: 'Left',
  props: {

  },
  components: {

  },
})
</script>

<style lang="less" scoped>
@import '../../../assets/css/global.less';
#expresslayout-left {
  display: flex;
  flex-direction: column;
  width: 260px;
  height: 100vh;
  background-color: #FAFAFA;
  color: #343333;

  .expresslayout-left-logo {
    width: 100%;
    height: 50px;
    line-height: 50px;
    text-align: center;
    vertical-align: middle;

    .logo-url {
      display: inline-block;
      width: 100%;
      height: 100%;
      overflow: hidden;
      .logo-title {
        display: inline-block;
        margin: 0;
        font-size: 22px;
        font-family: Roboto, sans-serif;
        color: rgb(51, 51, 51);
      }
    }
    img {
      vertical-align: middle;
    }
  }

  .scrollbar();

}
</style>