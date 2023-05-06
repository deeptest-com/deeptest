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
import {listScenario} from "@/views/plan/service";
import {listServe} from "@/services/serve";

const props = defineProps({
  submit: {
    type: Function,
    required: true,
  },
  cancel: {
    type: Function,
    required: true,
  },
})

const scenarios = ref([] as any[]);
const checkedScenarios = ref([] as number[]);
const isCheckAll = ref(false)

const listAllScenario = () => {
  listScenario().then((json) => {
    console.log('listScenario', json)
    scenarios.value = json.data
  })
}
listAllScenario()

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