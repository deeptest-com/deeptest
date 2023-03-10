<template>
  <div class="container">
    <div class="content">
      <div class="left tree">
        <div class="tag-filter-form">
          <a-input-search
              class="search-input"
              v-model:value="searchValue"
              placeholder="搜索接口分类"/>
          <!--          <div class="add-btn" @click="addApiTag">-->
          <!--            <PlusOutlined style="font-size: 16px;"/>-->
          <!--          </div>-->
        </div>
        <div style="margin: 0 8px;">
          <a-tree
              draggable
              showLine
              blockNode
              showIcon
              :expandedKeys="expandedKeys"
              :auto-expand-parent="autoExpandParent"
              @dragenter="onDragEnter"
              @drop="onDrop"
              @expand="onExpand"
              @select="selectTreeItem"
              :tree-data="treeData">
            <template #switcherIcon>
              <CaretDownOutlined/>
            </template>
            <template #title="nodeProps">
              <div>
                <span v-if="nodeProps.title.indexOf(searchValue) > -1">
                  {{ nodeProps.title.substr(0, nodeProps.title.indexOf(searchValue)) }}<span style="color: #f50">{{
                    searchValue
                  }}</span>{{ nodeProps.title.substr(nodeProps.title.indexOf(searchValue) + searchValue.length) }}
                </span>
                <span v-else>{{ nodeProps.title }}</span>
                <span class="more-icon">
                  <a-dropdown>
                       <EllipsisOutlined/>
                      <template #overlay>
                        <a-menu>
                          <a-menu-item key="0" @click="newCategorie(nodeProps)">
                             新建子分类
                          </a-menu-item>
                          <a-menu-item key="1" @click="deleteCategorie(nodeProps)">
                            删除分类
                          </a-menu-item>
                          <a-menu-item key="1" @click="editCategorie(nodeProps)">
                            编辑分类
                          </a-menu-item>
                        </a-menu>
                      </template>
                    </a-dropdown>
                </span>
              </div>
            </template>
          </a-tree>
        </div>
      </div>
      <div class="right">
        <!--  头部区域  -->
        <div class="top-action">
          <a-row type="flex" style="align-items: center;width: 100%">
            <a-col :span="4">
              <a-button class="action-new" type="primary" :loading="loading" @click="addApi">新建接口</a-button>
              <a-button class="action-import" type="primary" :disabled="!hasSelected" :loading="loading"
                        @click="importApi">
                导入
              </a-button>
            </a-col>
            <a-col :span="1"/>
            <a-col :span="4">
              <a-form-item label="创建人" style="margin-bottom: 0;">
                <a-select placeholder="请选择创建人">
                  <a-select-option value="admin"> admin</a-select-option>
                  <a-select-option value="superAdmin">super admin</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :span="1"/>
            <a-col :span="4">
              <a-form-item label="状态" style="margin-bottom: 0;">
                <a-select placeholder="请选择创建人" :options="interfaceStatusOpts"/>
              </a-form-item>
            </a-col>
            <a-col :span="2"/>
            <a-col :span="7" style="margin-right: 8px;">
              <a-input-search
                  placeholder="请输入你需要搜索的接口文档"
                  enter-button
                  @search="() => {

                  }"
              />
            </a-col>

          </a-row>


        </div>
        <a-table
            :row-selection="{
          selectedRowKeys: selectedRowKeys,
          onChange: onSelectChange
        }"
            :columns="columns"
            :data-source="data">
          <template #colTitle="{text,record}">
            <div class="customTitleColRender">
              <span>{{ text }}</span>
              <span class="edit" @click="editInterface(record)"><EditOutlined/></span>
            </div>
          </template>
          <template #colStatus="{record}">
            <div class="customTitleColRender">
              <span>{{ interfaceStatus.get(record.status) }}</span>
            </div>
          </template>
          <template #action="{record}">
            <div class="action-btns">
              <a-button type="link" @click="copy(record)">复制</a-button>
              <a-button type="link" @click="del(record)">删除</a-button>
              <a-button type="link" @click="disabled(record)">过时</a-button>
            </div>
          </template>
        </a-table>
      </div>
    </div>
    <!-- 编辑接口时，展开抽屉   -->
    <EditInterfaceDrawer
        :destroyOnClose="true"
        :interfaceId="editInterfaceId"
        :visible="drawerVisible"
        :key="clickTag"
        @refreshList="refreshList"
        @close="onCloseDrawer"/>
    <!--  创建接口 Tag  -->
    <CreateTagModal
        :visible="createTagModalvisible"
        :nodeInfo="currentNode"
        :mode="tagModalMode"
        @cancal="handleCancalCreateTag"
        @ok="handleCreateTag"/>
    <!--  创建新接口弹框  -->
    <CreateApiModal
        :visible="createApiModalvisible"
        @cancal="handleCancalCreateApi"
        @ok="handleCreateApi"/>

  </div>
