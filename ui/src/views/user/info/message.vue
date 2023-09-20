<template>
  <HomeLayout>
    <div class="message-main">
      <a-card>
        <template #title>
          <div>用户消息</div>
        </template>
        <template #extra></template>

        <div>
        </div>

      </a-card>

    </div>
  </HomeLayout>
</template>

<script setup lang="ts">
import {computed, onMounted, reactive, ref} from "vue";
import {useRouter} from 'vue-router';
import {useI18n} from "vue-i18n";

import {Form, notification} from 'ant-design-vue';
import {useStore} from "vuex";
import {StateType as UserStateType} from "@/store/user";
import HomeLayout from "@/layouts/HomeLayout.vue";

const useForm = Form.useForm;

const router = useRouter();
const {t} = useI18n();

const store = useStore<{User: UserStateType}>();
const modelRef = computed<any>(()=> store.state.User.currentUser);

onMounted(()=>{
  // store.dispatch("User/fetchCurrent");
})

const labelCol = {span: 4}
const wrapperCol = {span: 14}

</script>

<style lang="less" scoped>
.message-main {

}
</style>