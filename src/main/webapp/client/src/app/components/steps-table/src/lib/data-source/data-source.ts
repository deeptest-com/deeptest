import { Subject } from 'rxjs/Subject';
import { Observable } from 'rxjs/Observable';

export abstract class DataSource {

  protected onChangedSource = new Subject<any>();
  protected onCreatedSource = new Subject<any>();
  protected onUpdatedSource = new Subject<any>();
  protected onRemovedSource = new Subject<any>();

  abstract getAll(): Promise<any>;
  abstract getElements(): Promise<any>;

  refresh() {
    this.emitOnChanged('refresh');
  }

  load(data: Array<any>): Promise<any> {
    this.emitOnChanged('load');
    return Promise.resolve();
  }

  onCreated(): Observable<any> {
    return this.onCreatedSource.asObservable();
  }
  onUpdated(): Observable<any> {
    return this.onUpdatedSource.asObservable();
  }
  onRemoved(): Observable<any> {
    return this.onRemovedSource.asObservable();
  }

  onChanged(): Observable<any> {
    return this.onChangedSource.asObservable();
  }

  create(element: any): Promise<any> {
    this.emitOnCreated(element);
    this.emitOnChanged('create');
    return Promise.resolve();
  }

  remove(element: any): Promise<any> {
    this.emitOnRemoved(element);
    this.emitOnChanged('remove');
    return Promise.resolve();
  }

  update(element: any, values: any): Promise<any> {
    this.emitOnUpdated(element);
    this.emitOnChanged('update');
    return Promise.resolve();
  }

  empty(): Promise<any> {
    this.emitOnChanged('empty');
    return Promise.resolve();
  }

  protected emitOnCreated(element: any) {
    this.onCreatedSource.next(element);
  }
  protected emitOnRemoved(element: any) {
    this.onRemovedSource.next(element);
  }
  protected emitOnUpdated(element: any) {
    this.onUpdatedSource.next(element);
  }

  protected emitOnChanged(action: string) {
    this.getElements().then((elements) => this.onChangedSource.next({
      action: action,
      elements: elements
    }));
  }
}
