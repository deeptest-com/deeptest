<script lang="ts" setup>

import {ref, defineProps, defineEmits, watch, reactive, toRaw, computed, onMounted} from 'vue';
import {JSONSchemaDataTypes, schemaSettingInfo} from "./config";
import {cloneByJSON} from "@/utils/object";

const props = defineProps({
  value: {
    required: true,
    type: Object
  },
  refsOptions: {
    required: true,
    type: Array
  },
})
const emit = defineEmits(['change']);
const tabsList: any = ref([]);
const treeInfo: any = ref(null);
const visible: any = ref(false);

function changeType(tabsIndex: any, e: any) {
  let type = e.target.value;
  if (type === 'array') {
    if (tabsList.value.length === tabsIndex + 1) {
      tabsList.value.push(cloneByJSON(schemaSettingInfo));
      console.log(832, cloneByJSON(schemaSettingInfo));
    }
  } else {
    if (tabsIndex < tabsList.value.length) {
      tabsList.value.splice(tabsIndex + 1);
    }
  }
}

watch(() => {
  return tabsList.value
}, (newVal) => {
  console.log('tabsList value change: ', newVal);
}, {
  deep: true
})

onMounted(() => {
  console.log(832, 'open');
})

const typeOpts = ['string', 'number', 'boolean', 'array', 'object', 'integer'];

function initTabsList(types: any,treeInfo: any) {
  console.log(treeInfo,222);
  let tabsList: any = [];
  types.forEach((type: string) => {
    const defaultTabs: any = cloneByJSON(schemaSettingInfo);
    // 默认选中了类型
    if (typeOpts.includes(type)) {
      defaultTabs[0].active = typeOpts.includes(type);
      defaultTabs[0].value = type;
      const activeTabProps = defaultTabs[0].props.find((prop: any) => prop.value === type);
      activeTabProps.props.options.forEach((opt: any) => {
        opt.value = treeInfo[opt.name] || opt.value;
      })
      console.log('hello',defaultTabs[0],treeInfo,treeInfo['enum'],treeInfo['enum']?.length );
      // console.log('hello',activeTabProps);
    } else {
      defaultTabs[1].active = true;
      defaultTabs[1].value = type;
    }
    tabsList.push(defaultTabs)
  });

  // 如果是数组，还需加一项
  if (types[types.length - 1] === 'array') {
    const arrayItems: any = cloneByJSON(schemaSettingInfo);
    tabsList.push(arrayItems);
  }
  return tabsList;
}

function getValueFromTabsList(tabsList: any) {
  const res: any = {};
  tabsList.forEach((tabs: any) => {
    const activeTab = tabs.find((tab: any) => tab.active);
    res[activeTab.value] = {};
    const activeTabProps = activeTab?.props?.find((prop: any) => prop.value === activeTab.value);
    activeTabProps.props.options.forEach((opt: any) => {
      res[activeTab.value][opt.name] = opt.value;
    })
  })
  console.log(832 ,'res', res);
  return res;
}

watch(() => {
  return visible.value
}, (newVal: any) => {
  // 打开时，初始化数据
  if (newVal && props.value.type) {
    // 处理如 array[array[object]] 的场景
    const types = props.value.type.match(/\w+/g);
    tabsList.value = [...initTabsList(types,props.value)];
  }
  // 关闭了，触发change事件
  else {
    emit('change', getValueFromTabsList(tabsList.value));
  }
})

function selectTab(tabs: any, tabIndex: number) {
  tabs.forEach((tab: any, index: number) => {
    tab.active = tabIndex === index;
  })
}

</script>

<template>
  <a-popover :title="null"
             trigger="click"
             v-model:visible="visible"
             placement="left"
             :overlayClassName="'container'">
    <template #content>
      <div class="content" v-for="(tabs,tabsIndex) in tabsList" :key="tabsIndex">
        <div class="header">
          <div class="item"
               v-for="(tab,tabIndex) in tabs"
               @click="() => {
                  selectTab(tabs,tabIndex)
               }"
               :class="tab.active ? 'active' : ''"
               :key="tab.value">
            {{ tabsIndex === 0 ? tab.label : tab.subLabel }}
          </div>
        </div>
        <div class="main">
          <div class="item"
               v-for="(tab) in tabs"
               v-show="tab.active"
               :key="tab.value">
            <a-radio-group
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
                            v-if="opt.component === 'selectTag'"
                            v-model:value="opt.value"
                            mode="tags"
                            :placeholder="opt.placeholder"
                        />
                        <a-select
                            v-if="opt.component === 'select'"
                            v-model:value="opt.value"
                            :options="opt.options"
                            :placeholder="opt.placeholder"
                        />
                        <a-input
                            v-if="opt.component === 'input'"
                            v-model:value="opt.value"
                            :placeholder="opt.placeholder"
                        />
                        <a-input-number
                            v-if="opt.component === 'inputNumber'"
                            id="inputNumber"
                            :placeholder="opt.placeholder"
                            v-model:value="opt.value"
                        />
                        <a-switch
                            v-if="opt.component === 'switch'"
                            v-model:checked="opt.value"/>
                      </a-form-item>
                    </a-col>
                  </a-row>
                </div>
              </div>
            </a-form>
            <a-form :layout="'vertical'" v-if="tab.type === '$ref' && tab.active">
              <a-form-item
                  class="col-form-item"
                  :labelAlign="'right'"
                  :label="'请选择组件'">
                <a-select
                    label-in-value
                    :options="refsOptions"
                    placeholder="Select Components"
                    style="width: 100%"/>
              </a-form-item>
            </a-form>
          </div>
        </div>
      </div>
    </template>
    <a href="javascript:void(0)">{{ value.type }}</a>
  </a-popover>
</template>

<style lang="less" scoped>

.container {
  padding: 0;
}

.content {
  width: 600px;
  //height: 380px;
  //overflow-y: scroll;
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
  //padding: 8px 16px;
  //background-color: #f5f5f5;
  border-bottom: 1px solid #f5f5f5;
  display: flex;
  //justify-content: space-between;
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
    margin-top: 16px;
  }
}


</style>

