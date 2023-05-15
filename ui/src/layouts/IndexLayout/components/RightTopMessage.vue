<template>
    <router-link to="/" class="indexlayout-top-message">
      <a-badge
        class="indexlayout-top-message-badge"
        :count="message"
        :numberStyle="{ boxShadow: 'none', height: '15px', 'line-height': '15px' }"
      />
    </router-link>
</template>
<script lang="ts">
import { computed, ComputedRef, defineComponent, onMounted } from "vue";
import { useStore } from "vuex";
import { BellOutlined } from '@ant-design/icons-vue';
import { StateType as UserStateType } from "@/store/user";

interface RightTopMessageSetupData {
    message: ComputedRef<number>;
}

export default defineComponent({
    name: 'RightTopMessage',
    components: {
    },
    setup(): RightTopMessageSetupData {
        const store = useStore<{User: UserStateType}>();
        const message = computed<number>(()=> store.state.User.message);

        onMounted(()=>{
          store.dispatch("User/fetchMessage");
          // store.dispatch("User/fetchCurrent");
        })

        return {
            message
        }
    }
})
</script>