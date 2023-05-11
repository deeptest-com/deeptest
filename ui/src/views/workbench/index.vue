<template>
  <div class="workbench">
      <StatisticHeader :params="cardData" :type="1"/>
      <List style="margin-top:16px"/>
  </div>
</template>


<script setup lang="ts">
import {useRouter} from "vue-router";
import {ref,computed,onMounted} from "vue";
import { useStore } from "vuex";
import { StateType } from "@/views/home/store";
import { StateType as ProjectStateType } from "@/store/project";
import {setCache} from "@/utils/localCache";
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
setCache(settings.currProjectId, projectId.value);
console.log(`TODO: change default project to ${projectId.value} on server side`)
onMounted(() => {
  getHearderData()
});
const getHearderData = async (): Promise<void> => {
  await store.dispatch("Home/queryCard", {projectId:projectId.value||currProject.value.id});
  await store.dispatch("Home/queryPie", {projectId:projectId.value||currProject.value.id});
};
</script>

<style lang="less" scoped>
.workbench {
  padding: 16px;
}
</style>