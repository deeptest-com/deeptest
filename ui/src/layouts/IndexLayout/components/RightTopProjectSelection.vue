<template>
  <a-select
      ref="select"
      v-model:value="currProject"
      bordered="true"
      style="width: 160px"
      @focus="focus"
      @change="selectProject"
  >
    <a-select-option value="1">Porject A</a-select-option>
    <a-select-option value="lucy">Lucy</a-select-option>
    <a-select-option value="disabled" disabled>Disabled</a-select-option>
    <a-select-option value="Yiminghe">yiminghe</a-select-option>
  </a-select>
</template>
<script lang="ts">
import { computed, ComputedRef, defineComponent, onMounted } from "vue";
import { useStore } from "vuex";
import { StateType as UserStateType } from "@/store/user";

interface RightTopProjectSelection {
  projects: ComputedRef<number>;
  currProject: ComputedRef<number>;
  selectProject: () => void;

  message: ComputedRef<number>;
}

export default defineComponent({
    name: 'RightTopProjectSelection',
    components: {
    },
    setup(): RightTopProjectSelection {
        const store = useStore<{user: UserStateType}>();

        const projects = computed<number>(()=> store.state.user.message);
        const currProject = computed<number>(()=> store.state.user.message);

        const message = computed<number>(()=> store.state.user.message);

        onMounted(()=>{
            store.dispatch("user/fetchMessage");
          store.dispatch("user/fetchProject");
        })

      const selectProject = (): void => {
        console.log('selectProject')
      }

        return {
          selectProject,
          currProject,
          projects,
          message,
        }
    }
})
</script>