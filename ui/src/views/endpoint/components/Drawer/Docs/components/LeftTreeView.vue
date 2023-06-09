<!--
 左侧菜单树
-->
<template>
  <a-menu style="width: 100%;
          padding-top: 8px;"
          class="docs-menu"
          mode="inline">
    <a-menu-item v-for="item in items"
                 :key="item.id"
                 @click="select(item)"
                 :class="{'hide': item.endpointInfo && item.serveInfo && openKeysMap[item.serveId]}">
      <!-- ::::服务信息 -->
      <div class="menus-title" v-if="item.endpointList">
        <div class="icon" @click="(event) => {switchExpand(item,event)}">
          <RightOutlined class="expand-icon" :class="!openKeysMap[item.id] ? 'open' : ''"/>
        </div>
        <div class="left"><strong>{{ item?.name }}</strong></div>
      </div>
      <!-- ::::该服务下的所有接口列表 -->
      <div class="menus-item" v-if="item.endpointInfo && item.serveInfo">
        <div class="left">
          <a-tag :color="getMethodColor(item.method)">{{ item.method }}</a-tag>
          {{ item.name }}
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
  computed,
  watch,
} from 'vue';

import {DownOutlined, RightOutlined} from '@ant-design/icons-vue';
import {requestMethodOpts} from '@/config/constant';

const openKeysMap = ref<any>({});

function getMethodColor(method: any) {
  const item: any = requestMethodOpts.find((item: any) => {
    return item.value === method;
  });
  return item.color || '#04C495';
}


const props = defineProps({
  serviceList: {
    required: true,
    type: Object,
  },
})
const emit = defineEmits(['select']);

const items: any = ref([]);
watch(() => {
      return props.serviceList
    }, (newVal) => {
      items.value = newVal;
      newVal.forEach((item: any) => {
        if (item.endpointList) {
          openKeysMap.value[item.id] = false;
        }
      })
      console.log(items.value)
    }, {immediate: true}
)


const activeKey = ref([]);


function switchExpand(item, e) {
  e.stopPropagation();
  openKeysMap.value[item.id] = !openKeysMap.value[item.id];

}


function select(item) {
  console.log(item)
  emit('select', item);
}


</script>
<style lang="less" scoped>
.docs-menu {
  position: relative;
  //:deep(.ant-menu-submenu-title) {
  //  position: relative;
  //}
  :deep(.hide) {
    display: none !important;
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
    flex: 1;
    margin-right: 8px;
    //  添加省略号
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
}

.menus-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-left: 20px;

  .left {
    flex: 1;
    display: inline-block;
    margin-right: 8px;
    //margin-left: 16px;
    //  添加省略号
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .right {
    display: flex;
    align-items: center;
    width: 60px;
    justify-content: flex-end;

    .ant-tag {
      margin-left: 4px;
      margin-right: 4px;
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
