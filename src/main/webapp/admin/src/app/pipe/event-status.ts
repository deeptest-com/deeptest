import {Pipe, PipeTransform} from '@angular/core';

import {CONSTANT} from '../utils/constant';

@Pipe({name: 'eventStatus'})
export class EventStatusPipe implements PipeTransform {
    transform(s: string) : string {

        return CONSTANT.EventStatus[s];
    }
}
