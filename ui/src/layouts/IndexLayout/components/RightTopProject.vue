<template>
  <div class="indexlayout-top-project">
  <a-select
      ref="select"
      v-model:value="currProject.id"
      :bordered="true"
      style="width: 160px"
      @change="selectProject"
  >
    <a-select-option v-for="item in projects" :key="item.id" :value="item.id">{{ item.name }}</a-select-option>
  </a-select>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, onMounted} from "vue";
import {useStore} from "vuex";
import {StateType as UserStateType} from "@/store/user";
import {StateType as ProjectStateType} from "@/store/project";

interface RightTopProject {
  message: ComputedRef<number>;

  projects: ComputedRef<any>;
  currProject: ComputedRef<any>;
  selectProject: (value) => void;
}

export default defineComponent({
  name: 'RightTopProject',
  components: {},
  setup(): RightTopProject {
    const userStore = useStore<{ User: UserStateType }>();
    const projectStore = useStore<{ ProjectData: ProjectStateType }>();

    const message = computed<number>(() => userStore.state.User.message);
    const projects = computed<any>(() => projectStore.state.ProjectData.projects);
    const currProject = computed<any>(() => projectStore.state.ProjectData.currProject);

    userStore.dispatch("User/fetchMessage");
    projectStore.dispatch("ProjectData/fetchProject");

    const selectProject = (value): void => {
      console.log('selectProject', value)
      projectStore.dispatch('ProjectData/fetchProject', value);
    }

    return {
      message,

      currProject,
      projects,
      selectProject,
    }
  }
})
</script>