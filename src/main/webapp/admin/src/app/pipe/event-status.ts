import {Pipe, PipeTransform} from '@angular/core';

import {CONSTANT} from '../utils/constant';

@Pipe({name: 'eventStatus'})
export class EventStatusPipe implements PipeTransform {
    transform(s: string) : string {
        var status: string;
        
        if (s === 'not_start') {
            status = '未开始';
        } else if (s === 'register') {
            status = '报名中';
        } else if (s === 'sign'){
            status = '签到中';
        } else if (s === 'in_progress'){
            status = '进行中';
        } else if (s === 'end'){
            status = '已结束';
        } else if (s === 'cancel'){
            status = '已取消';
        }

        return status;
    }
}
