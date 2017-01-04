import {Pipe, PipeTransform} from '@angular/core';
import {DatePipe} from '@angular/common';

import {CONSTANT} from '../utils/constant';

@Pipe({name: 'iosDate'})
export class IosDatePipe implements PipeTransform {
    private datePipe: DatePipe = new DatePipe();

    transform(isoDate: string, pattern?: string): string {
      const date = new Date(isoDate);
      return this.datePipe.transform(date, pattern);
    }
}
