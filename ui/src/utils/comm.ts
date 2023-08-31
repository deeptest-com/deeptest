import {ProcessorCategory} from "@/utils/enum";
import {watch} from "vue";

export const getArrSelectItems = (arr) => {
    const ret : any[] = []

    for (const index in arr) {
        ret.push({label: arr[index], value: arr[index]})
    }

    return ret
}

export const getEnumSelectItems = (enumDef) => {
    const arr : any[] = []

    for (const item in enumDef) {
        arr.push({label: enumDef[item], value: item})
    }

    return arr
}

export const getEnumArr = (enumDef) => {
    const arr : any[] = []

    for (const item in enumDef) {
        arr.push(enumDef[item])
    }

    return arr
}

export const getResponseKey = (debugInfo) => {
    const key = `${debugInfo.debugInterfaceId}-${debugInfo.endpointInterfaceId}`
    console.log('getResponseKey', key)
    return key
}

export  function autoSave (data:any,fun:Function){
    watch(data, (val: any) => {
    if (!val) return;
    setTimeout( fun,1000)
  },{
    immediate: true
  });
}