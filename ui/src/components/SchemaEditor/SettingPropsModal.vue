<template>
  <div>
    <a-tabs
        :key="infoIndex"
        @change="(targetKey) => {
            changeTab(infoIndex,targetKey)
        }"
        v-for="(info,infoIndex) in moduleInfo"
        v-model:activeKey="activeKey[infoIndex]">

      <a-tab-pane v-for="(tab) in info" :key="tab.value"
                  :tab="infoIndex === 0 ? tab.label : tab.subLabel">

        <a-form :layout="'vertical'" v-if="tab.value === 'type'">
            <a-radio-group
              @change="(e) => {
                changeType(infoIndex,e)
              }"
              :value="selectedTypes[infoIndex]"
              button-style="solid">
            <a-radio-button
                v-for="item in tab.props"
                :key="item.value"
                :value="item.value">{{ item.label }}
            </a-radio-button>

          </a-radio-group>
          <div v-for="(item,itemIndex) in tab.props" :key="itemIndex">
            <div v-if="item.value === selectedTypes[infoIndex]">
              <div class="card-title">{{ item.props.label }}</div>
              <a-card
                  :bodyStyle="{padding:'0 16px 16px 16px'}"
                  :title="null">
                <a-row
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
                          @change="() => {}"
                      />
                      <a-select
                          v-if="opt.component === 'select'"
                          v-model:value="opt.value"
                          :options="opt.options"
                          :placeholder="opt.placeholder"
                          @change="() => {}"
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
              </a-card>
            </div>
          </div>
        </a-form>

        <a-form v-if="tab.value === 'components'">
          <a-select
              label-in-value
              placeholder="Select users"
              style="width: 100%"
          >
          </a-select>
        </a-form>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>
<script lang="ts" setup>
import {defineComponent, ref, defineProps, defineEmits, watch, reactive, toRaw, UnwrapRef, computed} from 'vue';
import {schemaSettingInfo} from './config';
const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
  value: {
    required: true,
    type: Object
  },
})

const emit = defineEmits(['ok', 'cancal']);

// 根据传入的数据生成模块数据
function genModuleInfo() {
  return schemaSettingInfo;
}

const moduleInfo: any = ref([genModuleInfo()]);
// 选中的类型，用于构造数据
const selectedTypes: any = ref(['string']);
const activeKey = ref(['type']);
const treeInfo:any = ref(null);

function changeTab(index, key) {
  console.log(index, key);
}

function changeType(index: any, e: any) {
  let type = e.target.value;
  selectedTypes.value[index] = type;
  if (type === 'array') {
    if (moduleInfo.value.length === index + 1) {
      moduleInfo.value.push(genModuleInfo());
      selectedTypes.value.push('string');
    }
  } else {
    if (index < moduleInfo.value.length) {
      moduleInfo.value.splice(index + 1);
    }
  }
}

watch(() => {
  return selectedTypes.value
}, (newVal) => {
  console.log(832, newVal[0]);
}, {
  immediate: true
})

watch(() => {
  return props.value
}, (newVal) => {
  treeInfo.value = newVal;
  console.log(moduleInfo.value)
})

watch(() => {
  return props.visible
}, () => {
  console.log(8329999, props.visible);
})


</script>
<style lang="less" scoped>

:deep(.ant-modal-body) {
  padding: 0;
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


</style>
