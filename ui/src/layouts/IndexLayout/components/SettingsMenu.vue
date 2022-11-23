<template>
  <a-menu
      :style="[{width: collapsed ? '55px' : '102px'}]"
      mode="vertical"
      @click="clickMenu"
  >
    <a-sub-menu key="settings-sub" popupClassName="settings-sub-menu">
      <template #icon>
        <SettingOutlined />
      </template>
      <template #title>
        <span class="title" :class="[{hide: collapsed}]">设置</span>
      </template>

<!--  <a-menu-item key="lang">
        <icon-svg type="language-outline" class="icon-svg" />
        <span>界面语言</span>
      </a-menu-item>-->
      <a-menu-item key="profile">
        <UserOutlined />
        <span>个人信息</span>
      </a-menu-item>
      <a-menu-item key="logout">
        <LogoutOutlined />
        <span>登出</span>
      </a-menu-item>
    </a-sub-menu>

  </a-menu>

  <SelectLang
      :isVisible="selectLangVisible"
      :onClose="closeSelectLang">
  </SelectLang>

</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, PropType, ref, toRefs} from "vue";
import {  } from '@ant-design/icons-vue';
import {SettingOutlined, DownOutlined, BellOutlined, UserOutlined, LogoutOutlined} from '@ant-design/icons-vue';
import {useStore} from "vuex";
import {StateType as GlobalStateType} from "@/store/global";
import {StateType as UserStateType} from "@/store/user";
import IconSvg from "@/components/IconSvg";
import {useI18n} from "vue-i18n";
import {useRouter} from "vue-router";
import SelectLang from "./RightTopSelectLang.vue";

export default defineComponent({
  name: 'SettingsMenu',
  props: {
  },
  components: {
    SelectLang,
    // IconSvg,
    SettingOutlined, UserOutlined, LogoutOutlined
  },
  setup(props) {
    const {t} = useI18n();
    const router = useRouter();

    const store = useStore<{ Global: GlobalStateType; User: UserStateType; }>();

    const collapsed = computed<boolean>(()=> store.state.Global.collapsed);

    const selectLangVisible = ref(false)
    const closeSelectLang = async (event: any) => {
      selectLangVisible.value = false
    }

    const clickMenu = async (e: any) => {
      console.log('clickMenu', e)
      const {key} = e;

      if (key === 'lang') {
        console.log('lang')
        selectLangVisible.value = true

      } else if (key === 'profile') {
        router.replace({path: '/user/profile'})

      } else if (key === 'logout') {
        const res: boolean = await store.dispatch('User/logout');
        if (res === true) {
          router.replace({
            path: '/user/login',
            query: {
              redirect: router.currentRoute.value.path,
              ...router.currentRoute.value.query
            }
          })
        }
      }
    }

    return {
      collapsed,
      selectLangVisible,
      closeSelectLang,
      clickMenu,
    }
  }
})
</script>

<style lang="less">
.indexlayout-left-menu-bottom {
  .ant-menu {
    background-color: transparent !important;

    .ant-menu-submenu-title {
      padding-left: 20px;
      padding-right: 0;

      color: #FFFFFF !important;
      opacity: 0.65;
      &:hover {
        color: #FFF  !important;
        opacity: 1;
      }

      .ant-menu-title-content {
        margin-left: 5px !important;
        .title {
          &.hide {
            display: none;
          }
        }
      }
      .ant-menu-submenu-arrow {
        display: none;
      }
    }
  }
}

.settings-sub-menu {
  .ant-menu-item {
    .icon-svg {
      display: inline-block !important;
      margin-right: 10px !important;
    }
  }
}

</style>