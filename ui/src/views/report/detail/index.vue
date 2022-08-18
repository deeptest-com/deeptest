<template>
  <div class="scenario-edit-main">
    <a-card :bordered="false">
      <template #title>
        <div>测试报告</div>
      </template>
      <template #extra>
        <a-button type="link" @click="() => back()">返回</a-button>
      </template>

      <div>
      </div>
    </a-card>
  </div>
</template>

<script lang="ts">
import {defineComponent, computed, ref, reactive, ComputedRef} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import { useI18n } from "vue-i18n";
import { Props, validateInfos } from 'ant-design-vue/lib/form/useForm';
import {message, Form, notification} from 'ant-design-vue';
const useForm = Form.useForm;
import {StateType} from "../store";
import {Scenario} from "@/views/scenario/data";

export default defineComponent({
    name: 'ScriptEditPage',
    setup() {
      const router = useRouter();
      const { t } = useI18n();

      const store = useStore<{ Scenario: StateType }>();
      const modelRef = computed<Partial<Scenario>>(() => store.state.Scenario.detailResult);

      const get = async (id: number): Promise<void> => {
        await store.dispatch('Scenario/getScenario', id);
      }
      const id = ref(+router.currentRoute.value.params.id)
      get(id.value)

      const back = ():void =>  {
        router.replace(`/scenario/index`)
      }

      return {
        labelCol: { span: 4 },
        wrapperCol: { span: 14 },
        id,
        modelRef,

        back,
      }
    }
})
</script>

<style lang="less" scoped>
.scenario-edit-main {

}
</style>
