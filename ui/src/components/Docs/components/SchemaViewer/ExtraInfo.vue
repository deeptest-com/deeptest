<script lang="ts" setup>
import {
  DeleteOutlined,
  InfoCircleOutlined,
  ReadOutlined,
  InfoCircleTwoTone
} from '@ant-design/icons-vue';
import {computed, defineProps, defineEmits, ref, watch} from "vue";

const props = defineProps(['value']);

const emit = defineEmits([])

const visible = ref(false);

</script>

<template>
  <div>
    <a-typography-text :strong="false" type="warning" style="margin-right: 6px;">
      {{
        value.deprecated ? 'deprecated' : 'deprecated'
      }}
    </a-typography-text>
    <a-typography-text :strong="false" type="warning" >
      {{
        value.required ? 'required' : 'required'
      }}
    </a-typography-text>
  </div>

</template>


