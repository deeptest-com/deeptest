<template>
    <a-table :row-selection="rowSelection" :columns="columns" :data-source="data" />
</template>
<script lang="ts">
import { defineComponent, computed, ref, unref } from 'vue';
import { Table } from 'ant-design-vue';

interface DataType {
    key: string | number;
    name: string;
    age: number;
    address: string;
}

const columns = [
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
];

const data: DataType[] = [];
for (let i = 0; i < 46; i++) {
    data.push({
        key: i,
        name: `Edward King ${i}`,
        age: 32,
        address: `London, Park Lane no. ${i}`,
    });
}

export default defineComponent({
    setup() {
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
        return {
            data,
            columns,
            selectedRowKeys,
            rowSelection,
        };
    },
});
</script>
  
  