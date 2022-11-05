<template>
  <div id="indexlayout-right-top" class="topNavEnable">
<!--    <RightTopWebsocket/>-->
  </div>
</template>
<script setup lang="ts">
import {defineComponent, PropType, toRefs} from "vue";
import {useI18n} from "vue-i18n";

    const {t} = useI18n();

</script>
<style lang="less">
#indexlayout-right-top {
  width: 100%;
  height: 50px;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
  z-index: 9;

  .indexlayout-right-top-top {
    display: flex;
    width: 100%;
    height: 50px;
    background-color: #222834;
    color: #c0c4cc;

    .indexlayout-flexible {
      width: 16px;
      height: 50px;
      line-height: 50px;
      text-align: center;
      cursor: pointer;
    }

    .indexlayout-left-logo {
      width: 150px;
      height: @headerHeight;
      line-height: @headerHeight;
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
          font-size: 16px;
          font-family: Roboto, sans-serif;
          color: #c0c4cc;
        }
      }

      img {
        vertical-align: middle;
      }
    }
  }
}
</style>