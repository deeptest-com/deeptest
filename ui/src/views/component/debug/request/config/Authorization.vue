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
          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>帮助</template>
            <QuestionCircleOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>

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
import { QuestionCircleOutlined, DeleteOutlined, PlusOutlined } from '@ant-design/icons-vue';

import EmptyPage from "@/components/others/empty.vue";
import RequestAuthorBasic from "./author/BasicAuthor.vue"
import RequestAuthorBearerToken from "./author/BearerToken.vue"
import RequestAuthorOAuth2 from "./author/OAuth2.vue"
import RequestAuthorApiKey from "./author/ApiKey.vue"
import {AuthorizationTypes, UsedBy} from "@/utils/enum";
import {getEnumSelectItems} from "@/views/interface1/service";

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

import {Param} from "@/views/component/debug/data";
import {StateType as Debug} from "@/views/component/debug/store";
const store = useStore<{  Debug: Debug }>();

const debugData = computed<any>(() => store.state.Debug.debugData);

const authorizationTypes = getEnumSelectItems(AuthorizationTypes)

const onParamChange = (idx) => {
  console.log('onParamChange', idx)

};

</script>

<style lang="less">
.authorization-main {
  height: 100%;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }

  .author-content {
    display: flex;
    height: 100%;

    .params {
      flex: 2;

      height: calc(100% - 28px);
      overflow-y: auto;
      .param {
        padding: 2px 3px;
        border-bottom: 1px solid #d9d9d9;

        .ant-col {
          border-right: 1px solid #d9d9d9;

          .label {
            display: inline-block;
            padding: 4px;
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