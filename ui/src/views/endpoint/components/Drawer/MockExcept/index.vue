<template>
  <!-- 高级mock 期望 -->
  <div class="container">
    <div class="top-action">
      <span class="enable">
        是否开启 &nbsp;&nbsp;
        <a-switch v-model:checked="checked" />
      </span>
      <a-button type="primary" @click="handleCreate">新建期望</a-button>
    </div>
    <div class="except-list">
      <div class="except-tip">
        <QuestionCircleOutlined />
        当您请求Mock接口时，会根据请求参数匹配的期望条件自动返回相应的结果。Mock请求地址：https://xx.xxx.xx/mock/txt
      </div>
      <a-table 
        :pagination="false"
        :rowKey="(record) => record.id" 
        :dataSource="expectList" 
        :columns="exceptColumns" 
        id="except-table">
        <template #mockName="{ record }">
          <div class="except-name">
            <HolderOutlined class="except-sort" style="margin-right: 6px" />
            <TooltipCell :tip="record.name" :text="record.name" :width="200" @edit="handleEdit(record)" /> 
          </div>
        </template>
      </a-table>
    </div>
  </div>

  <Detail v-if="open" @cancel="open = false" />
</template>
<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import Sortable from 'sortablejs';
import { QuestionCircleOutlined, HolderOutlined } from '@ant-design/icons-vue';
import { useStore } from 'vuex';

import TooltipCell from "@/components/Table/tooltipCell.vue";
import Detail from './detail.vue';

import { StateType as EndpointStateType } from '@/views/endpoint/store';
import { exceptColumns } from './index';

const store = useStore<{ Endpoint: EndpointStateType }>();

const expectList = computed(() => store.state.Endpoint.mockExpectList);
const open = ref(false);
const checked = ref(false);
const dataSource = ref([
  {
    id: '1',
    name: '创建期望1',
    enabled: false,
    creator: '管理员',
    createAt: '2023-09-07 12:00:04',
    updateAt: '2023-10-07 12:00:04',
  },
  {
    id: '2',
    name: '创建期望2',
    enabled: true,
    creator: '管理员',
    createAt: '2023-10-07 12:00:04',
    updateAt: '2023-11-07 12:00:04',
  },
]);

/**
 * 列表拖拽排序
 */
const initSortable = () => {
  const el = document.querySelector('#except-table tbody');
  Sortable.create(el, {
    handle: '.except-sort',
    animation: 150,
    sort: true,
    group: { name: 'name', pull: true, put: true },
    onEnd: (evt) => {
      const { newIndex, oldIndex } = evt;
      const source = dataSource.value[oldIndex];
      const destination = dataSource.value[newIndex];
      dataSource.value[newIndex] = source;
      dataSource.value[oldIndex] = destination;
      console.log('更新列表顺序', dataSource.value);
    }
  })
};

const handleEdit = (record) => {
  console.log('查看mock期望', record.id);
  // loading
  // open modal
  open.value = true;
}

const handleCreate = () => {
  open.value = true;
  store.commit('Endpoint/setMockExpectDetail', {});
}

onMounted(() => {
  initSortable();
  store.dispatch('Endpoint/getMockExpectList');
})
</script>
<style scoped lang="less">
.container {
  margin-top: 20px;
}

.top-action {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: flex-end;

  .enable {
    display: flex;
    align-items: center;
    margin-right: 20px;
  }
}

.except-list {
  .except-tip {
    margin: 10px 0;
    display: flex;
    align-items: center;

    .anticon {
      margin-right: 6px;
      color: #b0b0b0;
    }
  }
  :deep(.ant-table table) {
    .except-title {
      padding-left: 22px;
    }
    .except-name {
      display: flex;
      align-items: center;
      color: #447DFD;
      cursor: pointer;

      .anticon.anticon-holder {
        font-size: 16px;
        color: #000;
      }
    }

    .except-action {
      display: flex;
      align-items: center;

      .except-action-item:not(:last-child) {
        margin-right: 10px;
        cursor: pointer;
      }
    }
  }
}

</style>