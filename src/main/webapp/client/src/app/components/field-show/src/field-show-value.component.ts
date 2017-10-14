import { Component, Input, Output, EventEmitter } from '@angular/core';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';

@Component({
  selector: 'field-show-value',
  template: `
    <span *ngIf="casePropertyMap[prop]">{{casePropertyMap[prop][model[prop]]}}</span>
    <span *ngIf="!casePropertyMap[prop]">{{model[prop]}}</span>
  `,
})
export class FieldShowValueComponent {

  @Input() model: any = {};
  @Input() prop: string;

  public casePropertyMap: any = {};

  public constructor() {
    this.casePropertyMap = CONSTANT.CASE_PROPERTY_MAP;
  }

}
