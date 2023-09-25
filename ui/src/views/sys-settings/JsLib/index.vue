<template>
  <div class="js-lib-list-main">

  </div>
</template>
<script setup lang="ts">

import {computed, createVNode, onMounted, ref, watch} from 'vue';
import {useStore} from "vuex";
import {StateType as SysSettingStateType} from '../store';
import {useI18n} from "vue-i18n";
import {notifyError, notifySuccess} from "@/utils/notify";

const {t} = useI18n();

const store = useStore<{ SysSetting: SysSettingStateType }>();
const jsLibs = computed<any>(() => store.state.SysSetting.jsLibs);


</script>

<style scoped lang="less">
.js-lib-list-main {

}
</style>
