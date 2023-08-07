<template>
  <div class="dp-tree-context-submenu">
    <!-- only show these categories that has child processorTypes (length >= 1) -->

    <template v-if="processorTypes?.length > 1">
      <a-sub-menu trigger="['click']" class="menu-item" popupClassName="dp-tree-context-submenu"
                  @click.stop
                  :key="category.value" >

        <template #title>
          <span>{{t(category.label)}}</span>
        </template>

        <template v-for="(item) in processorTypes">
          <a-menu-item class="menu-item"
                       v-if="showSubMenuItem(entityType, category, item.label)"
                       :key="'add-'+mode+'-'+category.label+'-'+item.label">
            {{t(item.label)}}
          </a-menu-item>
        </template>
      </a-sub-menu>
    </template>

    <template v-if="processorTypes?.length === 1">
      <a-menu-item class="menu-item"
                   :key="'add-'+mode+'-'+category.label+'-'+processorTypes[0].label" >
        {{t(processorTypes[0].label)}}
      </a-menu-item>
    </template>

  </div>
</template>

<script setup lang="ts">
import {computed, defineProps} from "vue";
import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';
import {getProcessorTypeMap, showSubMenuItem} from "@/views/scenario/service";

const useForm = Form.useForm;

const props = defineProps<{
  category: any,
  entityType: string,
  mode: string,
}>()

const {t} = useI18n();

const processorTypeMap = getProcessorTypeMap()
const processorTypes = computed<any>(() => processorTypeMap[props.category.label])

</script>

<style lang="less">
</style>