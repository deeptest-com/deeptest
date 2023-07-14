<!--
 左侧菜单树
-->
<template>
  <a-menu style="width: 100%;
          height: auto;
          margin-bottom: 24px;
          padding-top: 8px;"
          class="docs-menu"
          :selectedKeys="selectedKeys"
          mode="inline">
    <a-menu-item v-for="item in items"
                 :key="`${item?.method || 'serve'}-${item.id}`"
                 @click="select(item)"
                 :class="{'hide': item.method && openKeysMap[item.serveId]}">
      <!-- ::::服务信息 -->
      <div class="menus-title" v-if="!item?.method">
        <div class="icon" @click="(event) => {switchExpand(item,event)}">
          <RightOutlined class="expand-icon" :class="!openKeysMap[item.id] ? 'open' : ''"/>
        </div>
        <div class="left"><strong>{{ item?.name }}</strong></div>
      </div>
      <!-- ::::该服务下的所有接口列表 -->
      <div class="menus-item" v-if="item?.method" :ref="(el) => {
          menuItemRefs[`${item.method}-${item.id}`] = el;
      }">
        <div class="left">
          <a-tag  class="method-tag" :color="getMethodColor(item.method)">{{ item.method }}</a-tag>
          <span :title="item.name">{{ item.name }}</span>
        </div>
      </div>
    </a-menu-item>
  </a-menu>
</template>

<script lang="ts" setup>
import {defineEmits, defineProps, ref, watch,} from 'vue';

import {RightOutlined} from '@ant-design/icons-vue';
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
  selectedKeys: {
    required: true,
    type: Array,
  }
})
const emit = defineEmits(['select']);

const items: any = ref([]);

watch(() => {
      return props.serviceList
    }, (newVal) => {
      items.value = newVal;
      newVal.forEach((item: any) => {
        if (!item.method) {
          openKeysMap.value[item.id] = false;
        }
      })

    }, {immediate: true}
)


function switchExpand(item, e) {
  e.stopPropagation();
  openKeysMap.value[item.id] = !openKeysMap.value[item.id];
}


function select(item) {
  emit('select', item);
}

const menuItemRefs = ref({})

watch(() => {
  return props.selectedKeys
},(newVal) => {
  //  选中的接口文档，滚动相应的位置
  if(menuItemRefs.value?.[`${newVal[0]}`]){
    menuItemRefs.value[`${newVal[0]}`].scrollIntoView({
      behavior: 'smooth',
      block: 'nearest',
      inline: 'nearest',
    });
  }
},{immediate: false});

</script>
<style lang="less" scoped>
.docs-menu {
  height: 100%;
  //margin-left: 1px;
  position: relative;
  border-right:none;
  //border-left: 1px solid #f0f0f0;
  //:deep(.ant-menu-submenu-title) {
  //  position: relative;
  //}

  &:before{
    content: '';
    position: absolute;
    top:0;
    right: 0;
    height: 100%;
    z-index: 99;
    background-color: #f0f0f0;
    width: 1px;
  }
  :deep(.hide) {
    display: none !important;
  }

  :deep(.ant-menu-item) {
    padding: 0 6px 0 8px!important;
    //padding: 0!important;
    margin: 0!important;
    border-radius: 4px;
    height: 36px;
    line-height: 36px;
    left: -6px;
  }
  :deep(.ant-menu-item-selected:after) {
    //right: 300px!important;
    display: none;
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
  height: 36px;
  line-height: 36px;

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

.method-tag{
  transform: scale(0.85);
  margin-right: 3px;
}

</style>
