<template>
  <a-drawer
      :placement="'right'"
      :width="1000"
      :closable="true"
      :visible="visible"
      class="drawer"
      wrapClassName="drawer-1"
      :bodyStyle="{padding:0,marginBottom:'60px'}"
      @close="onCloseDrawer">
    <!-- 头部信息  -->
    <template #title>
      <a-row type="flex" style="align-items: center;width: 100%">
        <a-col :span="8">
          <EditAndShowField :value="interfaceDetail.title" @update="updateTitle"/>
        </a-col>
      </a-row>
    </template>
    <!-- 基本信息 -->
    <InterfaceBasicInfo :interfaceDetail="interfaceDetail"/>
    <!-- 接口定义 -->
    <a-card
        style="width: 100%"
        title="接口定义"
        :tab-list="tabList"
        :active-tab-key="key"
        @tabChange="key => onTabChange(key, 'key')"
    >
      <div v-if="key === 'request'">
        <InterfaceForm  v-if="showMode === 'form'" />
        <div class="interface-code" v-if="showMode === 'code'">
          <MonacoEditor
              class="editor"
              :value="yamlCode"
              :language="'yaml'"
              :height="600"
              theme="vs"
              :options="{...MonacoOptions}"
              @change="() => {}"
          />
        </div>
      </div>
      <div v-else-if="key === 'run'">run content</div>
      <div v-else-if="key === 'mock'">mock content</div>
      <template #extra>
        <a-button :type="showMode === 'form' ? 'primary' : 'default'" @click="switchMode('form')">
          <template #icon>
            <BarsOutlined/>
          </template>
          图形
        </a-button>
        <a-button :type="showMode === 'code' ? 'primary' : 'default'" @click="switchMode('code')">
          <template #icon>
            <CodeOutlined/>
          </template>
          YAML
        </a-button>
      </template>
    </a-card>
    <!-- ::::接口提交按钮 -->
    <div class="drawer-btns">
      <a-space>
        <a-button type="primary" @click="save">保存</a-button>
        <a-button @click="cancal">取消</a-button>
      </a-space>
    </div>
  </a-drawer>
</template>

<script lang="ts" setup>

import {
  ref,
  defineProps,
  defineEmits,
  watch,
  computed,
  onUnmounted
} from 'vue';

import InterfaceBasicInfo from './InterfaceBasicInfo.vue';
import EditAndShowField from './EditAndShowField.vue';
import InterfaceForm from './InterfaceForm.vue';
import {requestMethodOpts, interfaceStatus, mediaTypesOpts, repCodeOpts} from '@/config/constant';
import {getInterfaceDetail, saveInterface, getYaml} from '../../service';
import {PlusOutlined, EditOutlined, CodeOutlined, BarsOutlined} from '@ant-design/icons-vue';
import {message} from 'ant-design-vue';
import contenteditable from 'vue-contenteditable';
import FieldItem from './FieldItem.vue'
import {momentUtc} from '@/utils/datetime';
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {MonacoOptions} from '@/utils/const';
import _default from "ant-design-vue/lib/color-picker";
import unmounted = _default.unmounted;
import SchemaEditor from '@/components/SchemaEditor/index.vue';
import {example2schema, schema2example} from "@/views/projectSetting/service";
import {useStore} from "vuex";

import {Interface, PaginationConfig} from "@/views/interface/data";
const store = useStore<{ Interface, ProjectGlobal }>();
const interfaceDetail = computed<Interface[]>(() => store.state.Interface.interfaceDetail);
const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
  interfaceId: {
    required: true,
  }
})
const emit = defineEmits(['ok', 'close','refreshList']);
const collapseActiveKey = ref(['1']);
const activeKey = ref('1');
const activeResCodeKey = ref('1');
const selectedMethod = ref('GET');
const selectedCode = ref('200');

function onCloseDrawer() {
  emit('close');
}

const selectedMethodDetail: any = ref(null);
const selectedCodeDetail: any = ref(null);


const showMode = ref('form');

const yamlCode = ref('');

