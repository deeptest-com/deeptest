<template>
  <div class="tree-container">
    <div class="tree-con">
      <div class="tag-filter-form">
        <a-input-search
            class="search-input"
            v-model:value="searchValue"
            placeholder="搜索接口分类"/>
        <div class="add-btn" @click="newCategorie(treeDataCategory?.[0])">
          <PlusOutlined style="font-size: 16px;"/>
        </div>
      </div>
      <div style="margin: 0 8px;">
        <a-tree
            class="deeptest-tree"
            draggable
            blockNode
            showIcon
            :expandedKeys="expandedKeys"
            :auto-expand-parent="autoExpandParent"
            @drop="onDrop"
            @expand="onExpand"
            @select="selectTreeItem"
            :tree-data="treeData">
          <template #switcherIcon>
            <CaretDownOutlined/>
          </template>
          <template #title="nodeProps">
            <div class="tree-title" :draggable="nodeProps.id === -1">
                <span class="tree-title-text" v-if="nodeProps.title.indexOf(searchValue) > -1">
                  {{ nodeProps.title.substr(0, nodeProps.title.indexOf(searchValue)) }}
                  <span style="color: #f50">{{ searchValue }}</span>
                  {{ nodeProps.title.substr(nodeProps.title.indexOf(searchValue) + searchValue.length) }}
                </span>
              <span class="tree-title-text" v-else>{{ nodeProps.title }}</span>
              <span class="more-icon" v-if="nodeProps.id !== -1">
                  <a-dropdown>
                       <MoreOutlined/>
                      <template #overlay>
                        <a-menu>
                          <a-menu-item key="0" @click="newCategorie(nodeProps)">
                             新建子分类
                          </a-menu-item>
                          <a-menu-item :disabled="nodeProps.id === -1" key="1" @click="deleteCategorie(nodeProps)">
                            删除分类
                          </a-menu-item>
                          <a-menu-item :disabled="nodeProps.id === -1" key="1" @click="editCategorie(nodeProps)">
                            编辑分类
                          </a-menu-item>
                        </a-menu>
                      </template>
                    </a-dropdown>
                </span>
            </div>
          </template>
        </a-tree>
        <div v-if="!treeData" class="nodata-tip">请点击上方按钮添加分类 ~</div>
      </div>
    </div>
    <!--  创建接口 Tag  -->
    <CreateCategoryModal
        :visible="createTagModalvisible"
        :nodeInfo="currentNode || {}"
        :mode="tagModalMode"
        @cancal="handleCancalTagModalCancal"
        @ok="handleTagModalOk"/>
  </div>
</template>
<script setup lang="ts">
import {
  computed, ref, onMounted,
  watch, defineEmits, defineProps
} from 'vue';
import {
  PlusOutlined,
  CaretDownOutlined,
  MoreOutlined
} from '@ant-design/icons-vue';
import {message, Modal} from 'ant-design-vue';
import CreateCategoryModal from '@/components/CreateCategoryModal/index.vue';
import {DropEvent} from 'ant-design-vue/es/tree/Tree';
import {useStore} from "vuex";
import {StateType as EndpointStateType} from "@/views/endpoint/store";
import {StateType as ProjectStateType} from "@/store/project";
import {setSelectedKey} from "@/utils/cache";

