<template>
    <div v-if="showContextMenu" :style="style" class="context-menu-main dp-context-menu">
      <a-menu @click="onMenuClick" mode="inline">
        <a-menu-item key="use-variable" class="menu-item">
          <EditOutlined />
          <span>使用变量</span>
        </a-menu-item>
      </a-menu>
    </div>
</template>

<script setup lang="ts">
import {defineProps, onMounted, onUnmounted, ref, watch} from "vue";
import { EditOutlined } from '@ant-design/icons-vue';
import {useI18n} from "vue-i18n";

const { t } = useI18n();

const props = defineProps({
  isShow:{
    type: Boolean,
    required: true
  },
  style:{
    type: Object,
    required: true
  },
  menuClick:{
    type: Function,
    required: true
  },
})

const onMenuClick = (e) => {
  props.menuClick(e.key)
}

const showContextMenu = ref(false)
watch(props, () => {
  console.log('watch props')
  showContextMenu.value = props.isShow
}, {deep: true})

const clearMenu = () => {
  console.log('clearMenu')
  showContextMenu.value = false
}
onMounted(() => {
  console.log('onMounted')
  document.addEventListener("click", clearMenu)
})
onUnmounted(() => {
  document.removeEventListener("click", clearMenu)
})

</script>

<style lang="less">
.context-menu-main {

}
</style>