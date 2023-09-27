<template>
  <div :class="['indexlayout-top-settings', theme]">
    <div class="user-info">

      <template v-if="isLyEnv">
        <!--  客户端下载 -->
        <a-dropdown placement="bottomRight" v-if="!isElectronEnv">
          <a class="indexlayout-top-usermenu ant-dropdown-link" style="margin-right: 4px;margin-left: 4px;">
            <DesktopOutlined type="top-right-web" class="top-right-icon-desktop"/>
            <span class="user-name">{{ '客户端下载' }}</span>
            <DownOutlined class="user-icon"/>
          </a>
          <template #overlay>
            <a-menu @click="downloadClient">
              <a-menu-item v-for="client in clientDownloadUrlOpts" :key="client.value">
                {{ client.label }}
              </a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>

        <!--  切换Agent -->
        <a-dropdown placement="bottomRight">
          <a class="indexlayout-top-usermenu ant-dropdown-link" style="margin-right: 6px;margin-left: 8px;">
            <IconSvg type="top-right-web" class="top-right-icon"/>
            <span class="user-name">{{ currentAgentLabel }}</span>
            <DownOutlined class="user-icon"/>
          </a>
          <template #overlay>
            <a-menu @click="changeAgentEnv">
              <a-menu-item v-for="agent in agentUrlOpts" :key="agent.value"
                           :style="agent.label === currentAgentLabel ? {color:'#1890ff','background-color': '#e6f7ff'} : {}">
                <a-tooltip placement="left" :title="agent.desc">
                  {{ agent.label }}
                </a-tooltip>
              </a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
      </template>

      <template v-else> <!-- 系统菜单 -->
        <a-dropdown placement="bottomRight">
          <a class="indexlayout-top-sysmenu ant-dropdown-link" style="margin-right: 6px;margin-left: 8px;">
            <SettingOutlined class="top-right-icon-desktop"/>
            <span class="user-name">系统</span>
            <DownOutlined class="user-icon"/>
          </a>
          <template #overlay>
            <a-menu @click="onSysMenuClick">
              <a-sub-menu key="agent-sub-menu" title="切换代理 &nbsp;">
                <a-menu-item v-for="agent in agentUrlOpts" :key="agent.value"
                             :style="agent.label === currentAgentLabel ? {color:'#1890ff','background-color': '#e6f7ff'} : {}">
                  <a-tooltip placement="left" :title="agent.desc">
                    {{ agent.label }}
                  </a-tooltip>
                </a-menu-item>
              </a-sub-menu>

              <a-menu-item key="userManage">
                用户管理
              </a-menu-item>

              <a-menu-item key="jslibManage">
                自定义代码库
              </a-menu-item>

              <a-menu-item key="download">
                下载客户端
              </a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
      </template>

      <!-- ::::用户信息 -->
      <a-dropdown placement="bottomRight">
        <a class="indexlayout-top-usermenu ant-dropdown-link" style="margin-right: 6px;margin-left: 8px;">
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
            <a-menu-item v-if="isLyEnv" key="management">
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
          <FullscreenOutlined v-if="isFullscreen"
                              :style="{'font-size': '14px','color':theme === 'white-theme' ? '#fff' : '#8A8A8A'}"/>
          <FullscreenExitOutlined v-if="!isFullscreen"
                                  :style="{'font-size': '14px','color':theme === 'white-theme' ? '#fff' : '#8A8A8A'}"/>
        </a-button>
      </a-tooltip>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, onMounted, ref} from "vue";
import {useStore} from "vuex";
import {getAgentLabel, getAgentUrl, getAgentUrlByValue, isElectronEnv} from '@/utils/agentEnv'
import {
  DownOutlined,
  SettingOutlined,
  UserOutlined,
  LogoutOutlined,
  FullscreenOutlined,
  FullscreenExitOutlined,
  DesktopOutlined, CheckOutlined,
} from '@ant-design/icons-vue';
import {useI18n} from "vue-i18n";
import IconSvg from "@/components/IconSvg";
import {CurrentUser, StateType as UserStateType} from "@/store/user";
import {useRouter} from "vue-router";
import {useFullscreen} from '@vueuse/core';
import {StateType as GlobalStateType} from "@/store/global";

