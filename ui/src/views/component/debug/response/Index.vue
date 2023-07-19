<template>
  <div class="response-main">
    <ResponseRenderer></ResponseRenderer>
  </div>
</template>

<script setup lang="ts">
import {onMounted} from "vue";
import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';
import {resizeHeight} from "@/utils/dom";
import ResponseMeta from './Meta.vue';
import ResponseRenderer from './Renderer.vue';

const useForm = Form.useForm;
const {t} = useI18n();

const sendRequest = (e) => {
  console.log('sendRequest')
};

onMounted(() => {
  console.log('onMounted')
  resizeHeight('content', 'top-panel', 'splitter-v', 'bottom-panel', 200, 200)
})

</script>

<style lang="less" scoped>
.response-main {
  height: 100%;
}
</style>