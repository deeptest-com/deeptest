import {requestMethodOpts} from "@/config/constant";

/**
 * 关于接口的一些定义和公共方法
 * */


/**
 * 获取接口方法的颜色
 * */
export const getMethodColor = (method:any) => {
    const item: any = requestMethodOpts.find((item: any) => {
        return item.value === method;
    });
    return item.color || '#04C495';
}
