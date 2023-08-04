<template>
  <a-popover :title="null"
             trigger="click"
             v-model:visible="visible"
             :overlayClassName="'data-type-setting-container'">
    <template #content>
      <div class="content" v-for="(tabs,tabsIndex) in tabsList" :key="tabsIndex" v-show="!(activeTabsIndex > 0 && tabsIndex > 0)">
        <div class="header">
          <div class="item"
               v-for="(tab,tabIndex) in tabs"
               @click="() => {
                 if(isDisabled){
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
                :disabled="isDisabled"
                v-model:value="tab.value"
                @change="(event) => changeType(tabsIndex, event)"
                button-style="solid">
              <a-radio-button
                  v-for="item in tab.props"
                  :key="item.value"
                  :value="item.value">{{ item.label }}
              </a-radio-button>
            </a-radio-group>
            <!-- ::::基本类型设置 -->
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
                            :disabled="isDisabled"
                            v-if="opt.component === 'selectTag'"
                            v-model:value="opt.value"
                            mode="tags"
                            :placeholder="opt.placeholder"
                        />
                        <a-select
                            v-if="opt.component === 'select'"
                            v-model:value="opt.value"
                            :disabled="isDisabled"
                            :options="opt.options"
                            :placeholder="opt.placeholder"
                        />
                        <a-input
                            v-if="opt.component === 'input'"
                            v-model:value="opt.value"
                            :disabled="isDisabled"
                            :placeholder="opt.placeholder"
                        />
                        <a-input-number
                            v-if="opt.component === 'inputNumber'"
                            id="inputNumber"
                            :disabled="isDisabled"
                            :placeholder="opt.placeholder"
                            v-model:value="opt.value"
                        />
                        <a-switch
                            v-if="opt.component === 'switch'"
                            :disabled="isDisabled"
                            v-model:checked="opt.value"/>
                      </a-form-item>
                    </a-col>
                  </a-row>
                </div>
              </div>
            </a-form>
            <!-- ::::引用类型设置 -->
            <a-form :layout="'vertical'" style="margin-bottom: 16px;" v-if="tab.type === '$ref' && tab.active">
              <a-form-item
                  class="col-form-item"
                  :labelAlign="'right'"
                  :label="'请选择组件'">
                <a-select
                    :options="refsOptions"
                    :disabled="isDisabled"
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
            <!-- ::::组合schema -->
            <a-form :layout="'vertical'" style="margin-bottom: 16px;" v-if="tab.type === 'combine' && tab.active">
              <a-form-item
                  class="col-form-item"
                  :labelAlign="'right'"
                  :label="'请选择复合关键字'">
                <a-select
                    :options="combineSchemaOpts"
                    :disabled="isDisabled"
                    @change="(e) => {
                      changeCombineType(tabsIndex,tabIndex,e);
                    }"
                    show-search
                    allowClear
                    @search="searchRefs"
                    :value="tab.value || null"
                    placeholder="Select an option below to combine your schemas"
                    style="width: 100%"/>
              </a-form-item>

              <div style="margin-top: 12px;margin-left: -14px;">
                <ul>
                  <li><a-typography-text type="secondary"><span class="form-item-info">all of：</span>根据所有子模式验证值</a-typography-text></li>
                  <li><a-typography-text type="secondary"><span class="form-item-info">one of:</span> 根据其中一个子模式验证值</a-typography-text></li>
                  <li><a-typography-text type="secondary"><span class="form-item-info">any of：</span>根据任意（一个或多个）子模式验证值</a-typography-text></li>
                </ul>
              </div>
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
import {schemaSettingInfo, typeOpts, combineSchemaOpts, combineTypes} from "./config";
import cloneDeep from "lodash/cloneDeep";
import {useStore} from "vuex";
import {StateType as ServeStateType} from "@/store/serve";
import debounce from "lodash.debounce";

const props = defineProps(['value', 'serveId', 'isRefChildNode', 'isRoot']);

const emit = defineEmits(['change']);
const tabsList: any = ref([]);
const visible: any = ref(false);

// 当前选中的顶层 tab index
/**
 * 这里备注下：Components 和 Combine Schemas 两种类型，都是通过 tabsList[0] 来控制的，所以这里的 tabsIndex 也是通过 tabsList[0] 来控制的
 * 另外，如果选中了Components 和 Combine Schemas 两种类型，则 TabLists[1,....] 则不需要展示了
 * */
