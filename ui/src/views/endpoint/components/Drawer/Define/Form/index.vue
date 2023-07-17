<template>
  <div class="content">
    <!-- 路径定义方式 -->
    <Path/>
    <!-- 请求方式定义 -->
    <Request/>
  </div>
</template>

<script lang="ts" setup>
import Path  from './Path.vue';
import Request  from './Request.vue';
</script>

<style lang="less" scoped>
.content {
  padding: 0;
}

</style>
