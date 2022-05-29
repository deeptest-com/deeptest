<template>
  <div class="dp-tree-context-menu">
    <a-menu @click="menuClick" mode="vertical">
      <template v-if="isRoot(treeNode.processorType)">
        <a-sub-menu @click.stop key="addProcessor" class="menu-item" popupClassName="dp-tree-context-submenu">
          <template #title>
            <FolderAddOutlined />
            <span>新建处理器</span>
          </template>
          <a-menu-item v-for="(item) in processorTypes" :key="'add'+item.value" class="menu-item">
            {{t(item.label)}}
          </a-menu-item>
        </a-sub-menu>

        <a-menu-item key="addInterface" class="menu-item">
          <FileAddOutlined />
          <span>新建请求</span>
        </a-menu-item>
      </template>

      <template v-if="isProcessor(treeNode.processorType)">
        <a-sub-menu @click.stop key="addProcessor" class="menu-item" popupClassName="dp-tree-context-submenu">
          <template #title>
            <FolderAddOutlined />
            <span>新建处理器</span>
          </template>
          <a-menu-item v-for="(item) in processorTypes" :key="'add'+item.value" class="menu-item">
            {{t(item.label)}}
          </a-menu-item>
        </a-sub-menu>

        <a-menu-item key="addInterface" class="menu-item">
          <FileAddOutlined />
          <span>新建请求</span>
        </a-menu-item>
      </template>

      <template v-if="isInterface(treeNode.processorType)">
        <a-sub-menu @click.stop key="addProcessor" class="menu-item" popupClassName="dp-tree-context-submenu">
          <template #title>
            <FolderAddOutlined />
            <span>新建父处理器</span>
          </template>
          <a-menu-item v-for="(item) in processorTypes" :key="'add'+item.value" class="menu-item">
            {{t(item.label)}}
          </a-menu-item>
        </a-sub-menu>
      </template>

      <template v-if="!isRoot(treeNode.processorType)">
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
import {FolderAddOutlined, FileAddOutlined, EditOutlined, CloseOutlined, PlusOutlined} from "@ant-design/icons-vue";
import {getEnumSelectItems} from "@/views/interface/service";
import {OAuth2ClientAuthenticationWay} from "@/views/interface/consts";
import {ProcessorType} from "@/utils/enum";

const useForm = Form.useForm;

const props = defineProps<{
  treeNode: Object,
  onMenuClick: Function,
}>()

const {t} = useI18n();

const processorTypes = getEnumSelectItems(ProcessorType)

const menuClick = (e) => {
  console.log('menuClick')
  const key = e.key
  const targetId = props.treeNode.id

  console.log(key, targetId)
  props.onMenuClick(key, targetId);
};

const isRoot = (type) => {
  return type === 'processor_root'
}
const isProcessor = (type) => {
  return type !==  'processor_interface' && type !== 'processor_root'
}
const isInterface = (type) => {
  return type ===  'processor_interface'
}

</script>

<style lang="less">
.dp-tree-context-menu {
  z-index: 9;
  .ant-menu {
    border: 1px solid #dedfe1;
    background: #f0f2f5;
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
  border: 1px solid #dedfe1;
  background: #f0f2f5;
  .menu-item {
    padding-left: 22px !important;
    height: 22px;
    line-height: 21px;
  }
}
</style>