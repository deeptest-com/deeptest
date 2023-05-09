<template>
    <div class="scenario-collapse">
        <div class="scenario-basicinfo">
            <div class="scenario-name">{{ record.scenarioName }}</div>
            <div class="scenario-priority">{{ record.scenarioPriority }}</div>
            <div class="scenario-status">{{ record.scenarioStatus === 0 ? '已完成' : '未完成' }}
            </div>
            <div class="scenario-rate">
                <div class="report-progress"
                    :style="`background: linear-gradient(90deg, #04C495 ${record.scenarioProgress}%, #FF6963 0);`">
                </div>
                通过率 {{ record.scenarioProgress }}%
            </div>
            <div class="scenario-action">
                <template v-if="expandActive">
                    <span style="cursor: pointer;" @click="$event => expandActive = false">收起 &nbsp;
                        <UpOutlined />
                    </span>
                </template>
                <template v-else>
                    <span style="cursor: pointer;" @click="$event => expandActive = true">展开 &nbsp;
                        <DownOutlined />
                    </span>
                </template>
            </div>
        </div>
        <div class="scenario-expandInfo" v-if="expandActive">
            扩展信息
        </div>
    </div>
</template>
<script lang="ts" setup>
import { defineProps, ref } from 'vue';
import { UpOutlined, DownOutlined } from '@ant-design/icons-vue';

defineProps(['record']);
const expandActive = ref(false);
</script>

<style scoped lang="less">
.scenario-basicinfo {
    display: flex;
    align-items: center;
    justify-content: space-between;
    border: 1px solid #E5E5E5;
    padding: 16px;

    div {
        overflow-wrap: break-word;
    }
}

.scenario-priority {
    font-weight: bold;
}

.scenario-status {
    display: flex;
    align-items: center;

    &:before {
        content: '';
        display: block;
        width: 6px;
        height: 6px;
        border-radius: 50%;
        background-color: #04C495;
        margin-right: 10px;
    }
}

.scenario-name {
    width: 333px;
    font-weight: bold;
}

.scenario-rate {
    width: 292px;
    display: flex;
    align-items: center;
    font-family: 'PingFang SC';
    font-style: normal;
    font-weight: 400;
    font-size: 14px;
    color: rgba(0, 0, 0, 0.85);
    padding: 0;

    .report-progress {
        width: 180px;
        height: 6px;
        border-radius: 41px;
        margin-right: 16px;
    }
}</style>
  
  