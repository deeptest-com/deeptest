<template>
  <div class="indexlayout-top-settings">
    <div class="msgs">
      <a @click="gotoMessage">
        <a-badge count="100" :overflowCount="99" show-zero
                 :numberStyle="{fontSize: '12px', minWidth: '16px', height: '16px', lineHeight: '15px', padding: '0 4px'}">
            <BellOutlined class="dp-light" :style="{ fontSize: '20px' }" />
        </a-badge>
      </a>
    </div>

    <a-dropdown>
      <a class="indexlayout-top-usermenu ant-dropdown-link">
        <SettingOutlined class="settings" />
        <DownOutlined/>
      </a>

      <template #overlay>
        <a-menu @click="onMenuClick">
          <a-menu-item key="profile">
            个人信息
          </a-menu-item>
          <a-menu-item key="logout">
            登出
          </a-menu-item>
        </a-menu>
      </template>
    </a-dropdown>
  </div>
</template>
<script lang="ts">
import {computed, defineComponent, ref} from "vue";
import {useStore} from "vuex";
import {DownOutlined, BellOutlined, SettingOutlined} from '@ant-design/icons-vue';

import {useI18n} from "vue-i18n";
import {CurrentUser, StateType as UserStateType} from "@/store/user";
import {useRouter} from "vue-router";

export default defineComponent({
  name: 'RightTopSettings',
  components: {
    DownOutlined,
    BellOutlined, SettingOutlined
  },
  setup() {
    const {t} = useI18n();
    const router = useRouter();
    const store = useStore<{ User: UserStateType }>();

    // 获取当前登录用户信息
    const currentUser = computed<CurrentUser>(() => store.state.User.currentUser);

    const selectLangVisible = ref(false)
    const closeSelectLang = async (event: any) => {
      selectLangVisible.value = false
    }

    const gotoMessage = () => {
      console.log('gotoMessage')
      router.replace({path: '/user/message'})
    }

    // 点击菜单
    const onMenuClick = (event: any) => {
      console.log('onMenuClick')

      const {key} = event;

      if (key === 'profile') {
        router.replace({path: '/user/profile'})
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
      }
    }

    return {
      t,
      currentUser,
      gotoMessage,
      onMenuClick,
      selectLangVisible,
      closeSelectLang,
    }
  }
})
</script>

<style lang="less" scoped>
  .msgs {
    text-align: left;
  }
  .settings {
    display: inline-block;
    margin-right: 3px;
  }
</style>