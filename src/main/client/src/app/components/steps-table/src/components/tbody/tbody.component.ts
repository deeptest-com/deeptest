import {Component, Input, Output, OnInit, EventEmitter, } from '@angular/core';

import { Grid } from '../../lib/grid';
import { Row } from '../../lib/data-set/row';
import { DataSource } from '../../lib/data-source/data-source';
import {Column} from "../../lib/data-set/column";

@Component({
  selector: '[ng2-st-tbody]',
  styleUrls: ['./tbody.component.scss'],
  templateUrl: './tbody.component.html',
})
export class Ng2SmartTableTbodyComponent implements OnInit {

  @Input() grid: Grid;
  @Input() source: DataSource;
  @Input() rowClassFunction: Function;

  @Input() upConfirm: EventEmitter<any>;
  @Input() downConfirm: EventEmitter<any>;

  @Input() createConfirm: EventEmitter<any>;
  @Input() saveConfirm: EventEmitter<any>;
  @Input() deleteConfirm: EventEmitter<any>;

  ngOnInit() {
    console.log('------', this.grid);
  }
  ngOnChanges() {

  }

  onCreate(event: any) {
    event.preventDefault();
    event.stopPropagation();

    this.grid.create(this.grid.getNewRow(), this.createConfirm);
  }

  onEdit(row: any) {
    event.preventDefault();
    event.stopPropagation();

    if (!this.grid.settings.canEdit) {
      return;
    }

    this.grid.edit(row);
  }

  onHoverRow($event, row):void {
    row.hover = true;
  }
  onOutRow($event, row):void {
    row.hover = false;
  }

}
