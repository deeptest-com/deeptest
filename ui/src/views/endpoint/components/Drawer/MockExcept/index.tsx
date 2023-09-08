import { CopyOutlined, DeleteOutlined, ExclamationCircleOutlined, HolderOutlined } from "@ant-design/icons-vue";
import TooltipCell from "@/components/Table/tooltipCell.vue";
import { Modal } from "ant-design-vue";
import { createVNode, ref } from "vue";

const handleChange = (record, e) => {
  console.log(record, e);
};

const handleClone = (record) => {
  console.log('克隆的对象', record);
};

const handleDelete = (record) => {
  Modal.confirm({
    title: '确认要删除该Mock用例吗',
    icon: createVNode(ExclamationCircleOutlined),
    okText: '确定',
    cancelText: '取消',
    onOk() {
      console.log('删除对象', record.id);
    },
  });
};

export const exceptColumns = [
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
    slots: { customRender: 'mockName' },
    // customRender({ record }) {
    //   return (
    //     <div class="except-name">
    //       <HolderOutlined class="except-sort" style={{ marginRight: '6px' }} />
    //       <TooltipCell tip={record.name} text={record.name} width={200} onEdit={(e) => handleEdit(record)} /> 
    //     </div>
    //   )
    // },
  },
  {
    title: '启用',
    dataIndex: 'disabled',
    key: 'disabled',
    customRender({ record }) {
      return (
        <a-switch checked={record.disabled} onChange={(e) => handleChange(record, e)} />
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
        <TooltipCell tip={record.createdAt} text={record.createdAt} width={180} />
      )
    },
  },
  {
    title: '更新时间',
    dataIndex: 'updatedAt',
    key: 'updatedAt',
    customRender({ record }) {
      return (
        <TooltipCell tip={record.updatedAt} text={record.updatedAt} width={180} />
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
            <span class="except-action-item" onClick={() => handleClone(record)}>
              <CopyOutlined />
            </span>
          </a-tooltip>
          <a-tooltip placement="top" title="删除">
            <span class="except-action-item" onClick={() => handleDelete(record)}>
              <DeleteOutlined />
            </span>
          </a-tooltip>
        </div>
      )
    }
  },
];

export const requestTabs = [{
  title: '请求头',
  type: 'header',
}, {
  title: '请求体',
  type: 'body',
}, {
  title: 'Query参数',
  type: 'queryParams',
}, {
  title: '路径参数',
  type: 'pathParams',
}];