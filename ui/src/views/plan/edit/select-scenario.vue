<template>
  <div class="select-scenario-main">
    <a-modal title="导入场景"
             :visible="isVisible"
             :onCancel="onCancel"
             class="select-scenario-modal"
             width="700px">

      <div class="header">
        <a-select
            v-model:value="serviceId"
            :dropdownMatchSelectWidth="false"
            :bordered="false">
          <a-select-option v-for="(item) in scenarios" :key="item.id" :value="item.id">
            {{item.name}}
          </a-select-option>
        </a-select>
      </div>

      <div class="body">

      </div>

      <template #footer>
        <a-button @click="onSubmit" type="primary">导入</a-button>
      </template>
    </a-modal>
  </div>
</template>
<script setup lang="ts">
import {defineProps, ref} from "vue";

const props = defineProps({
  isVisible: {
    type: Boolean,
    required: true
  },
  submit: {
    type: Function,
    required: true,
  },
  cancel: {
    type: Function,
    required: true,
  },
})

const serviceId = ref(0)
const scenarios = ref([]);

const onSubmit = () => {
  console.log('onSubmit')
  props.submit(serviceId, scenarios)
}

const onCancel = () => {
  console.log('onCancel')
  props.cancel()
}

</script>

<style lang="less" scoped>
.select-scenario-main {

}
</style>

<style lang="less">
.select-scenario-modal {
  .header {
    text-align: right;
  }
  .body {

  }
}
</style>