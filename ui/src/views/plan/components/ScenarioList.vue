<template>
    <div class="table-filter">
        <div class="left" v-if="showScenarioOperation">
            <a-button type="primary" @click="associateModalVisible = true">
                <template #icon><plus-outlined /></template>
                关联测试场景
            </a-button>
            <a-button type="default">批量移除</a-button>
        </div>
        <div class="right">
            <a-form-item label="优先级">
                <a-select ref="select" v-model:value="formState.priority" style="width: 120px" :options="options"
                    :field-names="{ label: 'name', value: 'id', options: 'children' }" @focus="focus"
                    @change="handleChange"></a-select>
            </a-form-item>
            <a-form-item label="创建人">
                <a-select ref="select" v-model:value="formState.creater" style="width: 120px" :options="options"
                    :field-names="{ label: 'name', value: 'id', options: 'children' }" @focus="focus"
                    @change="handleChange"></a-select>
            </a-form-item>
            <a-form-item>
                <a-input-search v-model:value="formState.keywords" placeholder="请输入需要搜索的用例名称" style="width: 220px" />
            </a-form-item>
        </div>
    </div>
    <a-table :row-selection="rowSelection" :columns="columns" :data-source="data" >
        <template #operation>
            <span>操作</span>
        </template>
    </a-table>
    <Associate :associate-modal-visible="associateModalVisible" @on-cancel="associateModalVisible = false" @on-ok="associateModalVisible = false"/>
</template>
<script lang="ts" setup>
import { computed, ref, unref, reactive, defineProps, watch } from 'vue';
import { PlusOutlined } from '@ant-design/icons-vue';
import Associate from './Associate.vue';
import { Table } from 'ant-design-vue';
import { any } from 'vue-types';

const props = defineProps({
    showScenarioOperation: {
        type: Boolean,
        default: true,
        required: false
    },
    list: {
        type: any,
        default: [],
        required: false
    }
})

interface DataType {
    key: string | number;
    name: string;
    age: number;
    address: string;
}

const associateModalVisible = ref(false);

const columns: any[] = reactive([
    {
        title: 'Name',
        dataIndex: 'name',
    },
    {
        title: 'Age',
        dataIndex: 'age',
    },
    {
        title: 'Address',
        dataIndex: 'address',
    },
]);

const data: DataType[] = [];
for (let i = 0; i < 46; i++) {
    data.push({
        key: i,
        name: `Edward King ${i}`,
        age: 32,
        address: `London, Park Lane no. ${i}`,
    });
}

const selectedRowKeys = ref<DataType['key'][]>([]); // Check here to configure the default column

const onSelectChange = (changableRowKeys: string[]) => {
    console.log('selectedRowKeys changed: ', changableRowKeys);
    selectedRowKeys.value = changableRowKeys;
};

const rowSelection = computed(() => {
    return {
        selectedRowKeys: unref(selectedRowKeys),
        onChange: onSelectChange,
        hideDefaultSelections: true,
        selections: [
            Table.SELECTION_ALL,
            Table.SELECTION_INVERT,
            Table.SELECTION_NONE,
            {
                key: 'odd',
                text: 'Select Odd Row',
                onSelect: changableRowKeys => {
                    let newSelectedRowKeys = [];
                    newSelectedRowKeys = changableRowKeys.filter((_key, index) => {
                        if (index % 2 !== 0) {
                            return false;
                        }
                        return true;
                    });
                    selectedRowKeys.value = newSelectedRowKeys;
                },
            },
            {
                key: 'even',
                text: 'Select Even Row',
                onSelect: changableRowKeys => {
                    let newSelectedRowKeys = [];
                    newSelectedRowKeys = changableRowKeys.filter((_key, index) => {
                        if (index % 2 !== 0) {
                            return true;
                        }
                        return false;
                    });
                    selectedRowKeys.value = newSelectedRowKeys;
                },
            },
        ],
    };
});

const options = ref<any>([
    {
        id: 'jack',
        name: 'Jack',
        children: [
            {
                id: 'small jack',
                name: 'samll Jack',
            },
        ],
    },
    {
        id: 'lucy',
        name: 'Lucy',
    },
    {
        id: 'disabled',
        name: 'Disabled',
        disabled: true,
    },
    {
        id: 'yiminghe',
        name: 'Yiminghe',
    },
]);

const formState = reactive({ priority: null, creater: null, keywords: '' });

const focus = () => {
    console.log('focus');
};

const handleChange = (value: string) => {
    console.log(`selected ${value}`);
};

watch(() => {
    return props.showScenarioOperation;
}, val => {
    if (val) {
        columns.push({
            title: 'operation',
            dataIndex: 'operation',
            key: 'operation',
            slots: {
                customRender: 'operation'
            }
        })
    }
}, { immediate: true })
</script>
<style scoped lang="less">
.table-filter {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 20px 0;

    .left, .right {
        display: flex;
        align-items: center;
 
        :deep(.ant-row.ant-form-item), :deep(.ant-btn) {
            margin-right: 20px;
            margin-bottom: 0;

            &:last-child {
                margin: 0;
            }
        }
    }
}
</style>
  
  