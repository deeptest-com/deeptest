<template>
  <div></div>
</template>

<script lang="ts">
import {defineComponent, onMounted, onBeforeUnmount} from "vue";
import { useI18n } from "vue-i18n";

import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import {notification} from "ant-design-vue";
import {NotificationKeyCommon} from "@/utils/const";
import {notifyError} from "@/utils/notify";

export default defineComponent({
  name: 'Notification',

  setup() {
    const { t } = useI18n();

    const notifyErr = (result: any) => {
      if (!result.httpCode) result.httpCode = 1000

      const msg = result.httpCode === 200 ? t('biz_'+result.resultCode) : t('http_'+result.httpCode)
      const desc = result.resultMsg ? result.resultMsg : ''

      if (result.resultCode !== 401) {
        notifyError(msg);
      }
    }

    onMounted(() => {
      console.log('onMounted')
      bus.on(settings.eventNotify, notifyErr);
    })
    onBeforeUnmount( () => {
      console.log('onBeforeUnmount')
      bus.off(settings.eventNotify, notifyErr);
    })

    return {
      t,
    }
  }
})
</script>