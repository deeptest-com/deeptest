<template>
  <div class="processor_group_default-main">
    <a-card :bordered="false">
      <a-form
          ref="formRef"
          :rules="rules"
          :model="formState"
          :label-col="{ span: 4 }"
          :wrapper-col="{ span: 16 }">
        <a-form-item label="分组名称" name="name">
          <a-input v-model:value="formState.name"/>
        </a-form-item>

        <a-form-item label="备注" name="comments">
          <a-input v-model:value="formState.comments"/>
        </a-form-item>

        <a-form-item :wrapper-col="{ span: 16, offset: 4 }">
          <a-button type="primary" @click.prevent="submit">保存</a-button>
          <a-button style="margin-left: 10px" @click="reset">重置</a-button>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import {computed, ref, watch} from "vue";
import {useStore} from "vuex";
import {StateType as ScenarioStateType} from "../../../../../store";
import {message} from "ant-design-vue";

const store = useStore<{ Scenario: ScenarioStateType; }>();
const nodeData: any = computed<boolean>(() => store.state.Scenario.nodeData);
const formState: any = ref({
  name: '',
  comments: '',
});
const formRef: any = ref(null);

watch(() => {
  return nodeData.value;
}, (val: any) => {
  debugger;
  if (!val) return;
  formState.value.name = val.name;
  formState.value.comments = val.comments;
});

const rules = {
  name: [
    {required: true, message: '请输入名称', trigger: 'blur'},
  ],
}

const submit = async () => {
  formRef.value
      .validate()
      .then(async () => {
        // 下面代码改成 await 的方式
        const res = await store.dispatch('Scenario/saveProcessor', {
          ...nodeData.value,
          name: formState.value.name,
          comments: formState.value.comments,
        });
        if (res === true) {
          message.success('保存成功');
        } else {
          message.error('保存失败');
        }
      })
      .catch(error => {
        console.log('error', error);
      });
};

const reset = () => {
  formRef.value.resetFields();
};


</script>

<style lang="less" scoped>
.processor_group_default-main {
  .icons {
    text-align: right;
    line-height: 32px;
  }
}
</style>
