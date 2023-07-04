<template>
  <div class="response-extractor-main">
    <div class="head">
      <a-row type="flex" class="extractor item">
        <a-col flex="50px">编号</a-col>
        <a-col flex="70px">来源</a-col>
        <a-col flex="90px">提取类型</a-col>
        <a-col flex="1">表达式</a-col>
        <a-col flex="100px" style="padding-left: 10px;">变量</a-col>
        <a-col flex="1">结果</a-col>
      </a-row>
    </div>

    <div class="items">
      <a-row v-for="(item, idx) in data" :key="idx" type="flex" class="item">
        <a-col flex="50px">{{idx + 1}}</a-col>
        <a-col flex="70px">{{ t(item.src) }}</a-col>
        <a-col flex="90px">{{ item.type ? t(item.type) : '' }}</a-col>
        <a-col flex="1"  style="width: 0; word-break: break-word;">
          <span v-if="item.src === ExtractorSrc.header">
            {{ item.key }}
          </span>
          <span v-if="item.src === ExtractorSrc.body">
            {{ item.type === ExtractorType.boundary ?
                `${item.boundaryStart}-${item.boundaryEnd}[${item.boundaryIndex}] ${item.boundaryIncluded}` :
                item.expression }}
          </span>
        </a-col>
        <a-col flex="100px" style="padding-left: 10px;">{{ item.variable }}</a-col>
        <a-col flex="1" :class="[item.result==='extractor_err'? 'dp-color-fail': '']"  style="width: 0; word-break: break-word;">
          {{item.result==='extractor_err'? t(item.result) : item.result}}
        </a-col>
      </a-row>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, inject} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {Form} from 'ant-design-vue';
import {
  CheckCircleOutlined,
  CloseCircleOutlined,
  DeleteOutlined,
  EditOutlined,
  PlusOutlined
} from '@ant-design/icons-vue';
import {ExtractorSrc, ExtractorType, UsedBy} from "@/utils/enum";
import {StateType as Debug} from "@/views/component/debug/store";

const usedBy = inject('usedBy') as UsedBy
const useForm = Form.useForm;
const {t} = useI18n();

const store = useStore<{  Debug: Debug }>();

const props = defineProps({
  data: {
    type: []
  }
})

</script>

<style lang="less">
</style>

<style lang="less" scoped>
.response-extractor-main {
  height: 100%;

  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .items {
    padding: 6px;
    height: calc(100% - 30px);
    overflow-y: auto;

    .item {
      .ant-col {
        padding: 0 3px;
        word-break: break-all;
      }
    }
  }
}
</style>