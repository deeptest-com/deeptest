<template>
  <div class="container">
    <div class="content">
      <div class="left tree">
        <div class="tag-filter-form">
          <a-input-search
              class="search-input"
              v-model:value="searchValue"
              placeholder="搜索接口分类"/>
          <div
              @click="addApiTag"
              class="add-btn"
              type="link">
            <PlusOutlined style="font-size: 12px;"/>
          </div>
        </div>
        <div style="margin: 0 8px;">
          <a-tree
              v-if="treeData"
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
          <div v-else style="padding: 4px 32px 0 0;">
            <a-alert
                message="提示"
                closable
                description="您还未有添加接口分类，请点击上方添加按钮添加！"
                type="info"
                show-icon
            />
          </div>
        </div>
      </div>
      <!--  头部搜索区域  -->
      <div class="right">
        <div class="top-action">
          <a-row type="flex" style="align-items: center;width: 100%">
            <a-col :span="4">
              <a-button class="action-new" type="primary" :loading="loading" @click="addApi">新建接口</a-button>
              <a-button class="action-import" type="primary" :disabled="!hasSelected" :loading="loading"
                        @click="importApi">
                批量操作
              </a-button>
            </a-col>
          </a-row>
        </div>
        <div class="top-search">
          <a-row type="flex" :gutter="16" style="width: 100%" justify="space-between" >
            <a-col :span="4">
              <a-form-item label="创建人" style="margin-bottom: 0">
                <a-select style="width: 140px;"   placeholder="请选择创建人">
                  <a-select-option value="admin"> admin</a-select-option>
                  <a-select-option value="superAdmin">super admin</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :span="4">
              <a-form-item label="状态" style="margin-bottom: 0;">
                <a-select  style="width: 140px;"  placeholder="请选择状态" :options="interfaceStatusOpts"/>
              </a-form-item>
            </a-col>
            <a-col :span="8">
              <a-form-item label="服务版本" style="margin-bottom: 0;">
                <a-select placeholder="选择服务" style="margin-right: 8px;width: 140px;" :options="interfaceStatusOpts"/>
                <a-select placeholder="选择服务版本" style="width: 140px;"  :options="serviceOptions"/>
              </a-form-item>
            </a-col>
            <a-col :span="8" >
              <a-input-search
                  style="width: 300px;display: flex;justify-content: end;"
                  placeholder="请输入关键词"
                  enter-button
                  @search="() => {

                  }"
              />
            </a-col>
          </a-row>
        </div>
        <a-table
            style="margin: 0 16px;"
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
              <!-- ::::todo 不同状态对应不同颜色 -->
              <a-tag color="red">{{ interfaceStatus.get(record.status) }}</a-tag>
            </div>
          </template>
          <template #colPath="{text}">
            <div class="customTitleColRender">
              <a-tag >{{ text }}</a-tag>
            </div>
          </template>
          <template #action="{record}">
            <a-dropdown @click="handleButtonClick">
              <MoreOutlined/>
              <template #overlay>
                <a-menu @click="handleMenuClick">
                  <a-menu-item key="1">
                    <a-button style="width: 80px" type="link" size="small" @click="copy(record)">复制</a-button>
                  </a-menu-item>
                  <a-menu-item key="2">
                    <a-button style="width: 80px" type="link" size="small"  @click="del(record)">删除</a-button>
                  </a-menu-item>
                  <a-menu-item key="3">
                    <a-button style="width: 80px" type="link" size="small"  @click="disabled(record)">过时</a-button>
                  </a-menu-item>
                </a-menu>
              </template>
            </a-dropdown>
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
        @cancal="handleCancalTagModalCancal"
        @ok="handleTagModalOk"/>
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
  InfoCircleOutlined,
  CaretDownOutlined,
  EllipsisOutlined,
  MoreOutlined
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
  newCategories,
} from './service';
import {
  getServeList
} from '../../views/projectSetting/service';
import CreateApiModal from './components/CreateApiModal.vue';
import CreateTagModal from './components/CreateTagModal.vue'
import EditInterfaceDrawer from './components/EditInterfaceDrawer.vue'
import {TreeDataItem, TreeDragEvent, DropEvent} from 'ant-design-vue/es/tree/Tree';


