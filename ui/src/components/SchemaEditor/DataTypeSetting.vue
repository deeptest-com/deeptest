<template>
  <a-popover :title="null"
             trigger="click"
             v-model:visible="visible"
             :overlayClassName="'data-type-setting-container'">
    <template #content>
      <div class="content" v-for="(tabs,tabsIndex) in tabsList" :key="tabsIndex">
        <div class="header">
          <div class="item"
               v-for="(tab,tabIndex) in tabs"
               @click="() => {
                 if(isRefChildNode){
                    return;
                 }
                  selectTab(tabs,tabIndex)
               }"
               :class="tab.active ? 'active' : ''"
               :key="tab.value">
            {{ tabsIndex === 0 ? tab.label : tab.subLabel }}
          </div>
        </div>
        <div class="main">
          <div class="item"
               v-for="(tab,tabIndex) in tabs"
               v-show="tab.active"
               :key="tab.value">
            <a-radio-group
                :size="'small'"
                class="select-type-btn"
                v-if="tab.active"
                :disabled="isRefChildNode"
                v-model:value="tab.value"
                @change="(event) => changeType(tabsIndex, event)"
                button-style="solid">
              <a-radio-button
                  v-for="item in tab.props"
                  :key="item.value"
                  :value="item.value">{{ item.label }}
              </a-radio-button>
            </a-radio-group>
            <a-form :layout="'vertical'" v-if="tab.type === 'type' && tab.active">
              <div v-for="(item,itemIndex) in tab.props" :key="itemIndex">
                <div v-if="item.value === tab.value">
                  <div class="card-title">{{ item.props.label }}</div>
                  <a-row
                      class="card-content"
                      type="flex"
                      justify="space-between"
                      align="top">
                    <a-col class="col" v-for="opt in item.props.options" :span="11" :key="opt.name">
                      <a-form-item
                          class="col-form-item"
                          :labelAlign="'right'"
                          :label="opt.label">
                        <a-select
                            :disabled="isRefChildNode"
                            v-if="opt.component === 'selectTag'"
                            v-model:value="opt.value"
                            mode="tags"
                            :placeholder="opt.placeholder"
                        />
                        <a-select
                            v-if="opt.component === 'select'"
                            v-model:value="opt.value"
                            :disabled="isRefChildNode"
                            :options="opt.options"
                            :placeholder="opt.placeholder"
                        />
                        <a-input
                            v-if="opt.component === 'input'"
                            v-model:value="opt.value"
                            :disabled="isRefChildNode"
                            :placeholder="opt.placeholder"
                        />
                        <a-input-number
                            v-if="opt.component === 'inputNumber'"
                            id="inputNumber"
                            :disabled="isRefChildNode"
                            :placeholder="opt.placeholder"
                            v-model:value="opt.value"
                        />
                        <a-switch
                            v-if="opt.component === 'switch'"
                            :disabled="isRefChildNode"
                            v-model:checked="opt.value"/>
                      </a-form-item>
                    </a-col>
                  </a-row>
                </div>
              </div>
            </a-form>
            <a-form :layout="'vertical'" style="margin-bottom: 16px;" v-if="tab.type === '$ref' && tab.active">
              <a-form-item
                  class="col-form-item"
                  :labelAlign="'right'"
                  :label="'请选择组件'">
                <a-select
                    :options="refsOptions"
                    :disabled="isRefChildNode"
                    @change="(e) => {
                      changeRef(tabsIndex,tabIndex,e);
                    }"
                    show-search
                    allowClear
                    @search="searchRefs"
                    :value="tab.value || null"
                    placeholder="Select Components"
                    style="width: 100%"/>
              </a-form-item>
            </a-form>
          </div>
        </div>
      </div>
    </template>
    <a href="javascript:void(0)">
      {{ typesLabel }}
      <LinkOutlined v-if="props?.value?.ref"/>
    </a>
  </a-popover>
</template>
<script lang="ts" setup>
import {ref, defineProps, defineEmits, watch, reactive, toRaw, computed, onMounted} from 'vue';
import {
  LinkOutlined
} from '@ant-design/icons-vue';
import {schemaSettingInfo, typeOpts} from "./config";
import cloneDeep from "lodash/cloneDeep";
import {useStore} from "vuex";
import {StateType as ServeStateType} from "@/store/serve";
import debounce from "lodash.debounce";

