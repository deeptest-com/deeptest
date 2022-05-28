<template>
  <div class="dp-tree-context-menu">
    <a-menu @click="menuClick" mode="vertical">
      <template v-if="isRoot(treeNode.processorType)">
        <a-menu-item key="addProcessor" class="menu-item">
          <PlusOutlined />
          <span>新建处理器</span>
        </a-menu-item>

        <a-menu-item key="addInterface" class="menu-item">
          <PlusOutlined />
          <span>新建请求</span>
        </a-menu-item>

        <a-sub-menu key="test" class="menu-item" popupClassName="dp-tree-context-submenu">
          <template #title>
            <PlusOutlined />
            <span>创建父处理器</span>
          </template>

          <a-menu-item class="menu-item">处理器1</a-menu-item>
          <a-menu-item class="menu-item">处理器2</a-menu-item>
        </a-sub-menu>
      </template>

      <template v-if="isProcessor(treeNode.processorType)">
        <a-menu-item key="addProcessor" class="menu-item">
          <PlusOutlined />
          <span>新建处理器</span>
        </a-menu-item>
        <a-menu-item key="addInterface" class="menu-item">
          <PlusOutlined />
          <span>新建请求</span>
        </a-menu-item>
      </template>

      <template v-if="isInterface(treeNode.processorType)">
        <a-menu-item key="addInterface" class="menu-item">
          <PlusOutlined />
          <span>新建父</span>
        </a-menu-item>
      </template>

      <template v-if="false">
        <a-menu-item key="rename" class="menu-item" v-if="treeNode.parentId > 0">
          <EditOutlined />
          <span>重命名</span>
        </a-menu-item>

        <a-menu-item key="add_brother_node" class="menu-item" v-if="treeNode.parentId > 0">
          <PlusOutlined />
          <span>创建兄弟节点</span>
        </a-menu-item>

        <a-menu-item key="add_child_node" class="menu-item" v-if="treeNode.isDir">
          <PlusOutlined />
          <span>创建子节点</span>
        </a-menu-item>

        <a-menu-item key="add_brother_dir" class="menu-item" v-if="treeNode.parentId > 0">
          <PlusOutlined />
          <span>创建兄弟目录</span>
        </a-menu-item>

        <a-menu-item key="add_child_dir" class="menu-item" v-if="treeNode.isDir">
          <PlusOutlined />
          <span>创建子目录</span>
        </a-menu-item>

        <a-menu-item key="remove" class="menu-item" v-if="treeNode.parentId > 0">
          <CloseOutlined />
          <span v-if="treeNode.isDir">删除目录</span>
          <span v-if="!treeNode.isDir">删除节点</span>
        </a-menu-item>
      </template>
    </a-menu>
  </div>
</template>

<script setup lang="ts">
import {defineComponent, defineProps, PropType, Ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';
import {EditOutlined, CloseOutlined, PlusOutlined} from "@ant-design/icons-vue";

const useForm = Form.useForm;

const props = defineProps<{
  treeNode: Object,
  onSubmit: Function,
}>()

const {t} = useI18n();

const menuClick = (e) => {
  console.log('menuClick', e, props.treeNode)
  const targetId = props.treeNode.id
  const key = e.key

  props.onSubmit(key, targetId);
};

const isRoot = (type) => {
  return type === 'processor_root'
}
const isProcessor = (type) => {
  return type.indexOf('processor_') > -1 && type !== 'processor_root'
}
const isInterface = (type) => {
  return type.indexOf('processor_') < 0
}

</script>

<style lang="less">
.dp-tree-context-menu {
  z-index: 9;
  .ant-menu {
    border: 1px solid #dedfe1;
    background: #f0f2f5;
    .menu-item, .menu-item .ant-menu-submenu-title {
      margin-bottom: 8px;
      margin-top: 4px;

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
  border: 1px solid #dedfe1;
  background: #f0f2f5;
  .menu-item {
    margin-bottom: 8px;
    margin-top: 8px;
    padding-left: 22px !important;
    height: 22px;
    line-height: 21px;
  }
}
</style>