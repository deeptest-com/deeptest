<template>
  <div class="right-top-update-main">
    <div></div>

    <a-modal title="升级提醒"
           :visible="isVisible"
           :onCancel="onCancel"
           :maskClosable="false"
           class="update-modal">
      <div>
        发现新的版本<b>{{newVersion}}</b>，请确定是否升级。
      </div>
      <div v-if="downloadingPercent > 0">
        <a-progress :percent="downloadingPercent" />
      </div>

      <template #footer>
        <a-button @click="update" type="primary">立即升级</a-button>
        <a-button @click="defer">明天提醒我</a-button>
        <a-button @click="skip">跳过这个版本</a-button>
      </template>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
  import {ref} from "vue";
  import settings from "@/config/settings";
  import {getCache, setCache} from "@/utils/localCache";

  const isVisible = ref(false)
  const currVersion  = ref(0)
  const newVersion = ref(0)
  const forceUpdate = ref(false)
  const downloadingPercent = ref(0)

  const isElectron = ref(!!window.require)
  let ipcRenderer = undefined as any

  if (isElectron.value && !ipcRenderer) {
    ipcRenderer = window.require('electron').ipcRenderer

    ipcRenderer.on(settings.electronMsgUpdate, async (event, data) => {
      console.log('update msg from electron', data)
      currVersion.value = data.currVersionStr
      newVersion.value = data.newVersionStr
      forceUpdate.value = data.forceUpdate

      const skippedVersion = await getCache(settings.skippedVersion);
      const ignoreUtil = await getCache(settings.ignoreUtil);
      if (skippedVersion === newVersion.value || Date.now() < ignoreUtil) return;

      isVisible.value = true
    })

    ipcRenderer.on(settings.electronMsgDownloading, async (event, data) => {
      console.log('downloading msg from electron', data);
      downloadingPercent.value = Math.round(data.percent * 100);
    })
  }

  const update  = () => {
    console.log('update')
    ipcRenderer.send(settings.electronMsgUpdate, {
      currVersion: currVersion.value,
      newVersion: newVersion.value,
      forceUpdate: forceUpdate.value
    })
  }
  const defer  = () => {
    console.log('defer')
    setCache(settings.ignoreUtil, Date.now() + 24 * 3600);
    isVisible.value = false
  }
  const skip  = () => {
    console.log('skip')
    setCache(settings.skippedVersion, newVersion.value);
    isVisible.value = false
  }

  const onCancel = () => {
    console.log('onCancel')
    isVisible.value = false
  }

</script>

<style lang="less">
.update-modal{
  .ant-modal-footer {
    text-align: center;
  }
}
</style>