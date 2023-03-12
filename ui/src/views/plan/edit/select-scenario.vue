<template>
  <div class="select-scenario-main">
    <a-modal title="导入场景"
             :visible="true"
             :onCancel="onCancel"
             class="select-scenario-modal"
             width="800px">

      <div class="header">
        <div class="space">
          <a-checkbox :checked="isCheckAll" @change="selectAll" />&nbsp;&nbsp;
        </div>
        <div class="left">
          选择所有
        </div>

        <div class="right">
          <a-select
              v-model:value="serveId"
              @change="selectServe"
              :dropdownMatchSelectWidth="false"
              :bordered="false">
            <a-select-option :key="0" :value="0">请选择</a-select-option>
            <a-select-option v-for="(item) in serves" :key="item.id" :value="item.id">
              {{ item.name }}
            </a-select-option>
          </a-select>
        </div>
      </div>

      <div class="body">
        <div class="scenario-list">
          <a-checkbox-group v-model:value="checkedScenarios">
            <div v-for="item in scenarios" :key="item.id" class="scenario-item">
              <div class="no">
                <a-checkbox :value="item.id" />
              </div>
              <div class="name">
                {{ item.name }}
              </div>
            </div>
          </a-checkbox-group>
        </div>
      </div>

      <template #footer>
        <a-button @click="onSubmit" type="primary">导入</a-button>
      </template>
    </a-modal>
  </div>
</template>
<script setup lang="ts">
import {defineProps, ref, watch} from "vue";
import {DeleteOutlined} from '@ant-design/icons-vue';
import {listScenario, listServe} from "@/views/plan/service";

const props = defineProps({
  scenariosInServe: {
    type: Array,
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

const serveId = ref(0)
const serves = ref([] as any[]);
const scenarios = ref([] as any[]);
const checkedScenarios = ref([] as number[]);
const isCheckAll = ref(false)

const loadServe = async () => {
  listServe().then((json) => {
    console.log('listServe', json)
    serves.value = json.data

    if (serves.value.length > 0) {
      serveId.value = serves.value[0].id
      selectServe()
    }
  })
}
loadServe()

const selectServe = async () => {
  if (serveId.value == 0) {
    scenarios.value = []
    return
  }

  listScenario(serveId.value).then((json) => {
    console.log('listScenario', json)
    scenarios.value = json.data

    props.scenariosInServe.forEach((item, index) => {
      checkedScenarios.value.push(+item.id)
    })
  })
}

const selectAll = () => {
  console.log('selectAll')

  checkedScenarios.value = []

  if (isCheckAll.value) {
    return
  }

  scenarios.value.forEach((item, index) => {
    checkedScenarios.value.push(+item.id)
  })
}

watch(checkedScenarios, () => {
  console.log('watch checkedScenarios', checkedScenarios.value)
  if (checkedScenarios.value.length == scenarios.value.length) {
    isCheckAll.value = true
  } else {
    isCheckAll.value = false
  }
})

const onSubmit = () => {
  console.log('onSubmit', checkedScenarios.value)
  props.submit(checkedScenarios.value)
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
    display: flex;
    padding: 0 5px;
    .space {
      width: 60px;
    }
    .left {
      flex: 1;
    }
    .right {
      width: 100px;
      text-align: right;
    }
  }

  .body {
    height: 200px;
    overflow-y: auto;

    .scenario-list {
      .scenario-item {
        display: flex;
        padding: 5px;

        .no {
          width: 60px;
        }

        .name {
          flex: 1;
        }

        .count {
          width: 100px;
        }

        .opt {
          width: 60px;
        }

        &:nth-child(even) {
          background-color: #fafafa;
        }
      }
    }
  }
}
</style>