</template>
<script setup lang="ts">
import {
  computed, reactive, toRefs, ref, onMounted,
  watch
} from 'vue';
import {ColumnProps} from 'ant-design-vue/es/table/interface';
import contenteditable from 'vue-contenteditable';
import {
  PlusOutlined,
  EditOutlined,
  CaretLeftOutlined,
  CaretDownOutlined,
  EllipsisOutlined
} from '@ant-design/icons-vue';
import {requestMethodOpts, interfaceStatus, interfaceStatusOpts} from '@/config/constant';
import {momentUtc} from '@/utils/datetime';
import {message, Modal} from 'ant-design-vue';
import {
  getInterfaceList,
  saveInterface,
  expireInterface,
  deleteInterface,
  copyInterface,
  getYaml,
  moveCategories,
  getCategories,
  deleteCategories,
  editCategories,
  newCategories
} from './service';
import CreateApiModal from './components/CreateApiModal.vue';
import CreateTagModal from './components/CreateTagModal.vue'
import EditInterfaceDrawer from './components/EditInterfaceDrawer.vue'


import {TreeDataItem, TreeDragEvent, DropEvent} from 'ant-design-vue/es/tree/Tree';

type Key = ColumnProps['key'];

// todo 待处理接口类型定义
interface DataType {
  key: Key;
  name: string;
  age: number;
  address: string;
}

/**
 * 表格数据
 * */
const columns = [
  {
    title: '序号',
    dataIndex: 'index',
  },
  {
    title: '接口名称',
    dataIndex: 'title',
    slots: {customRender: 'colTitle'},
  },
  {
    title: '状态',
    dataIndex: 'status',
    slots: {customRender: 'colStatus'},
  },
  {
    title: '创建人',
    dataIndex: 'createUser',
  },
  {
    title: '接口路径',
    dataIndex: 'path',
  },
  {
    title: '最近更新',
    dataIndex: 'updatedAt',
  },
  {
    title: '操作',
    key: 'operation',
    fixed: 'right',
    width: 100,
    slots: {customRender: 'action'},
  },
];

const searchValue = ref('');
const expandedKeys = ref<string[]>([]);
const autoExpandParent = ref<boolean>(false);
const treeData: any = ref(null)


const data = ref([]);

async function reloadList() {
  let res = await getInterfaceList({
    "prjectId": 0,
    "page": 1,
    "pageSize": 100,
    "status": 0,
    "userId": 0,
    // "title": "接口名称"
  });
  const {result, total} = res.data;
  result.forEach((item, index) => {
    item.index = index + 1;
    item.key = `${index + 1}`;
    item.updatedAt = momentUtc(item.updatedAt);
  })
  data.value = result;
  // TODO 待处理分页逻辑
}

