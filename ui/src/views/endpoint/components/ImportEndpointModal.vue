<template>
  <a-modal
      width="640px"
      :visible="visible"
      @ok="ok"
      @cancel="cancal"
      title="导入接口数据">
    <a-form
        ref="formRef"
        :model="formState"
        :rules="rules"
        :label-col="{ span: 5 }"
        :wrapper-col="{ span: 17 }">
      <a-form-item label="接口数据来源" name="driverType">
        <a-select
            style="width: 100%"
            v-model:value="formState.driverType"
            :options="driverTypeOpts"
            placeholder="请选择"/>
      </a-form-item>
      <a-form-item label="所属分类" name="categoryId">
        <a-tree-select
            @change="selectedCategory"
            :value="formState.categoryId"
            show-search
            :multiple="false"
            :treeData="treeData"
            style="width: 100%"
            :treeDefaultExpandAll="true"
            :replaceFields="{ title: 'name',value:'id'}"
            :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
            placeholder="请选择所属分类"
            allow-clear/>
      </a-form-item>
      <a-form-item label="数据同步方式" name="dataSyncType">
        <a-select
            style="width: 100%"
            v-model:value="formState.dataSyncType"
            :options="dataSyncTypeOpts"
            placeholder="请选择"/>
      </a-form-item>
      <a-form-item label="开启url导入" name="openUrlImport">
        <a-radio-group :disabled="disabled"
            :options="openUrlImportOpts"
            v-model:value="formState.openUrlImport"/>
      </a-form-item>
      <a-form-item label="上传文件" name="filePath" v-if="!formState.openUrlImport">
        <a-spin tip="上传中..." :spinning="uploading">
          <a-upload
              :fileList="fileList"
              accept=".json"
              :remove="handleRemove"
              @change="handleChangeFile"
              :before-upload="beforeUpload">
            <a-button>
              <upload-outlined></upload-outlined>
              点击上传文件
            </a-button>
          </a-upload>
        </a-spin>
      </a-form-item>
      <a-form-item label="swagger url" v-if="formState.openUrlImport" name="filePath">
      <a-input v-model:value="formState.filePath" />
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
  computed, watch,
} from 'vue';
import {useStore} from "vuex";
import {NewEndpointFormState} from "@/views/Endpoint/data";
import {InboxOutlined, UploadOutlined} from '@ant-design/icons-vue';

const store = useStore<{ Endpoint }>();
const treeDataCategory = computed<any>(() => store.state.Endpoint.treeDataCategory);

const driverTypeOpts = [

  {
    label: 'Postman',
    value: 'postman',
  },
  /*
  {
    label: 'Yapi',
    value: 'yapi',
  },
  */
  {
    label: 'Swagger',
    value: 'swagger',
  },
]

const dataSyncTypeOpts = [
  {
    label: '完全覆盖',
    value: 'full_cover',
  },
  {
    label: '复制新增',
    value: 'copy_add',
  }
]

const openUrlImportOpts = [
  {
    label: '是',
    value: true,
  },
  {
    label: '否',
    value: false,
  }
]
const treeData: any = computed(() => {
  const data = treeDataCategory.value;
  return data?.[0]?.children || [];
});

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
  selectedCategoryId: {
    required: true,
  }
})

const emit = defineEmits(['ok', 'cancal']);

const formRef = ref();

function ok() {
  if (uploading.value) {
    return;
  }
  formRef.value
      .validate()
      .then(() => {
        emit('ok', formState.value, () => {
          reset();
        });
      })
      .catch((error: ValidateErrorEntity) => {
        console.log('error', error);
      });
}

function cancal() {
  emit('cancal', formState.value);
  reset();
}

function selectedCategory(value) {
  formState.value.categoryId = value;
}

interface FileItem {
  uid: string;
  name?: string;
  status?: string;
  response?: string;
  url?: string;
  preview?: string;
  originFileObj?: any;
  file: string | Blob;
}

const fileList = ref<FileItem[]>([]);

const uploading = ref<boolean>(false);

const beforeUpload = (file) => {
  fileList.value = [file];
  return false;
};

function reset() {
  formRef.value.resetFields();
  fileList.value = [];
}

watch(() => {
  return fileList.value
}, async (newVal) => {
  if (newVal.length === 1) {
    uploading.value = true;
    const formData: any = new FormData();
    formData.append('file', newVal[0]);
    const res = await store.dispatch('Endpoint/upload', {
      file: formData,
    });
    if (res?.path) {
      formState.value.filePath = res.path;
    } else {
      // 没有上传成功
      fileList.value = [];
    }
    uploading.value = false;
  }
}, {
  immediate: false,
  deep: true
})

function handleChangeFile() {
  console.log('handleChangeFile', fileList.value)
}

function handleRemove() {
  // console.log('handleRemove', fileList.value);
  fileList.value = [];
}

const formState = ref({
  categoryId: null as any,
  driverType: null,
  "dataSyncType": null,   //数据同步方式 枚举值 full_cover：完全覆盖 copy_add：复制新增
  "openUrlImport": false,  //开启url导入
  "filePath": null, //文件路径
});

watch(() => {
  return props.visible
}, (newVal) => {
  // if(newVal) {
  //   reset();
  // }
  if (newVal && props.selectedCategoryId) {
    formState.value.categoryId = props.selectedCategoryId || null;

  }
}, {
  immediate: true
})

const rules = {
  categoryId: [{required: false}],
  driverType: [{required: true, message: '请选择接口数据来源'}],
  dataSyncType: [{required: true, message: '请选择数据同步方式'}],
  openUrlImport: [{required: false}],
  filePath: [{required: true, message: '请上传文件或输入url地址'}],
};

const disabled = computed(()=>{
  return !(formState.value.driverType == "swagger2" || formState.value.driverType == "swagger3")
})



</script>

<style lang="less" scoped>
.modal-btns {
  display: flex;
  justify-content: flex-end;
}
</style>
