import { DeleteOutlined } from "@ant-design/icons-vue";

const defaultData = {
  name: '',
  value: '',
  condition: ''
}

const conditionOptions = (type: string) => {
  const conditionArray = [
    !['body'].includes(type) && {
      name: '等于',
      value: '等于',
    }, 
    !['body'].includes(type) && {
      name: '不等于',
      value: '不等于',
    }, {
      name: '包含',
      value: '包含',
    }, {
      name: '不包含',
      value: '不包含',
  }];
  return [...conditionArray.filter((e: any) => ![undefined, null, false].includes(e))];
}

const Columns = (type: string, onChange: (...args: any[]) => void, onDelete: (...args: any[]) => void) => {
  const columnsArray = [
    {
      title: '参数名',
      dataIndex: 'name',
      key: 'name',
      customRender({ record }) {
        const handleInputChange = (e) => {
          record.name = e.target.value;
          onChange(type);
        }
        const handleSelectChange = (e) => {
          record.name = e;
          onChange(type);
        };
        return (
          <>
            {
              ['header', 'body'].includes(type) ? (
                <a-select value={record.name} options={[]} onChange={(e) => handleSelectChange(e)} />
              ) : (
                <a-input value={record.name} onChange={(e) => handleInputChange(e)} />
              )
            }
          </>
        )
      },
    },
    !['responseHeader'].includes(type) && {
      title: '比较',
      dataIndex: 'condition',
      key: 'condition',
      customRender({ record }) {
        const onChange = (e) => {
          console.log(e);
          record.condition = e;
        };
        return (
          <a-select value={record.condition} options={conditionOptions(type)} onChange={(e) => onChange(e)} />
        )
      },
    },
    {
      title: '参数值',
      dataIndex: 'value',
      key: 'value',
      customRender({ record }) {
        return (
          <a-input value={record.value} onChange={e => record.value = e} />
        )
      },
    },
    {
      title: '操作',
      dataIndex: 'enbaled',
      key: 'enbaled',
      width: 100,
      customRender({ record }) {
        const handleDelete = (e) => {
          console.log(e);
        };
        return (
          <div class="except-action" style={{ display: 'flex', alignItems:'center', justifyContent: 'center' }}>
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
  console.log(props);
  const { type, data, onChange, onDelete } = props;
  return (
    <a-table 
      class="mock-detail-response" 
      rowKey={(_record, index) => _record.idx} 
      columns={Columns(type, onChange, onDelete)} 
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


