<template>
  <div class="project-edit-main">
    <a-tabs>
      <a-tab-pane key="1" tab="基本信息">
        <EditPage></EditPage>
      </a-tab-pane>

      <a-tab-pane key="2" tab="邀请用户">
        <InvitePage></InvitePage>
      </a-tab-pane>

      <template #tabBarExtraContent>
        <a-button type="link" @click="() => back()">返回</a-button>
      </template>
    </a-tabs>
  </div>
</template>

<script setup lang="ts">
import EditPage from './edit.vue';
import InvitePage from './invite.vue';
import {useRouter} from "vue-router";
const router = useRouter();

const back = (): void => {
  router.replace(`/project/index`)
}

</script>

<style lang="less" scoped>
.project-edit-main {

}
</style>
