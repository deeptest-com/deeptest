<template>
  <div class="dp-tree-context-menu">
    <a-menu @click="menuClick" trigger="['click']" mode="vertical">
      <template v-if="isRoot(treeNode.entityCategory)">
        <a-sub-menu @click.stop key="addProcessor" trigger="['click']" class="menu-item" popupClassName="dp-tree-context-submenu">
          <template #title>
            <FolderAddOutlined />
            <span>新建处理器</span>
          </template>

          <template v-for="(category) in processorTypes" :key="category.value">
            <TreeContextSubMenu
                :processorTypes="processorTypeMap[category.label]"
                :category="category"
                mode="child" />
          </template>

        </a-sub-menu>

        <a-menu-item key="add-child-interface" class="menu-item">
          <FileAddOutlined />
          <span>新建请求</span>
        </a-menu-item>
      </template>

      <template v-if="isProcessor(treeNode.entityCategory)">
        <a-sub-menu @click.stop key="addProcessor" class="menu-item" popupClassName="dp-tree-context-submenu">
          <template #title>
            <FolderAddOutlined />
            <span>新建处理器</span>
          </template>
          <template v-for="(category) in processorTypes" :key="category.value">
            <TreeContextSubMenu
                :processorTypes="processorTypeMap[category.label]"
                :category="category"
                mode="child"/>
          </template>
        </a-sub-menu>

        <a-menu-item key="add-child-interface" class="menu-item">
          <FileAddOutlined />
          <span>新建请求</span>
        </a-menu-item>
      </template>

      <template v-if="isInterface(treeNode.entityCategory)">
        <a-sub-menu @click.stop key="addProcessor" class="menu-item" popupClassName="dp-tree-context-submenu">
          <template #title>
            <FolderAddOutlined />
            <span>新建父处理器</span>
          </template>
          <template v-for="(category) in processorTypes" :key="category.value">
            <TreeContextSubMenu
                :processorTypes="processorTypeMap[category.label]"
                :category="category"
                mode="parent"/>
          </template>
        </a-sub-menu>
      </template>

      <template v-if="!isRoot(treeNode.entityCategory)">
        <a-menu-item key="rename" class="menu-item">
          <EditOutlined />
          <span>重命名</span>
        </a-menu-item>
        <a-menu-item key="remove" class="menu-item">
          <CloseOutlined />
          <span>删除</span>
        </a-menu-item>
      </template>

    </a-menu>
  </div>
</template>

<script setup lang="ts">
import {defineComponent, defineProps, PropType, Ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';
import {isRoot, isProcessor, isInterface} from '../../service'
import {FolderAddOutlined, FileAddOutlined, EditOutlined, CloseOutlined, PlusOutlined} from "@ant-design/icons-vue";

import {getProcessorTypeMap, getProcessorCategories} from "@/views/scenario/service";
import TreeContextSubMenu from "./TreeContextSubMenu.vue";

const useForm = Form.useForm;

const props = defineProps<{
  treeNode: any,
  onMenuClick: Function,
}>()

const {t} = useI18n();

const processorTypes = getProcessorCategories()
const processorTypeMap = getProcessorTypeMap()

const menuClick = (e) => {
  console.log('menuClick')
  const key = e.key
  const targetId = props.treeNode.id

  console.log(key, targetId)
  props.onMenuClick(key, targetId);
};

</script>

<style lang="less">
.dp-tree-context-menu {
  .ant-menu {
    border: 1px solid #dedfe1;
    background-color: #fff !important;

    .menu-item, .menu-item .ant-menu-submenu-title {
      padding-left: 12px !important;
      height: 22px;
      line-height: 21px;
    }
    .menu-item .ant-menu-submenu-title {
      padding-left: 0 !important;
    }
  }
}

.dp-tree-context-submenu {
  .ant-menu.ant-menu-sub.ant-menu-vertical {
    overflow-y: hidden;
  }

  .ant-menu-submenu.menu-item {
    margin: 0px !important;
  }

  .menu-item {
    margin-top: 5px !important;
    margin-bottom: 5px !important;
    height: 22px;
    line-height: 21px;

    .ant-menu-submenu-title {
      height: 22px;
      line-height: 21px;
    }
  }
}
</style>