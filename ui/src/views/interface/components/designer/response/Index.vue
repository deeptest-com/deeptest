<template>
  <div class="response-main">
    <ResponseMeta></ResponseMeta>
    <ResponseRenderer></ResponseRenderer>
  </div>
</template>

<script setup lang="ts">
import {computed, ComputedRef, defineComponent, onMounted, PropType, Ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';
import {resizeHeight} from "@/utils/dom";
import {useStore} from "vuex";
import {StateType} from "@/views/interface/store";
import ResponseMeta from './Meta.vue';
import ResponseRenderer from './Renderer.vue';

const useForm = Form.useForm;

    const {t} = useI18n();
    // const store = useStore<{ Interface: StateType }>();
    // const responseData = computed<any>(() => store.state.Interface.responseData);

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