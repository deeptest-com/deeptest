import {defineComponent,ref, defineProps, defineEmits, computed, watch, createVNode} from 'vue';
import {vOnClickOutside} from '@vueuse/components';

export default defineComponent({
    name: 'Select',
    props: {
        placeholder: {
            type: String,
            default: 'small'
          },
          value: {
            type: [],
            required: true,
          },
          width: {
            type: String,
            default: '180px'
          },
          options: {
            type: [],
            default: [],
          },
    },
    emits: ['change'],
    setup(props, {emit}) {

        const visible = ref(false)

        const options = computed<any[]>(() => props.options)
        
        const values:any = ref(props?.value || [])
        
        const optionsMap = computed(() => {
          const map = new Map()
          options.value.forEach((item) => {
            map.set(item.value, item.label)
          })
          return map
        })
        
        const maxTagPlaceholder = (omittedValues: any[]) => {
          let res = ""
          omittedValues.forEach((item: { label: string; }) => {
            res += res ? "," + item.label : item.label
          })
        /*
          return createVNode('tooltip', {
            placement: 'top',
            title: res,
            overlayClassName:'dp-select-tag-tooltip'
          }, {
            default: () => {
              return `+${omittedValues.length}...`
            },
          })
          */
         return <a-tooltip placement='top' title={res} >{omittedValues.length}...</a-tooltip>
        
        }
        
        const change = (e: any) => {
          values.value = e
          emit('change', e)
        }
        
        const focus = () => {
          visible.value = true
        }
        
        const close = (key: any) => {
          values.value = values.value.filter((arrItem: any) => arrItem != key)
        }
        
        function canClose(e: { target: Node | null; }) {
          const indexlayout = document.getElementById('indexlayout');
          if (indexlayout != null && indexlayout.contains(e.target)) {
            visible.value = false
          }
        }

        const select = () => {
            return       <a-select
            mode={'multiple'}
            maxTagCount="1"
            allowClear
            onChange={change(this)}
            placeholder="placeholder"
            options={options}
            style="width: 180px;"
            value={values}
            OnFocus={focus()}
            onBlur={blur()}
            maxTagPlaceholder={maxTagPlaceholder}
            v-on-click-outside={canClose}/>
        }


        const tag = () => {
            const tags = values.value.map((item: any,key: any) => <a-tag key={key} closable OnClose={close(item)}>{optionsMap.value.get(item) }</a-tag>)
            return tags
        }

        const tags = () => {
            return <a-popover visible={visible.value && values.value?.length}
            placement={'top'}
            trigger="click"
            autoAdjustOverflow={false}
            overlayClassName="dp-select-tooltip" 
            content = {tag()}
            />
        }


        return ()=>
                <div>
                    {select()}
                </div>
    }
})