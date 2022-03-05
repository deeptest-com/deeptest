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
          <span class="curr-method">{{ modelData.method }}</span>
          <DownOutlined />
        </a-button>
      </a-dropdown>
    </div>
    <div class="url">
      <a-input v-model:value="modelData.url" class="dp-bg-light" />
    </div>
    <div class="send">
      <a-dropdown-button type="primary" trigger="click" @click="sendRequest">
        发送
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
                  <a-input @click.stop v-model:value="modelData.name" />
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
import { DownOutlined, UndoOutlined, SaveOutlined, LinkOutlined, CheckOutlined } from '@ant-design/icons-vue';
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {StateType} from "@/views/interface/store";
import {Methods} from "@/views/interface/consts";

interface RequestSenderSetupData {
  modelData: ComputedRef;
  methods: string[]

  selectMethod: (e) => void;
  sendRequest: (e) => void;
  clearAll: (e) => void;
  saveName: (e) => void;
  copyLink: (e) => void;
  saveAs: (e) => void;
  none: (e) => void;
}

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
  setup(props): RequestSenderSetupData {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const modelData = computed<any>(() => store.state.Interface.modelResult);

    const methods = Methods;

    const selectMethod = (e) => {
      console.log('selectMethod', e)
    };
    const sendRequest = (e) => {
      console.log('sendRequest', e)
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


    return {
      modelData,
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
    width: 100px;
  }
}

</style>