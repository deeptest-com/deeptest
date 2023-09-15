import { DeleteOutlined } from "@ant-design/icons-vue";

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
];

// 不同的类型 下拉选择项也不同
const CompareWayOfString = ['equal', 'notEqual', 'contain', 'notContain', 'regularMatch'];

const CompareWayOfBodies = ['equal', 'notEqual', 'contain', 'notContain'];

const conditionOptions = (type: string, record: any, optionsMap: any, selectType?: string) => {
  if (type === 'requestBodies') {
    return ['fullText'].includes(selectType || '') ? conditionArray.filter(condition => CompareWayOfBodies.includes(condition.value)) : conditionArray;
  }
  if (['requestQueryParams', 'requestPathParams'].includes(type)) {
    return conditionArray;
  }
  const options = optionsMap['header'] || [];
  if (record.name) {
    const selectHeaderType = options.find(option => option.name === record.name)?.type;
    return selectHeaderType === 'string' ? conditionArray.filter(condition => CompareWayOfString.includes(condition.value)) : conditionArray;
  }
  return conditionArray;
}

const Columns = (opts: { type: string, onColumnChange: (...args: any[]) => void, onDelete: (...args: any[]) => void, optionsMap: any, selectType: string}) => {
  const { type, onColumnChange, onDelete, optionsMap, selectType } = opts;
  const columnsArray = [
    // 请求头渲染参数名
    ['requestHeaders'].includes(type) && {
      title: '参数名',
      dataIndex: 'name',
      key: 'name',
      customRender({ record }) {
        const handleSelectChange = (e) => {
          record.name = e.replace(/\s/ig, '');
          onColumnChange(type);
        };
        return (
          <>
            <a-auto-complete
              value={record.name} 
              options={(optionsMap['header'] || []).map(option => ({ label: option.name, value: option.name }))} 
              filter-option={false}
              onChange={(e) => handleSelectChange(e)} />
          </>
        )
      },
    },
    // 请求体渲染参数名
    ['requestBodies'].includes(type) && selectType !== 'fullText' && {
      title() {
        return (
          <>
            {selectType === 'keyValue' ? '参数名' : 'XPath表达式'}
          </>
        )
      },
      dataIndex: 'name',
      key: 'name',
      customRender({ record }) {
        const handleInputChange = (e) => {
          record.name = e.target.value.replace(/\s/ig, '');
          onColumnChange(type);
        }
        return (
          <>
            <a-input 
              value={record.name} 
              onChange={(e) => handleInputChange(e)} />
          </>
        )
      },
    },
    ['responseHeaders', 'requestQueryParams', 'requestPathParams'].includes(type)&& {
      title: '参数名',
      dataIndex: 'name',
      key: 'name',
      customRender({ record }) {
        const handleInputChange = (e) => {
          record.name = e.target.value.replace(/\s/ig, '');;
          onColumnChange(type);
        }
        return (
          <>
            <a-input 
              value={record.name} 
              onChange={(e) => handleInputChange(e)} />
          </>
        )
      },
    },
    !['responseHeaders'].includes(type) && {
      title: '比较',
      dataIndex: 'compareWay',
      key: 'compareWay',
      width: 180,
      customRender({ record }) {
        const onConditionChange = (e) => {
          record.compareWay = e;
          onColumnChange(type);
        };
        const options = conditionOptions(type, record, optionsMap, selectType);
        if (record.name && !record.compareWay) {
          record.compareWay = options[0].value;
        }
        return (
          <a-select value={record.compareWay} options={options} onChange={(e) => onConditionChange(e)} />
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
          onColumnChange(type);
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
  const { type, data } = props;
  return (
    <a-table
      class="mock-detail-response"
      rowKey={(_record, index) => _record.idx}
      columns={Columns(props)}
      dataSource={type === 'requestBodies' ? data[type].filter(e => e.selectType === props.selectType) : data[type]}
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


