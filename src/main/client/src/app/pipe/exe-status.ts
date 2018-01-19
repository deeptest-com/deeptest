import {Pipe, PipeTransform} from '@angular/core';

import {CONSTANT} from '../utils/constant';

@Pipe({name: 'exeStatus'})
export class ExeStatusPipe implements PipeTransform {
    transform(str: string) : string {
        var status: string;

        if (str == 'not_start') {
            status = '未开始';
        } else if (str == 'in_progress') {
            status = '执行中';
        } else if (str == 'end') {
          status = '已完成';
        }

        return status;
    }
}
