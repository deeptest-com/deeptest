<template>
    <div class="indexlayout-right-footer">
      <div class="footer-links">
        <a href="https://github.com/lqsong/admin-antd-vue"
           target="_blank">Github</a>
      </div>
      <div class="footer-copyright">Copyright © 2022 deeptest.com, All Rights Reserved</div>
    </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
export default defineComponent({
    name: 'RightFooter',
    components: {
    }
})
</script>
<style lang="less" scoped>
.indexlayout-right-footer {
  overflow: hidden;
  text-align: center;
  font-size: 14px;
  color: #808695;
  .footer-links {
    margin-top: 3px;
    a {
      margin-right: 30px;
      color: #808695;
      text-decoration: none;
      &:last-child {
        margin-right: 0;
      }
      &:hover {
        color: #515a6e;
      }
    }
  }
  .footer-copyright {
    margin-bottom: 5px;
  }
}
</style>