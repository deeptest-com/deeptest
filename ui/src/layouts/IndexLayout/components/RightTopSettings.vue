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
        </a-menu>
      </template>
    </a-dropdown>
  </div>
</template>
<script lang="ts">
import {computed, defineComponent, ref} from "vue";
import {useStore} from "vuex";
import {DownOutlined, BellOutlined,} from '@ant-design/icons-vue';

import {useI18n} from "vue-i18n";
import {CurrentUser, StateType as UserStateType} from "@/store/user";

export default defineComponent({
  name: 'RightTopSettings',
  components: {
    DownOutlined, BellOutlined,
  },
  setup() {
    const {t} = useI18n();
    const store = useStore<{ User: UserStateType }>();

    // 获取当前登录用户信息
    const currentUser = computed<CurrentUser>(() => store.state.User.currentUser);

    const selectLangVisible = ref(false)
    const closeSelectLang = async (event: any) => {
      selectLangVisible.value = false
    }

    // 点击菜单
    const onMenuClick = async (event: any) => {
      const {key} = event;
      if (key === 'msg') {
        console.log('msg')
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

</style>