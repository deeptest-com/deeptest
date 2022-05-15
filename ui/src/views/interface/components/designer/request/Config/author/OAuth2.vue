<template>
  <div class="author-basic-main author-content">
    <div class="params">
      <a-row class="param">
        <a-col flex="160px">
          <span class="label">Access Token</span>
        </a-col>
        <a-col flex="1">
          <a-select
              v-model:value="interfaceData.oauth20.accessToken"
              :options="accessTokens"
              size="small"
              :bordered="false"
              style="width: calc(100% - 10px)"
          >
          </a-select>
        </a-col>
      </a-row>
      <a-row class="param">
        <a-col flex="160px"></a-col>
        <a-col flex="1">
          <a-input v-model:value="interfaceData.oauth20.accessToken" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>

      <a-row class="param">
        <a-col flex="160px">
          <span class="label">Header Prefix</span>
        </a-col>
        <a-col flex="1">
          <a-input v-model:value="interfaceData.oauth20.headerPrefix" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>

      <br />

      <a-row class="param">
        <a-col flex="1" class="dp-right">
          <a-button @click="generateToken" type="link" size="small">
            <span class="curr-method">生成新令牌</span>
          </a-button>
        </a-col>
      </a-row>

      <a-row class="param">
        <a-col flex="160px">
          <span class="label">Token Name</span>
        </a-col>
        <a-col flex="1">
          <a-input v-model:value="interfaceData.oauth20.name" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>

      <a-row class="param">
        <a-col flex="160px">
          <span class="label">Grant Type</span>
        </a-col>
        <a-col flex="1">
          <a-select
              v-model:value="interfaceData.oauth20.grantType"
              :options="oauth2GrantTypes"
              size="small"
              :bordered="false"
              style="width: calc(100% - 10px)"
          >
          </a-select>
        </a-col>
      </a-row>

      <a-row class="param">
        <a-col flex="160px">
          <span class="label">Callback URL</span>
        </a-col>
        <a-col flex="1">
          <a-input v-model:value="interfaceData.oauth20.callbackUrl" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>

      <a-row class="param">
        <a-col flex="160px">
          <span class="label">Authentication URL</span>
        </a-col>
        <a-col flex="1">
          <a-input v-model:value="interfaceData.oauth20.authURL" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>
      <a-row class="param">
        <a-col flex="160px">
          <span class="label">Access Token URL</span>
        </a-col>
        <a-col flex="1">
          <a-input v-model:value="interfaceData.oauth20.accessTokenURL" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>

      <a-row class="param">
        <a-col flex="160px">
          <span class="label">Client ID</span>
        </a-col>
        <a-col flex="1">
          <a-input v-model:value="interfaceData.oauth20.clientID" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>
      <a-row class="param">
        <a-col flex="160px">
          <span class="label">Client Secret</span>
        </a-col>
        <a-col flex="1">
          <a-input v-model:value="interfaceData.oauth20.clientSecret" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>

      <a-row class="param">
        <a-col flex="160px">
          <span class="label">Scope</span>
        </a-col>
        <a-col flex="1">
          <a-input v-model:value="interfaceData.oauth20.scope" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>

      <a-row class="param">
        <a-col flex="160px">
          <span class="label">State</span>
        </a-col>
        <a-col flex="1">
          <a-input v-model:value="interfaceData.oauth20.state" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>

      <a-row class="param">
        <a-col flex="160px">
          <span class="label">Client Authentication</span>
        </a-col>
        <a-col flex="1">
          <a-select
              v-model:value="interfaceData.oauth20.clientAuthentication"
              :options="oauth2ClientAuthWays"
              size="small"
              :bordered="false"
              style="width: calc(100% - 10px)"
          >
          </a-select>
        </a-col>
      </a-row>

    </div>
    <div class="tips">
      <div class="dp-light">授权头将会在你发送请求时自动生成。</div>
      <div class="dp-link-primary">了解更多 <ArrowRightOutlined /></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, defineComponent, onBeforeUnmount, onMounted, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {ArrowRightOutlined, DeleteOutlined, PlusOutlined, QuestionCircleOutlined} from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import {Interface} from "@/views/interface/data";
import {genOAuth2AccessToken, getEnumSelectItems} from "@/views/interface/service";
import {AuthorizationTypes, OAuth2ClientAuthenticationWay, OAuth2GrantTypes} from "@/views/interface/consts";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import {WsMsg} from "@/types/data";

const {t} = useI18n();
const store = useStore<{ Interface: StateType }>();
const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

const oauth2GrantTypes = getEnumSelectItems(OAuth2GrantTypes)
const oauth2ClientAuthWays = getEnumSelectItems(OAuth2ClientAuthenticationWay)

const accessTokens = ref([])

const generateToken = () => {
  console.log('generateToken', interfaceData.value.oauth20)

  genOAuth2AccessToken(interfaceData.value.oauth20).then((result) => {
    console.log(result)
    if (result.code === 0) {
      window.open(result.data.url, '_blank');
    }
  })
}

onMounted(() => {
  console.log('onMounted')
  bus.on(settings.eventWebSocketMsg, OnWebSocketMsg);
})
onBeforeUnmount( () => {
  bus.off(settings.eventWebSocketMsg, OnWebSocketMsg);
})

const OnWebSocketMsg = (data: any) => {
  console.log('OnWebSocketMsg in OAuth2', data.msg)

  const jsn = JSON.parse(data.msg) as WsMsg
  console.log(jsn)
  if (jsn.access_token) {
    interfaceData.value.oauth20.accessToken = jsn.access_token

    if (jsn.token_type === 'bearer') {
      interfaceData.value.oauth20.headerPrefix = 'Bearer'
    }
  }
}

</script>

<style lang="less" scoped>
.author-basic-main {
}

</style>