async function loadCategories() {
  let res = await getCategories({
    currProjectId: 1,
    serveId: 1,
    moduleId: 2
  });
  if (res.code === 0 && res.data) {
    // const data = [res.data];
    const data = [
      {
        "id": '1',
        "name": "根节点",
        "desc": "",
        "parentId": '0',
        "children": [
          {
            "id": '1-1',
            "name": "目录1-1",
            "desc": "",
            "parentId": '1',
            "children": [
              {
                "id": '1-1-1',
                "name": "目录1-1-1",
                "desc": "",
                "parentId": '1-1',
                "children": null,
              },
              {
                "id": '1-1-2',
                "name": "目录1-1-2",
                "desc": "",
                "parentId": '1-1',
                "children": null,
              }
            ]
          },
          {
            "id": '1-2',
            "name": "目录1-2",
            "desc": "",
            "parentId": '1',
            "children": [
              {
                "id": '1-2-1',
                "name": "目录1-2-1",
                "desc": "",
                "parentId": '1-2',
                "children": null,
              },
              {
                "id": '1-2-2',
                "name": "目录1-2-2",
                "desc": "",
                "parentId": '1-2',
                "children": [
                  {
                    "id": '1-2-2-1',
                    "name": "目录1-2-2-1",
                    "desc": "",
                    "parentId": '1-2-2',
                    "children": null,
                  },
                  {
                    "id": '1-2-2-2',
                    "name": "目录1-2-2-2",
                    "desc": "",
                    "parentId": '1-2-2',
                    "children": null,
                  }
                ]
              }
            ]
          }
        ],
        "slots": {
          "icon": "icon"
        }
      }
    ]

    // eslint-disable-next-line no-inner-declarations
    function fn(arr: any) {
      if (!Array.isArray(arr)) {
        return;
      }
      arr.forEach((item, index) => {
        item.key = item.id;
        item.title = item.name;
        if (Array.isArray(item.children)) {
          fn(item.children)
        }
      });
    }

    fn(data);
    treeData.value = data;
  }
}