const props = defineProps({
  value: {
    required: true,
    type: Object
  },
  serveId: {
    required: true,
    type: Array
  },
  isRefChildNode: {
    required: true,
    type: Boolean
  }
})
const emit = defineEmits(['change']);
const tabsList: any = ref([]);
const visible: any = ref(false);
// 返回，如何展示类型
const typesLabel: any = computed(() => {
  let {type, types} = props.value || {};
  type = props?.value?.name || type || '';
  if (!type) {
    return 'null';
  }
  const labels = Array.isArray(types) ? [...types, type] : [type];
  const result = labels.reduceRight((acc, cur, index) => {
    if (index === labels.length - 1) {
      return [cur];
    }
    return [cur, acc];
  }, []);
  return JSON.stringify(result).replace(/[",]/g, '').replace(/^\[/, '').replace(/\]$/, '');
});

function changeType(tabsIndex: any, e: any) {
  let type = e.target.value;
  if (type === 'array') {
    if (tabsList.value.length === tabsIndex + 1) {
      tabsList.value.push(cloneDeep(schemaSettingInfo));
    }
  } else {
    if (tabsIndex < tabsList.value.length) {
      tabsList.value.splice(tabsIndex + 1);
    }
  }
}

// ref 组件
function changeRef(tabsIndex, tabIndex, e) {
  tabsList.value[tabsIndex][tabIndex].value = e;
  // 选中的是ref，则需要隐藏其他的选择
  if (e) {
    tabsList.value.splice(tabsIndex + 1);
  }
}

function selectTab(tabs: any, tabIndex: number) {
  tabs.forEach((tab: any, index: number) => {
    tab.active = tabIndex === index;
  })
  // 切换成普通选择模式时，如果是选中的是数组，则需要添加一个tab
  if(tabIndex === 0 && tabs[tabIndex].value === 'array' && tabsList.value.length === 1){
    tabsList.value.push(cloneDeep(schemaSettingInfo));
  }
}

function initTabsList(types: any, treeInfo: any) {
  let tabsList: any = [];
  types.forEach((type: string) => {
    const defaultTabs: any = cloneDeep(schemaSettingInfo);
    if (typeOpts.includes(type)) {
      defaultTabs[0].active = typeOpts.includes(type);
      defaultTabs[0].value = type;
      const activeTabProps = defaultTabs[0].props.find((prop: any) => prop.value === type);
      activeTabProps?.props?.options?.forEach((opt: any) => {
        opt.value = treeInfo[opt.name] || opt.value;
      })
    } else {
      defaultTabs[0].active = false;
      defaultTabs[0].value = treeInfo?.type || 'string';
      defaultTabs[1].active = true;
      defaultTabs[1].value = treeInfo?.ref;
    }
    tabsList.push(defaultTabs)
  });
  // 如果是数组，还需加一项
  if (types[types.length - 1] === 'array') {
    const arrayItems: any = cloneDeep(schemaSettingInfo);
    tabsList.push(arrayItems);
  }
  return tabsList;
}

function getValueFromTabsList(tabsList: any) {
  const result: any = [];
  // debugger;
  tabsList.forEach((tabs: any) => {
    let activeTab = tabs.find((tab: any) => tab.active);
    // debugger
    // 如果 activeTab.type === '$ref'，则说明是引用类型, 还需要判断是否有值，没有值还是展示基本类型
    if (activeTab.type === '$ref' && !activeTab.value) {
      activeTab = tabs[0];
    }
    let res: any = {};
    if (activeTab.type === '$ref') {
      const selectedRef: any = refsOptions.value.find((ref: any) => ref.value === activeTab.value);
      res = {
        type: selectedRef?.type || '',
        ref: activeTab.value || '',
        name: selectedRef?.name || '',
        content:null
      };
    } else {
      res = {
        type: activeTab.value
      };
      const activeTabProps = activeTab?.props?.find((prop: any) => prop.value === activeTab.value);
      activeTabProps?.props?.options?.forEach((opt: any) => {
        res[opt.name] = opt.value;
      })
    }
    result.push(res);
  })
  return result;
}


const store = useStore<{ Endpoint, ServeGlobal: ServeStateType }>();
const refsOptions:any = ref([]);

async function searchRefs(keyword) {
  //TODO 加缓存，否则会重复拿数据
  debounce(async () => {
    refsOptions.value = await store.dispatch('Endpoint/getAllRefs', {
      "serveId": props.serveId,
      page:1,
      name:keyword,
      pageSize:20
    });
  }, 500)();
}

onMounted(async () => {
  // await searchRefs('');
})

watch(() => {return visible.value}, async (newVal: any) => {

  if(visible.value){
    await searchRefs('');
  }

  let {type, types} = props.value || {};
  // ref 优先级高于 type，如果是 ref，则优先取 ref值判断类型
  type =  props.value?.ref || type;
  const allTypes = [...(types || []), type];
  // 打开时，初始化数据
  if (newVal && (props.value.type || props.value.ref)) {
    tabsList.value = [...initTabsList(allTypes, props.value)];
  }
  // 关闭了，触发change事件
  else {
    // 仅选择类型改变了才触发change事件
    // 需要兼容 选择ref 的场景
    const value = getValueFromTabsList(tabsList.value);
    // 如果是 ref, 则直接返回, ref的优先级高于 type
    const newTypes = value.map((item: any) => item.ref || item.type);
    if (JSON.stringify(allTypes) !== JSON.stringify(newTypes)) {
      emit('change', value);
    }
  }
  }, {immediate: false})

</script>

<style lang="less" scoped>

.container {
  padding: 0;
}


.content {
  width: 480px;
  overflow-y: scroll;
}

:deep(.ant-input-number) {
  width: 100%
}

:deep(.ant-form-item-label) {
  label {
    font-weight: bold;
  }
}

.card-title {
  font-weight: bold;
  margin: 12px 0 8px 0;
}

.col {
  //margin-bottom: 8px;
}

.col-form-item {
  margin-bottom: 8px;
}

.header {
  border-bottom: 1px solid #f5f5f5;
  display: flex;

  .item {
    margin-right: 16px;
    cursor: pointer;
    height: 30px;
    line-height: 30px;
    font-weight: bold;

    &.active {
      color: #1890ff;
    }
  }
}

.main {
  .item {
    //margin-top: 16px;
    .select-type-btn {
      margin-top: 16px;
    }
  }
}

</style>

<style lang="less">

.data-type-setting-container {
  .ant-popover-inner {
    max-height: 480px;
    overflow-y: scroll;
  }
}
</style>

