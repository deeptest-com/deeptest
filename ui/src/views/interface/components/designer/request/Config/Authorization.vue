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
              v-model:value="interfaceData.authorizationType"
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

      <template v-if="!interfaceData.authorizationType" class="none">
        <EmptyPage desc="无授权信息"></EmptyPage>
      </template>

      <div v-if="interfaceData.authorizationType === 'basicAuth'" class="content">
        <div class="params">
          <a-row class="param">
            <a-col flex="1">
              <a-input v-model:value="interfaceData.basicAuth.username" placeholder="用户名" class="dp-bg-input-transparent" />
            </a-col>
          </a-row>
          <a-row class="param">
            <a-col flex="1">
              <a-input v-model:value="interfaceData.basicAuth.password" placeholder="密码" class="dp-bg-input-transparent" />
            </a-col>
          </a-row>
        </div>
        <div class="tips">
          <div class="dp-light">授权头将会在你发送请求时自动生成。</div>
          <div class="dp-link-primary">了解更多 <ArrowRightOutlined /></div>
        </div>
      </div>

      <div v-if="interfaceData.authorizationType === 'bearerToken'" class="content">
        <div class="params">
          <a-row class="param">
            <a-col flex="1">
              <a-input v-model:value="interfaceData.bearerToken.username" placeholder="用户名" class="dp-bg-input-transparent" />
            </a-col>
          </a-row>
        </div>
        <div class="tips">
          <div class="dp-light">授权头将会在你发送请求时自动生成。</div>
          <div class="dp-link-primary">了解更多 <ArrowRightOutlined /></div>
        </div>
      </div>

      <div v-if="interfaceData.authorizationType === 'oAuth20'" class="content">
        <div class="params">
          <a-row class="param">
            <a-col flex="1">
              <a-input v-model:value="interfaceData.oAuth20.key"
                       placeholder="Key" class="dp-bg-input-transparent" />
            </a-col>
          </a-row>
          <a-row class="param">
            <a-col flex="1">
              <a-input v-model:value="interfaceData.oAuth20.oidcDiscoveryURL"
                       placeholder="OpenID Connect Discovery URL" class="dp-bg-input-transparent" />
            </a-col>
          </a-row>
          <a-row class="param">
            <a-col flex="1">
              <a-input v-model:value="interfaceData.oAuth20.authURL"
                       placeholder="Authentication URL" class="dp-bg-input-transparent" />
            </a-col>
          </a-row>
          <a-row class="param">
            <a-col flex="1">
              <a-input v-model:value="interfaceData.oAuth20.accessTokenURL"
                       placeholder="Access Token URL" class="dp-bg-input-transparent" />
            </a-col>
          </a-row>
          <a-row class="param">
            <a-col flex="1">
              <a-input v-model:value="interfaceData.oAuth20.clientID"
                       placeholder="Client ID" class="dp-bg-input-transparent" />
            </a-col>
          </a-row>
          <a-row class="param">
            <a-col flex="1">
              <a-input v-model:value="interfaceData.basicAuth.scope"
                       placeholder="Scope" class="dp-bg-input-transparent" />
            </a-col>
          </a-row>
          <a-row class="param">
            <a-col flex="1">
              <a-button class="dp-bg-light">
                <span class="curr-method">生成令牌</span>
              </a-button>
            </a-col>
          </a-row>

        </div>
        <div class="tips">
          <div class="dp-light">授权头将会在你发送请求时自动生成。</div>
          <div class="dp-link-primary">了解更多 <ArrowRightOutlined /></div>
        </div>
      </div>

      <div v-if="interfaceData.authorizationType === 'apiKey'" class="content">
        <div class="params">
          <a-row class="param">
            <a-col flex="1">
              <a-input v-model:value="interfaceData.basicAuth.username" placeholder="用户名" class="dp-bg-input-transparent" />
            </a-col>
          </a-row>
          <a-row class="param">
            <a-col flex="1">
              <a-input v-model:value="interfaceData.basicAuth.password" placeholder="取值" class="dp-bg-input-transparent" />
            </a-col>
          </a-row>
          <a-row class="param">
            <a-col flex="80px">传递方式</a-col>
            <a-col flex="1">
              <a-select
                  v-model:value="interfaceData.apiKey.transferMode"
                  size="small"
                  :dropdownMatchSelectWidth="false"
                  :bordered="false"
              >
                <a-select-option value="headers">请求头</a-select-option>
                <a-select-option value="queryParams">查询参数</a-select-option>
              </a-select>
            </a-col>
          </a-row>
        </div>
        <div class="tips">
          <div class="dp-light">授权头将会在你发送请求时自动生成。</div>
          <div class="dp-link-primary">了解更多 <ArrowRightOutlined /></div>
        </div>
      </div>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined, DeleteOutlined, PlusOutlined, ArrowRightOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import {Interface} from "@/views/interface/data";

import EmptyPage from "@/components/others/empty.vue";

export default defineComponent({
  name: 'RequestAuthorization',
  components: {
    EmptyPage,
    QuestionCircleOutlined, DeleteOutlined, PlusOutlined, ArrowRightOutlined,
  },
  setup(props) {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

    const authorizationTypes = ref([
      {value: '', label: 'None'},
      {value: 'basicAuth', label: 'Basic Auth'},
      {value: 'bearerToken', label: 'Bearer Token'},
      {value: 'oAuth20', label: 'OAuth 2.0'},
      {value: 'apiKey', label: 'API Key'},
    ])

    const onParamChange = (idx) => {
      console.log('onParamChange', idx)

    };

    return {
      interfaceData,
      authorizationTypes,
      onParamChange,
    }
  }
})

</script>

<style lang="less" scoped>
.authorization-main {
  height: 100%;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }

  .content {
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