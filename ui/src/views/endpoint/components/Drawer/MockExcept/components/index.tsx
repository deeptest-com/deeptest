import { DeleteOutlined } from "@ant-design/icons-vue";

const defaultData = {
  name: '',
  value: '',
  condition: ''
}

const conditionOptions = (type: string) => {
  const conditionArray = [
    {
      "label": "等于",
      "value": "equal"
    },
    {
      "label": "不等于",
      "value": "notEqual"
    },
    {
      "label": "大于",
      "value": "greaterThan"
    },
    {
      "label": "大于或等于",
      "value": "greaterThanOrEqual"
    },
    {
      "label": "小于",
      "value": "lessThan"
    },
    {
      "label": "小于或等于",
      "value": "lessThanOrEqual"
    },
    {
      "label": "包含",
      "value": "contain"
    },
    {
      "label": "不包含",
      "value": "notContain"
    },
    {
      "label": "正则匹配",
      "value": "regularMatch"
    },
    {
      "label": "存在",
      "value": "exist"
    },
    {
      "label": "不存在",
      "value": "notExist"
    }
  ];

  const bodyConditionArray = [
    {
      "label": "等于",
      "value": "equal"
    },
    {
      "label": "不等于",
      "value": "notEqual"
    },
    {
      "label": "包含",
      "value": "contain"
    },
    {
      "label": "不包含",
      "value": "notContain"
    },
  ];
  return ['requestBodies'].includes(type) ? bodyConditionArray : conditionArray;
}

const Columns = (type: string, onColumnChange: (...args: any[]) => void, onDelete: (...args: any[]) => void, options: any) => {
  const columnsArray = [
    {
      title: '参数名',
      dataIndex: 'name',
      key: 'name',
      customRender({ record }) {
        const handleInputChange = (e) => {
          record.name = e.target.value;
          onColumnChange(type);
        }
        const handleSelectChange = (e) => {
          record.name = e;
          onColumnChange(type);
        };
        return (
          <>
            {
              (['requestHeaders', 'requestBodies'].includes(type) && options[type === 'requestBodies' ? 'body' : 'header']) ? (
                <a-select value={record.name} options={options[type === 'requestBodies' ? 'body' : 'header'].map(option => ({ label: option, value: option }))} onChange={(e) => handleSelectChange(e)} />
              ) : (
                <a-input value={record.name} onChange={(e) => handleInputChange(e)} />
              )
            }
          </>
        )
      },
    },
    !['responseHeaders'].includes(type) && {
      title: '比较',
      dataIndex: 'compareWay',
      key: 'compareWay',
      customRender({ record }) {
        const onConditionChange = (e) => {
          record.compareWay = e;
        };
        return (
          <a-select value={record.compareWay} options={conditionOptions(type)} onChange={(e) => onConditionChange(e)} />
        )
      },
    },
    {
      title: '参数值',
      dataIndex: 'value',
      key: 'value',
      customRender({ record }) {
        const handleChange = (e) => {
          record.value = e.target.value;
        }
        return (
          <a-input value={record.value} onChange={e => handleChange(e)} />
        )
      },
    },
    {
      title: '操作',
      dataIndex: 'enbaled',
      key: 'enbaled',
      width: 100,
      customRender({ record, index }) {
        return (
          <div class="except-action" style={{ display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
            <a-tooltip placement="top" title="删除">
              <span class="except-action-item" onClick={() => onDelete(record, type)}>
                <DeleteOutlined />
              </span>
            </a-tooltip>
          </div>
        )
      },
    },
  ];
  return [...columnsArray.filter((e: any) => ![undefined, null, false].includes(e))];
}

const List = (props) => {
  const { type, data, onColumnChange, onDelete, optionsMap } = props;
  return (
    <a-table
      class="mock-detail-response"
      rowKey={(_record, index) => _record.idx}
      columns={Columns(type, onColumnChange, onDelete, optionsMap)}
      dataSource={data[type]}
      pagination={false}
      bordered />
  )
};

export const MockData = (props) => {
  return (
    <div style="max-height: 200px; overflow-y: scroll">
      <List {...props} />
    </div>
  )
};


