<template>
  <div class="indexlayout-top-settings">

    <div class="msgs">
      <a @click="gotoMessage">
        <a-badge count="100" :overflowCount="99" show-zero
                 :numberStyle="{fontSize: '12px', minWidth: '12px', height: '10px', lineHeight: '10px', padding: '0 4px','box-shadow':'none'}">
          <BellOutlined  class="dp-light user-icon" :style="{ fontSize: '16px',color:'#fff',left:'-2px',position:'relative' }"/>
        </a-badge>
      </a>
    </div>

    <div class="user-info">
      <a-dropdown>
        <a class="indexlayout-top-usermenu ant-dropdown-link">
          <UserOutlined class="user-icon"/>
          <span class="user-name">{{ currentUser.name }}</span>
          <DownOutlined class="user-icon"/>
        </a>
        <template #overlay>
          <a-menu @click="onMenuClick">
            <a-menu-item key="profile">
              <SettingOutlined class="settings"/>个人信息
            </a-menu-item>
            <a-menu-item key="logout">
              <LogoutOutlined/> 登出
            </a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
    </div>


  </div>
</template>
<script lang="ts">
import {computed, defineComponent, ref} from "vue";
import {useStore} from "vuex";
import {DownOutlined, BellOutlined, SettingOutlined, UserOutlined,LogoutOutlined} from '@ant-design/icons-vue';

import {useI18n} from "vue-i18n";
import {CurrentUser, StateType as UserStateType} from "@/store/user";
import {useRouter} from "vue-router";

export default defineComponent({
  name: 'RightTopSettings',
  components: {
    DownOutlined,
    BellOutlined, SettingOutlined, UserOutlined,LogoutOutlined
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

      // console.log(currentUser.value);

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
.user-info {
  .user-name {
    margin-left: 4px;
    margin-right: 4px;
    display: inline-block;
    color: #FFFFFF;
  }
  .user-info{
    color: #FFFFFF;
  }
  .user-icon{
    color: #FFFFFF;
    //font-size: 18px;
    //margin-left: 4px;
  }
  margin-right: 24px;

}

.msgs {
  text-align: center;
  margin-right: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  top:2px;
}

.settings {
  display: inline-block;
  margin-right: 16px;
}
</style>
