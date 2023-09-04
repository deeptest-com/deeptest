<template>
  <a-modal
      width="640px"
      :visible="visible"
      @ok="ok"
      @cancel="cancal"
      title="发布文档">
    <a-form
        ref="formRef"
        :model="formState"
        :rules="rules"
        :label-col="{ span: 5 }"
        :wrapper-col="{ span: 17 }">

      <a-form-item label="是否新建版本" name="isNewVersion">
        <a-space direction="vertical">
          <a-radio-group style="width: 400px" v-model:value="formState.isNewVersion" button-style="solid"
                         :options="pubOptions"/>
        </a-space>
      </a-form-item>


      <a-form-item label="版本号" name="version" v-if="formState.isNewVersion">
        <a-input style="width: 400px" v-model:value="formState.version"/>
      </a-form-item>

      <a-form-item label="描述" name="name" v-if="formState.isNewVersion">
        <a-input style="width: 400px" v-model:value="formState.name"/>
      </a-form-item>

      <a-form-item label="版本号" name="selectedVersion" v-if="!formState.isNewVersion">
        <a-select
            v-model:value="formState.selectedVersion"
            style="width: 400px"
            placeholder="请选择版本"
            :options="versionOptions"
        ></a-select>
      </a-form-item>

    </a-form>
  </a-modal>
</template>
<script lang="ts" setup>
import {ValidateErrorEntity} from 'ant-design-vue/es/form/interface';
import {
  reactive,
  ref,
  UnwrapRef,
  defineProps,
  defineEmits,
  onMounted,
  computed, watch,
} from 'vue';
import {useStore} from "vuex";
import {message, notification} from "ant-design-vue";
import {NewEndpointFormState} from "@/views/Endpoint/data";
import {InboxOutlined, UploadOutlined} from '@ant-design/icons-vue';
import {notifySuccess} from "@/utils/notify";

const store = useStore<{ Endpoint }>();

// 是否新建版本
const pubOptions = [
  {
    value: 1,
    label: '是'
  },
  {
    value: 0,
    label: '否'
  }
];
const versionOptions: any = ref([]);

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
  endpointIds: {
    required: true,
    type: Array,
    default: () => [],
  },
})

const emit = defineEmits(['ok', 'cancal']);

const formRef = ref();

async function ok() {
  formRef.value
      .validate()
      .then(async () => {
        let name: any = '';
        let version: any = '';
        if (formState.value.isNewVersion) {
          name = formState.value.name;
          version = formState.value.version;
        } else {
          versionOptions.value.forEach((item: any) => {
            if (item.value === formState.value.selectedVersion) {
              name = item.name;
              version = item.value;
            }
          });
        }
        const res = await store.dispatch('Docs/publishDocument',
            {
              name,
              version,
              endpointIds: props.endpointIds.map((item: any) => Number(item)),
            }
        );
        if (res) {
          notifySuccess('发布成功');
          window.location.href = `/#/docs/index?documentId=${res.data}`
        }
        reset();
        emit('ok');
      })
      .catch((error: ValidateErrorEntity) => {
        console.log('error', error);
      });
}

function cancal() {
  emit('cancal', formState.value);
  reset();
}

function reset() {
  formRef.value.resetFields();
}


const formState = ref({
  "name": null,
  "isNewVersion": 1,
  "version": null,
  "selectedVersion": null,
});

watch(() => {
  return props.visible
}, async (newVal) => {
  if (newVal) {
    const res = await store.dispatch('Docs/getVersionList');
    versionOptions.value = res.map((item: any) => {
      return {
        ...item,
        label: item.name ? `${item.version}（${item.name}）` : item.version,
        value: item.version,
      }
    });
  }
}, {
  immediate: true
})

const rules = computed(() => {

  if (formState.value.isNewVersion) {
    return {
      isNewVersion: [{required: true}],
      name: [{required: false, message: '请输入版本描述信息'}],
      version: [{required: true, message: '请输入正确版本号信息，示例 1.0.0 ', pattern: /^\d{1,2}\.\d{1,2}\.\d{1,2}$/}],
    }
  } else {
    return {
      isNewVersion: [{required: true}],
      selectedVersion: [{required: true, message: '请选择文档版本'}],
    }
  }

});


</script>

<style lang="less" scoped>
.modal-btns {
  display: flex;
  justify-content: flex-end;
}
</style>