async function switchMode(val) {
  showMode.value = val;
  // 需求去请求YAML格式
  if (val === 'code') {
    let res = await getYaml(interfaceDetail.value);
    yamlCode.value = res.data;
  }
}

function paramsNameChange(val) {
  // todo 待解析，联动接口字段
  // var a = 'api/user/{id}/{detailID}'
  // 解析path 中的参数
  // let parsePathReg = /\{(\w+)\}/g
  // let path = interfaceDetail.value.path;
  // let params = path.match(parsePathReg);
  // if (val) {
  //   params.push(`{${val}}`)
  // }
  // // todo 需要处理，几个表单项的联动场景
  // console.log(832, params, val);
  // interfaceDetail.value.path = path.replace(`{${data.name}}`, '');
}



function updateTitle(title) {
  // ::::todo dispatch一个东西
  // interfaceDetail.title = title
}



const tabList = [
  {
    key: 'request',
    tab: '定义',
    slots: {tab: 'customRenderRequest'},
  },
  // {
  //   key: 'response',
  //   tab: '响应定义',
  //   slots: {tab: 'customRenderResponse'},
  // },
  {
    key: 'run',
    tab: '调试',
    slots: {tab: 'customRenderRun'},

  },
  {
    key: 'mock',
    tab: 'Mock',
    slots: {tab: 'customRenderMock'},
  },
];

const key = ref('request');

const onTabChange = (value: string, type: string) => {
  if (type === 'key') {
    key.value = value;
  }
};



// 取消
async function cancal() {
  emit('close');
}

// 保存
async function save() {
  let res = await saveInterface(interfaceDetail.value);
  if (res.code === 0) {
    message.success('保存成功');
    emit('close');
    emit('refreshList')
  }
}

const loading = ref(false);

/**
 * 打开窗口时，需要重新获取
 * */
watch(() => {
  return props.visible;
}, async (newVal) => {
  if (newVal) {
    // // todo 默认选中第一个有值的method ，临时方案，应该高亮展示一些场景
    // if (interfaceDetail.value.interfaces[0]?.method) {
    //   selectedMethod.value = interfaceDetail.value.interfaces[0].method;
    //   selectedMethodDetail.value = interfaceDetail.value.interfaces[0];
    // }
    // if (selectedMethodDetail.value?.responseBodies[0]?.code) {
    //   selectedCode.value = selectedMethodDetail.value?.responseBodies[0]?.code;
    //   selectedCodeDetail.value = selectedMethodDetail.value?.responseBodies[0];
    // }
  } else {
    // interfaceDetail.value = null;
    // selectedMethodDetail.value = null;
    // selectedCodeDetail.value = null;
  }
}, {
  immediate: true
})

</script>

<style lang="less" scoped>

.drawer {
  margin-bottom: 60px;

  .title {
    width: auto;

    .ant-input-affix-wrapper {
      width: auto;
      border: none;

      &:focus {
        border: none;
        outline: none;
        box-shadow: none;
      }
    }

    input {
      width: auto;
      border: none;

      &:focus {
        border: none;
        border: none;
        outline: none;
        box-shadow: none;
      }
    }
  }


}
.drawer-btns {
  background: #ffffff;
  border-top: 1px solid rgba(0, 0, 0, 0.06);
  position: absolute;
  bottom: 0;
  right: 0;
  width: 100%;
  height: 60px;
  display: flex;
  justify-content: flex-end;
  align-items: center;
  margin-right: 16px;
}

:deep(.ant-drawer-body) {
  padding: 0 !important;
  border: 1px solid red;
  margin-bottom: 60px;
}

.drawer {

}

.interfaceName {
  min-width: 1em;
  &:focus {
    outline: none;
  }
  &:hover,
  &:focus {
    outline: none;
    //border-bottom: 1px solid rgba(0, 0, 0, 0.65);;
  }
}
.schema {
  margin-left: 20px;
  width: 50%;
  height: 100%;
  overflow-y: auto;
  overflow-x: hidden;
  border: 1px solid rgba(0, 0, 0, .1);
  border-radius: 8px;
  padding: 12px;
}

</style>
