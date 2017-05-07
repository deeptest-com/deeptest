import {Pipe, PipeTransform} from '@angular/core';

import {CONSTANT} from '../utils/constant';

@Pipe({name: 'markError'})
export class MarkErrorPipe implements PipeTransform {
    transform(text: string) : string {
        if (text.indexOf('Fail') > -1) {
          text = '<div class="Fail">' + text + '</div>'
        } else if (text.indexOf('Pass') > -1 || text.indexOf('Success') > -1) {
          text = '<div class="Pass">' + text + '</div>'
        }

        return text;
    }
}
