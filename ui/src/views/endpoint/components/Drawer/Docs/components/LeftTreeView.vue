<!--
 项目基本信息
-->
<template>
  <a-menu style="width: 100%;
          padding-top: 8px;"
          class="docs-menu"
          mode="inline">

    <a-menu-item v-for="item in items"
                 :key="item.id"
                 
                 :class="{'hide': item.interfaces && openKeysMap[item.serveId]}">
      <!-- ::::服务信息 -->
      <div class="menus-title" v-if="item.endpointList">
        <div class="icon" @click="(event) => {switchExpand(item,event)}">
          <RightOutlined class="expand-icon" :class="!openKeysMap[item.id] ? 'open' : ''"/>
        </div>
        <div class="left">{{ item.name }}</div>
      </div>
      <!-- ::::该服务下的所有接口列表 -->
      <div class="menus-item" v-if="item.interfaces">
        <div class="left">
          {{ item.title }}
        </div>
        <div class="right">
          <a-tag color="pink">post</a-tag>
          <a-tag color="red">get</a-tag>
          <a-tag color="orange">GET</a-tag>
        </div>
      </div>
    </a-menu-item>

  </a-menu>

</template>

<script lang="ts" setup>
import {
  ref,
  defineProps,
  defineEmits,
  computed, watch,
} from 'vue';
import {MailOutlined, QqOutlined, AppstoreOutlined, SettingOutlined} from '@ant-design/icons-vue';

import {DownOutlined, RightOutlined} from '@ant-design/icons-vue';

const openKeysMap = ref<any>({});


const handleClick = (e: Event) => {
  console.log('click', e);
};

const titleClick = (e: Event) => {
  console.log('titleClick', e);
};

const props = defineProps({
  serviceList: {
    required: true,
    type: Object,
  },
})

const items: any = ref([]);
watch(() => {
      return props.serviceList
    }, (newVal) => {
      items.value = newVal;
      newVal.forEach((item: any) => {
        if(item.endpointList){
          openKeysMap.value[item.id] = false;
        }
      })
  console.log(items.value)
    }, {immediate: true}
)


const activeKey = ref([]);

const emit = defineEmits(['ok', 'close', 'refreshList']);

function switchExpand(item, e) {
  e.stopPropagation();
  openKeysMap.value[item.id] = !openKeysMap.value[item.id];
  console.log('openKeysMap', openKeysMap.value)
}


</script>
<style lang="less" scoped>
.docs-menu {
  position: relative;
  //:deep(.ant-menu-submenu-title) {
  //  position: relative;
  //}
  :deep(.hide) {
    display: none!important;
  }
  :deep(.ant-menu-item) {
    padding: 0 6px !important;
  }
}

.menus-title {
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: relative;
  left: -4px;

  .icon {
    width: 24px;
    justify-content: center;
    height: 16px;
    display: flex;
    align-items: center;
    cursor: pointer;
  }

  .left {
    display: flex;
    flex: 1;
    align-items: center;
  }
}

.menus-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-left: 20px;

  .left {
    display: flex;
    align-items: center;
  }

  .right {
    display: flex;
    align-items: center;

    .ant-tag {
      margin-left: 4px;
      margin-right: 0px;
    }
  }
}

.expand-icon {
  font-size: 12px;
  transition: transform 0.3s;

  &.open {
    transform: rotate(90deg);
  }
}


</style>
