import {Pipe, PipeTransform} from '@angular/core';

import {CONSTANT} from '../utils/constant';

@Pipe({name: 'isActive'})
export class IsActivePipe implements PipeTransform {
    transform(isActive: boolean) : string {
        var status: string;

        if (isActive) {
          status = '活动';
        } else {
          status = '关闭';
        }

        return status;
    }
}
