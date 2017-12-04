import { Subject } from 'rxjs/Subject';
import { Observable } from 'rxjs/Observable';
import { EventEmitter } from '@angular/core';

import { Deferred, getDeepFromObject } from './helpers';
import { Column } from './data-set/column';
import { Row } from './data-set/row';
import { DataSet } from './data-set/data-set';
import { DataSource } from './data-source/data-source';

export class Grid {

  source: DataSource;
  settings: any;
  dataSet: DataSet;

  constructor(source: DataSource, settings: any) {
    this.setSettings(settings);
    this.setSource(source);
  }

  getNewRow(): Row {
    return this.dataSet.newRow;
  }

  setSettings(settings: Object) {
    this.settings = settings;
    this.dataSet = new DataSet([], this.getSetting('columns'));

    if (this.source) {
      this.source.refresh();
    }
  }

  getDataSet(): DataSet {
    return this.dataSet;
  }

  setSource(source: DataSource) {
    this.source = this.prepareSource(source);

    this.source.onCreated().subscribe((data) => {
      console.log('source.onCreated');
    });
    this.source.onDeleted().subscribe((data) => {
      console.log('source.onDeleted');
    });
    this.source.onSaved().subscribe((data) => {
      console.log('source.onSaved');

      const changedRow = this.dataSet.findRowByData(data);
      changedRow.setData(data);
    });

    this.source.onChanged().subscribe((changes) => {
      console.log('source.onChanged');

      this.processDataChange(changes);
    });
  }

  getSetting(name: string, defaultValue?: any): any {
    return getDeepFromObject(this.settings, name, defaultValue);
  }

  getColumns(): Array<Column> {
    return this.dataSet.getColumns();
  }

  getRows(): Array<Row> {
    return this.dataSet.getRows();
  }

  edit(row: Row) {
    row.isInEditing = true;
  }

  up(curr: Row, confirmEmitter: EventEmitter<any>) {
    const deferred = new Deferred();
    deferred.promise.then(() => {
      this.source.up(curr.getData());
    }).catch((err) => {});

    confirmEmitter.emit({
      data: curr.getData(),
      source: this.source,
      confirm: deferred
    });

  }
  down(curr: Row, confirmEmitter: EventEmitter<any>) {
    const deferred = new Deferred();
    deferred.promise.then((newData) => {
      this.source.down(curr.getData());
    }).catch((err) => {});

    confirmEmitter.emit({
      data: curr.getData(),
      source: this.source,
      confirm: deferred
    });

  }

  create(row: Row, confirmEmitter: EventEmitter<any>, curr?: Row) {
    const deferred = new Deferred();
    deferred.promise.then((newData) => {
      newData = newData ? newData : row.getNewData();

      this.source.create(newData, curr?curr.getData(): null).then(() => {
        let newRow = this.dataSet.findRowByData(newData);

        newRow.isInEditing = true;
        newRow.isNew = true;
      });
    }).catch((err) => {
      // doing nothing
    });

    confirmEmitter.emit({
      newData: row.getNewData(),
      source: this.source,
      confirm: deferred,
    });

  }

  save(row: Row, confirmEmitter: EventEmitter<any>) {
    const deferred = new Deferred();
    deferred.promise.then((newData) => {
      newData = newData ? newData : row.getNewData();

      this.source.save(row.getData(), newData).then(() => {
        row.isInEditing = false;
      });
    }).catch((err) => {
      // doing nothing
    });

    confirmEmitter.emit({
      data: row.getData(),
      newData: row.getNewData(),
      source: this.source,
      confirm: deferred,
    });
  }

  delete(row: Row, confirmEmitter: EventEmitter<any>) {
    const deferred = new Deferred();
    deferred.promise.then(() => {
      this.source.delete(row.getData());
    }).catch((err) => {
      // doing nothing
    });

    if(!!confirmEmitter) {
      confirmEmitter.emit({
        data: row.getData(),
        source: this.source,
        confirm: deferred,
      });
    }

  }

  processDataChange(changes: any) {
    if (this.shouldProcessChange(changes)) {
      this.dataSet.setData(changes['elements']);
    }
  }

  shouldProcessChange(changes: any): boolean {
    if (['up', 'down', 'create', 'delete', 'save', 'refresh', 'load'].indexOf(changes['action']) !== -1) {
      return true;
    }

    return false;
  }

  prepareSource(source: any): DataSource {

    source.refresh();
    return source;
  }

}
