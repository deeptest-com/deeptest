import { createVNode, ref } from "vue";
import { Modal } from "ant-design-vue";
import { CopyOutlined, DeleteOutlined, ExclamationCircleOutlined, HolderOutlined } from "@ant-design/icons-vue";
import { useStore } from "vuex";

import TooltipCell from "@/components/Table/tooltipCell.vue";
import { momentUtc } from "@/utils/datetime";

const handleChange = async (record, e, store) => {
  record.disabled = !e;
  // await store.commit('Global/setSpinning', true);
  const result = await store.dispatch('Endpoint/disabledMockExpect', {
    id: record.id,
    disabled: !e,
  })
  if (result) {
    record.disabled = !e;
  }
  // await store.commit('Global/setSpinning', false);
};

const handleClone = async (record, store) => {
  await store.commit('Global/setSpinning', true);
  await store.dispatch('Endpoint/cloneMockExpect', {
    id: record.id,
  })
  await store.commit('Global/setSpinning', false);
};

const handleDelete = (record, store) => {
  Modal.confirm({
    title: '确认要删除该Mock用例吗',
    icon: createVNode(ExclamationCircleOutlined),
    okText: '确定',
    cancelText: '取消',
    async onOk() {
      await store.commit('Global/setSpinning', true);
      await store.dispatch('Endpoint/deleteMockExpect', {
        id: record.id,
      })
      await store.commit('Global/setSpinning', false);
    },
  });
};

export const exceptColumns = (store) => [
  {
    title() {
      return (
        <div class="except-title">
          用例名称
        </div>
      )
    },
    dataIndex: 'name',
    key: 'name',
    width: 200,
    slots: { customRender: 'mockName' },
  },
  {
    title: '启用',
    dataIndex: 'disabled',
    key: 'disabled',
    customRender({ record }) {
      return (
        <a-switch checked={!record.disabled} onChange={(e) => handleChange(record, e, store)} />
      )
    },
  },
  {
    title: '创建人',
    dataIndex: 'createUser',
    key: 'createUser',
  },
  {
    title: '创建时间',
    dataIndex: 'createdAt',
    key: 'createdAt',
    customRender({ record }) {
      return (
        <TooltipCell tip={momentUtc(record.createdAt)} text={momentUtc(record.createdAt)} width={180} />
      )
    },
  },
  {
    title: '更新时间',
    dataIndex: 'updatedAt',
    key: 'updatedAt',
    customRender({ record }) {
      return (
        <TooltipCell tip={momentUtc(record.updatedAt)} text={momentUtc(record.updatedAt)} width={180} />
      )
    },
  },
  {
    title: '操作',
    dataIndex: 'action',
    key: 'action',
    width: 80,
    customRender({ record }) {
      return (
        <div class="except-action">
          <a-tooltip placement="top" title="克隆">
            <span class="except-action-item" onClick={() => handleClone(record, store)}>
              <CopyOutlined />
            </span>
          </a-tooltip>
          <a-tooltip placement="top" title="删除">
            <span class="except-action-item" onClick={() => handleDelete(record, store)}>
              <DeleteOutlined />
            </span>
          </a-tooltip>
        </div>
      )
    }
  },
];

export const requestTabs = [{
  title: '查询参数',
  type: 'requestQueryParams',
}, {
  title: '路径参数',
  type: 'requestPathParams',
}, {
  title: '请求体',
  type: 'requestBodies',
}, {
  title: '请求头',
  type: 'requestHeaders',
}];