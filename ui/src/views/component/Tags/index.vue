<template>
<a-select
      mode="tags"
      v-model:value="values"
      :options="options"
      :size="size"
      placeholder="请选择标签"
      style="width: 200px"
      @blur="blur"
    ></a-select>
</template>

<script setup lang="ts">
import { ref,defineProps } from 'vue';

const props = defineProps({
    size:{
        type:String,
        default:'small'
    },
    
    options:{
        type:[],
        default:[],
    },

    selectValues:{
        type:[],
        default:[],
    },
    ok:{
        type: Function,
    }
    
})

const values = ref(props.selectValues)
const options = ref(props.options)


const blur = ()=>{
    if (props.ok) {
        props.ok(values.value);
    }
}


</script>