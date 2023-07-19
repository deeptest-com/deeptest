<template>
<a-select
      mode="tags"
      v-model:value="values"
      :options="options"
      :size="size"
      placeholder="请选择标签"
      style="width: 100px"
      @change="change"
    ></a-select>
</template>

<script setup lang="ts">

import {useStore} from "vuex";
import { ref,defineProps,computed, watch } from 'vue';
const store = useStore<{ Endpoint }>();

const props = defineProps({
    size:{
        type:String,
        default:'small'
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

const change = async ()=>{
  //console.log("saveTags",values)
  await store.dispatch('Endpoint/updateEndpointTag',{id:props.record.id,tagNames:values.value,projectId:props.record.projectId});
}

watch(()=>{return props.record?.tags},(newVal)=>{
    values.value = [...newVal]
})


</script>