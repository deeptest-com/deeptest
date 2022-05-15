<template>
  <div class="mock-oauth2-callback">
    <a-card :bordered="false">
      <template #title>
        OAuth2认证回调模拟
      </template>
      <template #extra>
      </template>

      <div class="dp-center">
        <div>AccessTokenURL: {{accessTokenURL}}</div>
        <div>ClientId: {{clientId}}</div>
        <div>ClientSecret: {{clientSecret}}</div>
        <div>Code: {{code}}</div>

        <br />

        <a-button type="primary" @click="getAccessToken">获取访问令牌</a-button>

        <br />

        <template v-if="accessToken">
          <div>AccessToken: {{accessToken}}</div>
          <div>TokenType: {{tokenType}}</div>
        </template>

      </div>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import {computed, defineComponent, ref} from "vue";
import {getUrlKey} from "@/utils/url";
import {getOAuth2AccessToken} from "@/services/mock";

const url = window.location.href

const accessTokenURL = getUrlKey('accessTokenURL', url)
const clientId = getUrlKey('clientId', url)
const clientSecret = getUrlKey('clientSecret', url)
const code = getUrlKey('code', url)
console.log(accessTokenURL, clientId, clientSecret, code)

const accessToken = ref('')
const tokenType = ref('')

const getAccessToken = () => {
  console.log('getAccessToken')

  getOAuth2AccessToken({accessTokenURL, clientId, clientSecret, code}).then(
      (jsn) => {
        console.log(jsn.data)

        accessToken.value = jsn.data.access_token
        tokenType.value = jsn.data.token_type
      }
  )
}

const onMounted = () => {
  console.log('onMounted')
}

</script>

<style lang="less" scoped>
.mock-oauth2-callback {

}
</style>