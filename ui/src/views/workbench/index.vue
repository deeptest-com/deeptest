<template>
  <div class="workbench">
      <StatisticHeader/>
      <List style="margin-top:16px"/>
  </div>
</template>


<script setup lang="ts">
import {useRouter} from "vue-router";
import {ref} from "vue";
import {setCache} from "@/utils/localCache";
import settings from "@/config/settings";
import StatisticHeader from "@/components/StatisticHeader/index.vue";
import List from "./component/List/index.vue";

const router = useRouter();
const projectId = ref(+router.currentRoute.value.params.id)
setCache(settings.currProjectId, projectId.value);
console.log(`TODO: change default project to ${projectId.value} on server side`)

</script>

<style lang="less" scoped>
.workbench {
  padding: 16px;
}
</style>