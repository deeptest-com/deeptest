import {Pipe, PipeTransform} from '@angular/core';

import {CONSTANT} from '../utils/constant';

@Pipe({name: 'percentNumb'})
export class PercentPipe implements PipeTransform {
    transform(numb: number) : string {
        return (numb * 100).toFixed(2) + '%';
    }
}
