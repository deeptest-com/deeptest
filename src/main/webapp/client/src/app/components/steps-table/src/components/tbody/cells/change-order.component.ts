import {Component, Input, Output, EventEmitter, OnChanges, ChangeDetectionStrategy } from '@angular/core';

import { Grid } from '../../../lib/grid';
import { Row } from '../../../lib/data-set/row';
import { DataSource } from '../../../lib/data-source/data-source';

@Component({
  selector: 'ng2-st-tbody-change-order',
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `

    <a href="#" class="ng2-smart-action" (click)="onUp($event)">
      <span class="ion-arrow-up-a ionic-icon link near"></span>
    </a>
    <a href="#" class="ng2-smart-action" (click)="onDown($event)">
      <span class="ion-arrow-down-a ionic-icon link near"></span>
    </a>
    
  `,
})
export class TbodyChangeOrderComponent implements OnChanges {

  @Input() grid: Grid;
  @Input() row: Row;
  @Input() source: DataSource;

  @Input() upConfirm: EventEmitter<any>;
  @Input() downConfirm: EventEmitter<any>;

  onUp(event: any) {
    event.preventDefault();
    event.stopPropagation();

    // this.grid.up(this.grid.getNewRow(), this.upConfirm, this.row);
  }

  onDown(event: any) {
    event.preventDefault();
    event.stopPropagation();

    // this.grid.down(this.grid.getNewRow(), this.downConfirm, this.row);
  }

  ngOnChanges(){

  }

}
