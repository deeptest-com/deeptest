import {watch,ref} from "vue";

const isFirst  = ref(true);
/**
 * 用于自动保存的 场景树编辑
 * */
function useScenarioAutoSave(data: any, fun: Function) {
    watch(data, (val: any,oldValue) => {
       // console.log('83222', val,oldValue);
        if(!oldValue){
            return;
        }
        if (!val) return;
        if(isFirst.value){
            isFirst.value = false;
            return;
        }
        fun()
        //setTimeout(fun, 1000)
        // debugger;

    }, {
        immediate: false
    });
}

export {
    isFirst,
    useScenarioAutoSave
}
