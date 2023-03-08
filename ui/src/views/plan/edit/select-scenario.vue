<template>
  <div class="select-scenario-main">
    <a-modal title="导入场景"
             :visible="isVisible"
             :onCancel="onCancel"
             class="select-scenario-modal"
             width="800px">

      <div class="header">
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

      <div class="body">
        <div class="scenario-list">
          <div v-for="(item, idx) in scenarios" :key="item.id" class="scenario-item">
            <div class="no">
              {{ idx + 1 }}
            </div>
            <div class="name">
              {{ item.name }}
            </div>
            <div class="opt">
              <span>
                <DeleteOutlined @click="removeScenario(item)" class="dp-primary"/>
              </span>
            </div>
          </div>
        </div>
      </div>

      <template #footer>
        <a-button @click="onSubmit" type="primary">导入</a-button>
      </template>
    </a-modal>
  </div>
</template>
<script setup lang="ts">
import {defineProps, ref} from "vue";
import {DeleteOutlined} from '@ant-design/icons-vue';
import {listScenario, listServe} from "@/views/plan/service";

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

const serveId = ref(0)
const serves = ref([]);
const scenarios = ref([]);

const loadServe = async () => {
  listServe().then((json) => {
    console.log('listServe', json)
    serves.value = json.data
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
  })
}

const onSubmit = () => {
  console.log('onSubmit')
  props.submit(serveId, scenarios)
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
    height: 300px;
    overflow-y: auto;

    .scenario-list {
      padding: 16px 0;

      .scenario-item {
        display: flex;
        padding: 5px;

        .no {
          width: 100px;
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