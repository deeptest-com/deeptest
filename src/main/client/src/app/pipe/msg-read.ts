import {Pipe, PipeTransform} from '@angular/core';

import {CONSTANT} from '../utils/constant';

@Pipe({name: 'msgRead'})
export class MsgReadPipe implements PipeTransform {
    transform(isRead: boolean) : string {
        var status: string;

        if (isRead) {
            status = '已读';
        } else {
            status = '未读';
        }

        return status;
    }
}
