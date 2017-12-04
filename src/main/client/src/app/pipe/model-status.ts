import {Pipe, PipeTransform} from '@angular/core';

import {CONSTANT} from '../utils/constant';

@Pipe({name: 'modelStatus'})
export class ModelStatusPipe implements PipeTransform {
    transform(disabled: boolean) : string {
        var status: string;

        if (disabled) {
            status = '禁用';
        } else {
            status = '启用';
        }

        return status;
    }
}
