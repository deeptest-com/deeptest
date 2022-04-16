<template>
  <div class="indexlayout-top-settings">
    <a-dropdown>
      <a @click="e => e.preventDefault()"
         class="indexlayout-top-usermenu ant-dropdown-link">
        <span>设置</span>
        <DownOutlined/>
      </a>

      <template #overlay>
        <a-menu @click="onMenuClick">
          <a-menu-item key="msg">
            <BellOutlined :style="{ fontSize: '16px' }" />
            <span>消息（0）</span>
          </a-menu-item>
          <a-menu-item key="lang">
            <icon-svg type="language-outline" class="icon" />
            <span>界面语言</span>
          </a-menu-item>
          <a-menu-item key="info">
            <UserOutlined />
            <span>个人信息</span>
          </a-menu-item>
          <a-menu-item key="logout">
            <LogoutOutlined />
            <span>{{ t('index-layout.topmenu.logout') }}</span>
          </a-menu-item>
        </a-menu>
      </template>
    </a-dropdown>

    <SelectLang
        :isVisible="selectLangVisible"
        :onClose="closeSelectLang">
    </SelectLang>

  </div>
</template>
<script lang="ts">
import {computed, defineComponent, ref} from "vue";
import {useStore} from "vuex";
import {useRouter} from "vue-router";
import {useI18n} from "vue-i18n";
import {DownOutlined, BellOutlined, UserOutlined, LogoutOutlined} from '@ant-design/icons-vue';
import IconSvg from "@/components/IconSvg";
import {CurrentUser, StateType as UserStateType} from "@/store/user";

import SelectLang from "./RightTopSelectLang.vue";

export default defineComponent({
  name: 'RightTopSettings',
  components: {
    SelectLang, IconSvg,
    DownOutlined, BellOutlined, UserOutlined, LogoutOutlined,
  },
  setup() {
    const {t} = useI18n();
    const router = useRouter();
    const store = useStore<{ user: UserStateType }>();

    // 获取当前登录用户信息
    const currentUser = computed<CurrentUser>(() => store.state.user.currentUser);

    const selectLangVisible = ref(false)
    const closeSelectLang = async (event: any) => {
      selectLangVisible.value = false
    }

    // 点击菜单
    const onMenuClick = async (event: any) => {
      const {key} = event;
      if (key === 'msg') {
        console.log('msg')

      } else if (key === 'lang') {
        console.log('lang')
        selectLangVisible.value = true

      } else if (key === 'info') {
        console.log('info')

      } else if (key === 'logout') {
        const res: boolean = await store.dispatch('user/logout');
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
      t,
      currentUser,
      onMenuClick,
      selectLangVisible,
      closeSelectLang,
    }
  }
})
</script>

<style lang="less" scoped>
.icon {
  display: inline-block;
  margin-right: 8px;
}
</style>