watch(() => {
  return searchValue.value
}, (newVal) => {

  // const expanded = treeData.value
  //     .map((item: any) => {
  //       if ((item.title as string).indexOf(value) > -1) {
  //         return getParentKey(item.key as string, treeData.value);
  //       }
  //       return null;
  //     })
  //     .filter((item, i, self) => item && self.indexOf(item) === i);

  console.log(832, newVal)

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
    console.log(832, parentIds);
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

function selectTreeItem(selectedKeys) {
  console.log(832, selectedKeys);
  // ::::TODO 发送请求

}


function findNodeById(id, tree) {
  if (tree.id === id) {
    return tree;
  }
  if (tree.children) {
    for (const child of tree.children) {
      const node = findNodeById(id, child);
      if (node) {
        return node;
      }
    }
  }
  return null;
}

function deleteNodeById(id, tree: any) {
  if (!tree.children) {
    return;
  }
  tree.children = tree.children.filter(child => child.id !== id);

  if (tree.children) {
    for (const child of tree.children) {
      deleteNodeById(id, child);
    }
  }
}

const currentNode = ref(null);
const tagModalMode = ref('new');

// 删除分类
async function deleteCategorie(node) {
  Modal.confirm({
    title: () => 'Are you sure delete this task?',
    content: () => 'Some descriptions',
    okText: () => 'Yes',
    okType: 'danger',
    cancelText: () => 'No',
    onOk: async () => {
      const res = await deleteCategories({
        id: node.id
      });
      if (res.data.code === 0) {
        deleteNodeById(node.id, treeData.value[0]);
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

function moveNode(fromId, toId, tree) {
  if (tree.id === fromId) {
    const nodeToMove = findNodeById(fromId, tree);
    const parent = findNodeById(toId, tree);
    if (parent) {
      parent.children.push(nodeToMove);
      deleteNodeById(fromId, tree);
    }
    return;
  }
  if (tree.children) {
    for (const child of tree.children) {
      moveNode(fromId, toId, child);
    }
  }
}

function swapNodes(firstId, secondId, tree) {
  if (tree.children) {
    const firstNode = findNodeById(firstId, tree);
    const secondNode = findNodeById(secondId, tree);
    if (firstNode && secondNode) {
      const firstIndex = tree.children.findIndex(node => node.id === firstId);
      const secondIndex = tree.children.findIndex(node => node.id === secondId);
      if (firstIndex !== -1 && secondIndex !== -1) {
        tree.children[firstIndex] = secondNode;
        tree.children[secondIndex] = firstNode;
      }
    }
    for (const child of tree.children) {
      swapNodes(firstId, secondId, child);
    }
  }
}


function onDrop(info: DropEvent) {
  console.log(832, info.dragNode);
  console.log(832, info.node);
  const dragKey = info.dragNode.eventKey;
  const dropKey = info.node.eventKey;
  // ::::根据是否同级，如果同级则移动。
  moveNode(dragKey, dropKey, treeData.value[0]);
  swapNodes(dragKey, dropKey, treeData.value[0])
  console.log(832, dragKey, dropKey);
}

const onDragEnter = (info: TreeDragEvent) => {
  // console.log(832,info);
  console.log(832);
  // expandedKeys 需要展开时
  // expandedKeys.value = info.expandedKeys
};


onMounted(async () => {
  await reloadList();
  await loadCategories();
})


async function refreshList() {
  await reloadList();
}

// const selectedRowKeys: Key[] = ref([]);
const selectedRowKeys = ref<Key[]>([]);
const loading = false;

// 是否批量选中了
// const hasSelected = computed(() => state.selectedRowKeys.length > 0);
const hasSelected = false;

// 抽屉是否打开
const drawerVisible = ref<boolean>(false);

const onSelectChange = (keys: Key[], rows: any) => {
  console.log('selectedRowKeys changed: ', keys, rows);
  selectedRowKeys.value = [...keys];
};

const editInterfaceId = ref('');


const clickTag = ref(0);

/**
 * 接口编辑
 * */
function editInterface(record) {
  console.log('editInterface');
  editInterfaceId.value = record.id;
  drawerVisible.value = true;
  clickTag.value++;
}

/**
 * 批量操作
 * */
function batchHandle() {
  console.log('batchHandle')
}

async function copy(record: any) {
  let res = await copyInterface(record.id);
  if (res.code === 0) {
    message.success('复制成功');
    await reloadList();

  }
}

async function disabled(record: any) {
  let res = await expireInterface(record.id);
  if (res.code === 0) {
    message.success('置为无效成功');
    await reloadList();
  }
}

/**
 * 删除接口
 * */
async function del(record: any) {
  let res = await deleteInterface(record.id);
  if (res.code === 0) {
    message.success('删除成功');
    await reloadList();
  }
}

/**
 * 接口导入逻辑
 * */
function importApi() {
  console.log('导入')
}

/**
 * 关闭抽屉
 * */
function onCloseDrawer() {
  drawerVisible.value = false;
}


const createTagModalvisible = ref(false);
const createApiModalvisible = ref(false);

/**
 * 添加接口分类
 * */
function addApiTag() {
  createTagModalvisible.value = true;
}

/**
 * 添加接口
 * */
function addApi() {
  createApiModalvisible.value = true;
}


async function handleCreateApi(data) {
  let res = await saveInterface({
    "serveId": 1,
    "path": data.path,
    "title": data.title,
  });
  createApiModalvisible.value = false;
  if (res.code === 0) {
    message.success('新建接口成功');
    await reloadList();
  }
}

function handleCancalCreateApi() {
  createApiModalvisible.value = false;
}

async function handleCreateTag(obj) {
  const res = await editCategories({
    "id": obj.id,
    "name": obj.name,
    "desc": obj.description,
    "parent": obj.parentId
  });
  if (res.code === 0) {
    createTagModalvisible.value = false;
    message.success('修改成功');
  } else {
    message.success('修改失败');
  }
}

function handleCancalCreateTag() {
  createTagModalvisible.value = false;
}


</script>

<style scoped lang="less">
.container {
  margin: 16px;
  background: #ffffff;

  //:deep(.ant-tree-switcher-noop) {
  //  display: none;
  //}
}

.tag-filter-form {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 60px;

  .search-input {
    margin-left: 8px;
    margin-right: 8px;
  }

  .add-btn {
    margin-left: 12px;
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

.more-icon {
  position: absolute;
  right: 8px;
}


</style>