const props = defineProps({
  theme: {
    required: false,
    type: String
  }
})

const {t} = useI18n();
const router = useRouter();
const store = useStore<{ User: UserStateType, Global: GlobalStateType }>();
const {isFullscreen, enter, exit, toggle} = useFullscreen();
// 获取当前登录用户信息
const currentUser = computed<CurrentUser>(() => store.state.User.currentUser);

// 获取当前可以切换的 Agent 地址
const agentUrlOpts = computed(() => {
  const opts = store.state.Global.configInfo?.agentUrlOpts;
  if (opts?.length > 0) {
    if (!isElectronEnv) {
      return opts.filter((item) => item.value !== 'local');
    }
    return opts;
  }
  return [];
});

const selectLangVisible = ref(false)
const closeSelectLang = async (event: any) => {
  selectLangVisible.value = false
}

const gotoMessage = () => {
  router.replace({path: '/message'})
}

// 点击菜单
const onMenuClick = (event: any) => {
  console.log('onMenuClick')
  // console.log(currentUser.value);
  const {key} = event;

  if (key === 'profile') {
    router.replace({path: '/profile'})
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
    router.replace({path: '/user-manage'})
  }
}

// 系统菜单
const onSysMenuClick = (event: any) => {
  console.log('onSysMenuClick', event)
  const {key, keyPath} = event;

  if (key === 'userManage') { //
    router.replace({path: '/sys-setting/user-manage'})

  } else if (key === 'jslibManage') { //
    router.replace({path: '/sys-setting/jslib'})

  } else if (key === 'download') {
    window.open('https://deeptest.com/setup.html');

  } else if (keyPath[0] === 'agent-sub-menu') {
    const url = getAgentUrlByValue(agentUrlOpts.value, key);
    window.localStorage.setItem('dp-cache-agent-value', key);
    window.localStorage.setItem('dp-cache-agent-url', url);
    window.location.reload();
  }
}

function changeAgentEnv(event: any) {
  const {key} = event;
  const url = getAgentUrlByValue(agentUrlOpts.value, key);
  window.localStorage.setItem('dp-cache-agent-value', key);
  window.localStorage.setItem('dp-cache-agent-url', url);
  window.location.reload();
}

// 下载客户端
function downloadClient(event: any) {
  if(event?.key){
    window.open(event.key);
  }
}

const isLyEnv = process?.env?.VUE_APP_DEPLOY_ENV === 'ly';
const clientDownloadUrlOpts = computed(() => {
  if (!isLyEnv) {
    return []
  }
  const clientVersion = store.state.Global.clientVersion;
  const url = process?.env?.VUE_APP_API_STATIC;
  return [
    {
      label: 'Windows 客户端',
      desc: 'Windows 客户端',
      value: `${url}/LeyanAPI/${clientVersion}/win64/LeyanAPI.zip`
    },
    {
      label: 'macOS 客户端',
      desc: 'macOS 客户端',
      value: `${url}/LeyanAPI/${clientVersion}/darwin/LeyanAPI.zip`
    }
  ];
});

const currentAgentLabel = computed(() => {
  return getAgentLabel(agentUrlOpts.value);
})

const onManagementClick = () => {
  router.replace({path: '/user-manage'})
}

onMounted(async () => {
    const list = await store.dispatch('Global/getConfigByKey', {key: 'agentUrlOpts'});

    console.log('===-==', agentUrlOpts.value)

    // 如果没有缓存，根据当前环境选择一个默认值
    if (!window.localStorage.getItem('dp-cache-agent-value')) {
      const agentValue = isElectronEnv ? 'local' : 'test';
      const url = getAgentUrlByValue(list, agentValue);
      window.localStorage.setItem('dp-cache-agent-value', agentValue);
      window.localStorage.setItem('dp-cache-agent-url', url);
    }

    // 获取客户端最新版本号
    await store.dispatch('Global/getClientVersion');
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
  .indexlayout-top-sysmenu {
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

.top-right-icon {
  transform: scale(1.2);
}
.top-right-icon-desktop{
  margin-right: 2px;
}
</style>
