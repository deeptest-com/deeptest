import { Input, Output, EventEmitter, Component, OnInit, ElementRef, Inject } from '@angular/core';

import {GlobalState} from '../../../global.state';

import { CONSTANT } from '../../../utils/constant';

@Component({
  selector: '[table-row]',
  templateUrl: './table-row.html'
})
export class TableRowComponent implements OnInit {
  orgId: number;
  @Input()
  public model: any;
  @Input()
  public maxLevel: number;
  @Input()
  public keywords: string;

  counter = Array;

  constructor(private _state:GlobalState, private el: ElementRef) {
    this.orgId = CONSTANT.CURR_ORG_ID;

  }

  public ngOnInit(): void {

  }

}
