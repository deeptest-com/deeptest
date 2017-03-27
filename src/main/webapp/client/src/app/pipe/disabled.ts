import {Pipe, PipeTransform} from '@angular/core';

import {CONSTANT} from '../utils/constant';

@Pipe({name: 'disabled'})
export class DisabledPipe implements PipeTransform {
    transform(disabled: boolean) : string {
        var status: string;

        if (disabled) {
          status = '禁用';
        } else {
          status = '启动';
        }

        return status;
    }
}
