import { Component, Input, EventEmitter, OnChanges } from '@angular/core';

import { Grid } from '../../../lib/grid';
import { Row } from '../../../lib/data-set/row';

import * as _ from 'lodash';

@Component({
  selector: 'ng2-st-tbody-create-cancel',
  template: `
    <a href="#" class="ng2-smart-action ng2-smart-action-edit-save" (click)="onSave($event)">
      <i class="fa fa-check"></i>
    </a>
    <a href="#" class="ng2-smart-action ng2-smart-action-edit-cancel" (click)="onCancelEdit($event)">
      <i class="fa fa-times"></i>
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

    if (!_.trim(this.row.getNewData().opt)) {
      confirm('操作步骤不能为空!');
      return false;
    }

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
