<template>
    <div class="select-server">
        <a-form-item>
            <a-select
                v-model:value="currServe.id"
                :placeholder="'请选择服务'"
                :bordered="true"
                style="width: 334px"
                @focus="getServeList"
                @change="selectServe">
                <a-select-option v-for="item in serves" :key="item.id" :value="item.id">{{ item.name }}</a-select-option>
            </a-select>
        </a-form-item>
    </div>
</template>
<script setup lang="ts">
import { computed } from "vue";
import { useStore } from "vuex";
import { StateType as ServeStateType } from "@/store/serve";
import {raw} from "body-parser";

const store = useStore<{ ServeGlobal: ServeStateType }>();
const currServe = computed<any>(() => store.state.ServeGlobal.currServe);
const serves = computed<any>(() => store.state.ServeGlobal.serves);

const selectServe = (value): void => {
    store.dispatch('ServeGlobal/changeServe', value);
}

async function getServeList() {
  // 需要重新更新可选服务列表
  await store.dispatch("ServeGlobal/fetchServe");
}

</script>
<style scoped lang="less">
.select-server {
    padding-left: 20px;
    height: 100%;
    display: flex;
    align-items: center;
    :deep(.ant-row.ant-form-item) {
        margin: 0;
    }
}
</style>