import {StateType as ProjectStateType} from "@/store/project";
import {useStore} from "vuex";
const store = useStore<{ ProjectGlobal: ProjectStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);


const createTagModalvisible = ref(false);
const createApiModalvisible = ref(false);
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
    slots: {customRender: 'colPath'},
  },
  {
    title: '所属服务',
    dataIndex: 'serveName',
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
/*************************************************
 * :::: 左侧区域按接口分类搜索树形结构 逻辑  start
 ************************************************/
const searchValue = ref('');
const expandedKeys:any = ref<string[]>([]);
const autoExpandParent = ref<boolean>(true);
const treeData: any = ref(null)
const data = ref([]);

async function loadCategories() {
  let res = await getCategories({
    currProjectId: currProject.value.id,
    serveId: 1,
    moduleId: 2
  });
  if (res.code === 0 && res.data) {
    const allKeys:any = [];
    // eslint-disable-next-line no-inner-declarations
    function fn(arr: any) {
      if (!Array.isArray(arr)) {
        return;
      }
      arr.forEach((item, index) => {
        item.key = item.id;
        allKeys.push(item.id);
        item.title = item.name;
        if (Array.isArray(item.children)) {
          fn(item.children)
        }
      });

    }

    fn([res.data]);
    treeData.value = [res.data];
    expandedKeys.value = [...new Set(allKeys)];
  } else {
    treeData.value = null;
  }
}

watch(
    () => {
      return searchValue.value
    }, (newVal) => {
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
  // ::::TODO 发送请求 待确定参数
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
    title: () => '提示',
    content: () => '确定，删除该分类吗？',
    okText: () => '确定',
    okType: 'danger',
    cancelText: () => '取消',
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

async function handleTagModalOk(obj) {
  // 修改
  if (tagModalMode.value === 'edit') {
    const res = await editCategories({
      "id": obj.id,
      "name": obj.name,
      "desc": obj.description,
      "parent": obj.parentId
    });
    if (res.code === 0) {
      createTagModalvisible.value = false;
      // 修改数据
      let targetNode = findNodeById(obj.id, treeData.value[0]);
      targetNode = {
        ...targetNode,
        "name": obj.name,
        "desc": obj.description,
      };
      message.success('修改成功');
    } else {
      message.success('修改失败');
    }
    //  创建
  } else {
    const res = await newCategories({
      "name": obj.name,
      "Mode": "child",
      "targetId": obj.id,
      "projectId": currProject.value.id,
      "serveId": 0,
      "moduleId": "2"
    });
    if (res.code === 0) {
      let targetNode = findNodeById(obj.id, treeData.value[0]);
      targetNode.children.push({
        ...res.data
      })
      createTagModalvisible.value = false;
      message.success('新建成功成功');
    } else {
      message.success('新建失败失败');
    }
  }
}

function handleCancalTagModalCancal() {
  createTagModalvisible.value = false;
}

function addApiTag() {
  createTagModalvisible.value = true;
}

async function onDrop(info: DropEvent) {
  // console.log(info);
  const dropKey = info.node.eventKey;
  const dragKey = info.dragNode.eventKey;
  const dropPos = info.node.pos.split('-');
  const dropPosition = info.dropPosition - Number(dropPos[dropPos.length - 1]);
  const res = await moveCategories({
    "currProjectId": currProject.value.id,
    "dragKey": 10,
    "dropKey": 7,
    "dropPos": 1
  });
  if (res.code !== 0) {
    return;
  }
  const loop = (data: TreeDataItem[], key: string, callback: any) => {
    data.forEach((item, index, arr) => {
      if (item.key === key) {
        return callback(item, index, arr);
      }
      if (item.children) {
        return loop(item.children, key, callback);
      }
    });
  };
  const data = [...treeData.value];
  // Find dragObject
  let dragObj: TreeDataItem = {};
  loop(data, dragKey, (item: TreeDataItem, index: number, arr: TreeDataItem[]) => {
    arr.splice(index, 1);
    dragObj = item;
  });
  if (!info.dropToGap) {
    // Drop on the content
    loop(data, dropKey, (item: TreeDataItem) => {
      item.children = item.children || [];
      // where to insert 示例添加到尾部，可以是随意位置
      item.children.push(dragObj);
    });
  } else if (
      (info.node.children || []).length > 0 && // Has children
      info.node.expanded && // Is expanded
      dropPosition === 1 // On the bottom gap
  ) {
    loop(data, dropKey, (item: TreeDataItem) => {
      item.children = item.children || [];
      // where to insert 示例添加到尾部，可以是随意位置
      item.children.unshift(dragObj);
    });
  } else {
    let ar: TreeDataItem[] = [];
    let i = 0;
    loop(data, dropKey, (item: TreeDataItem, index: number, arr: TreeDataItem[]) => {
      ar = arr;
      i = index;
    });
    if (dropPosition === -1) {
      ar.splice(i, 0, dragObj);
    } else {
      ar.splice(i + 1, 0, dragObj);
    }
  }
  treeData.value = data;
}

/*************************************************
 * :::::::: 左侧区域按接口分类搜索树形结构 逻辑  send
 ************************************************/


/*************************************************
 * ::::表格筛选区域 逻辑 start
 ************************************************/

const serviceOptions = ref([]);
async function getServersList() {
  // 请求服务列表
  let res = await getServeList({
    projectId: currProject.value.id,
    "page": 0,
    "pageSize": 100,
  });
  if (res.code === 0) {
    res.data.result.forEach((item) => {
      item.label = item.name;
      item.value = item.id;
    })
    serviceOptions.value = res.data.result;
  }
}



/*************************************************
 * ::::表格筛选区域 逻辑 end
 ************************************************/


/*************************************************
 * ::::表格区域 逻辑 start
 ************************************************/
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
    "projectId": currProject.value.id,
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

async function reloadList() {
  let res = await getInterfaceList({
    "projectId": currProject.value.id,
    "page": 1,
    "pageSize": 100,
    // "status": 0,
    // "userId": 0,
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

/*************************************************
 * ::::表格区域 逻辑 end
 ************************************************/



// 实时监听项目切换，如果项目切换了则重新请求数据
watch(() => {
  return currProject.value;
}, async (newVal) => {
  await reloadList();
  await loadCategories();
  await getServersList();
}, {
  immediate: true
})


async function refreshList() {
  await reloadList();
}


</script>
<style scoped lang="less">
.container {
  margin: 16px;
  background: #ffffff;
  min-height: calc(100vh - 80px);
}

.tag-filter-form {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 60px;

  .search-input {
    margin-left: 8px;
    //margin-right: 8px;
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
    height: calc(100vh - 80px);
  }

  .right {
    flex: 1
  }
}

.action-new {
  margin-right: 8px;
}

.top-search {
  height: 60px;
  display: flex;
  align-items: center;
  margin-left: 16px;
  margin-bottom: 8px;
}

.top-action {
  height: 60px;
  display: flex;
  align-items: center;
  margin-left: 16px;
  margin-top: 8px;
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


::v-deep{
  .ant-alert-info{
    padding: 12px;
  }
  .ant-alert-icon{
    font-size: 14px;
    position: relative;
    top: 4px;
    left: 8px;
  }
  .ant-alert-message{
    font-size: 14px;
  }
  .ant-alert-description{
    font-size: 12px;
  }
}

</style>
