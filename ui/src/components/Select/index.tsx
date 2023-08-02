import {defineComponent,ref, defineProps, defineEmits, computed, watch, createVNode} from 'vue';
import { vOnClickOutside } from '@vueuse/components'
import { SelectTypes } from 'ant-design-vue/es/select';

export default defineComponent({
    name: 'Select',
    props: ['placeholder', 'value', 'width', 'options'],
    emits: ['change'],
    setup(props, {emit}) {

        const visible = ref(true)

        const options1 = computed<SelectTypes['options']>(() => {
          const res: SelectTypes['options'] = []
            props.options.forEach((item:any)=>{
              res.push({label:item.label,value:item.value})
            })
          return res
        })
        
        const values:any = computed(()=>props?.value || [])
        
        const optionsMap = computed(() => {
          const map = new Map()
          options1?.value?.forEach((item) => {
            map.set(item.value, item.label)
          })
          return map
        })
        
        const maxTagPlaceholder = (omittedValues: any[]) => {
          //debugger
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
          //values.value = e
          //debugger
          console.log("change",e)
          emit('change', e)
        }
        
        const focus = () => {
          visible.value = true
          console.log(visible.value,)
        }

        const blur = (e :any) => {
          //debugger
        }
        
        const close = (key: any) => {
          console.log(key)
          //debugger
          const value = values.value.filter((arrItem: any) => arrItem != key)
          emit('change', value)
        }
        
        function canClose(e :any) {
          debugger
        //  const indexlayout = document.getElementById('indexlayout');
       //   if (indexlayout != null && indexlayout.contains(e.target)) {
       //     visible.value = false
         // }
        }



        const tag = () => {
          if (values.value?.length > 0) {
            //debugger
            return values.value.map(
              (item: any,key: any) => 
              <a-tag  key={key} 
              closable onClose={()=>close(item)}>{optionsMap.value.get(item)}
              </a-tag>)
          }
           
        }

        

        return {
          change,
          options1,
          values,
          maxTagPlaceholder,
          props,
          visible,
          tag,
          canClose,
          focus,
          blur
        }
        
    },
    render(){
      return (<>
     
      <a-popover visible={this.visible && this.values?.length > 0}
            placement={'top'}
            trigger="click"
            autoAdjustOverflow={false}
            overlayClassName="dp-select-tooltip" 
            content = {this.tag()}
            />
      <a-select
            mode={'multiple'}
            maxTagCount={1}
            allowClear
            onChange={this.change}
            placeholder={this.props.placeholder}
            options={this.options1 ||[]}
            style={"width: 180px;"}
            value={this.values}
            onFocus={this.focus}
            maxTagPlaceholder={this.maxTagPlaceholder}
            onBlur={this.blur}
            v-on-click-outside={this.canClose}
            />
      </>)
    }      
})