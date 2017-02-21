import {Pipe, PipeTransform} from '@angular/core';

import {CONSTANT} from '../utils/constant';

@Pipe({name: 'orderStatus'})
export class OrderStatusPipe implements PipeTransform {
    transform(order: any) : string {
        if (!order) {
            return;
        }

        var s = order.status;
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


@Pipe({name: 'booleanToCn'})
export class BooleanToCn implements PipeTransform {
    transform(bl: any) : string {
        if (bl == true) {
            return '是';
        } else {
            return '否';
        }
    }
}

@Pipe({name: 'serviceName'})
export class ServiceNamePipe implements PipeTransform {
  transform(code: any) : string {
    var name;

    if (code === 'taxi') {
      name = '叫车';
    } else if (code === 'accommodation') {
      name = '住宿';
    } else if (code === 'food'){
      name = '餐饮';
    } else if (code === 'wifi'){
      name = 'WIFI';
    } else if (code === 'shopping'){
      name = '购物';
    } else if (code === 'print'){
      name = '打印';
    }

    return name;
  }
}
@Pipe({name: 'aroundName'})
export class AroundNamePipe implements PipeTransform {
  transform(code: any) : string {
    if (code === 'food') {
      name = '美食';
    } else if (code === 'accommodation') {
      name = '住宿';
    } else if (code === 'transportation'){
      name = '交通';
    } else if (code === 'tour'){
      name = '游览';
    } else if (code === 'shopping'){
      name = '购物';
    } else if (code === 'entertainment'){
      name = '娱乐';
    }

    return name;
  }
}
