<template>
    <div style:="width: 200px">
    <a-tag style="margin:2px;"
    :key="tag" v-for="(tag,index) in values"
    closable @close="close(index)"
    >{{tag}}</a-tag>
    <PlusCircleOutlined  @click="showSelect=true" style="padding-top:3px;margin-left:5px"/>
    <a-select 
        v-if="showSelect"
        v-model:value="tag"
        show-search
        placeholder="请添加标签"
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
import { ref,defineProps,defineEmits,computed, watch } from 'vue';
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

    values:{
        type: [],
        required: true,
    }
    
})

const emits = defineEmits('updateTags')

const values = ref(props?.values||[])

const options = computed(()=> 
   props.options.filter(
        arrItem =>  values.value.indexOf(arrItem.value) == -1
        )   
)

const showSelect = ref(false)
const tag = ref()
const searchValue = ref()

const updateTags = (tags) => {
     emits('updateTags',tags)
}

function canColse() {
  if(isOpen.value){
    return;
  }

  showSelect.value = false;
  tag.value = undefined
}

const isOpen = ref(false);

function dropdownVisibleChange(open) {
  isOpen.value = open;
}

const filterOption = (input: string, option: any) => {
  return option.value.toLowerCase().indexOf(input.toLowerCase()) >= 0;

};


const handleChange = async (value: string) => {
    value = value.trim()
    if (!value) {
        //debugger;
        return 
    }

   values.value = Array.from(new Set([...values.value,value]))
   await updateTags(values.value)
   tag.value = undefined
 };

const handleBlur = () => {
  console.log('blur');
  showSelect.value=false
  tag.value = undefined
};
const handleFocus = () => {
  console.log('focus');
};

const search = (va)=>{
    searchValue.value = va
}

const enter = (value) => {
    if (value.code=="Enter" ){
        console.log('enter',searchValue.value);
        handleChange(searchValue.value)
    }
}


const close = (index)=>{
    console.log("colse",index)
    values.value.splice(index, 1)
    updateTags(values.value)
}

watch(()=>{return props.values},(newVal)=>{
    values.value = [...new Set(newVal)]
})



</script>