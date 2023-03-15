<template>
  <div class="indexlayout-top-project">
  <a-select
      ref="select"
      v-model:value="currProject.id"
      :bordered="true"
      style="width: 280px;margin-left: 16px;"
      @change="selectProject"
  >
    <a-select-option v-for="item in projects" :key="item.id" :value="item.id">{{ item.name }}</a-select-option>
  </a-select>
  </div>
</template>

<script setup lang="ts">
import {computed, ComputedRef, defineComponent, onMounted} from "vue";
import {useStore} from "vuex";
import router from '@/config/routes';
import {StateType as UserStateType} from "@/store/user";
import {StateType as ProjectStateType} from "@/store/project";
import {StateType as EnvironmentStateType} from "@/store/environment";

const store = useStore<{ User: UserStateType, ProjectGlobal: ProjectStateType, Environment: EnvironmentStateType }>();

const message = computed<number>(() => store.state.User.message);
const projects = computed<any>(() => store.state.ProjectGlobal.projects);
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);

store.dispatch("User/fetchMessage");
store.dispatch("ProjectGlobal/fetchProject");

const selectProject = (value): void => {
  console.log('selectProject', value)
  store.dispatch('ProjectGlobal/changeProject', value);
  store.dispatch('Environment/getEnvironment', {id: 0, projectId: value})

  if(router.currentRoute.value.path.indexOf('/scenario/') > -1) {
    router.replace('/scenario/index')
  }
}

</script>
