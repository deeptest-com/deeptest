<template >
       <a-tooltip :visible="visible" style=" background-color: #fafafa">
        <template #title>
            <a-tag  :key="key" v-for="(item,key) in values"  closable @close="close(item)">{{optionsMap.get(item)}}</a-tag>
        </template>
        <a-select
            mode="multiple"
            :maxTagCount="1"
            allowClear
            @change="change"
            :placeholder="placeholder"
            :options="options"
            style="width: 180px;"
            :value="values"
            @focus="focus" 
            @blur="blur"
            :maxTagPlaceholder="maxTagPlaceholder"
            @dropdownVisibleChange="dropdownVisibleChange"
            v-on-click-outside="canClose"
            :open = "open"
            >
            
        </a-select>
        </a-tooltip>
       
</template>

<script type="ts" setup>

import { ref,defineProps,defineEmits,computed, watch } from 'vue';
import { vOnClickOutside } from '@vueuse/components';
const props = defineProps({
    placeholder:{
        type:String,
        default:'small'
    },
    value:{
        type: [],
        required: true,
    },
    width:{
        type:String,
        default:'180px'
    },
    options:{
        type:[],
        default:[],
    },
    
})

const visible = ref(false)

const options = computed(()=> props.options )

const emits = defineEmits('change')

const values = ref(props?.value || [])

const open = ref(false)

const isOpen = ref(false)

const optionsMap = computed(()=>{
   let map= new Map() 
    options.value.forEach((item)=>{
        map.set(item.value,item.label)
    })
    return map
})

const maxTagPlaceholder = (omittedValues)=>{
    let res = ""
    omittedValues.forEach((item)=>{
            res += res? ","+item.label:item.label
    })

    return <a-tooltip  placement="top" title={res}>+{omittedValues.length}...</a-tooltip>
    
}

const change = (e)=>{
    values.value = e
    emits('change',e)
}

const focus = ()=>{
    visible.value = true
    open.value = true
}

const blur = ()=>{
    if (open.value) {
        return 
    }

   // isOpen.value = visible.value
    //visible.value = false
}

const close = (key)=>{
   console.log("close",key)
   //isOpen.value = true
   values.value =  values.value.filter(arrItem => arrItem != key )
}

const dropdownVisibleChange = (open) => {
    console.log("dropdownVisibleChange",open)
    isOpen.value = open
}

function canClose() {
  
  if(!visible.value || isOpen.value ){
    return;
  }
  

  visible.value = false
  open.value = false
  
}

</script>
<style lang="less" scoped>

.ant-tooltip-inner {
    min-width: 30px;
    min-height: 32px;
    padding: 6px 8px;
    color: #fff;
    text-align: left;
    text-decoration: none;
    word-wrap: break-word;
    background-color: #fafafa;
    border-radius: 2px;
    box-shadow: 0 3px 6px -4px rgba(0, 0, 0, 0.12), 0 6px 16px 0 rgba(0, 0, 0, 0.08), 0 9px 28px 8px rgba(0, 0, 0, 0.05);
}



</style>
