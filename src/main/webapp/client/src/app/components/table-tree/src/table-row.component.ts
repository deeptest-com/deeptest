import { Input, Output, EventEmitter, Component, OnInit, ElementRef, Inject } from '@angular/core';

import {GlobalState} from '../../../global.state';

import { CONSTANT } from '../../../utils/constant';

@Component({
  selector: '[table-row]',
  template: require('./table-row.html')
})
export class TableRowComponent implements OnInit {
  @Input()
  public model: any;
  @Input()
  public maxLevel: number;
  @Input()
  public keywords: string;

  counter = Array;

  constructor(private _state:GlobalState, private el: ElementRef) {
    let that = this;

  }

  public ngOnInit(): void {

  }

}
