<template>
  <div class="indexlayout-main-conent">
    <a-card :bordered="false">
      <template #title>
        <div>
          <a-button type="primary" @click="() => record()" id="com-deeptest-record" class="act-btn">录制</a-button>
          <a-button @click="() => playback()" class="act-btn">播放</a-button>
        </div>
      </template>
      <template #extra>
        <a-button type="link" @click="() => back()">返回</a-button>
      </template>

      <div class="script">
        <div class="title">测试步骤</div>
        <div class="desc">
          <div v-for="(step, index) in script.steps" :key="index" class="step">
            <div class="cmd">{{step.action}} {{step.selector}}  {{step.value}}</div>
            <div class="capture" style="border: 3px cornflowerblue;"><img v-if="stepImage[step.uuid]" :src="stepImage[step.uuid]"></div>
          </div>
        </div>
      </div>

      <br /> <!--WebSocket Test-->
      <div>
        <div><a-input id="input" type="text" v-model:value="wsMsg.in" /></div>
        <div><a-button id="sendBtn" @click="sendWs">Send</a-button></div>
        <div>
          <pre>{{ wsMsg.out }}</pre>
        </div>
      </div>

    </a-card>
  </div>
</template>

<script lang="ts">
import {defineComponent, onMounted, onBeforeUnmount, getCurrentInstance, ComputedRef, Ref, ref, reactive} from "vue";
import { useStore } from 'vuex';
import {StateType as ListStateType} from "@/views/script/store";
import {useRouter} from "vue-router";
import {ActionRecordStart, EventNameJson, EventNodeIdJson, EventNameImg, EventNodeIdImg, ScopeDeeptest, ActionRecordedMsg} from "@/utils/const";
import {WebSocket, WsEventName} from "@/services/websocket";
import {getToken} from "@/utils/localToken";
import {ScriptItem, StepImage, StepItem} from "@/views/script/data";
import {createStep} from "@/views/script/service";

interface DesignScriptPageSetupData {
  script: ScriptItem;
  stepImage: any;

  loading: Ref<boolean>;
  getScript:  (current: number) => Promise<void>;
  record: () => void;
  playback: () => void;
  back: () => void;

  wsMsg: any,
  sendWs: () => void;
}

export default defineComponent({
    name: 'ScriptEditPage',
    setup(): DesignScriptPageSetupData {
      const router = useRouter();
      const store = useStore<{ ListScript: ListStateType}>();

      const script = reactive<ScriptItem>({steps: []})
      const stepImage = reactive({})
      const loading = ref<boolean>(true)

      const id = +router.currentRoute.value.params.id
      console.log('id', id)
      const getScript = async (id: number): Promise<void> => {
        loading.value = true;
        // await store.dispatch('ListScript/getScript', {
        // });
        loading.value = false;
      }

      const record = ():void =>  {
        console.log('record')

        window.postMessage({
          scope: ScopeDeeptest,
          content: {
            act: ActionRecordStart,
          }
        }, "*");
      }
      const playback = ():void =>  {
        console.log('playback')
      }

      const back = ():void =>  {
        router.push(`/~/script/list`)
      }

      let init = true;
      let wsMsg = reactive({in: '', out: ''});

      let room: string | null = ''
      getToken().then((token) => {
        room = token
      })

      const sendWs = () => {
        console.log('sendWs');
        WebSocket.sentMsg(room, wsMsg.in);
        wsMsg.out = wsMsg.out + 'client: ' + wsMsg.in + '\n';
      };

      const { proxy } = getCurrentInstance() as any;
      WebSocket.init(proxy)

      onMounted(() => {
        const eventNodeJson = document.getElementById(EventNodeIdJson)
        if (eventNodeJson) {
          eventNodeJson.addEventListener(EventNameJson, function () {
            const msg: StepItem = JSON.parse(eventNodeJson.innerText);
            console.log('====', msg);

            script.steps.push({
              uuid: msg.uuid,
              action: msg.action,
              selector: msg.selector,
              value: msg.value? msg.value:'',
            })
          });
        }

        const eventNodeImg = document.getElementById(EventNodeIdImg)
        if (eventNodeImg) {
          eventNodeImg.addEventListener(EventNameImg, function () {
            const msg: StepImage = JSON.parse(eventNodeImg.innerText);
            console.log('====', script.steps);

            if (msg.uuid) {
              stepImage[msg.uuid] = msg.image

              const step = Object.assign({scriptId: id}, script.steps[script.steps.length - 1])
              createStep(step, msg.image)
            }
          });
        }

        getScript(1);
        if (init) {
          proxy.$sub(WsEventName, (data) => {
            console.log(data[0].msg);
            wsMsg.out = wsMsg.out + 'server: ' + data[0].msg + '\n';
            console.log(wsMsg.out);
          });
          init = false;
        }
      });
      onBeforeUnmount(() => {
        proxy.$unsub(WsEventName, () => {
          console.log('unsub event ' + WsEventName);
        });
      });

      return {
        script,
        stepImage,
        loading,
        getScript,
        record,
        playback,
        back,
        wsMsg,
        sendWs,
      }
    }
})
</script>

<style lang="less" scoped>
  .act-btn {
    margin-right: 20px;
  }

  .script {
    .title {
      font-weight: bolder;
    }
    .desc {
      .step {
        display: flex;
        .cmd {
          flex: 1;
        }
        .capture {
          width: 600px;

          img {
            height: 50px;
          }
        }
      }
    }
  }
</style>
