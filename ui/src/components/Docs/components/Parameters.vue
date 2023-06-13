<!--
 参数展示组件，适用于 Path 参数、Query 参数、Header 参数、Cookie 参数
-->
<template>
  <a-list v-if="data?.length" item-layout="horizontal" :data-source="data" :bordered="false" :split="false">
    <template #renderItem="{ item }">
      <a-list-item>
        <a-list-item-meta>
          <template #title>
            <div class="title">
              <div class="name">
                <a-typography-text :strong="true"> {{ item.name }}</a-typography-text>
              </div>
              <div class="type">
                <a-typography-text type="secondary">{{ item.type }}</a-typography-text>
              </div>
              <a-divider class="divider" v-if="item.required"/>
              <div>
                <div class="required" v-if="item.deprecated" style="margin-right: 6px;">
                  <a-typography-text :strong="false" type="warning">{{
                      item.deprecated ? 'deprecated' : ''
                    }}
                  </a-typography-text>
                </div>
                <div class="required" v-if="item.required">
                  <a-typography-text :strong="false" type="warning">{{
                      item.required ? 'required' : ''
                    }}
                  </a-typography-text>
                </div>
              </div>
            </div>
          </template>
          <template #description>
            <div class="description">
              <div class="description-item"
                   v-for="option in item.options"
                   :key="option.label">
                <div v-if="option.label !== 'description'" class="label">
                  <a-typography-text type="secondary"> {{ option.label }}：</a-typography-text>
                </div>
                <div class="value">
                  <a-typography-text v-if="option.label === 'description'"> {{
                      option.value
                    }}
                  </a-typography-text>
                  <a-typography-text v-else type="secondary"> {{ option.value }}</a-typography-text>
                </div>
              </div>
            </div>
          </template>
        </a-list-item-meta>
      </a-list-item>
    </template>
  </a-list>
</template>

<script lang="ts" setup>
import {
  ref,
  defineProps,
  defineEmits,
  computed, watch,
} from 'vue';

import {DownOutlined, RightOutlined} from '@ant-design/icons-vue';

const props = defineProps({
  items: {
    required: true,
    type: Object,
  },
})

interface DataItem {
  title: string;
  name: string;
  type?: string;
  required?: boolean;
  options?: any[];
}

const data = computed(() => {
  const res: any = [];
  props?.items?.forEach((item) => {
    const options: any = [];
    Object.keys(item).forEach((key) => {
      if (key !== 'name' && key !== 'type' && key !== 'required') {
        // 为 0 的时候不显示
        if (item[key] === 0) {
          // options.push({
          //   label: key,
          //   value: item[key],
          // })
        } else if (item[key]) {
          options.push({
            label: key,
            value: item[key],
          })
        }
      }
    })
    res.push({
      name: item?.name,
      type: item?.type,
      required: item?.required,
      options
    })
  })
  return res;
});

const emit = defineEmits(['ok', 'close', 'refreshList']);

const expand = ref(true);

function switchExpand() {
  expand.value = !expand.value;
}

watch(() => {
  return props.items
}, (newVal) => {
  // console.log(8322222, newVal)
})


</script>
<style lang="less" scoped>

.title {
  display: flex;
  align-items: center;
  font-weight: normal;


  .name {
    margin-right: 4px;
  }

  .divider {
    flex: 1;
    width: auto !important;
    min-width: auto !important;
    margin: 0 8px;
  }
}

.description {
  display: flex;
  flex-direction: column;

  .description-item {
    display: flex;

    .label {
      width: auto;
    }

    .value {
      flex: 1;
    }
  }
}

</style>
