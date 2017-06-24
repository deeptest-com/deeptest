import { Subject } from 'rxjs/Subject';
import { Observable } from 'rxjs/Observable';
import { EventEmitter } from '@angular/core';

import { Deferred, getDeepFromObject } from './helpers';
import { Column } from './data-set/column';
import { Row } from './data-set/row';
import { DataSet } from './data-set/data-set';
import { DataSource } from './data-source/data-source';

export class Grid {

  createFormShown: boolean = false;

  source: DataSource;
  settings: any;
  dataSet: DataSet;

  onSelectRowSource = new Subject<any>();

  constructor(source: DataSource, settings: any) {
    this.setSettings(settings);
    this.setSource(source);
  }

  isMultiSelectVisible(): boolean {
    return this.getSetting('selectMode') === 'multi';
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
    this.source.onRemoved().subscribe((data) => {
      console.log('source.onRemoved');
    });

    this.source.onUpdated().subscribe((data) => {
      console.log('source.onUpdated');

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

  selectRow(row: Row) {
    this.dataSet.selectRow(row);
  }

  multipleSelectRow(row: Row) {
    this.dataSet.multipleSelectRow(row);
  }

  onSelectRow(): Observable<any> {
    return this.onSelectRowSource.asObservable();
  }

  edit(row: Row) {
    row.isInEditing = true;
  }

  up(row: Row, confirmEmitter: EventEmitter<any>, curr?: Row) {
    console.log(row, curr);

    const deferred = new Deferred();
    deferred.promise.then((newData) => {
      // newData = newData ? newData : row.getNewData();
      //
      // this.source.create(newData).then(() => {
      //   this.dataSet.createNewRow(curr);
      // });
    }).catch((err) => {});

    confirmEmitter.emit({
      newData: row.getNewData(),
      source: this.source,
      confirm: deferred
    });

  }
  down(row: Row, confirmEmitter: EventEmitter<any>, curr?: Row) {
    console.log(row, curr);

    const deferred = new Deferred();
    deferred.promise.then((newData) => {
      // newData = newData ? newData : row.getNewData();
      //
      // this.source.create(newData).then(() => {
      //   this.dataSet.createNewRow(curr);
      // });
    }).catch((err) => {});

    confirmEmitter.emit({
      newData: row.getNewData(),
      source: this.source,
      confirm: deferred
    });

  }

  create(row: Row, confirmEmitter: EventEmitter<any>, curr?: Row) {
    console.log(row, curr);

    const deferred = new Deferred();
    deferred.promise.then((newData) => {
      newData = newData ? newData : row.getNewData();

      this.source.create(newData).then(() => {
        this.dataSet.createNewRow(curr);
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
      if (deferred.resolve.skipEdit) {
        row.isInEditing = false;
      } else {
        this.source.update(row.getData(), newData).then(() => {
          row.isInEditing = false;
        });
      }
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
      this.source.remove(row.getData());
    }).catch((err) => {
      // doing nothing
    });

    confirmEmitter.emit({
      data: row.getData(),
      source: this.source,
      confirm: deferred,
    });

  }

  processDataChange(changes: any) {
    if (this.shouldProcessChange(changes)) {
      this.dataSet.setData(changes['elements']);
    }
  }

  shouldProcessChange(changes: any): boolean {
    if (['create', 'remove', 'refresh', 'load'].indexOf(changes['action']) !== -1) {
      return true;
    }

    return false;
  }

  prepareSource(source: any): DataSource {

    source.refresh();
    return source;
  }

}
