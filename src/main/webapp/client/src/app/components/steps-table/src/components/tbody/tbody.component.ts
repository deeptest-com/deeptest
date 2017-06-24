import {Component, Input, Output, EventEmitter, } from '@angular/core';

import { Grid } from '../../lib/grid';
import { Row } from '../../lib/data-set/row';
import { DataSource } from '../../lib/data-source/data-source';
import {Column} from "../../lib/data-set/column";

@Component({
  selector: '[ng2-st-tbody]',
  styleUrls: ['./tbody.component.scss'],
  templateUrl: './tbody.component.html',
})
export class Ng2SmartTableTbodyComponent {

  @Input() grid: Grid;
  @Input() source: DataSource;
  @Input() rowClassFunction: Function;

  @Input() upConfirm: EventEmitter<any>;
  @Input() downConfirm: EventEmitter<any>;

  @Input() createConfirm: EventEmitter<any>;
  @Input() saveConfirm: EventEmitter<any>;
  @Input() deleteConfirm: EventEmitter<any>;

  ngOnChanges() {

  }
}
