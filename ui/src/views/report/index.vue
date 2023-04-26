<template>
    <div class="report-container">
        <router-view></router-view>
    </div>
</template>
<script setup></script>
<style lang="less" scoped>
.report-container {
    margin: 16px;
    background: #ffffff;
    min-height: calc(100vh - 92px);
    overflow: hidden;
    padding: 20px;
    box-sizing: border-box;
}
</style>