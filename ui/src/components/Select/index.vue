<template>
  <div :ref="selectRef">
    <a-popover :visible="visible && values?.length"
               :placement="'top'"
               trigger="click"
               :autoAdjustOverflow="false"
               overlayClassName="dp-select-tooltip"
               style=" background-color: #fafafa">
      <template #content>
        <a-tag :key="key" v-for="(item,key) in values" closable @close="close(item)">{{ optionsMap.get(item) }}</a-tag>
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
          :open="open"
      >

      </a-select>
    </a-popover>
  </div>

  v-for=xx




</template>

<script type="ts" setup>

import {ref, defineProps, defineEmits, computed, watch, createVNode} from 'vue';
import {vOnClickOutside} from '@vueuse/components';

const props = defineProps({
  placeholder: {
    type: String,
    default: 'small'
  },
  value: {
    type: [],
    required: true,
  },
  width: {
    type: String,
    default: '180px'
  },
  options: {
    type: [],
    default: [],
  },

})

const visible = ref(false)

const selectRef = ref(null);

const options = computed(() => props.options)

const emits = defineEmits('change')

const values = ref(props?.value || [])

const open = ref(false)

const isOpen = ref(false)

const optionsMap = computed(() => {
  let map = new Map()
  options.value.forEach((item) => {
    map.set(item.value, item.label)
  })
  return map
})

const maxTagPlaceholder = (omittedValues) => {
  let res = ""
  omittedValues.forEach((item) => {
    res += res ? "," + item.label : item.label
  })

  return createVNode('a-tooltip', {
    placement: 'top',
    title: res,
  }, {
    default: () => {
      return `${+omittedValues.length}...`
    },
  })

}

const change = (e) => {
  values.value = e
  emits('change', e)
}

const focus = () => {
  visible.value = true
  open.value = true
}

const blur = () => {
  if (open.value) {
    return
  }

  // isOpen.value = visible.value
  //visible.value = false
}

const close = (key) => {
  console.log("close", key);
  //isOpen.value = true
  // debugger;
  if(selectRef?.value){
    // debugger;
    selectRef.value.focus();
  }
  values.value = values.value.filter(arrItem => arrItem != key)
}

const dropdownVisibleChange = (open) => {
  console.log("dropdownVisibleChange", open)
  isOpen.value = open
}

function canClose(e) {
  const indexlayout = document.getElementById('indexlayout');
  if (indexlayout.contains(e.target)) {
    visible.value = false
    open.value = false
  }
}

</script>
<style lang="less">

.dp-select-tooltip{
  .ant-tooltip-arrow{
    display: none;
  }
  .ant-popover-inner-content {
    padding: 6px 2px 6px 12px;
    min-height: 24px;
    max-height: 64px;
    overflow-y: scroll;
    max-width: 320px;
    background-color: #fff;
    .ant-tag{
      margin-bottom: 3px;
      padding: 0 3px;
    }
  }

}


</style>
