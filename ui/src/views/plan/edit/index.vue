<template>
  <div class="plan-edit-main">
    <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
      <a-form-item label="名称">
        <a-input v-if="fieldName==='name'"
                 v-model:value="modelRef.name"
                 @focusout="saveName"
                 @pressEnter="saveName" />

        <span v-else>
          {{ modelRef.name }}
          <edit-outlined class="editable-cell-icon" @click="editField('name')"/>
        </span>
      </a-form-item>

      <a-form-item label="描述">
        <a-input v-if="fieldName==='desc'"
                 v-model:value="modelRef.desc"
                 @focusout="saveDesc"
                 @pressEnter="saveDesc" />

        <span v-else>
              {{ modelRef.desc }}
              <edit-outlined class="editable-cell-icon" @click="editField('desc')"/>
            </span>
      </a-form-item>

      <a-form-item label="场景">
        <div class="scenario-list">
          <div class="scenario-item">
            <div class="no"></div>
            <div class="name"></div>
            <div class="count"></div>
            <div class="opt">
              <span @click="selectScenario()" class="dp-link-primary">导入场景</span>
            </div>
          </div>

          <div v-for="(item, idx) in scenarios" :key="item.id" class="scenario-item">
            <div class="no">
              {{idx+1}}
            </div>
            <div class="name">
              {{item.name}}
            </div>
            <div class="count">
              {{item.interfaceCount}}
            </div>
            <div class="opt">
              <span>
                <DeleteOutlined @click="removeScenario(item)" class="dp-primary"/>
              </span>
            </div>
          </div>
        </div>
      </a-form-item>
    </a-form>

    <SelectScenario
        v-if="modalVisible"
        :scenariosInServe="scenarios"
        :submit="addScenarioToServe"
        :cancel="() => modalVisible = false"
    />
  </div>
</template>

<script setup lang="ts">
import {defineProps, PropType, reactive, ref} from "vue";
import {EditOutlined, DeleteOutlined} from '@ant-design/icons-vue';
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {useI18n} from "vue-i18n";
import {Form, notification} from 'ant-design-vue';
import {StateType} from "../store";
import {addScenarios, get, getDetail, removeScenarioFromPlan} from "@/views/plan/service";
import SelectScenario from "./select-scenario.vue"
import {NotificationKeyCommon} from "@/utils/const";

const useForm = Form.useForm;

const router = useRouter();
const {t} = useI18n();

const props = defineProps({
  modelId: {
    type: Number,
    required: true
  },
  categoryId: {
    type: Number,
    required: true
  },
  onFieldSaved: {
    type: Function as PropType<() => void>,
    required: true
  }
})

const store = useStore<{ Plan: StateType }>();
const modelRef = ref({} as any)
const scenarios = ref([])

const fieldName = ref('')
const modalVisible = ref(false);

const getData = (id: number) => {
  if (id === 0) {
    modelRef.value = {}
    return
  }

  getDetail(id).then((json) => {
    if (json.code === 0) {
      modelRef.value = json.data
      scenarios.value = json.data.scenarios
    }
  })
}
getData(props.modelId)

const editField = (field) => {
  console.log('edit')
  fieldName.value = field
}

const selectScenario = () => {
  console.log('selectScenario')
  modalVisible.value = true
}
const addScenarioToServe = (scenarios) => {
  console.log('addScenarios', props.modelId, scenarios)
  addScenarios(props.modelId, scenarios).then((json) => {
    if (json.code === 0) {
      modalVisible.value = false
      getData(props.modelId)
    }
  })
}

const removeScenario = (item) => {
  console.log('removeScenario')
  removeScenarioFromPlan(props.modelId, item.id).then((json) => {
    if (json.code === 0) {
      getData(props.modelId)
    }
  })
}

const saveName = () => {
  console.log('saveName')
  if (!modelRef.value.name) return
  saveModel()
}
const saveDesc = () => {
  console.log('saveDesc')
  saveModel()
}

const saveModel = async () => {
  console.log('saveModel');
  store.dispatch('Plan/savePlan', modelRef.value).then((res) => {
    console.log('res', res)
    fieldName.value = ''
    if (res === true) {
      props.onFieldSaved()
    }
  })
};

const labelCol = {span: 3}
const wrapperCol = {span: 21}

</script>

<style lang="less" scoped>
.plan-edit-main {
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
</style>
