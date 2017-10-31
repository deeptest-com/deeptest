import { Component, Input, Output, EventEmitter } from '@angular/core';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';

@Component({
  selector: 'field-show-value',
  styleUrls: ['./value.scss'],
  templateUrl: './field-show-value.html',
})
export class FieldShowValueComponent {

  @Input() model: any = {};
  @Input() prop: string;
  @Input() valType: string;
  @Input() valFormat: string;

  public casePropertyMap: any = {};

  public constructor() {
    this.casePropertyMap = CONSTANT.CASE_PROPERTY_MAP;
  }

}
