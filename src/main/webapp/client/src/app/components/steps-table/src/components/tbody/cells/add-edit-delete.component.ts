import {Component, Input, Output, EventEmitter, OnChanges, ChangeDetectionStrategy } from '@angular/core';

import { Grid } from '../../../lib/grid';
import { Row } from '../../../lib/data-set/row';
import { DataSource } from '../../../lib/data-source/data-source';

@Component({
  selector: 'ng2-st-tbody-add-edit-delete',
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <a href="#" *ngIf="isActionAdd" class="ng2-smart-action ng2-smart-action-edit-edit"
       [innerHTML]="addRowButtonContent" (click)="onAdd($event)"></a>
    
    <a href="#" *ngIf="isActionEdit" class="ng2-smart-action ng2-smart-action-edit-edit"
        [innerHTML]="editRowButtonContent" (click)="onEdit($event)"></a>
    
    <a href="#" *ngIf="isActionDelete" class="ng2-smart-action ng2-smart-action-delete-delete"
        [innerHTML]="deleteRowButtonContent" (click)="onDelete($event)"></a>
  `,
})
export class TbodyAddEditDeleteComponent implements OnChanges {

  @Input() grid: Grid;
  @Input() row: Row;
  @Input() source: DataSource;
  @Input() deleteConfirm: EventEmitter<any>;
  @Input() editConfirm: EventEmitter<any>;

  @Output() add = new EventEmitter<any>();
  @Output() edit = new EventEmitter<any>();
  @Output() delete = new EventEmitter<any>();
  @Output() editRowSelect = new EventEmitter<any>();

  isActionAdd: boolean;
  isActionEdit: boolean;
  isActionDelete: boolean;
  addRowButtonContent: string;
  editRowButtonContent: string;
  deleteRowButtonContent: string;

  onAdd(event: any) {
    event.preventDefault();
    event.stopPropagation();

    this.editRowSelect.emit(this.row);

    if (this.grid.getSetting('mode') === 'external') {
      this.add.emit({
        source: this.source,
      });
    } else {
      this.grid.edit(this.row);
    }
  }

  // onAdd(event: any) {
  //   event.preventDefault();
  //   event.stopPropagation();
  //   if (this.grid.getSetting('mode') === 'external') {
  //     this.create.emit({
  //       source: this.source,
  //     });
  //   } else {
  //     this.grid.createFormShown = true;
  //   }
  // }

  onEdit(event: any) {
    event.preventDefault();
    event.stopPropagation();

    this.editRowSelect.emit(this.row);

    if (this.grid.getSetting('mode') === 'external') {
      this.edit.emit({
        data: this.row.getData(),
        source: this.source,
      });
    } else {
      this.grid.edit(this.row);
    }
  }

  onDelete(event: any) {
    event.preventDefault();
    event.stopPropagation();

    if (this.grid.getSetting('mode') === 'external') {
      this.delete.emit({
        data: this.row.getData(),
        source: this.source,
      });
    } else {
      this.grid.delete(this.row, this.deleteConfirm);
    }
  }

  ngOnChanges(){
    this.isActionAdd = this.grid.getSetting('actions.add');
    this.isActionEdit = this.grid.getSetting('actions.edit');
    this.isActionDelete = this.grid.getSetting('actions.delete');
    this.addRowButtonContent = this.grid.getSetting('add.addButtonContent');
    this.editRowButtonContent = this.grid.getSetting('edit.editButtonContent');
    this.deleteRowButtonContent = this.grid.getSetting('delete.deleteButtonContent');
  }
}
