<!--
 项目基本信息
-->
<template>
  <div class="docs-header">
    <div class="logo">
      <ReadOutlined style="font-size: 20px"/>
      <span class="logo-name">
        {{ `${title} - 接口文档` }}
      </span>
    </div>
    <a-popover
        :autoAdjustOverflow="false"
        :overlayClassName="'deeptest-docs-search-popover'"
        placement="bottomLeft"
        :arrowPointAtCenter="false"
        :visible="visible"
        :overlayStyle="{}">
      <template #content>
        <div class="select-content">
          <a-list item-layout="horizontal" :data-source="data" v-if="data?.length > 0">
            <template #renderItem="{ item }">
              <a-list-item @click="selectItem(item)" class="list-item">
                <a-list-item-meta>
                  <template #title>
                    <span class="title" :title="item.name">
                      <a-tag class="method-tag" v-if="item.method" :color="getMethodColor(item.method)">{{
                          item.method
                        }}</a-tag>
                      {{ item.name }}
                    </span>
                  </template>
                  <template #description>
                    <a class="description" :title="item.url || item.description" href="javascript:void (0)"
                       v-if="item.description || item.url">{{ item.url || item.description }}</a>
                  </template>
                  <template #avatar>
                    <CloudServerOutlined v-if="!item.method" style="font-size: 20px;margin-top: 2px;"/>
                    <ReadOutlined v-else style="font-size: 20px;margin-top: 2px;"/>
                  </template>
                </a-list-item-meta>
              </a-list-item>
            </template>
          </a-list>
          <a-empty v-else :image="Empty.PRESENTED_IMAGE_SIMPLE" :description="'请输入合适的关键词搜索文档'"/>
        </div>
      </template>
      <div class="search">
        <span class="left-divider"/>
        <SearchOutlined class="icon"/>
        <a-input class="search-input"
                 @focus="focus"
                 @blur="blur"
                 v-model:value="keywords"
                 ref="searchInputRef"
                 placeholder="输入关键词搜索文档..."/>
        <span class="search-shortcut" v-show="!isFocus">{{ shortCutText }}</span>
      </div>
    </a-popover>
    <div class="space"/>
    <div class="action">
      <a-dropdown class="version-info" style="width: 100px;" placement="bottomCenter">
        <a-button :size="'small'">
          文档版本：Latest
          <DownOutlined/>
        </a-button>
        <template #overlay>
          <a-menu>
            <a-menu-item>
              <span class="version-text">v1.0.1</span>
            </a-menu-item>
            <a-menu-item>
              <span class="version-text">v1.0.1</span>
            </a-menu-item>
            <a-menu-item>
              <span class="version-text">v1.2.1</span>
            </a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>

      <a-dropdown class="version-info" style="width: 100px;" placement="bottomLeft">
        <a-button :size="'small'" type="text">
          <template #icon>
            <ShareAltOutlined class="action-item"/>
          </template>
          分享
        </a-button>
        <template #overlay>
          <a-menu>
            <a-menu-item>
              <span class="version-text">分享文档</span>
            </a-menu-item>
            <a-menu-item>
              <span class="version-text">关闭分享</span>
            </a-menu-item>
            <a-menu-item>
              <span class="version-text">复制链接</span>
            </a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>

      <a-tooltip placement="bottom">
        <template #title>全屏</template>
        <a-button type="text" class="share-btn">
          <FullscreenOutlined style="font-size: 14px"/>
        </a-button>
      </a-tooltip>

    </div>
  </div>
</template>

<script lang="ts" setup>
import {
  ref,
  defineProps,
  defineEmits,
  computed, watch,
} from 'vue';
import { Empty } from 'ant-design-vue';
import {
  DownOutlined,
  RightOutlined,
  SearchOutlined,
  ShareAltOutlined,
  CloudServerOutlined,
  ReadOutlined,
  FullscreenOutlined,
  FullscreenExitOutlined
} from '@ant-design/icons-vue';
import {useMagicKeys} from '@vueuse/core'
import {getCodeColor, getMethodColor} from "../hooks/index"
import debounce from "lodash.debounce";

const searchInputRef: any = ref(null);
const isMac = navigator.platform.toUpperCase().indexOf('MAC') >= 0;

const shortCutText = ref(isMac ? '⌘ K' : 'Ctrl K');


