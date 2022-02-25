<template>
  <a-menu @click="menuClick" mode="inline">
    <a-menu-item key="add_brother_node" class="menu-item" v-if="treeNode.parentId > 0">
      <PlusOutlined/>
      <span>创建兄弟节点</span>
    </a-menu-item>
    <a-menu-item key="add_child_node" class="menu-item" v-if="treeNode.isDir">
      <PlusOutlined/>
      <span>创建子节点</span>
    </a-menu-item>

    <a-menu-item key="add_brother_dir" class="menu-item" v-if="treeNode.parentId > 0">
      <PlusOutlined/>
      <span>创建兄弟目录</span>
    </a-menu-item>
    <a-menu-item key="add_child_dir" class="menu-item" v-if="treeNode.isDir">
      <PlusOutlined/>
      <span>创建子目录</span>
    </a-menu-item>

    <a-menu-item key="remove" class="menu-item" v-if="treeNode.parentId > 0">
      <CloseOutlined/>
      <span v-if="treeNode.isDir">删除目录</span>
      <span v-if="!treeNode.isDir">删除节点</span>
    </a-menu-item>
  </a-menu>
</template>

<script lang="ts">
import {defineComponent, PropType, Ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';
import {CloseOutlined, PlusOutlined} from "@ant-design/icons-vue";

const useForm = Form.useForm;

interface TreeContextMenuSetupData {
  menuClick: (e) => void;
}

export default defineComponent({
  name: 'TreeContextMenu',
  props: {
    treeNode: {
      type: Object,
      required: true
    },
    onSubmit: {
      type: Function as PropType<(selectedKey: string, targetId: number) => void>,
      required: true
    }
  },
  components: {
    PlusOutlined, CloseOutlined,
  },
  setup(props): TreeContextMenuSetupData {
    const {t} = useI18n();

    const menuClick = (e) => {
      console.log('menuClick', e, props.treeNode)
      const targetId = props.treeNode.id
      const key = e.key

      props.onSubmit(key, targetId);
    };

    return {
      menuClick
    }
  }
})
</script>