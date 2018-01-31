import {Component, Input, Output, EventEmitter, OnChanges, ChangeDetectionStrategy } from '@angular/core';

import { Grid } from '../../../lib/grid';
import { Row } from '../../../lib/data-set/row';
import { DataSource } from '../../../lib/data-source/data-source';

@Component({
  selector: 'ng2-st-tbody-change-order',
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `

    <span *ngIf="isFirst" href="#" class="ng2-smart-action link transparent">
      <span class="fa fa-arrow-up ionic-icon near"></span>
    </span>
    <span *ngIf="!isFirst" href="#" class="ng2-smart-action link" (click)="onUp($event)">
      <span class="fa fa-arrow-up ionic-icon near"></span>
    </span>
    
    <span *ngIf="isLast" href="#" class="ng2-smart-action link transparent">
      <span class="fa fa-arrow-down ionic-icon near"></span>
    </span>
    <span *ngIf="!isLast" href="#" class="ng2-smart-action link" (click)="onDown($event)">
      <span class="fa fa-arrow-down ionic-icon near"></span>
    </span>
    
  `,
})
export class TbodyChangeOrderComponent implements OnChanges {

  @Input() grid: Grid;
  @Input() row: Row;
  @Input() source: DataSource;

  @Input() upConfirm: EventEmitter<any>;
  @Input() downConfirm: EventEmitter<any>;

  @Input() isFirst: boolean;
  @Input() isLast: boolean;

  onUp(event: any) {
    event.preventDefault();
    event.stopPropagation();

    this.grid.up(this.row, this.upConfirm);
  }

  onDown(event: any) {
    event.preventDefault();
    event.stopPropagation();

    this.grid.down(this.row, this.downConfirm);
  }

  ngOnChanges(){

  }

}
