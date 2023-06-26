<template>
  <div :class="['indexlayout-top-settings', theme]">
    <div class="user-info">
      <a-dropdown placement="bottomRight">
        <a class="indexlayout-top-usermenu ant-dropdown-link">
          <UserOutlined class="user-icon"/>
          <span class="user-name">{{ currentUser.name }}</span>
          <DownOutlined class="user-icon"/>
        </a>
        <template #overlay>
          <a-menu @click="onMenuClick">
            <a-menu-item key="profile">
              <SettingOutlined class="settings"/>
              个人信息
            </a-menu-item>
            <a-menu-item key="management">
              <SettingOutlined class="settings"/>
              用户管理
            </a-menu-item>
            <a-menu-item key="logout">
              <LogoutOutlined/>
              登出
            </a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>

      <a-tooltip placement="bottom" @click="toggle">
        <template #title>{{ isFullscreen ? '退出全屏' : '全屏' }}</template>
        <a-button type="text" class="share-btn">
          <FullscreenOutlined v-if="isFullscreen" :style="{'font-size': '14px','color':theme === 'white-theme' ? '#fff' : '#8A8A8A'}"/>
          <FullscreenExitOutlined v-if="!isFullscreen" :style="{'font-size': '14px','color':theme === 'white-theme' ? '#fff' : '#8A8A8A'}"/>
        </a-button>
      </a-tooltip>
    </div>
    <!-- <div class="msgs">
      <a @click="gotoMessage">
        <a-badge count="100" :overflowCount="99" show-zero
                 :numberStyle="{fontSize: '12px', minWidth: '12px', height: '10px', lineHeight: '10px', padding: '0 4px','box-shadow':'none'}">
          <MailOutlined  class="dp-light user-icon" :style="{ fontSize: '16px',left:'-2px',position:'relative' }"/>
        </a-badge>
      </a>
    </div> -->
    <!--    <div class="user-management" @click="onManagementClick">-->
    <!--      <a-tooltip placement="bottom" title="用户管理">-->
    <!--        <icon-svg type="setting" :style="{width: '20px', height: '20px', fill: `${theme === 'white-theme' ? '#fff' : '#8A8A8A' }`}" />-->
    <!--      </a-tooltip>-->
    <!--    </div>-->
  </div>
</template>
<script lang="ts">
import {computed, defineComponent, ref} from "vue";
import {useStore} from "vuex";

import {
  DownOutlined,
  SettingOutlined,
  UserOutlined,
  LogoutOutlined,
  FullscreenOutlined,
  FullscreenExitOutlined
} from '@ant-design/icons-vue';
import {useI18n} from "vue-i18n";
// import IconSvg from "@/components/IconSvg";
import {CurrentUser, StateType as UserStateType} from "@/store/user";
import {useRouter} from "vue-router";
import {useClipboard, useFullscreen} from '@vueuse/core';

export default defineComponent({
  name: 'RightTopSettings',
  components: {
    DownOutlined,
    SettingOutlined, UserOutlined, LogoutOutlined,
    // IconSvg
    FullscreenOutlined, FullscreenExitOutlined
  },
  props: {
    theme: {
      required: false,
      type: String
    }
  },
  setup() {
    const {t} = useI18n();
    const router = useRouter();
    const store = useStore<{ User: UserStateType }>();
    const {isFullscreen, enter, exit, toggle} = useFullscreen();
    // 获取当前登录用户信息
    const currentUser = computed<CurrentUser>(() => store.state.User.currentUser);

    const selectLangVisible = ref(false)
    const closeSelectLang = async (event: any) => {
      selectLangVisible.value = false
    }

    const gotoMessage = () => {
      console.log('gotoMessage')
      router.replace({path: '/user-manage/message'})
    }

    // 点击菜单
    const onMenuClick = (event: any) => {
      console.log('onMenuClick')

      // console.log(currentUser.value);

      const {key} = event;

      if (key === 'profile') {
        router.replace({path: '/user-manage/profile'})
      } else if (key === 'logout') {
        store.dispatch('User/logout').then((res) => {
          if (res === true) {
            router.replace({
              path: '/user/login',
              query: {
                redirect: router.currentRoute.value.path,
                ...router.currentRoute.value.query
              }
            })
          }
        })
      } else if (key === 'management') {
        router.replace({path: '/user-manage/index'})
      }
    }

    const onManagementClick = () => {
      router.replace({path: '/user-manage/index'})
    }

    return {
      t,
      currentUser,
      gotoMessage,
      onMenuClick,
      selectLangVisible,
      closeSelectLang,
      onManagementClick,
      toggle,
      isFullscreen
    }
  }
})
</script>

<style lang="less" scoped>

.indexlayout-top-settings {
  color: #FFFFFF;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  padding-right: 16px;

  &.white-theme {
    .msgs {
      .user-icon {
        color: #fff;
      }
    }

    .user-info {
      .user-icon {
        color: #fff;
      }

      .user-name {
        color: #fff;
      }
    }
  }

  .msgs {
    width: 40px;

    .user-icon {
      color: #8A8A8A;
    }
  }

  .indexlayout-top-usermenu {
    color: #c0c4cc;
  }
}

.user-management {
  display: flex;
  align-items: center;
  cursor: pointer;
}

.user-info {
  .user-name {
    margin-left: 4px;
    margin-right: 4px;
    display: inline-block;
    color: #8A8A8A;
  }

  .user-info {
    color: #8A8A8A;
  }

  .user-icon {
    color: #8A8A8A;
    //font-size: 18px;
    //margin-left: 4px;
  }

  //margin-right: 8px;

}

.msgs {
  text-align: center;
  margin-right: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  top: 2px;
}

.settings {
  display: inline-block;
  margin-right: 16px;
}
</style>
