import {Pipe, PipeTransform} from '@angular/core';

import {CONSTANT} from '../utils/constant';

@Pipe({name: 'fieldType'})
export class FieldTypePipe implements PipeTransform {
  map: any = {'number': '数字', 'string': '字符串', 'text': '多行文本',
              'radio': '单选按钮', 'checkbox': '多选框', 'dropdown': '下拉菜单', 'multi_select': '多选菜单',
              'date': '日期', 'url': '网址',
              'user': '用户', 'version': '版本',
              'steps': '测试步骤', 'results': '测试结果'
             };

  transform(s: string) : string {
      return this.map[s];
  }
}

@Pipe({name: 'fieldApplyTo'})
export class FieldApplyToPipe implements PipeTransform {
  map: any = {'test_case': '测试用例', 'test_result': '测试结果'};

  transform(s: string) : string {
    return this.map[s];
  }
}

@Pipe({name: 'fieldFormat'})
export class FieldFormatPipe implements PipeTransform {
  map: any = {'rich_text': '富文本', 'plain_text': '纯文本'};

  transform(s: string) : string {
    if (!s) {
      return 'N/A';
    }
    return this.map[s];
  }
}

@Pipe({name: 'trueOrFalse'})
export class TrueOrFalsePipe implements PipeTransform {
  map: any = {'true': '是', 'false': '否'};

  transform(b: boolean) : string {
    if (!b) {
      return '否';
    }

    return this.map['' + b];
  }
}

@Pipe({name: 'disableOrNot'})
export class DisableOrNotPipe implements PipeTransform {
  map: any = {'true': '禁用', 'false': '启动'};

  transform(disabled: boolean) : string {
    return this.map['' + disabled];
  }
}
