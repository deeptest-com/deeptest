import { Component, Input, Output, SimpleChange, EventEmitter, OnChanges } from '@angular/core';

import { Grid } from './lib/grid';
import { DataSource } from './lib/data-source/data-source';
import { Row } from './lib/data-set/row';
import { deepExtend } from './lib/helpers';
import { LocalDataSource } from './lib/data-source/local/local.data-source';

@Component({
  selector: 'steps-table',
  styleUrls: ['./steps-table.component.scss'],
  templateUrl: './steps-table.component.html',
})
export class StepsTableComponent implements OnChanges {

  @Input() source: any;
  @Input() settings: Object = {};

  @Output() custom = new EventEmitter<any>();

  @Output() upConfirm = new EventEmitter<any>();
  @Output() downConfirm = new EventEmitter<any>();

  @Output() createConfirm = new EventEmitter<any>();
  @Output() saveConfirm = new EventEmitter<any>();
  @Output() deleteConfirm = new EventEmitter<any>();

  grid: Grid;
  defaultSettings: Object = {
    columns: {}
  };

  ngOnChanges(changes: { [propertyName: string]: SimpleChange }) {
    if (this.grid) {

      if (changes['settings']) {
        this.grid.setSettings(this.prepareSettings());
      }
      if (changes['source']) {
        this.source = this.prepareSource();

        this.grid.setSource(this.source);
      }
    } else {
      this.initGrid();
    }
  }

  initGrid() {
    this.source = this.prepareSource();

    this.grid = new Grid(this.source, this.prepareSettings());
  }

  prepareSource(): DataSource {
    if (this.source instanceof DataSource) {
      return this.source;
    } else if (this.source instanceof Array) {
      return new LocalDataSource(this.source);
    }

    return new LocalDataSource();
  }

  prepareSettings(): Object {
    return deepExtend({}, this.defaultSettings, this.settings);
  }

}
