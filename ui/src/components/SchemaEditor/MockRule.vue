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
  const options = store.state.Endpoint?.mockExpressions || [];
  if (!keywords.value) {
    return options
  }
  return options.filter((item) => {
    return item.label.includes(keywords.value) || item.name.includes(keywords.value);
  });
});


const props = defineProps(['tree']);

const value = ref('');

watch(() => {
      return props.tree?.mockType;
    }, () => {
      value.value = props.tree?.mockType;
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
        :size="'small'"
        placeholder="请选择mock规则"
        style="width: 200px"
        :default-active-first-option="false"
        :show-arrow="false"
        :filter-option="false"
        :not-found-content="null"
        :options="mockExpressions"
        @search="handleSearch"
        @change="handleChange"
    >
      <template #option="{ value, label }">
        <span class="select-item" :key="value">
            <span class="left">
              @{{ label }}
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

  .select-item {
    display: flex;
    width: 360px;
    justify-content: space-between;
    align-items: center;
  }
}
</style>


