<template>
  <a-drawer
      :placement="'right'"
      :width="1000"
      :closable="true"
      :visible="visible"
      class="drawer"
      wrapClassName="drawer-1"
      :bodyStyle="{padding:0}"
      @close="onCloseDrawer"
  >
    <template #title>
      <div> title="[接口编号] 接口名称"</div>
    </template>

    <!-- 基本信息  -->
    <a-card
        class="card-baseInfo"
        :bordered="false"
        title="基本信息">
      <template #extra><a href="#">more</a></template>
      <a-descriptions :title="null">
        <a-descriptions-item label="创建人">Zhou Maomao</a-descriptions-item>
        <a-descriptions-item label="创建时间">2022.11.11 11:22:232</a-descriptions-item>
        <a-descriptions-item label="状态">Hangzhou, Zhejiang</a-descriptions-item>
        <a-descriptions-item label="服务版本">empty</a-descriptions-item>
        <a-descriptions-item label="最近更新">2022.11.11 11:22:232</a-descriptions-item>
      </a-descriptions>
    </a-card>

    <!-- 接口定义 -->
    <a-card
        style="width: 100%"
        title="接口定义"
        :tab-list="tabList"
        :active-tab-key="key"
        @tabChange="key => onTabChange(key, 'key')"
    >
      <div v-if="key === 'request'">
        <a-collapse :activeKey="collapseActiveKey">
          <a-collapse-panel key="1" >
            <template #header>
              <a-form-item label="路径">
                <a-input style="width: 800px">
                  <template #addonBefore>
                    <a-select :value="'http://localhost:3000'" style="width: 200px">
                      <a-select-option value="http://localhost:3000">http://localhost:3000</a-select-option>
                      <a-select-option value="http://localhost:3001">http://localhost:3001</a-select-option>
                    </a-select>
                  </template>
                  <template #addonAfter>
                    <a-button type="primary">
                      <template #icon><PlusOutlined /></template>
                      路径参数
                    </a-button>
                  </template>
                </a-input>
              </a-form-item>
            </template>
            <p>{{ 121212 }}</p>
          </a-collapse-panel>

          <a-collapse-panel key="2" >
            <template #header>
              <a-form-item label="请求方式">
                <a-tabs v-model:activeKey="activeKey">
                  <a-tab-pane key="1" tab="Tab 1">Content of Tab Pane 1</a-tab-pane>
                  <a-tab-pane key="2" tab="Tab 2" force-render>Content of Tab Pane 2</a-tab-pane>
                  <a-tab-pane key="3" tab="Tab 3">Content of Tab Pane 3</a-tab-pane>
                </a-tabs>
              </a-form-item>
            </template>
            <p>{{ 121212 }}</p>
          </a-collapse-panel>


        </a-collapse>



      </div>
      <div v-else-if="key === 'response'">response content</div>
      <div v-else-if="key === 'run'">run content</div>
      <div v-else-if="key === 'mock'">mock content</div>



      <template #extra>
        <a href="#">More</a>
      </template>


      <!--      {{ contentList[key] }}-->


    </a-card>




    <div class="drawer-btns">
      <a-space>
        <a-button type="primary">保存</a-button>
        <a-button>取消</a-button>
      </a-space>
    </div>




  </a-drawer>
</template>

<script lang="ts" setup>
import {ValidateErrorEntity} from 'ant-design-vue/es/form/interface';
import {defineComponent, reactive, ref, toRaw, UnwrapRef, defineProps, defineEmits, watch} from 'vue';
import {requestMethodOpts} from '@/config/constant';
import { PlusOutlined } from '@ant-design/icons-vue';
const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  }
})

const emit = defineEmits(['ok', 'close']);

const collapseActiveKey = ref('1');
const activeKey = ref('1');

function onCloseDrawer() {
  emit('close');
}


interface FormState {
  name: string;
  remark: string | undefined;
}


const tabList = [
  {
    key: 'request',
    tab: '请求定义',
    slots: {tab: 'customRenderRequest'},
  },
  {
    key: 'response',
    tab: '响应定义',
    slots: {tab: 'customRenderResponse'},
  },
  {
    key: 'run',
    tab: '运行调试',
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
  console.log(value, type);
  if (type === 'key') {
    key.value = value;
  }
};


const formRef = ref();

const formState: UnwrapRef<FormState> = reactive({
  name: '接口类型1',
  remark: '用户信息相关',
});

const rules = {
  name: [
    {required: true, message: '请输入接口名称', trigger: 'blur'},
    {min: 3, max: 50, message: '最长多少个字符', trigger: 'blur'},
  ],
  path: [{required: true, message: 'Please select Activity zone', trigger: 'change'}],
  tag: [{required: true, message: 'Please select activity resource', trigger: 'change'}],
};


</script>

<style lang="less" scoped>
.drawer-btns {
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
}

.drawer {


}

.card-baseInfo {
  width: 100%;
}

</style>
