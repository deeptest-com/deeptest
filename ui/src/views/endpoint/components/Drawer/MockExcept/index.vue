<template>
  <!-- 高级mock 期望 -->
  <div class="container">
    <div class="mock-top">
      <div class="top-tabs">
        <a-button 
          v-for="(item, index) in tabs"
          :type="activityKey === item.key ? 'primary' : 'link'"
          :key="index"
          @click="activityKey = item.key">
          {{ item.value }}
        </a-button>
      </div>
      <div class="top-action">
        <span class="enable">
          是否开启 &nbsp;&nbsp;
          <a-switch v-model:checked="checked" @change="handleChange" />
        </span>
        <a-button type="primary" @click="handleCreate">新建期望</a-button>
      </div>
    </div>
    <div class="except-list" v-if="activityKey === 'mock'">
      <div class="except-tip">
        <QuestionCircleOutlined />
        当您请求Mock接口时，会根据请求参数匹配的期望条件自动返回相应的结果。Mock请求地址：https://xx.xxx.xx/mock/txt
      </div>
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
    <div v-else>
      脚本内容
    </div>
  </div>

  <Detail v-if="open" @cancel="open = false" />
</template>
<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import Sortable from 'sortablejs';
import { QuestionCircleOutlined, HolderOutlined } from '@ant-design/icons-vue';
import { useStore } from 'vuex';
import EditAndShowField from '@/components/EditAndShow/index.vue';
import Detail from './detail.vue';

import { StateType as EndpointStateType } from '@/views/endpoint/store';
import { exceptColumns } from './index';
import { message } from 'ant-design-vue';

const store = useStore<{ Endpoint: EndpointStateType }>();

const expectList = computed(() => store.state.Endpoint.mockExpectList);
const loading = computed(() => store.state.Endpoint.mockExpectLoading);
const endpointDetail = computed(() => store.state.Endpoint.endpointDetail);
const open = ref(false);
const checked = ref(!endpointDetail.value.advancedMockDisabled);
const activityKey = ref('mock');

const tabs = [{
  value: '期望',
  key: 'mock'
}, {
  value: '脚本',
  key: 'sh'
}];

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

const handleChange = async (value) => {
  await store.dispatch('Endpoint/updateMockStatus', {
    advancedMockDisabled: !value,
  })
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

.mock-top {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
}

.top-action {
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