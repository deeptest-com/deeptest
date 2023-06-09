import {
    ref,
    defineProps,
    defineEmits,
    computed,
    watch,
} from 'vue';

import {requestMethodOpts,responseCodes} from '@/config/constant';


function getMethodColor(method: any) {
    const item: any = requestMethodOpts.find((item: any) => {
        return item.value === method;
    });
    return item.color || '#04C495';
}

function getCodeColor(code) {
    const item: any = responseCodes.find((item: any) => {
        return item.value === code;
    });
    return item.color;
}


export {
    getMethodColor,
    getCodeColor
}