const store = useStore<{ Endpoint: EndpointStateType, ProjectGlobal: ProjectStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const treeDataCategory = computed<any>(() => store.state.Endpoint.treeDataCategory);
const treeDataMapCategory = computed<any>(() => store.state.Endpoint.treeDataMapCategory);
const createTagModalvisible = ref(false);
const props = defineProps({
  serveId: {
    required: false,
    type: Number || String,
  },
})
const searchValue = ref('');
const expandedKeys = ref<string[]>([]);
const autoExpandParent = ref<boolean>(false);
const treeData: any = computed(() => {
  const data = treeDataCategory.value;
  if(!data?.[0]?.id){
    return null;
  }
  data[0].children = data[0].children || [];
  function fn(arr: any) {
    if (!Array.isArray(arr)) {
      return;
    }
    arr.forEach((item) => {
      item.key = item.id;
      item.title = item.name;
      if (Array.isArray(item.children)) {
        fn(item.children)
      }
    });
  }
  fn(data);
  const children = data?.[0]?.children;
  //  末尾如果没有未分类，需要主动 push 未分类,未分类的 id 为 -1
  if (children?.length === 0 || (children?.length && children[children.length - 1]?.id != -1)) {
    children.push({
      id: -1,
      key: -1,
      title: '未分类',
      name: '未分类',
      parentId: data[0]?.id,
      children: []
    })
  }
  return data?.[0]?.children || null;
});

async function loadCategories() {
  await store.dispatch('Endpoint/loadCategory');
  expandAll();
}

watch(() => {
  return currProject.value;
}, async (newVal) => {
  if (newVal?.id) {
    await loadCategories();
  }
}, {
  immediate: true
})

watch(
    () => {
      return searchValue.value
    },
    (newVal) => {
      // 打平树形结构
      function flattenTree(tree) {
        const nodes: Array<any> = [];

        function traverse(node) {
          nodes.push(node);
          if (node.children) {
            node.children.forEach(traverse);
          }
        }

        traverse(tree);
        return nodes;
      }

      const flattenTreeList = flattenTree(treeData.value[0]);

      function findParentIds(nodeId, tree) {
        let current: any = tree.find(node => node.id === nodeId);
        let parentIds: Array<string> = [];
        while (current && current.parentId) {
          parentIds.unshift(current.parentId); // unshift 方法可以将新元素添加到数组的开头
          current = tree.find(node => node.id === current.parentId);
        }
        return parentIds;
      }

      let parentKeys: any = [];
      for (let i = 0; i < flattenTreeList.length; i++) {
        let node = flattenTreeList[i];
        if (node.title.includes(newVal)) {
          parentKeys.push(node.parentId);
          parentKeys = parentKeys.concat(findParentIds(node.parentId, flattenTreeList));
        }
      }
      parentKeys = [...new Set(parentKeys)];
      expandedKeys.value = parentKeys;
      autoExpandParent.value = true;
    });

const onExpand = (keys: string[]) => {
  expandedKeys.value = keys;
  autoExpandParent.value = false;
};

// 展开所有
function expandAll() {
  const keys: any = [];
  const data = treeDataCategory.value;

  function fn(arr: any) {
    if (!Array.isArray(arr)) {
      return;
    }
    arr.forEach((item, index) => {
      keys.push(item.id);
      if (Array.isArray(item.children)) {
        fn(item.children)
      }
    });
  }
  fn(data);
  expandedKeys.value = keys;
}

let selectedKeys = ref<number[]>([]);
const emit = defineEmits(['select']);

function selectTreeItem(keys, e) {
  selectedKeys.value = keys;
  setSelectedKey('category-endpoint', currProject.value.id, selectedKeys.value[0]);
  // 如果没有选中的节点，就默认选中根节点
  emit('select', selectedKeys.value?.[0] ? selectedKeys.value[0] : null);
  const selectedData = treeDataMapCategory.value[selectedKeys.value[0]]
  store.dispatch('Endpoint/getCategoryNode', selectedData);
}

const currentNode = ref(null);
// 新建或者修改
const tagModalMode = ref('new');

// 删除分类
async function deleteCategorie(node) {
  Modal.confirm({
    title: () => '确定删除该分类吗？',
    content: () => '删除后所有所有子分类都会被删除',
    okText: () => '确定',
    okType: 'danger',
    cancelText: () => '取消',
    onOk: async () => {
      const res = await store.dispatch('Endpoint/removeCategoryNode', node.id);
      if (res) {
        message.success('删除成功');
      } else {
        message.success('删除失败');
      }
    },
    onCancel() {
      console.log('Cancel');
    },
  });

}

// 新建分类
function newCategorie(node) {
  if (!node) {
    return;
  }
  tagModalMode.value = 'new';
  createTagModalvisible.value = true;
  currentNode.value = node;
}

//编辑分类
function editCategorie(node) {
  tagModalMode.value = 'edit';
  createTagModalvisible.value = true;
  currentNode.value = node;
}

async function handleTagModalOk(obj) {
  obj = Object.assign(currentNode.value, obj);
  // 修改
  if (tagModalMode.value === 'edit') {
    const res = await store.dispatch('Endpoint/saveCategory', {
      id: obj.id,
      name: obj.name,
      desc: obj.desc,
    });
    if (res?.code === 0) {
      createTagModalvisible.value = false;
      message.success('修改分类成功');
    } else {
      message.error('修改分类失败，请重试~');
    }
  }
  // 新建
  else if (tagModalMode.value === 'new') {
    const res = await store.dispatch('Endpoint/createCategoryNode', {
      "name": obj.name,
      "desc": obj.desc,
      "mode": "child",
      "targetId": obj.id,
      type: "endpoint",
      "projectId": currProject.value.id,
      "serveId": props.serveId || null, // 仅在接口管理模块下有
    });
    if (res?.code === 0) {
      createTagModalvisible.value = false;
      message.success('新建分类成功');
    } else {
      message.error('修改分类失败，请重试~');
    }
  }
}

function handleCancalTagModalCancal() {
  createTagModalvisible.value = false;
}


async function onDrop(info: DropEvent) {
  const dropKey = info.node.eventKey;
  const dragKey = info.dragNode.eventKey;
  const pos = info.node.pos.split('-');
  const dropPosition = info.dropPosition - Number(pos[pos.length - 1]);
  // 未分类不让移动
  if (dragKey === -1) {
    message.warning('未分类不能移动');
    return;
  }
  if (dropKey === -1) {
    message.warning('其他分类不能移动到未分类下');
    return;
  }
  const res = await store.dispatch('Endpoint/moveCategoryNode', {
    "currProjectId": currProject.value.id,
    "dragKey": dragKey, // 移动谁
    "dropKey": dropKey,  // 移动那儿
    "dropPos": dropPosition // 0 表示移动到目标节点的子节点，-1 表示移动到目标节点的前面， 1表示移动到目标节点的后面
  });
  if (res) {
    // 移动到目标节点的子节点，则需要展开目标节点
    if (dropPosition === 0) {
      expandedKeys.value = [...new Set([...expandedKeys.value, dropKey])];
    }
    message.success('移动成功');
  } else {
    message.error('移动失败');
  }
}

onMounted(async () => {
  await loadCategories();
  expandAll();
})

</script>

<style scoped lang="less">
.container {
  margin: 16px;
  background: #ffffff;
}

.tag-filter-form {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 50px;
  margin-top: 8px;
  .search-input {
    margin-left: 16px;
    margin-right: 8px;
  }

  .add-btn {
    margin-left: 2px;
    margin-right: 16px;
    cursor: pointer;
  }
}

.content {
  display: flex;
  width: 100%;

  .left {
    width: 300px;
    border-right: 1px solid #f0f0f0;
  }

  .right {
    flex: 1
  }
}

.action-new {
  margin-right: 8px;
}

.top-action {
  height: 60px;
  display: flex;
  align-items: center;
  margin-left: 16px;

  .ant-btn {
    margin-right: 16px;
  }
}

.action-btns {
  display: flex;
}

.customTitleColRender {
  display: flex;

  .edit {
    margin-left: 8px;
    cursor: pointer;
  }
}

.form-item-con {
  display: flex;
  justify-content: center;
  align-items: center;
}

.tree-title {
  position: relative;

  .tree-title-text {
    display: inline-block;
    width: calc(100% - 24px);
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  //&:hover{
  //  .more-icon {
  //    background-color: #f5f5f5;
  //  }
  //}
  .more-icon {
    position: absolute;
    right: -8px;
    width: 20px;
  }
}

.nodata-tip {
  margin-top: 8px;
  text-align: center;
}

</style>
