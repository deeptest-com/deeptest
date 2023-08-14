<template>
  <div class="log-content" v-if="detailList.length">
    <div class="log-content-content"> ===
      <a-descriptions bordered :size="'small'"  >
        <a-descriptions-item  v-for="detail in detailList" :key="detail.label" :label="detail.label">
          {{detail.value}}
        </a-descriptions-item>
      </a-descriptions>
    </div>

    <div class="log-content-btn" @click="showLogDetail">
      更多详情
      <RightOutlined/>
    </div>

    <LogContentDrawer
        :data="data"
        :visible="visible"
        @onClose="visible = false"/>
  </div>
</template>
<script setup lang="ts">
import {defineProps, h, defineEmits, computed, toRefs, ref} from 'vue';
import LogContentDrawer from './LogContentDrawer/index.vue';
import {RightOutlined, LoadingOutlined, ExclamationCircleOutlined, CheckCircleOutlined} from '@ant-design/icons-vue';
import {responseCodes} from '@/config/constant';
import {formatWithSeconds} from '@/utils/datetime';

enum StatusMap {
  'pass' = '通过',
  'expires' = '过期',
  'fail' = '失败'
}

enum ClassMap {
  'pass' = 'endpoint-success',
  'expires' = 'endpoint-expires',
  'fail' = 'endpoint-error',
  'loading' = 'endpoint-loading'
}

const props = defineProps({
  data: {
    type: Object,
    required: true
  }
});

const emits = defineEmits([]);

const detailList = computed(() => {
  const detail = JSON.parse(props.data.detail || '{}')
  const list:any = [];
  Object.entries(detail).forEach(([label,value]:any) => {
    list.push({
      label,
      value
    })
  })
  return list;
})


const logDetail:any = ref(null);
const visible = ref(false);

function showLogDetail() {
    visible.value = true;
}


</script>
<style scoped lang="less">


.log-content {
  .log-content-content {
    background-color: #fff;
    padding: 8px 16px;
  }

  .log-content-btn {
    margin-top: 8px;
    text-align: center;
    cursor: pointer;
    line-height: 22px;
    font-size: 14px;
  }

}

</style>
