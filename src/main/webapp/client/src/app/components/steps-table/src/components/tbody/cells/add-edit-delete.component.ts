import {Component, Input, Output, EventEmitter, OnChanges, ChangeDetectionStrategy } from '@angular/core';

import { Grid } from '../../../lib/grid';
import { Row } from '../../../lib/data-set/row';
import { DataSource } from '../../../lib/data-source/data-source';

@Component({
  selector: 'ng2-st-tbody-add-edit-delete',
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <a href="#" class="ng2-smart-action ng2-smart-action-edit-edit" (click)="onCreate($event)">
      <i class="ion-plus"></i>
    </a>
    
    <a href="#" class="ng2-smart-action ng2-smart-action-edit-edit" (click)="onEdit($event)">
      <i class="ion-edit"></i>
    </a>
    
    <a href="#" class="ng2-smart-action ng2-smart-action-delete-delete" (click)="onDelete($event)">
      <i class="ion-trash-a"></i>
    </a>
  `,
})
export class TbodyAddEditDeleteComponent implements OnChanges {

  @Input() grid: Grid;
  @Input() row: Row;
  @Input() source: DataSource;

  @Input() createConfirm: EventEmitter<any>;
  @Input() deleteConfirm: EventEmitter<any>;
  @Input() editConfirm: EventEmitter<any>;

  onCreate(event: any) {
    event.preventDefault();
    event.stopPropagation();

    this.grid.create(this.grid.getNewRow(), this.createConfirm, this.row);
  }

  onEdit(event: any) {
    event.preventDefault();
    event.stopPropagation();

    this.grid.edit(this.row);
  }

  onDelete(event: any) {
    event.preventDefault();
    event.stopPropagation();

    this.grid.delete(this.row, this.deleteConfirm);
  }

  ngOnChanges(){

  }
}
