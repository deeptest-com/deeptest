import { Component, Input, EventEmitter, OnChanges } from '@angular/core';

import { Grid } from '../../../lib/grid';
import { Row } from '../../../lib/data-set/row';

@Component({
  selector: 'ng2-st-tbody-create-cancel',
  template: `
    <a href="#" class="ng2-smart-action ng2-smart-action-edit-save" (click)="onSave($event)">
      <i class="ion-checkmark"></i>
    </a>
    <a href="#" class="ng2-smart-action ng2-smart-action-edit-cancel" (click)="onCancelEdit($event)">
      <i class="ion-close"></i>
    </a>
  `,
})
export class TbodyCreateCancelComponent implements OnChanges {

  @Input() grid: Grid;
  @Input() row: Row;
  @Input() saveConfirm: EventEmitter<any>;
  @Input() deleteConfirm: EventEmitter<any>;

  onSave(event: any) {
    event.preventDefault();
    event.stopPropagation();

    this.grid.save(this.row, this.saveConfirm);
  }

  onCancelEdit(event: any) {
    event.preventDefault();
    event.stopPropagation();

    if (this.row.isNew) {
      this.grid.delete(this.row, this.deleteConfirm);
    } else {
      this.row.isInEditing = false;
    }
  }

  ngOnChanges() {

  }
}
