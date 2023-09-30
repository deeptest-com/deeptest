<template>
  <div class="authorization-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">
          <span>
            授权类型
          </span>

          <a-select
              ref="authorizationType"
              v-model:value="debugData.authorizationType"
              :options="authorizationTypes"
              size="small"
              :dropdownMatchSelectWidth="false"
              :bordered="false"
          >
          </a-select>
        </a-col>
        <a-col flex="80px" class="dp-right">
          <Tips section="why_interface" title="授权头将会在你发送请求时自动生成。" />

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>全部清除</template>
            <DeleteOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>新增</template>
            <PlusOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>
        </a-col>
      </a-row>
    </div>

    <template v-if="!debugData.authorizationType" class="none">
      <EmptyPage desc="无授权信息"></EmptyPage>
    </template>

    <RequestAuthorBasic v-if="debugData.authorizationType === 'basicAuth'"></RequestAuthorBasic>
    <RequestAuthorBearerToken v-if="debugData.authorizationType === 'bearerToken'"></RequestAuthorBearerToken>
    <RequestAuthorApiKey v-if="debugData.authorizationType === 'apiKey'"></RequestAuthorApiKey>
<!--    <RequestAuthorOAuth2 v-if="debugData.authorizationType === 'oAuth2'"></RequestAuthorOAuth2>-->
  </div>
</template>

<script setup lang="ts">
import {computed, inject} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { DeleteOutlined, PlusOutlined, ArrowRightOutlined } from '@ant-design/icons-vue';

import EmptyPage from "@/components/others/empty.vue";
import Tips from "@/components/Tips/index.vue";
import RequestAuthorBasic from "./author/BasicAuthor.vue"
import RequestAuthorBearerToken from "./author/BearerToken.vue"
import RequestAuthorApiKey from "./author/ApiKey.vue"
import {AuthorizationTypes, UsedBy} from "@/utils/enum";
import {StateType as Debug} from "@/views/component/debug/store";
import {getEnumSelectItems} from "@/utils/comm";

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();
const store = useStore<{  Debug: Debug }>();

const debugData = computed<any>(() => store.state.Debug.debugData);

const authorizationTypes = getEnumSelectItems(AuthorizationTypes)

const onParamChange = (idx) => {
  console.log('onParamChange', idx)

};

</script>

<style lang="less">
.authorization-main {
  overflow-y: scroll;
  height: 100%;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }

  .author-content {
    display: flex;
    height: calc(100% - 30px);

    .params {
      flex: 2;

      height: calc(100% - 28px);
      overflow-y: auto;
      .param {
        border-bottom: 1px solid #d9d9d9;
        height: 32px;

        .ant-col {
          border-right: 1px solid #d9d9d9;
          display: flex;
          align-items: center;
          justify-content: flex-start;

          &:last-child {
            border-right: 0;
          }

          input {
            margin-top: 1px;
          }
        }
      }
    }

    .tips {
      flex: 1;
      padding: 10px;
    }
  }

}

</style>
