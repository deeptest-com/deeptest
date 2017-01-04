import {Pipe, PipeTransform} from '@angular/core';

import {CONSTANT} from '../utils/constant';

@Pipe({name: 'cny'})
export class CurrencyCnyPipe implements PipeTransform {
    transform(str: string) : string {
        if (!str) {
            return '';
        } 
        return str.replace('CNY', '￥ ');
    }
}