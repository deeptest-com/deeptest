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
    <InterfaceBasicInfo :interfaceDetail="interfaceDetail" @changeStatus="changeStatus"/>
    <!-- 接口设计区域 -->
    <a-card
        style="width: 100%"
        title="接口设计"
        :tab-list="tabList"
        :bodyStyle="{padding:'16px'}"
        :active-tab-key="key"
        @tabChange="key => onTabChange(key, 'key')">
      <div v-if="key === 'request'">
        <InterfaceDefine/>
      </div>
      <div v-else-if="key === 'run'">
        <!-- ::::TODO 调试模块 -->
      </div>
      <div v-else-if="key === 'mock'">mock content</div>
    </a-card>
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
  computed,
} from 'vue';
import InterfaceBasicInfo from './InterfaceBasicInfo.vue';
import EditAndShowField from './EditAndShowField.vue';
import InterfaceDefine from './InterfaceDefine.vue';
import {useStore} from "vuex";
import {Interface} from "@/views/interface/data";

const store = useStore<{ Interface, ProjectGlobal, ServeGlobal }>();
const interfaceDetail = computed<Interface>(() => store.state.Interface.interfaceDetail);

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
})
const emit = defineEmits(['ok', 'close', 'refreshList']);

function onCloseDrawer() {
  emit('close');
}

async function changeStatus(status) {
  await store.dispatch('Interface/updateStatus',
      {id:interfaceDetail.value.id, status: status}
  );
  await store.dispatch('Interface/getInterfaceDetail', {id: interfaceDetail.value.id});
}

async function updateTitle(title) {
  await store.dispatch('Interface/updateInterfaceDetail',
      {...interfaceDetail.value, title: title}
  );
  await store.dispatch('Interface/getInterfaceDetail', {id: interfaceDetail.value.id});
}

const tabList = [
  {
    key: 'request',
    tab: '定义',
    slots: {tab: 'customRenderRequest'},
  },
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

async function cancal() {
  emit('close');
}

async function save() {
  await store.dispatch('Interface/updateInterfaceDetail',
      {...interfaceDetail.value}
  );
  emit('close');
}


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
  z-index: 99;
}

</style>
