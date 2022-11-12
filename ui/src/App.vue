<template>
    <a-config-provider :locale="antdLocales">
      <router-view></router-view>
      <Notification></Notification>
    </a-config-provider>
</template>
<script lang="ts">
import { defineComponent, computed, onMounted } from "vue";
import { antdMessages } from "@/config/i18n";
import { setHtmlLang } from "@/utils/i18n";
import { useI18n } from "vue-i18n";
import Notification from "./components/others/Notification.vue";

export default defineComponent({
  name: 'App',
  components: {
    Notification,
  },
  setup() {
    const { locale } = useI18n();
    const antdLocales = computed(()=> antdMessages[locale.value]);

    onMounted(() => {
      setHtmlLang(locale.value);
    })

    return {
      antdLocales
    }
  }
})
</script>