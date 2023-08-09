<template>
  <div class="tree-main">
    <div class="tree-filters">
      <a-select 
        style="margin-right: 20px; width: 100%"
        :bordered="true"
        :placeholder="'请选择服务'"
        v-model:value="serveId"
        @change="selectServe">
        <a-select-option v-for="item in serves" :key="item.id" :value="item.id">{{ item.name }}</a-select-option>
      </a-select>
      <a-input-search
        style="width: 100%"
        placeholder="请输入关键词"
        enter-button
        v-model:value="searchValue"
      />
    </div>

    <div class="tree-container">
      <a-tree
          class="deeptest-tree"
          showIcon
          checkable
          :tree-data="treeData"
          :replaceFields="fieldNames"
          @check="onChecked"
          :defaultExpandAll="true"
      >

        <template #switcherIcon>
          <CaretDownOutlined/>
        </template>

        <template #title="nodeProps">
          <span v-if="nodeProps.dataRef.type == 'dir' || nodeProps.dataRef.type == ''"><FolderOpenOutlined/> {{nodeProps.dataRef.name+' ('+nodeProps.dataRef.count+')'}}</span>
          <span v-if="nodeProps.dataRef.type == 'endpoint'"><ApiOutlined /> {{nodeProps.dataRef.name}}</span>
          <span v-if="nodeProps.dataRef.type == 'case'"><ShareAltOutlined /> {{nodeProps.dataRef.name}}
            <a-tag class="method-tag" :color="getMethodColor(nodeProps.dataRef.method || 'GET', nodeProps.dataRef.disable)">{{
                      nodeProps.dataRef.method || "GET"
                    }}</a-tag>
          </span>
          <!--
                      <div class="tree-title" :draggable="nodeProps?.dataRef?.id === -1">
                          <span class="tree-title-text" v-if="nodeProps?.dataRef?.name.indexOf(searchValue) > -1">
                            <span>{{ nodeProps?.dataRef?.name.substr(0, nodeProps?.dataRef?.name.indexOf(searchValue)) }}</span>
                            <span style="color: #f50">{{ searchValue }}</span>
                            <span>{{ nodeProps?.dataRef?.name.substr(nodeProps?.dataRef?.name.indexOf(searchValue) + searchValue.length) }}</span>
                          </span>
                        <span class="tree-title-text" v-else>{{ nodeProps?.dataRef?.name }}</span>
                      </div>
                        --->
        </template>
  
      </a-tree>

      <div v-if="!treeData.length" class="nodata-tip">
        <div class="empty-container">
          <img src="@/assets/images/empty.png" alt="">
          <span>暂无数据</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, onMounted, ref, watch} from 'vue';
import {CaretDownOutlined,FolderOpenOutlined,ApiOutlined,ShareAltOutlined} from '@ant-design/icons-vue';
import {useStore} from "vuex";

import {StateType as ProjectStateType} from "@/store/project";
import {StateType as ServeStateType} from "@/store/serve";

import {listServe} from "@/services/serve";
import {getSelectedTreeNode,filterByKeyword} from "@/utils/tree";
import {getMethodColor} from "@/utils/dom";
import cloneDeep from "lodash/cloneDeep";


const store = useStore<{ Endpoint: any, ProjectGlobal: ProjectStateType, ServeGlobal: ServeStateType, DiagnoseInterface }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const currServe = computed<any>(() => store.state.ServeGlobal.currServe);

const treeData = computed<any>(() => {
  const children = cloneDeep(store.state.Endpoint.caseTree);
  if (children?.length > 0) {
    return [...filterByKeyword(children, searchValue.value, 'name')];
  }
  return []
})
const treeDataMap = computed<any>(() => store.state.Endpoint.caseTreeMap);


const props = defineProps({
  selectInterfaces: {
    type: Function,
    required: true,
  },
})

const fieldNames = {
  title: 'name',
  key: 'id',
}

const serves = ref([] as any[]);
const serveId = ref(0)


const onChecked = (checkedKeys, e) => {
  console.log('onChecked', checkedKeys,treeDataMap.value)  
  const selectedNodes = getSelectedTreeNode(checkedKeys, treeDataMap.value)
  props.selectInterfaces(selectedNodes)

  console.log('selectedNodes', selectedNodes)
}


const loadServe = async () => {
  listServe().then((json) => {
    serves.value = json.data.serves

    if (serves.value.length > 0) {
      serveId.value = serves.value[0].id
    }
  })
}

onMounted(() => {
  loadServe()
})

const searchValue = ref('');
const expandedKeys = ref<number[]>([]);

async function loadTreeData(serveId:number) {
  if (currProject?.value?.id > 0 ) {
    await store.dispatch('Endpoint/getCaseTree', {currProjectId: currProject.value.id, serveId: serveId});
   // expandAll();
  }
}


const selectServe = () => {
  console.log('selectServe', serveId.value)
  loadTreeData(serveId.value)
}


watch((currServe.value), async (newVal) => {
  console.log('watch currProject', currProject?.value.id, currServe?.value.id)
  await loadTreeData(currServe?.value.id);
}, {
  immediate: true
})


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
    margin-bottom: 20px;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .tree-container {
    background: #ffffff;
    max-height: 400px;
    overflow-y: scroll;

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

.empty-container {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 40px;

    img {
        width: 90px;
        height: 90px;
        margin-bottom: 8px;
    }

    span {
        font-size: 12px;
        line-height: 20px;
        text-align: center;
        color: rgba(0, 0, 0, 0.46);
    }
}
</style>
