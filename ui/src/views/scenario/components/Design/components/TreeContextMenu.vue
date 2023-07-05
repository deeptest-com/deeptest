<template>
  <div class="dp-tree-context-menu" :set="mode = treeNode.isDir ? 'child' : 'brother'">
    <a-menu @click="menuClick" trigger="['click']" mode="vertical">
      {{void (
        menuItems = getMenu(treeNode.entityCategory),
        isInterface = treeNode.entityCategory === ProcessorCategory.ProcessorInterface
       )}}

      <template v-if="isInArray(ProcessorAction.ActionAddProcessor, menuItems) ||
                      isInArray(ProcessorAction.ActionInInterface, menuItems)">
        <a-sub-menu @click.stop key="addProcessor" trigger="['click']" class="menu-item" popupClassName="dp-tree-context-submenu">
          <template #title>
            <FolderAddOutlined />
            <span>新建处理器</span>
          </template>

          <template v-for="(category) in processorCategories" :key="category.value">
            <TreeContextSubMenu
                v-if="showMenuItem(isInterface, category.label)"
                :category="category"
                :isInterface="isInterface"
                mode="child" />
          </template>
        </a-sub-menu>
      </template>

      <template v-if="isInArray(ProcessorAction.ActionImportInterface, menuItems)">
        <a-sub-menu @click.stop key="addInterface" trigger="['click']" class="menu-item" popupClassName="dp-tree-context-submenu">
          <template #title>
            <FolderAddOutlined />
            <span>添加请求</span>
          </template>

          <a-menu-item key="add-child-interface-fromDefine" class="menu-item">
            <span>从接口管理模块</span>
          </a-menu-item>
          <a-menu-item key="add-child-interface-fromTest" class="menu-item">
            <span>从接口调试模块</span>
          </a-menu-item>
        </a-sub-menu>
      </template>

      <template v-if="isInArray(ProcessorAction.ActionEdit, menuItems)">
        <a-menu-item key="edit" class="menu-item">
          <EditOutlined />
          <span>编辑</span>
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
import {defineProps} from "vue";
import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';
import {FolderAddOutlined, EditOutlined, CloseOutlined} from "@ant-design/icons-vue";

import {isInArray} from "@/utils/array";
import {ProcessorAction, ProcessorCategory} from "@/utils/enum";
import {getProcessorCategories, getMenu, showMenuItem} from "@/views/scenario/service";
import TreeContextSubMenu from "./TreeContextSubMenu.vue";

const useForm = Form.useForm;

const props = defineProps<{
  treeNode: any,
  onMenuClick: Function,
}>()

const {t} = useI18n();

const processorCategories = getProcessorCategories()

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
