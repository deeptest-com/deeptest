<script lang="ts" setup>

import {computed, defineProps, defineEmits, ref, watch} from "vue";
const props = defineProps(['value','deprecated','required']) ;
const emit = defineEmits([])
const visible = ref(false);
</script>

<template>
  <div>
    <a-typography-text v-if="deprecated" :strong="false" type="warning" style="margin-right: 6px;">
      deprecated
    </a-typography-text>
    <a-typography-text v-if="required" :strong="false" type="warning" >
      required
    </a-typography-text>
  </div>
</template>


