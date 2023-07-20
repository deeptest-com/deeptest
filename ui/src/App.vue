<template>
    <a-config-provider :locale="antdLocales">
      <router-view></router-view>
      <Notification></Notification>
    </a-config-provider>
</template>
<script lang="ts">
import {defineComponent, computed, onMounted, watch} from "vue";
import { antdMessages } from "@/config/i18n";
import { setHtmlLang } from "@/utils/i18n";
import { useI18n } from "vue-i18n";
import Notification from "./components/others/Notification.vue";
import renderfeedback from "@/utils/feedback";
import {useStore} from "vuex";

import { StateType as UserStateType, CurrentUser } from "@/store/user";

export default defineComponent({
  name: 'App',
  components: {
    Notification,
  },
  setup() {
    const { locale } = useI18n();
    const antdLocales = computed(()=> antdMessages[locale.value]);

    // 以下代码仅适用于乐研环境，其他环境删除即可
    const store = useStore<{User: UserStateType}>();
    const currentUser = computed<CurrentUser>(()=> store.state.User.currentUser);
    watch(() => {
      return currentUser.value
    },(newVal) => {
      if(newVal?.username){
        // 渲染乐研评论反馈系统
        renderfeedback(currentUser);
      }
    },{immediate:true})

    onMounted(() => {
      setHtmlLang(locale.value);
    })

    return {
      antdLocales
    }
  }
})
</script>
