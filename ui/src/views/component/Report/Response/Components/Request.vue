<template>
    <div class="request">
        <ConBoxTitle title="请求URL"></ConBoxTitle>
        <div class="request-url">
            {{ data?.url || '' }}
        </div>
        <ConBoxTitle title="Header"></ConBoxTitle>
        <div class="request-headers">
            <a-table :dataSource="data?.headers || []" :columns="columns" />
        </div>
        <ConBoxTitle title="Body"></ConBoxTitle>
        <div class="request-body">
            <span>Body类型: {{ data?.bodyType || '' }}</span>
        </div>
        <ConBoxTitle title="请求代码"></ConBoxTitle>
        <div class="language-selector">
            <div :class="['language-item', selectIndex === index ? 'active' : '']" v-for="(item, index) in languageOptions" :key="index">
                <Icon :type="item.icon" :fill="item.fill" />
                {{ item.label }}
            </div>
            <div class="editor-wrapper">
                <MonacoEditor
                    class="editor"
                    :value="codeValue"
                    :language="'powershell'"
                    theme="vs"
                    :options="editorOptions" />
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import { defineProps, watch, h, ref } from 'vue';
import ConBoxTitle from '@/components/ConBoxTitle/index.vue';
import IconSvg from '@/components/IconSvg';
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import { MonacoOptions } from "@/utils/const";

const props = defineProps({
    data: {
        type: Object
    }
});

const columns = [
    {
        title: 'Name',
        dataIndex: 'name',
    },
    {
        title: 'Value',
        dataIndex: 'value',
    }
];

const languageOptions = [
    {
        icon: 'powershell',
        label: 'Shell',
        fill: 'rgb(3, 169, 244)'
    },
    {
        icon: 'javascript',
        label: 'Javascript',
        fill: 'rgb(255, 202, 40)'
    },
    {
        icon: 'java',
        label: 'Java',
        fill: 'rgb(39, 38, 54)'
    },
    {
        icon: 'swift',
        label: 'Swift',
        fill: 'rgb(254, 94, 47)'
    },
    {
        icon: 'php',
        label: 'Php',
        fill: 'rgb(3, 169, 244)'
    },
    {
        icon: 'Python',
        label: 'Python',
        fill: ''
    },
    {
        icon: 'http',
        label: 'Http',
        fill: ''
    }
];

const Icon = (props) => {
    return h(IconSvg, {
        type: props.type,
        style: {
            fill: props.fill
        }
    })
};

const selectIndex = ref(0);
const editorOptions = ref(MonacoOptions);
const codeValue = `curl --location --request ${props.data?.method} ${props.data?.url}
--header 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36'
--header 'Accept: */*'
--header 'Cache-Control: no-cache'
--header 'Host: localhost:8000'
--header 'Connection: keep-alive'
`;

watch(() => {
    return props.data;
}, val => {
    console.log(7777777,val);
}, {
    immediate: true,
    deep: true
})
</script>
<style scoped lang="less">
.request {
    padding-top: 20px;

    .request-url,
    .request-headers,
    .request-body {
        padding: 10px;
    }
}

.language-selector {
    margin-top: 20px;

    .language-item {
        display: inline-flex;
        width: 70px;
        height: 64px;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        border: 1px solid #e9dddd;
        border-radius: 8px;
        margin-right: 20px;
        cursor: pointer;

        &:hover {
            border: 1px solid rgb(3, 169, 244);
        }

        &.active {
            border: 1px solid rgb(3, 169, 244);
        }
    }
}

.editor-wrapper {
    padding-top: 20px;
    height: 400px;
    overflow-x: hidden;
    overflow-y: hidden;

    &>div {
        height: 100%;
    }

}
</style>
