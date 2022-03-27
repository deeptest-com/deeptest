<template>
  <div class="sender-main">
    <div class="methods">
      <a-dropdown trigger="click">
        <template #overlay>
          <a-menu @click="selectMethod">
            <a-menu-item v-for="(item) in methods" :key="item">{{ item }}</a-menu-item>
          </a-menu>
        </template>
        <a-button class="dp-bg-light">
          <span class="curr-method">{{ requestData.method }}</span>
          <DownOutlined />
        </a-button>
      </a-dropdown>
    </div>
    <div class="url">
      <a-input v-model:value="requestData.url" class="dp-bg-light" />
    </div>
    <div class="send">
      <a-dropdown-button type="primary" trigger="click" @click="sendRequest">
        <span>发送</span>

        <template #overlay>
          <a-menu>
            <a-menu-item @click="clearAll" key="clearAll">
              <UndoOutlined />
              全部清除
            </a-menu-item>
          </a-menu>
        </template>
        <template #icon><DownOutlined /></template>
      </a-dropdown-button>
    </div>
    <div class="save">
      <a-dropdown-button trigger="click" @click="sendRequest" class="dp-bg-light">
        <SaveOutlined />
        保存
        <template #overlay>
          <a-menu>
            <a-menu-item @click.prevent="none" key="copyLink" class="edit-name">
              <div class="dp-edit-interface-name">
                <div class="left">
                  <a-input @click.stop v-model:value="requestData.name" />
                </div>
                <div class="right">
                  <CheckOutlined @click.stop="saveName" class="save-button" />
                </div>
              </div>
            </a-menu-item>

            <a-menu-item @click="copyLink" key="copyLink">
              <LinkOutlined />
              复制链接
            </a-menu-item>

            <a-menu-item @click="saveAs" key="saveAs">
              <LinkOutlined />
              另存为
            </a-menu-item>
          </a-menu>
        </template>
        <template #icon><DownOutlined /></template>
      </a-dropdown-button>
    </div>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, PropType} from "vue";
import { notification, message } from 'ant-design-vue';
import { DownOutlined, UndoOutlined, SaveOutlined, LinkOutlined, CheckOutlined } from '@ant-design/icons-vue';
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {StateType} from "@/views/interface/store";
import {Methods} from "@/views/interface/consts";
import {regxUrl} from "@/utils/validation";

export default defineComponent({
  name: 'RequestSender',
  props: {
    onSend: {
      type: Function as PropType<() => void>,
      required: true
    }
  },
  components: {
    DownOutlined, UndoOutlined, SaveOutlined, LinkOutlined, CheckOutlined,
  },
  setup(props) {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const requestData = computed<any>(() => store.state.Interface.requestData);

    const methods = Methods;

    const selectMethod = (val) => {
      console.log('selectMethod', val.key)
      requestData.value.method = val.key
    };
    const sendRequest = (e) => {
      requestData.value.params = requestData.value.params.filter((param) => {
        return !param.disabled && !!param.name
      })
      requestData.value.headers = requestData.value.headers.filter((param) => {
        return !param.disabled && !!param.name
      })

      console.log('sendRequest', requestData.value)
      validateInfo()
      props.onSend()
    };
    const clearAll = (e) => {
      console.log('clearAll', e)
    };
    const saveName = (e) => {
      console.log('saveName', e)
      e.preventDefault();
    };
    const copyLink = (e) => {
      console.log('copyLink', e)
    };
    const saveAs = (e) => {
      console.log('saveAs', e)
    };
    const none = (e) => {
      console.log('none', e)
      e.preventDefault()
    };

    const validateInfo = () => {
      let msg = ''
      if (!requestData.value.url) {
        msg = '请求地址不能为空'
      } else if (!regxUrl.test(requestData.value.url)) {
        msg = '请求地址格式错误'
      }

      if (msg) {
        notification.warn({
          message: msg,
          placement: 'bottomLeft'
        });
      }
    };

    return {
      requestData,
      methods,
      selectMethod,
      sendRequest,
      clearAll,
      saveName,
      copyLink,
      saveAs,
      none,
    }
  }
})

</script>

<style lang="less">
.dp-edit-interface-name {
  display: flex;
  .left {
    flex: 1;
  }
  .right {
    width: 30px;
    padding-left: 10px;
    .save-button {
      vertical-align: -5px
    }
  }
}
</style>

<style lang="less" scoped>
.sender-main {
  display: flex;
  padding: 0 18px 0 0px;
  .methods {
    width: 100px;
    .curr-method {
      width: 55px;
    }
  }
  .url {
    flex: 1;
  }
  .send {
    width: 100px;
  }
  .save {
    width: 93px;
  }
}

</style>