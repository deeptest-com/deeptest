<template>
  <div class="test-interface-design-main">
    <div class="tabs">
      <a-tabs v-model:activeKey="activeKey" type="editable-card" @edit="onEdit">
        <a-tab-pane v-for="pane in panes" :key="pane.key" :tab="pane.title">
          <UrlInput />

        </a-tab-pane>
      </a-tabs>
    </div>

    <div class="selection">
      <EnvSelection />
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, ref} from 'vue';
import {useStore} from "vuex";

import {StateType as ProjectStateType} from "@/store/project";
import {StateType as TestInterfaceStateType} from '../store';
import {StateType as ServeStateType} from "@/store/serve";

import EnvSelection from './env-selection.vue'
import UrlInput from './url-input.vue'

const store = useStore<{ TestInterface: TestInterfaceStateType, ProjectGlobal: ProjectStateType, ServeGlobal: ServeStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const currServe = computed<any>(() => store.state.ServeGlobal.currServe);

const treeData = computed<any>(() => store.state.TestInterface.treeData);
const treeDataMap = computed<any>(() => store.state.TestInterface.treeDataMap);

const panes = ref([
  { title: 'Tab 1', content: 'Content of Tab 1', key: '1' },
  { title: 'Tab 2', content: 'Content of Tab 2', key: '2' },
  { title: 'Tab 3', content: 'Content of Tab 3', key: '3' },
]);

const activeKey = ref(panes.value[0].key);

const newTabIndex = ref(0);

const callback = (key: string) => {
  console.log(key);
};

const add = () => {
  activeKey.value = `newTab${++newTabIndex.value}`;
  panes.value.push({ title: 'New Tab', content: 'Content of new Tab', key: activeKey.value });
};

const remove = (targetKey: string) => {
  let lastIndex = 0;
  panes.value.forEach((pane, i) => {
    if (pane.key === targetKey) {
      lastIndex = i - 1;
    }
  });
  panes.value = panes.value.filter(pane => pane.key !== targetKey);
  if (panes.value.length && activeKey.value === targetKey) {
    if (lastIndex >= 0) {
      activeKey.value = panes.value[lastIndex].key;
    } else {
      activeKey.value = panes.value[0].key;
    }
  }
};

const onEdit = (targetKey: string | MouseEvent, action: string) => {
  if (action === 'add') {
    add();
  } else {
    remove(targetKey as string);
  }
};

</script>

<style lang="less">
.test-interface-design-main {
  .tabs {
    .ant-tabs-card {
      .ant-tabs-extra-content {
        margin-right: 160px;
      }
    }
  }
  .selection {
    position: absolute;
    top: 16px;
    right: 16px;
  }
}
</style>

<style scoped lang="less">
.test-interface-design-main {
  padding: 16px;
}

</style>