const props = defineProps({
  items: {
    required: true,
    type: Object,
  },
  data: {
    required: true,
    type: Object,
  }
})
const data: any = ref([]);

const emit = defineEmits(['select']);

const expand = ref(true);
const keys = useMagicKeys()
const CtrlK = keys['Ctrl+K'];
const cmdK = keys['Command+K'];

function switchExpand() {
  expand.value = !expand.value;
}

const title = computed(() => {
  return props.data?.[0]?.value
})

function selectItem(item) {
  emit('select', item?.value);
  keywords.value = null;
}

const isFocus = ref(false);
const keywords = ref(null);

function focus() {
  isFocus.value = true;
}

const visible = computed(() => {
  // console.log(searchInputRef?.value?.isFocused())
  return keywords.value || isFocus.value;
})

function blur() {
  isFocus.value = false;
}


watch(CtrlK, (v) => {
  if (!isMac) {
    searchInputRef.value.focus();
  }
})

watch(cmdK, (v) => {
  if (isMac) {
    searchInputRef.value.focus();
  }
})


function keywordsChange(newVal) {
  if (newVal && props?.items?.length) {
    let lists: any = [];
    const keyword = newVal.toLowerCase();
    props?.items.forEach((item) => {
      const c1 = item.name && item.name.toLowerCase().includes(keyword);
      const c2 = item.method && item.method.toLowerCase().includes(keyword);
      const c3 = item.url && item.url.toLowerCase().includes(keyword);
      const c4 = item.description && item.description.toLowerCase().includes(keyword);
      if (c1 || c2 || c3 || c4) {
        lists.push({
          name: item.name,
          method: item.method,
          url: item.url,
          description: item.description,
          value:item
        })
      }
    })
    data.value = [...lists];
  } else {
    data.value = [];
  }

}

watch(() => {
  return keywords.value
}, (newVal: any) => {
  debounce(keywordsChange, 200)(newVal);
});


</script>
<style lang="less" scoped>
.docs-header {
  display: flex;
  justify-content: space-between;
  height: 56px;
  align-items: center;

  .logo {
    width: 294px;
    padding-left: 18px;
    display: flex;
    //justify-content: space-between;
    align-items: center;

    .logo-name {
      font-weight: bold;
      margin-left: 12px;
      font-size: 16px;
    }


  }

  .search {
    width: 280px;
    display: flex;
    align-items: center;
    position: relative;

    .search-input, .search-input:active, .search-input:hover, .search-input:focus {
      width: 180px;
      border: none !important;
      outline: none;
      box-shadow: none;
    }

    .left-divider {
      border-left: 1px solid rgba(0, 0, 0, 0.06) !important;
      width: 1px;
      padding-right: 12px;
      height: 18px;
    }

    .icon {
      opacity: 0.6;
    }

    .search-shortcut {
      opacity: 0.9;
      color: #ced4d9;
      background-color: rgba(150, 150, 150, 0.06);
      border-color: rgba(100, 100, 100, 0.2);
      border-radius: 4px;
      width: 50px;
      height: 22px;
      line-height: 22px;
      //display: inline-block;
      padding: 0px 8px;
      font-size: 12px;
      text-align: center;
      display: flex;
      align-items: center;
      justify-content: center;
    }

    //:deep(.ant-input-affix-wrapper-focused) {
    //  border: none !important;
    //  outline: none;
    //}
  }

  .space {
    flex: 1;
  }

  .action {
    width: 200px;
    margin-right: 20px;
    display: flex;
    justify-content: flex-end;
    align-items: center;

    .action-item {
      cursor: pointer;
    }
  }
}

.version-text {
  display: inline-block;
  padding: 0 8px;
}

.version-info {
  margin-right: 6px;
  margin-left: 6px;
  cursor: pointer;
}


.select-content {
  width: 450px;
  max-height: 400px;
  overflow-y: scroll;
  border-radius: 6px;
  background: #fff;
  padding: 12px 24px;
}

.list-item {
  cursor: pointer;
}

.share-btn {
  text-align: center;
  align-items: center;
  justify-content: center;
  display: flex;
  width: 32px;
  height: 32px;
}

.method-tag {
  transform: scale(0.85);
  margin-right: 3px;
}

.title, .description {
  //  超出一行加省略号
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  width: 100%;
  display: inline-block;
}
</style>
