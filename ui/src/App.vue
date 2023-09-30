<template>
    <a-config-provider :locale="antdLocales">
      <router-view></router-view>
      <Notification></Notification>
    </a-config-provider>
</template>
<script lang="ts">
import {defineComponent, computed, onMounted, watch, ref} from "vue";
import { antdMessages } from "@/config/i18n";
import { setHtmlLang } from "@/utils/i18n";
import { useI18n } from "vue-i18n";
import Notification from "./components/others/Notification.vue";
import renderfeedback from "@/utils/feedback";
import {useStore} from "vuex";

import { StateType as UserStateType, CurrentUser } from "@/store/user";
import settings from "@/config/settings";
import {getCache} from "@/utils/localCache";
import {getAgentUrlByValue, isElectronEnv} from "@/utils/agentEnv";
import {isLeyan} from "@/utils/comm";

export default defineComponent({
  name: 'App',
  components: {
    Notification,
  },
  setup() {
    const { locale } = useI18n();
    const antdLocales = computed(()=> antdMessages[locale.value]);

    const isLyEnv = isLeyan();
    // NOTICE: 以下代码仅适用于乐研环境，其他环境删除即可
    const store = useStore<{User: UserStateType}>();
    const currentUser = computed<CurrentUser>(()=> store.state.User.currentUser);
    watch(() => {
      return currentUser.value
    },(newVal) => {
      // 仅乐研环境才会接入
      if(newVal?.username && isLyEnv){
        // 渲染乐研评论反馈系统
        renderfeedback(currentUser);
      }
    },{immediate:true})


    /*************************************************
     * ::::::::::: 以下代码仅适用于 Electron 环境 ::::::::::
     ************************************************/
    if (isElectronEnv && window?.require('electron')?.ipcRenderer && isLyEnv) {
      const ipcRenderer = window.require('electron').ipcRenderer
      // 更新本地占用的端口号
      ipcRenderer.on(settings.electronMsgUsePort, async (event, data) => {
        console.log('use port msg from electron', event,data);
        window.localStorage.setItem('dp-cache-agent-local-port',data?.agentPort || '');
      })
    }

    onMounted(() => {
      setHtmlLang(locale.value);
    })

    return {
      antdLocales
    }
  }
})
</script>
