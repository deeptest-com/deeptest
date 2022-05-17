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

          <br />

          <a-button type="primary" @click="useAccessToken">使用访问令牌</a-button>
        </template>

        <template v-if="error">
          <div>Error: {{error}}</div>
          <div>ErrorDescription: {{errorDescription}}</div>
          <div>ErrorUri: {{errorUri}}</div>
        </template>

      </div>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import {computed, defineComponent, ref} from "vue";
import {getUrlKey} from "@/utils/url";
import {getOAuth2AccessToken, useOAuth2AccessToken} from "@/services/mock";
import {useStore} from "vuex";
import {StateType as ProjectStateType} from "@/store/project";

const projectStore = useStore<{ ProjectData: ProjectStateType }>();
const currProject = computed<any>(() => projectStore.state.ProjectData.currProject);

const url = window.location.href
const interfaceId = getUrlKey('interfaceId', url)
const name = getUrlKey('name', url)
const accessTokenURL = getUrlKey('accessTokenURL', url)
const clientId = getUrlKey('clientId', url)
const clientSecret = getUrlKey('clientSecret', url)
const code = getUrlKey('code', url)

const accessToken = ref('')
const tokenType = ref('')

const error = ref('')
const errorDescription = ref('')
const errorUri = ref('')

const getAccessToken = () => {
  console.log('getAccessToken')

  getOAuth2AccessToken({accessTokenURL, clientId, clientSecret, code}).then(
      (jsn) => {
        console.log(jsn.data)

        accessToken.value = jsn.data.access_token
        tokenType.value = jsn.data.token_type

        error.value = jsn.data.error
        errorDescription.value = jsn.data.error_description
        errorUri.value = jsn.data.error_uri
      }
  )
}

const useAccessToken = () => {
  console.log('useOAuth2AccessToken')
  useOAuth2AccessToken(name, accessToken.value, tokenType.value, interfaceId, currProject.value.id)
}

const onMounted = () => {
  console.log('onMounted')
}

</script>

<style lang="less" scoped>
.mock-oauth2-callback {

}
</style>