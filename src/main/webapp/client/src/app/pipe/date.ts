import {Pipe, PipeTransform} from '@angular/core';

import {CONSTANT} from '../utils/constant';

@Pipe({name: 'myDatePath'})
export class DatePipe implements PipeTransform {
    transform(str: string, external: any) : string {
        return str;
    }
}
