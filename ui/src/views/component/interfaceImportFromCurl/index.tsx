import {defineComponent,ref, reactive} from 'vue';
import {Form} from 'ant-design-vue';
import  './index.less';


export default defineComponent({
    name: 'Select',
    props: [],
    emits: ['onCancel','onFinish'],
    setup(props, {emit}) {

        const useForm = Form.useForm;
        const modelRef = reactive({content: ''} as any);
        const showError = ref(false)
        
        const validateCurl = async (rule: any, value: string,callback: any) => {
            if (value === '') {
              showError.value = true
              return Promise.reject("请输Curl请求")
            } else {
              if (!rule.pattern.test(value)){
                showError.value = true
                return Promise.reject("不是合法的cURL请求，请重试。")
              }
              showError.value = false
              return Promise.resolve();
            }
          };

        const rulesRef = reactive({
            content: [
              {required: true,  message: '',validator:validateCurl, trigger: 'change',pattern:/curl\s+.*\s+.*/},
            ],
          });

        const { validate, validateInfos} = useForm(modelRef, rulesRef);

        const onfinish = () => {
            validate().then(() => {
                emit('onFinish',modelRef.content)
            }).catch((error) => {
              console.log('error', error)
            })
          }
          
        const onCancel = () => {
            //debugger
            emit('onCancel')
          }

        const error = () => {
            if (showError.value) 
            return <a-alert message="不是合法的cURL请求，请重试。" type={'error'} show-icon />
        }   
        
        return {
           modelRef,
           showError,
           onfinish,
           onCancel,
           validateInfos,
           error
        }
        
    },
    render(){
      return (<>
        <a-modal
            title="请输入curl命令"
            destroy-on-close={true}
            visible={true}
            onCancel={this.onCancel}
            onOk={this.onfinish}
            width="1000px"
       >
            <a-form >
                <a-form-item>
                <a-textarea size={'large'} v-model={[this.modelRef.content,'value']}/>
                </a-form-item>
                <a-form-item>
                {this.error()}
                </a-form-item>
            </a-form>
      </a-modal>

      </>)
    }      
})