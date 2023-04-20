<template>
  <div class="indexlayout-top-project">
    <a-select
        v-model:value="currProject.id"
        :bordered="true"
        style="width: 280px;margin-left: 16px;"
        @change="selectProject">
      <a-select-option v-for="item in projects" :key="item.id" :value="item.id">{{ item.name }}</a-select-option>
    </a-select>
  </div>
</template>

<script setup lang="ts">
import {computed, watch, ref, onMounted} from "vue";
import {useStore} from "vuex";
import {useRoute} from "vue-router";
import router from '@/config/routes';
import {StateType as UserStateType} from "@/store/user";
import {StateType as ProjectStateType} from "@/store/project";
import {StateType as ServeStateType} from "@/store/serve";
import {StateType as EnvironmentStateType} from "@/store/environment";

const store = useStore<{ User: UserStateType,
  ProjectGlobal: ProjectStateType, ServeGlobal: ServeStateType, Environment: EnvironmentStateType }>();

const route = useRoute();

const message = computed<number>(() => store.state.User.message);
const projects = computed<any>(() => store.state.ProjectGlobal.projects);
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);

store.dispatch("User/fetchMessage");
store.dispatch("ProjectGlobal/fetchProject");
store.dispatch("ServeGlobal/fetchServe");

const selectProject = (value): void => {
  console.log('selectProject', value)
  window.localStorage.setItem('currentProjectId', value);
  store.dispatch('ProjectGlobal/changeProject', value);
  store.dispatch('Environment/getEnvironment', {id: 0, projectId: value});

  // 项目切换后，需要重新更新可选服务列表
  store.dispatch("ServeGlobal/fetchServe");

  if(router.currentRoute.value.path.indexOf('/scenario/') > -1) {
    router.replace('/scenario/index')
  }
}

</script>
