<template>
  <div class="js-lib-list-main">
1
  </div>
</template>
<script setup lang="ts">

import {computed, createVNode, onMounted, ref, watch} from 'vue';
import {useStore} from "vuex";
import {StateType as ProjectStateType} from "@/store/project";
import {StateType as ProjectSettingStateType} from '../../store';
import {useI18n} from "vue-i18n";
import {notifyError, notifySuccess} from "@/utils/notify";

const {t} = useI18n();

const store = useStore<{ ProjectGlobal: ProjectStateType, ProjectSetting: ProjectSettingStateType, Project }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);


</script>

<style scoped lang="less">
.js-lib-list-main {

}
</style>
