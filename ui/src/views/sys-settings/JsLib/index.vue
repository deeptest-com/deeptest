<template>
  <div class="js-lib-list-main">
    <a-card :bordered="false">
      <template #title>
        <a-button type="primary" @click="() => onEdit(0)">新建自定义库</a-button>
      </template>
      <template #extra>
        <a-input-search
            @change="onSearch"
            @search="onSearch"
            v-model:value="queryParams.keywords"
            placeholder="输入名称搜索"
            class="search" />
      </template>

      <EmptyComp>
        <template #content>
          <a-table :data-source="models" :columns="jslibColumns" :rowKey="(_record, index) => _record.id">
            <template #name="{ text, record }">
              <div class="record-name">
                <EditAndShowField :custom-class="'custom-serve show-on-hover'" placeholder="请输入自定义脚本库名称"
                                  :value="text || ''"
                                  @update="(e: string) => updateName(e, record)"
                                  @edit="onEdit(record.id)"/>
              </div>
            </template>

            <template #customStatus="{ record }">
              <a-tag :color="disabledStatusTagColor.get(record.disabled ? 1 : 0)">
                {{ disabledStatus.get(record.disabled ? 1 : 0) }}
              </a-tag>
            </template>

            <template #typesFile="{ text }">
             {{ text }}
            </template>

            <template #scriptFile="{ text }">
              {{ text }}
            </template>

            <template #createUser="{record}">
              <div class="customTagsColRender">
                {{ record.createUser }}
              </div>
            </template>

            <template #createdAt="{ record }">
              <span>{{ momentUtc(record.createdAt) }}</span>
            </template>

            <template #updatedAt="{ record }">
              <span>{{ momentUtc(record.updatedAt) }}</span>
            </template>

            <template #operation="{ record }">
              <a-dropdown>
                <MoreOutlined/>
                <template #overlay>
                  <a-menu>
                    <a-menu-item key="1">
                      <span class="dp-link operation" @click="onDisable(record)">
                        <span v-if="!record.disabled">禁用</span>
                        <span v-else>解禁</span>
                      </span>
                    </a-menu-item>

                    <a-menu-item key="2">
                      <span class="dp-link operation" @click="onDelete(record)">删除</span>
                    </a-menu-item>
                  </a-menu>
                </template>
              </a-dropdown>
            </template>
          </a-table>
        </template>
      </EmptyComp>

      <!-- 编辑抽屉 -->
      <EditDrawer :modelId="modelId"
              :visible="drawerVisible"
              :onClose="onClose" />
    </a-card>
  </div>
</template>

<script setup lang="ts">
import {computed, createVNode, onMounted, reactive, ref, watch} from 'vue';
import {useStore} from "vuex";
import {ExclamationCircleOutlined, MoreOutlined} from '@ant-design/icons-vue';
import {StateType as SysSettingStateType} from '../store';
import {useI18n} from "vue-i18n";
import {disabledStatus, disabledStatusTagColor} from "@/config/constant"
import {notifyError, notifySuccess} from "@/utils/notify";
import {Modal} from "ant-design-vue";
import {momentUtc} from '@/utils/datetime';

import {jslibColumns} from '../config';
import EditAndShowField from '@/components/EditAndShow/index.vue';
import EmptyComp from '@/components/TableEmpty/index.vue';
import EditDrawer from './drawer.vue';
import debounce from "lodash.debounce";

const {t} = useI18n();

const store = useStore<{ SysSetting: SysSettingStateType }>();
const models = computed<any>(() => store.state.SysSetting.jslibModels);

const drawerVisible = ref(false);
const modelId = ref(0);

const queryParams = ref<any>({
  keywords: '',
})
const onSearch = debounce(() => {
  list();
}, 500)

list()
function list() {
  console.log('list', queryParams.value)
  store.dispatch('SysSetting/listJslib', queryParams.value)
}

function updateName(value: string, record: any) {
  store.dispatch('SysSetting/updateJslibName', {
    id: record.id,
    name: value
  });
}

const onEdit = (id) => {
  console.log('onEdit')
  modelId.value = id;
  drawerVisible.value = true;
}

async function onDelete(record: any) {
  Modal.confirm({
    title: '确认要删除该自定义脚本库吗',
    icon: createVNode(ExclamationCircleOutlined),
    onOk() {
      store.dispatch('SysSetting/deleteJslib', record.id);
    }
  })
}

async function onDisable(record: any) {
  store.dispatch('SysSetting/disableJslib', record.id);
}

function onClose() {
  drawerVisible.value = false;
}

</script>

<style scoped lang="less">
.js-lib-list-main {
  margin: 20px;

  .search {
    width: 270px; margin-left: 16px
  }

  .operation {
    text-align: center;
    display: inline-block;
    width: 80px;
  }

  .record-name {
    width: 120px;
  }
}
</style>
