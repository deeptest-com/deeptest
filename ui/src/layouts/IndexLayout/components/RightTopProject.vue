<template>
  <div class="indexlayout-top-project">
    <a-dropdown :overlayClassName="'dp-project-switch-container'"
                :visible="dropdownVisible"
                @click="dropdownVisible = !dropdownVisible"
                :overlayStyle="{width:'300px'}">
      <a-button class="header">
        <span class="before-icon icon"><img :src="getProjectLogo(currProject?.logo)" alt=""></span>
        <span class="header-text" :title="currProject.name">{{ currProject.name }}</span>
        <DownOutlined class="after-icon"/>
      </a-button>
      <template #overlay>
        <a-menu class="menu">
          <a-menu-item key="filter" class="menu-item filter">
            <a-input-search allowClear v-model:value="keyword" placeholder="搜索项目名称"/>
          </a-menu-item>
          <a-menu-item class="menu-scroll menu-item" key="menu-items">
            <div key="recently" class="menu-scroll-item recently">
              最近访问的项目
            </div>
            <div class="menu-scroll-item"
                 :class="{'first':index===0}" v-for="(item,index) in myRecentProject"
                 @click="() => {
                  selectProject(item)
                 }"
                 :key="'recently' + item.id + Math.random()">
              <span class="icon"><img :src="getProjectLogo(item?.logo)" alt=""></span>
              <span class="text" :title="item.name">{{ item.name }}</span>
            </div>
            <div class="menu-scroll-item my" key="my">
              我参与的项目
            </div>
            <div class="menu-scroll-item"
                 :class="{'first':index===0}"
                 @click="() => {
                  selectProject(item)
                 }"
                 v-for="(item,index) in myProject"
                 :key="'my' + item.id + Math.random()">
              <span class="icon"><img :src="getProjectLogo(item?.logo)" alt=""></span>
              <span class="text">{{ item.name }}</span>
            </div>
          </a-menu-item>
          <a-menu-item key="footer" class="menu-item footer">
            <a-button type="link" :size="'small'" @click="newProject">
              <PlusOutlined/>
              新建项目
            </a-button>
            <a-button type="link" :size="'small'" @click="viewAllProject">
              浏览所有项目
            </a-button>
          </a-menu-item>
        </a-menu>
      </template>
    </a-dropdown>

    <!-- 创建项目弹窗 -->
    <CreateProjectModal
        :visible="createProjectModalVisible"
        @update:visible="createProjectModalVisible = false"
        @handleSuccess="handleCreateSuccess"
    />

  </div>
</template>

<script setup lang="ts">
import {computed, watch, ref, onMounted, onUnmounted, unref} from "vue";
import {useStore} from "vuex";
import {useRoute} from "vue-router";
import router from '@/config/routes';
import {StateType as UserStateType} from "@/store/user";
import {StateType as ProjectStateType} from "@/store/project";
import {StateType as ServeStateType} from "@/store/serve";
import {StateType as EnvironmentStateType} from "@/store/environment";
import {
  DownOutlined,
  PlusOutlined,
} from '@ant-design/icons-vue';
import CreateProjectModal from "@/components/CreateProjectModal/index.vue";
import {getProjectLogo} from "@/components/CreateProjectModal";

const store = useStore<{
  User: UserStateType,
  ProjectGlobal: ProjectStateType, ServeGlobal: ServeStateType, Environment: EnvironmentStateType
}>();

const route = useRoute();
const createProjectModalVisible = ref(false);
const message = computed<number>(() => store.state.User.message);
const projects = computed<any>(() => store.state.ProjectGlobal.projects);
const recentProjects = computed<any>(() => store.state.ProjectGlobal.recentProjects);
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);

store.dispatch("User/fetchMessage");
store.dispatch("ProjectGlobal/fetchProject");

const keyword = ref('');
const dropdownVisible = ref(false);

const myProject = computed(() => {
  return projects.value.filter((item: any) => item.name.includes(keyword.value));
});
const myRecentProject = computed(() => {
  return recentProjects.value.filter((item: any) => item.name.includes(keyword.value));
});

function viewAllProject() {
  router.push({path:'',query:{type:'all'}});
}

function newProject() {
  createProjectModalVisible.value = true;
  dropdownVisible.value = false;
}

const handleCreateSuccess = async () => {
  createProjectModalVisible.value = false;
  await store.dispatch("ProjectGlobal/fetchProject");
};

const selectProject = async (item): Promise<void> => {
  dropdownVisible.value = false;
  window.localStorage.setItem('currentProjectId', item.id);
  const { path, params }: any = unref(router.currentRoute);
  router.replace(path.replace(params.projectNameAbbr, item.shortName));
  await store.dispatch('ProjectGlobal/changeProject', item.id);
  await store.dispatch('Environment/getEnvironment', {id: 0, projectId: item.id});

  // 项目切换后，需要重新更新可选服务列表
  await store.dispatch("ServeGlobal/fetchServe");

  // 更新左侧菜单以及按钮权限
  await store.dispatch('Global/getPermissionList', { projectId: item.id });
}

const handleClickOut = (event) => {
  try {
    if (!dropdownVisible.value) {
      return;
    }
    const target1: any = document.querySelector('.dp-project-switch-container');
    const target2: any = document.querySelector('.header.ant-dropdown-trigger.ant-dropdown-open');
    if (!target1.contains(event.target) && !target2.contains(event.target)) {
      dropdownVisible.value = false;
    }
  } catch (e) {
    console.log('handleClickOut', e);
  }
};

onMounted(() => {
  document.addEventListener('click', handleClickOut);
});
onUnmounted(() => {
  document.removeEventListener('click', handleClickOut);
});

/**
 * fixed: 保证选中项目不为空之后，再fetchServe列表，避免由于currProjectId=NaN导致服务列表返回异常问题
 */
watch(() => {return currProject.value;}, (val: any) => {
  if (val.id) {
    store.dispatch("ServeGlobal/fetchServe");
  }
})

watch(dropdownVisible, (value) => {
  if (value)  {
    keyword.value = ""
  }
})

</script>

<style lang="less" scoped>
@dropdown-width: 300px;
.indexlayout-top-project{
  height: 32px;
  line-height: 32px;
}

.header {
  width: @dropdown-width;
  position: relative;
  text-align: left;
  padding-left: 30px;
  padding-right: 24px;

  .before-icon {
    position: absolute;
    left: 4px;
    top: 4px;
  }
  .header-text{
    width: 98%;
    display: inline-block;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis; /* 在文本溢出时显示省略号 */
  }

  .after-icon {
    position: absolute;
    top: 8px;
    right: 8px;
  }
}

.icon {
  width: 24px;
  height: 24px;
  display: inline-block;

  img {
    width: 24px;
    height: 24px;
  }
}


</style>
