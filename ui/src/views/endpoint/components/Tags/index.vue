<template>
    <a-tag 
    :key="index" v-for="(tag,index) in values"
    closable @close="close(index)"
    >{{tag}}</a-tag>
    <PlusCircleOutlined  @click="showSelect=true"/>
    
    <div v-if="showSelect">
    <a-select
        v-model:value="tag"
        show-search
        placeholder="请选择标签"
        style="width: 200px"
        :options="options"
        :filter-option="filterOption"
        @focus="handleFocus"
        @blur="handleBlur"
        @change="handleChange"
        v-on-click-outside="canColse"
        @dropdownVisibleChange="dropdownVisibleChange"
        @inputKeyDown="enter"
        @search="search"
     ></a-select>
     </div>

</template>


<script setup lang="ts">

import {useStore} from "vuex";
import { ref,defineProps,computed, watch } from 'vue';
import { PlusCircleOutlined } from '@ant-design/icons-vue';
import { vOnClickOutside } from '@vueuse/components';
const store = useStore<{ Endpoint }>();

const props = defineProps({
    size:{
        type:String,
        default:'small'
    },
    width:{
        type:String,
        default:'120px'
    },
    options:{
        type:[],
        default:[],
    },

    record:{
        type: Object,
        required: true,
    }
    
})

const values = ref(props.record?.tags)
const options = computed(()=>props.options)
const showSelect = ref(false)
const tag = ref('')
const searchValue = ref()


function canColse() {
  if(isOpen.value){
    return;
  }

  showSelect.value = false;
}

const isOpen = ref(false);

function dropdownVisibleChange(open) {
  isOpen.value = open;
}

const filterOption = (input: string, option: any) => {
  return option.value.toLowerCase().indexOf(input.toLowerCase()) >= 0;
};


const handleChange = (value: string) => {
  values.value = Array.from(new Set([...values.value,value]))
  console.log(`selected ${value}`);
  updateTags(values.value)
};

const handleBlur = () => {
  console.log('blur');
  showSelect.value=false
};
const handleFocus = () => {
  console.log('focus');
};

const search = (va)=>{
    searchValue.value = va
    console.log('search',va);
}

const enter = (value) => {
    if (value.code=="Enter"){
        console.log('enter',value,searchValue.value);
        handleChange(searchValue.value)
    }
}

const updateTags = (tags)=>{
    store.dispatch('Endpoint/updateEndpointTag', {
      id:props.record.id,tagNames:tags,projectId:props.record.projectId
    });
}

const close = (index)=>{
    console.log("colse",index)
    values.value.splice(index, 1)
    updateTags(values.value)
}

watch(()=>{return props.record?.tags},(newVal)=>{
    values.value = [...newVal]
})



</script>