const activeTabsIndex = computed(() => {
  if(!Array.isArray(tabsList.value?.[0])){
    return -1;
  }
  return tabsList.value?.[0]?.findIndex((item: any, index: any) => {
    return item.active;
  });
});
const isDisabled: any = computed(() => {
  return props.isRefChildNode && !props.isRoot;
})

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

// ref 组件
function changeCombineType(tabsIndex, tabIndex, e) {
  tabsList.value[tabsIndex][tabIndex].value = e;
}


function selectTab(tabs: any, tabIndex: number) {
  tabs.forEach((tab: any, index: number) => {
    tab.active = tabIndex === index;
  })
  // 切换成普通选择模式时，如果是选中的是数组，则需要添加一个tab
  if (tabIndex === 0 && tabs[tabIndex].value === 'array' && tabsList.value.length === 1) {
    tabsList.value.push(cloneDeep(schemaSettingInfo));
  }
  console.log('832 tabsList', tabsList.value)
  // 切换到 组件 Tab 或者 组合schema Tab 时，需要清空其他的tab
  // if(tabIndex === 2 || tabIndex === 1){
  //   tabsList.value.splice(1);
  // }
}

/**
 * 初始化tabsList数据
 * */
function initTabsList(types: any, treeInfo: any) {
  let tabsList: any = [];
  types.forEach((type: string) => {
    const defaultTabs: any = cloneDeep(schemaSettingInfo);
    // 基本类型，即第一个tab
    if (typeOpts.includes(type)) {
      defaultTabs[0].active = true;
      defaultTabs[0].value = type;
      const activeTabProps = defaultTabs[0].props.find((prop: any) => prop.value === type);
      activeTabProps?.props?.options?.forEach((opt: any) => {
        opt.value = treeInfo[opt.name] || opt.value;
      })

      defaultTabs[1].active = false;
      defaultTabs[1].value = treeInfo?.ref;

      defaultTabs[2].active = false;
      defaultTabs[2].value = 'allOf';

      //  组合类型,即，第三个tab
    } else if (combineTypes.includes(type)) {
      defaultTabs[0].active = false;
      defaultTabs[0].value = treeInfo?.type || 'string';

      defaultTabs[1].active = false;
      defaultTabs[1].value = treeInfo?.ref;

      defaultTabs[2].active = true;
      defaultTabs[2].value = type;
      //  引用类型，即，选中第二个tab时
    } else {

      defaultTabs[0].active = false;
      defaultTabs[0].value = treeInfo?.type || 'string';

      defaultTabs[1].active = true;
      defaultTabs[1].value = treeInfo?.ref;

      defaultTabs[2].active = false;
      defaultTabs[2].value = 'allOf';
    }
    tabsList.push(defaultTabs)
  });
  // 如果是数组，还需加一项，渲染需要
  if (types[types.length - 1] === 'array') {
    const arrayItems: any = cloneDeep(schemaSettingInfo);
    tabsList.push(arrayItems);
  }
  return tabsList;
}

function getValueFromTabsList(tabsList: any) {
  const result: any = [];

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
        content: null
      };
    } else if (activeTab.type === 'combine') {
      res = {
        type: activeTab.value,
      };
      res[activeTab.value] = [];
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
const refsOptions: any = ref([]);

async function searchRefs(keyword) {
  //TODO 加缓存，否则会重复拿数据
  debounce(async () => {
    refsOptions.value = await store.dispatch('Endpoint/getAllRefs', {
      "serveId": props.serveId,
      page: 1,
      name: keyword,
      pageSize: 20
    });
  }, 500)();
}


watch(() => {
  return visible.value
}, async (newVal: any) => {

  // 打开时，初始化数据
  if (visible.value) {
    await searchRefs('');
  }

  let {type, types} = props.value || {};
  // ref 优先级高于 type，如果是 ref，则优先取 ref值判断类型
  type = props.value?.ref || type;
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
    margin-bottom: 16px;
    .select-type-btn {
      margin-top: 16px;
    }
  }
}
.form-item-info{
  display: inline-block;
  text-align: left;
  font-weight: bold;
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

