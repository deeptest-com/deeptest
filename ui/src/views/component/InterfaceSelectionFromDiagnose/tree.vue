<template>
  <div class="tree-main">
    <div class="tree-filters">
      <a-select 
        style="width: 100%;margin-bottom: 20px" 
        :bordered="true"
        :placeholder="'请选择服务'"
        v-model:value="serveId"
        @change="selectServe">
        <a-select-option v-for="item in serves" :key="item.id" :value="item.id">{{ item.name }}</a-select-option>
      </a-select>
      <a-input-search
        style="display: flex;justify-content: end;width: 100%;margin-bottom: 16px; "
        placeholder="请输入关键词"
        enter-button
        v-model:value="searchValue"/>
    </div>

    <div class="tree-container">
      <a-tree
          class="deeptest-tree"
          showIcon
          :checkable="true"
          :expandedKeys="expandedKeys"
          :auto-expand-parent="autoExpandParent"
          :tree-data="treeData"
          @check="onChecked"
          :replace-fields="replaceFields">

        <template #switcherIcon>
          <CaretDownOutlined/>
        </template>

        <template #title="nodeProps">
          <div class="tree-title" :draggable="nodeProps.dataRef.id === -1">
              <span class="tree-title-text" v-if="nodeProps.dataRef.title.indexOf(searchValue) > -1">
                <span>{{ nodeProps.dataRef.title.substr(0, nodeProps.dataRef.title.indexOf(searchValue)) }}</span>
                <span style="color: #f50">{{ searchValue }}</span>
                <span>{{ nodeProps.dataRef.title.substr(nodeProps.dataRef.title.indexOf(searchValue) + searchValue.length) }}</span>
              </span>
            <span class="tree-title-text" v-else>{{ nodeProps.dataRef.title }}</span>
          </div>
        </template>
      </a-tree>

      <div v-if="!treeData.length" class="nodata-tip">
        <Empty />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, onMounted, ref, watch} from 'vue';
import {CaretDownOutlined,} from '@ant-design/icons-vue';
import {useStore} from "vuex";

import Empty from "@/components/Empty";
import {StateType as ProjectStateType} from "@/store/project";
import {StateType as DiagnoseInterfaceStateType} from '@/views/diagnose/store';
import {StateType as ServeStateType} from "@/store/serve";

import {listServe} from "@/services/serve";
import {filterTree, getSelectedTreeNode} from "@/utils/tree";
import {isInArray} from "@/utils/array";

const store = useStore<{ DiagnoseInterface: DiagnoseInterfaceStateType, ProjectGlobal: ProjectStateType, ServeGlobal: ServeStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const currServe = computed<any>(() => store.state.ServeGlobal.currServe);

const treeData = computed<any>(() => store.state.DiagnoseInterface.treeData);
const treeDataMap = computed<any>(() => store.state.DiagnoseInterface.treeDataMap);

const props = defineProps({
  selectInterfaces: {
    type: Function,
    required: true,
  },
})

const serves = ref([] as any[]);
const serveId = ref(0)

const onChecked = (checkedKeys, e) => {
  console.log('onChecked', checkedKeys, e.checkedNodes)

  const selectedNodes = getSelectedTreeNode(checkedKeys, treeDataMap.value)
  props.selectInterfaces(selectedNodes)

  console.log('selectedNodes', selectedNodes)
}
const getChildren = (node, mp) => {
  mp[node.id] = true

  if (node.children) {
    node.children.forEach((child, index) => {
      getChildren(child, mp)
    })
  }
}

const loadServe = async () => {
  listServe().then((json) => {
    serves.value = json.data.serves

    if (serves.value.length > 0) {
      serveId.value = serves.value[0].id
    }
  })
}
loadServe()

const replaceFields = {key: 'id'};
const searchValue = ref('');
const expandedKeys = ref<number[]>([]);
const autoExpandParent = ref<boolean>(false);

async function loadTreeData() {
  if (currProject?.value?.id > 0 && currServe?.value?.id > 0) {
    await store.dispatch('DiagnoseInterface/loadTree', {projectId: currProject.value.id, serveId: currServe.value.id});
    expandAll();
  }
}

async function getServeServers() {
  await store.dispatch('DiagnoseInterface/getServeServers', {
    id: currServe.value.id,
  })
}

const selectServe = () => {
  console.log('selectServe', serveId.value)
}

watch((currProject), async (newVal) => {
  console.log('watch currProject', currProject?.value.id, currServe?.value.id)
  await loadTreeData();
  await getServeServers()
}, {
  immediate: true
})

watch(searchValue, (newVal) => {
  expandedKeys.value = filterTree(treeData.value, newVal)
  console.log('searchValue', expandedKeys.value)
  autoExpandParent.value = true;
});

// 展开所有
function expandAll() {
  const keys: any = [];
  const data = treeData.value;

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

onMounted(async () => {
  console.log('onMounted')
})

</script>

<style scoped lang="less">
.tree-main {
  .tree-filters {
    margin-bottom: 0;
  }

  .tree-container {
    background: #ffffff;
    max-height: 400px;
    overflow-y: hidden;

    .tree-title {
      position: relative;

      .tree-title-text {
        display: inline-block;
        white-space: nowrap;
      }

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
  }
}

</style>
