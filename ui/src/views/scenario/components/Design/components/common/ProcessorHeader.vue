<!-- ::::场景执行器详情：header 头部信息  -->
<template>
  <span>{{'自定义请求'}}</span>
  <EditAndShow placeholder="修改标题"
                    :value="detailResult?.title || ''"
                    @update="updateTitle"/>
  <a-divider style="    margin: 8px 0;"/>
</template>

<script setup lang="ts">
import {computed, defineProps, inject, ref} from "vue";
import {useStore} from "vuex";
import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Scenario} from "@/views/scenario/store";
import EditAndShow from "@/components/EditAndShow/index.vue";

const store = useStore<{ Debug: Debug, Scenario: Scenario }>();
const nodeData: any = computed<boolean>(() => store.state.Scenario.nodeData);



// 更新标题
async function updateTitle(title) {
  debugger;
  // await store.dispatch('Scenario/saveScenario',
  //     {id: detailResult.value.id, name: title}
  // );
  // emit('refreshList');
}


</script>

<style lang="less">
.processor-edit-fullscreen {

}
</style>

