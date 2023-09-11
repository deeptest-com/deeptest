<template>
  <!-- 高级mock 期望 -->
  <div class="except-list">
    <a-table 
      :pagination="false"
      :rowKey="(record) => record.id" 
      :dataSource="expectList"
      :loading="loading" 
      :columns="exceptColumns(store)" 
      id="except-table">
      <template #mockName="{ record }">
        <div class="except-name">
          <HolderOutlined class="except-sort" style="margin-right: 6px" />
          <EditAndShowField 
            :custom-class="'custom-endpoint show-on-hover'"
            :value="record.name"
            placeholder="期望名称"
            @update="(val) => handleUpdateName(val, record)"
            @edit="handleEdit(record)"/>
        </div>
      </template>
    </a-table>
  </div>

  <Detail v-if="open" @cancel="open = false" />
</template>
<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { HolderOutlined } from '@ant-design/icons-vue';
import { useStore } from 'vuex';
import { message } from 'ant-design-vue';
import Sortable from 'sortablejs';

import EditAndShowField from '@/components/EditAndShow/index.vue';
import Detail from './detail.vue';

import { StateType as EndpointStateType } from '@/views/endpoint/store';
import { exceptColumns } from './index';

const store = useStore<{ Endpoint: EndpointStateType }>();
const expectList = computed(() => store.state.Endpoint.mockExpectList);
const loading = computed(() => store.state.Endpoint.mockExpectLoading);
const open = ref(false);

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
      const source = expectList.value[oldIndex];
      const destination = expectList.value[newIndex];
      expectList.value[newIndex] = source;
      expectList.value[oldIndex] = destination;
      store.dispatch('Endpoint/sortMockExpect', expectList.value.map(e => Number(e.id)));
    }
  })
};

const handleUpdateName = async (value, record) => {
  await store.dispatch('Endpoint/updateMockExpectName', {
    id: record.id,
    name: value,
  })
  message.success('修改mock期望名称成功');
}

const handleEdit = async (record) => {
  store.commit('Global/setSpinning', true);
  await store.dispatch('Endpoint/getMockExpectDetail', {
    id: record.id,
  });
  store.commit('Global/setSpinning', false);
  open.value = true;
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

.except-list {
  margin-top: 20px;
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