<template>
  <div class="body-schema">
    <SchemaEditor :value="content"/>
  </div>
</template>
<script lang="ts" setup>
import {
  CloseOutlined,
  EditOutlined,
  DeleteOutlined,
  InfoCircleOutlined,
  PlusOutlined,
} from '@ant-design/icons-vue';
import {computed, defineProps, defineEmits, ref, watch} from "vue";
import SchemaEditor from './schema';
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {MonacoOptions} from '@/utils/const';

const props = defineProps(['info']) ;
const examples: any = ref([]);

const content: any = ref({
  "type": "object",
  "deprecated": false,
  "minProperties": 1,
  // "ref": "#dddd",
  "properties": {
    "name1": {
      "type": "integer",
      "enum": ['a', 'b', 'c'],
      "format": null,
      "behavior": null,
      "default": "",
      "example": "example",
      "minimum": "",
      "maximum": "",
      "maxLength": "",
      "multipleOf": "",
      "exclusiveMin": false,
      "exclusiveMax": false,
      "deprecated": false
    },
    "name2": {
      "type": "string",
      "description": "description 12121212",
    },
    "obj1": {
      "type": "object",
      "example":"example 1 3 44",
      "properties": {
        "name1": {
          "type": "integer",
          "enum": ['a', 'b', 'c'],
          "format": null,
          "behavior": null,
          "default": "",
          "example": "example",
          "minimum": "",
          "maximum": "",
          "maxLength": "",
          "multipleOf": "",
          "exclusiveMin": false,
          "exclusiveMax": false,
          "deprecated": false
        },
        "name2": {
          "type": "string"
        },
      }
    }
  }
});

</script>


<style lang="less" scoped>
@import "var.less";
.body-schema{

}

</style>

