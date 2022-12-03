<template>
  <div class="formurlencoded-main">
    <div class="dp-param-grid">
      <div class="params">
        <a-row v-for="(item, idx) in interfaceData.bodyFormUrlencoded" :key="idx" type="flex" class="param">
          <a-col flex="1">
            <a-input v-model:value="item.name" @change="onFormUrlencoded(idx)" class="dp-bg-input-transparent" />
          </a-col>
          <a-col flex="1">
            <a-input v-model:value="item.value" @change="onFormUrlencoded(idx)" class="dp-bg-input-transparent" />
          </a-col>
          <a-col flex="80px" class="dp-right dp-icon-btn-container">
            <a-tooltip v-if="!item.disabled" @click="disable(idx)" overlayClassName="dp-tip-small">
              <template #title>禁用</template>
              <CheckCircleOutlined class="dp-icon-btn dp-trans-80" />
            </a-tooltip>

            <a-tooltip v-if="item.disabled" @click="disable(idx)" overlayClassName="dp-tip-small">
              <template #title>启用</template>
              <CloseCircleOutlined class="dp-icon-btn dp-trans-80 dp-light" />
            </a-tooltip>

            <a-tooltip @click="remove(idx)" overlayClassName="dp-tip-small">
              <template #title>移除</template>
              <DeleteOutlined class="dp-icon-btn dp-trans-80"/>
            </a-tooltip>

            <a-tooltip @click="insert(idx)" overlayClassName="dp-tip-small">
              <template #title>插入</template>
              <PlusOutlined class="dp-icon-btn dp-trans-80"/>
            </a-tooltip>
          </a-col>
        </a-row>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, ComputedRef, defineComponent, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined, DeleteOutlined, PlusOutlined, CheckCircleOutlined, CloseCircleOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import {Param, Interface, BodyFormUrlEncodedItem} from "@/views/interface/data";

const {t} = useI18n();
const store = useStore<{ Interface: StateType }>();
const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

const onFormUrlencoded = (idx) => {
  console.log('onFormUrlencoded', idx)
  if (interfaceData.value.bodyFormUrlencoded.length <= idx + 1
        && (interfaceData.value.bodyFormUrlencoded[idx].name !== '' || interfaceData.value.bodyFormUrlencoded[idx].value !== '')) {
    interfaceData.value.bodyFormUrlencoded.push({} as BodyFormUrlEncodedItem)
  }
}

const add = () => {
  console.log('add')
  interfaceData.value.bodyFormUrlencoded.push({} as BodyFormUrlEncodedItem)
}
const removeAll = () => {
  console.log('removeAll', interfaceData.value.bodyFormUrlencoded)
  interfaceData.value.bodyFormUrlencoded = [{} as BodyFormUrlEncodedItem]
}

const disable = (idx) => {
  console.log('enable', idx)
  interfaceData.value.bodyFormUrlencoded[idx].disabled = !interfaceData.value.bodyFormUrlencoded[idx].disabled
}
const remove = (idx) => {
  console.log('remove')
  interfaceData.value.bodyFormUrlencoded.splice(idx, 1)
  add()
}
const insert = (idx) => {
  console.log('insert')
  interfaceData.value.bodyFormUrlencoded.splice(idx+1, 0, {} as Param)
}

</script>

<style lang="less" scoped>
.formurlencoded-main {
  height: 100%;
}

</style>