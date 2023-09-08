<script lang="ts" setup>
import {
  ArrowUpOutlined,
  ArrowDownOutlined,
  CopyOutlined,
} from '@ant-design/icons-vue';
import {computed, defineProps, defineEmits, ref, watch} from "vue";
// import EditAndShowSelect from "@/components/EditAndShowSelect/index.vue";
import {useStore} from "vuex";
import {StateType as ServeStateType} from "@/store/serve";

const store = useStore<{ Endpoint, ServeGlobal: ServeStateType }>();

const mockExpressions = computed(() => {
  let options = store.state.Endpoint?.mockExpressions || [];
  // 如果有类型，则需要过滤相同类型的数据
  if (props.tree?.type) {
    options = options.filter((item) => {
      return item.type && item.type === props.tree?.type;
    })
  }
  if (keywords.value) {
    options = options.filter((item) => {
      return item.label.includes(keywords.value) || item.name.includes(keywords.value);
    });
  }
  // console.log('83222 options', options)
  return [...options];
});


const props = defineProps(['tree', 'readonly']);

const value = ref('');


watch(() => {
      return props.tree?.['x-mock-type'];
    }, (newVal) => {
      value.value = newVal || null;
    }, {
      immediate: true
    }
);


const emit = defineEmits(['update']);

const keywords = ref('');
// 搜索
const handleSearch = (val: string) => {
  keywords.value = val;
};

const handleChange = (val: string) => {
  value.value = val;
  emit('update', val);
};

</script>

<template>
  <div class="container">
    <a-select
        v-model:value="value"
        show-search
        :disabled="props.readonly"
        :size="'small'"
        placeholder="请选择mock规则"
        style="width: 200px"
        :default-active-first-option="false"
        :show-arrow="false"
        :filter-option="false"
        :dropdownStyle="{minWidth: '300px'}"
        :not-found-content="null"
        :options="mockExpressions"
        @search="handleSearch"
        @change="handleChange"
    >
      <template #option="{ value, label, name, id }">
        <span class="select-item" :key="value || id">
            <span class="left" :title="label">
              {{ label }}
            </span>
             <span class="right" :title="name">
              {{ name }}
            </span>
        </span>
      </template>
    </a-select>
  </div>

</template>

<style lang="less" scoped>
.container {
  display: inline-block;
  margin-left: 16px;

  :deep(.ant-select-selector) {
    border: 1px solid transparent;
  }

}

.select-item {
  display: flex;
  justify-content: space-between;
  align-items: center;

  .left {
    max-width: 200px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
  }

  .right {
    max-width: 80px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
  }
}

</style>


