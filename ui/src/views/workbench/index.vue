<template>
  <div class="workbench">
      <StatisticHeader :params="cardData" :type="1"/>
      <List style="margin-top:16px"/>
  </div>
</template>


<script setup lang="ts">
import {useRouter} from "vue-router";
import {ref,computed,onMounted, nextTick,watch} from "vue";
import { useStore } from "vuex";
import { StateType } from "@/views/home/store";
import { StateType as ProjectStateType } from "@/store/project";
import {setCache,getCache} from "@/utils/localCache";
import settings from "@/config/settings";
import StatisticHeader from "@/components/StatisticHeader/index.vue";
import List from "./component/List/index.vue";
const store = useStore<{
  Home: StateType,
  ProjectGlobal: ProjectStateType;
}>();
const router = useRouter();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const cardData = computed<any>(() => store.state.Home.cardData);
const projectId = ref(+router.currentRoute.value.params.id)
// setCache(settings.currProjectId, projectId.value);

console.log(`TODO: change default project to ${projectId.value} on server side`)
// onMounted(async () => {

// });

const getHearderData = async (): Promise<void> => {
  await store.dispatch("Home/queryCard", {projectId:currProject.value.id});
  await store.dispatch("Home/queryPie", {projectId:currProject.value.id});
};
watch(
  () => {
    return currProject.value;
  },
  (val: any) => {
    console.log("~------currProject---", val);
    if (val.id) {
      getHearderData()
    }
  },
  {
    immediate: true,
  }
);

</script>

<style lang="less" scoped>
.workbench {
  height: 100%;
  overflow-y: scroll;
}
</style>