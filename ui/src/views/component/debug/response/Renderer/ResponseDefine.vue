<template>
    <div class="content">
    <div class="switch">
        <span class="title">响应校验</span>
        <a-tooltip placement="topLeft">
            <template #title>
            <span>校验接口返回的数据格式、结构及状态码是否符合接口定义</span>
            </template>
        <QuestionCircleOutlined />
      </a-tooltip>&nbsp;
      <a-switch size="small" v-model:checked="formState.open" @change="change"/>
   </div>
   <div class="codes">
    <a-select
        v-model:value="formState.code"
        :options="options"
        @change="select"
        :size="'small'"
        style="width: 80px;"

        ></a-select>
   </div>
<div style="clear:right;">
    <span></span>
</div>
</div>

</template>

<script lang="ts" setup>

import {ref,defineProps, defineEmits,computed,reactive} from 'vue';
const props = defineProps(['codes','code','open'])
const emits = defineEmits(['change'])
import {QuestionCircleOutlined} from '@ant-design/icons-vue';

const options = computed(
    ()=>props.codes.map(pro => ({ value: pro }))
)


const formState = reactive({"open":props.open,"code":props.code})
const change = (e)=> {
    formState.open = e
    console.log("responseDefine",formState)
    emits("change",formState)
}

const select = () =>{
    console.log("responseDefine",formState)
    emits("change",formState)
}


</script>

<style lang="less" scoped>
.content {
    height: 25px;


    .switch {
        float: left;

        .title {
         padding-left: 2px;
         font-weight: bold;
        }
        .ant-switch {
            margin-left: 3px;
            margin-bottom: 3px;
        }
    }

    .codes {
        float: right;
    }
}
</style>