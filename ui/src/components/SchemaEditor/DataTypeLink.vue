<script lang="ts" setup>
import {
  ArrowUpOutlined,
  ArrowDownOutlined,
  CopyOutlined,
} from '@ant-design/icons-vue';
import {computed,ref} from "vue";
import vClickOutside from "v-click-outside"

const props = defineProps<{
  type: boolean,
}>();

const emit = defineEmits<{
  (e: 'showSettingPropsModal',target:any): void,
  (e: 'clickOutside'): void,
}>()

const aRef = ref(null);


</script>

<template>
  <a href="javascript:void(0)"
     :ref="aRef"
     v-click-outside="emit('clickOutside')"
     @click="emit('showSettingPropsModal',aRef)"
     :class="[type, 'setDataTypeAction']"
  >{{type}}</a>
</template>


