<template>
  <div class="author-basic-main author-content">
    <div class="params">
      <a-row class="param">
        <a-col flex="1">
          <a-button class="dp-bg-light">
            <span class="curr-method">生成令牌</span>
          </a-button>
        </a-col>
      </a-row>

      <a-row class="param">
        <a-col flex="160px">
          <span class="label">Token</span>
        </a-col>
        <a-col flex="1">
          <a-input v-model:value="interfaceData.oAuth20.key"
                   placeholder="Key" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>

      <a-row class="param">
        <a-col flex="160px">
          <span class="label">OpenID Connect Discovery URL</span>
        </a-col>
        <a-col flex="1">
          <a-input v-model:value="interfaceData.oAuth20.oidcDiscoveryURL" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>

      <a-row class="param">
        <a-col flex="160px">
          <span class="label">Authentication URL</span>
        </a-col>
        <a-col flex="1">
          <a-input v-model:value="interfaceData.oAuth20.authURL" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>
      <a-row class="param">
        <a-col flex="160px">
          <span class="label">Access Token URL</span>
        </a-col>
        <a-col flex="1">
          <a-input v-model:value="interfaceData.oAuth20.accessTokenURL" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>
      <a-row class="param">
        <a-col flex="160px">
          <span class="label">Client ID</span>
        </a-col>
        <a-col flex="1">
          <a-input v-model:value="interfaceData.oAuth20.clientID" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>
      <a-row class="param">
        <a-col flex="160px">
          <span class="label">Scope</span>
        </a-col>
        <a-col flex="1">
          <a-input v-model:value="interfaceData.basicAuth.scope" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>
    </div>
    <div class="tips">
      <div class="dp-light">授权头将会在你发送请求时自动生成。</div>
      <div class="dp-link-primary">了解更多 <ArrowRightOutlined /></div>
    </div>
  </div>
</template>

<script lang="ts">
import {computed, defineComponent, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {ArrowRightOutlined, DeleteOutlined, PlusOutlined, QuestionCircleOutlined} from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import {Interface} from "@/views/interface/data";

export default defineComponent({
  name: 'RequestAuthorOAuth2',
  components: {
    ArrowRightOutlined,
  },
  setup(props) {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

    return {
      interfaceData,
    }
  }
})

</script>

<style lang="less" scoped>
.author-basic-main {
}

</style>