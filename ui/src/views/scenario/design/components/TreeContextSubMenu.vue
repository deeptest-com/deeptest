<template>
  <div class="dp-tree-context-submenu">
    <template v-if="processorTypes?.length > 1">
      <a-sub-menu @click.stop :key="category.value" trigger="['click']" class="menu-item" popupClassName="dp-tree-context-submenu">
        <template #title>
          <span>{{t(category.label)}}</span>
        </template>
        <a-menu-item v-for="(item) in processorTypes"
                     :key="'add-'+mode+'-'+category.label+'-'+item.label" class="menu-item">
          {{t(item.label)}}
        </a-menu-item>
      </a-sub-menu>
    </template>

    <template v-if="processorTypes?.length === 1">
      <a-menu-item :key="'add-'+mode+'-'+category.label+'-'+processorTypes[0].label" class="menu-item">
        {{t(processorTypes[0].label)}}
      </a-menu-item>
    </template>

  </div>
</template>

<script setup lang="ts">
import {defineProps} from "vue";
import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';

const useForm = Form.useForm;

const props = defineProps<{
  processorTypes: any,
  category: any,
  mode: string,
}>()

const {t} = useI18n();

</script>

<style lang="less">
</style>