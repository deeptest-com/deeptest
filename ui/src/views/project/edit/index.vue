<template>
  <div class="indexlayout-main-conent">
    <a-card :bordered="false">
      <template #title>
        <div>编辑项目</div>
      </template>
      <template #extra>
        <a-button type="link" @click="() => back()">返回</a-button>
      </template>

      <div>表单</div>
    </a-card>
  </div>
</template>

<script lang="ts">
import {defineComponent, onMounted, onBeforeUnmount, getCurrentInstance, ComputedRef, Ref, ref, reactive} from "vue";
import { useStore } from 'vuex';
import {StateType as ListStateType} from "@/views/script/store";
import {useRouter} from "vue-router";
import {ActionRecordStart, EventNameJson, EventNodeIdJson, EventNameImg, EventNodeIdImg, ScopeDeeptest, ActionRecordedMsg} from "@/utils/const";
import {WebSocket, WsEventName} from "@/services/websocket";
import {getToken} from "@/utils/localToken";
import {ScriptItem, StepImage, StepItem} from "@/views/script/data";
import {createStep} from "@/views/script/service";

interface DesignScriptPageSetupData {
  back: () => void;
}

export default defineComponent({
    name: 'ScriptEditPage',
    setup(): DesignScriptPageSetupData {
      const router = useRouter();
      const back = ():void =>  {
        router.push(`/~/project/list`)
      }

      return {
        back,
      }
    }
})
</script>

<style lang="less" scoped>

</style>
