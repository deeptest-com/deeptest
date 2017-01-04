import {Pipe, PipeTransform} from '@angular/core';

import {CONSTANT} from '../utils/constant';

@Pipe({name: 'orderStatus'})
export class OrderStatusPipe implements PipeTransform {
    transform(s: string) : string {
        var status;
        
        if (s === 'INIT') {
            status = '未支付';
        } else if (s === 'PAYING') {
            status = '支付中';
        } else if (s === 'PAID'){
            status = '已支付';
        } else if (s === 'SHIPPING'){
            status = '发货中';
        } else if (s === 'RECEIVED'){
            status = '已收货';
        } else if (s === 'RATED'){
            status = '已评价';
        } else if (s === 'CANCEL'){
            status = '已取消';
        } else if (s === 'PAY_FEATURE'){
            status = '支付错误';
        } else if (s === 'SHIPPING_FEATURE'){
            status = '快递问题';
        }

        return status;
    }
}