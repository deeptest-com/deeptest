import { Subject } from 'rxjs/Subject';
import { Observable } from 'rxjs/Observable';

export abstract class DataSource {

  protected onCreatedSource = new Subject<any>();
  protected onSavedSource = new Subject<any>();
  protected onRemovedSource = new Subject<any>();

  protected onChangedSource = new Subject<any>();

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
  onSaved(): Observable<any> {
    return this.onSavedSource.asObservable();
  }
  onRemoved(): Observable<any> {
    return this.onRemovedSource.asObservable();
  }

  onChanged(): Observable<any> {
    return this.onChangedSource.asObservable();
  }

  create(element: any, curr: any): Promise<any> {
    this.emitOnCreated(element);
    this.emitOnChanged('create');
    return Promise.resolve();
  }

  remove(element: any): Promise<any> {
    this.emitOnRemoved(element);
    this.emitOnChanged('remove');
    return Promise.resolve();
  }

  save(element: any, values: any): Promise<any> {
    this.emitOnSaved(element);
    this.emitOnChanged('save');
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
  protected emitOnSaved(element: any) {
    this.onSavedSource.next(element);
  }

  protected emitOnChanged(action: string) {
    this.getElements().then((elements) => this.onChangedSource.next({
      action: action,
      elements: elements
    }));
  }
}
