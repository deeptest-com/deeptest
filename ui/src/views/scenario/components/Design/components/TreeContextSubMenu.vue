<template>
  <div class="dp-tree-context-submenu">
    <template v-if="processorTypes?.length > 1">
      <a-sub-menu trigger="['click']" class="menu-item" popupClassName="dp-tree-context-submenu"
                  @click.stop
                  :key="category.value" >

        <template #title>
          <span>{{t(category.label)}}</span>
        </template>

        <template v-for="(item) in processorTypes" :key="'add-'+mode+'-'+category.label+'-'+item.label">
          <a-menu-item class="menu-item"
                       v-if="showSubMenuItem(item.label)">
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
import {getProcessorTypeMap} from "@/views/scenario/service";
import {isInArray} from "@/utils/array";
import {ProcessorCategory, ProcessorCookie} from "@/utils/enum";

const useForm = Form.useForm;

const props = defineProps<{
  category: any,
  mode: string,
  isInterface: boolean,
}>()

const {t} = useI18n();

const processorTypeMap = getProcessorTypeMap()
const processorTypes = computed<any>(() => processorTypeMap[props.category.label])

const showSubMenuItem = (type) => {
  if (props.isInterface &&
      props.category.label === ProcessorCategory.ProcessorCookie && !isInArray(type, [ProcessorCookie.Get,])) {
    return false
  }

  return true
}

</script>

<style lang="less">
</style>