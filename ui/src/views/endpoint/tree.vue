<template>
  <div class="tree-container">
    <div class="tree-con">
      <div class="tag-filter-form">
        <a-input-search
            class="search-input"
            v-model:value="searchValue"
            placeholder="搜索接口分类"/>
        <div class="add-btn" @click="addApiTag">
          <PlusOutlined style="font-size: 16px;"/>
        </div>
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
    <!--  创建接口 Tag  -->
    <CreateTagModal
        :visible="createTagModalvisible"
        :nodeInfo="currentNode"
        :mode="tagModalMode"
        @cancal="handleCancalTagModalCancal"
        @ok="handleTagModalOk"/>
  </div>
</template>
<script setup lang="ts">
import {
  computed, reactive, toRefs, ref, onMounted,
  watch
} from 'vue';
import {
  PlusOutlined,
  CaretDownOutlined,
  EllipsisOutlined
} from '@ant-design/icons-vue';
import {message, Modal} from 'ant-design-vue';

import CreateTagModal from './components/CreateTagModal.vue'


import {TreeDataItem, TreeDragEvent, DropEvent} from 'ant-design-vue/es/tree/Tree';


const createTagModalvisible = ref(false);


const searchValue = ref('');
const expandedKeys = ref<string[]>([]);
const autoExpandParent = ref<boolean>(false);
const treeData: any = ref(null)


const data = ref([]);


async function loadCategories() {
  // let res = await getCategories({
  //   currProjectId: 1,
  //   serveId: 1,
  //   moduleId: 2
  // });
  // if (res.code === 0 && res.data) {


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

watch(
    () => {
      return searchValue.value
    },
    (newVal) => {
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
      // const res = await deleteCategories({
      //   id: node.id
      // });
      // if (res.data.code === 0) {
      //   deleteNodeById(node.id, treeData.value[0]);
      //   message.success('删除成功');
      // } else {
      //   message.success('删除失败');
      // }
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
  // if (tagModalMode.value === 'edit') {
  //   const res = await editCategories({
  //     "id": obj.id,
  //     "name": obj.name,
  //     "desc": obj.description,
  //     "parent": obj.parentId
  //   });
  //   if (res.code === 0) {
  //     createTagModalvisible.value = false;
  //     // 修改数据
  //     let targetNode = findNodeById(obj.id, treeData.value[0]);
  //     targetNode = {
  //       ...targetNode,
  //       "name": obj.name,
  //       "desc": obj.description,
  //     };
  //     message.success('修改成功');
  //   } else {
  //     message.success('修改失败');
  //   }
  //   //  创建
  // } else {
  //   const res = await newCategories({
  //     "name": obj.name,
  //     "Mode": "child",
  //     "targetId": obj.id,
  //     "projectId": 0,
  //     "serveId": 0,
  //     "moduleId": "2"
  //   });
  //   if (res.code === 0) {
  //     let targetNode = findNodeById(obj.id, treeData.value[0]);
  //     targetNode.children.push({
  //       ...res.data
  //     })
  //     createTagModalvisible.value = false;
  //     message.success('新建成功成功');
  //   } else {
  //     message.success('新建失败失败');
  //   }
  // }
}

function handleCancalTagModalCancal() {
  createTagModalvisible.value = false;
}


/**
 * 添加接口分类
 * */
function addApiTag() {
  createTagModalvisible.value = true;
}


async function onDrop(info: DropEvent) {
  // console.log(info);
  // const dropKey = info.node.eventKey;
  // const dragKey = info.dragNode.eventKey;
  // const dropPos = info.node.pos.split('-');
  // const dropPosition = info.dropPosition - Number(dropPos[dropPos.length - 1]);
  //
  // const res = await moveCategories({
  //   "currProjectId": 1,
  //   "dragKey": 10,
  //   "dropKey": 7,
  //   "dropPos": 1
  // });
  //
  // if (res.code !== 0) {
  //   return;
  // }
  //
  // const loop = (data: TreeDataItem[], key: string, callback: any) => {
  //   data.forEach((item, index, arr) => {
  //     if (item.key === key) {
  //       return callback(item, index, arr);
  //     }
  //     if (item.children) {
  //       return loop(item.children, key, callback);
  //     }
  //   });
  // };
  // const data = [...treeData.value];
  // // Find dragObject
  // let dragObj: TreeDataItem = {};
  // loop(data, dragKey, (item: TreeDataItem, index: number, arr: TreeDataItem[]) => {
  //   arr.splice(index, 1);
  //   dragObj = item;
  // });
  // if (!info.dropToGap) {
  //   // Drop on the content
  //   loop(data, dropKey, (item: TreeDataItem) => {
  //     item.children = item.children || [];
  //     // where to insert 示例添加到尾部，可以是随意位置
  //     item.children.push(dragObj);
  //   });
  // } else if (
  //     (info.node.children || []).length > 0 && // Has children
  //     info.node.expanded && // Is expanded
  //     dropPosition === 1 // On the bottom gap
  // ) {
  //   loop(data, dropKey, (item: TreeDataItem) => {
  //     item.children = item.children || [];
  //     // where to insert 示例添加到尾部，可以是随意位置
  //     item.children.unshift(dragObj);
  //   });
  // } else {
  //   let ar: TreeDataItem[] = [];
  //   let i = 0;
  //   loop(data, dropKey, (item: TreeDataItem, index: number, arr: TreeDataItem[]) => {
  //     ar = arr;
  //     i = index;
  //   });
  //   if (dropPosition === -1) {
  //     ar.splice(i, 0, dragObj);
  //   } else {
  //     ar.splice(i + 1, 0, dragObj);
  //   }
  // }
  // treeData.value = data;
}


onMounted(async () => {
  await loadCategories();
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
