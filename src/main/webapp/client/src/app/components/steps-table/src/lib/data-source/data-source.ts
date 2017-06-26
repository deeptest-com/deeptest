import { Subject } from 'rxjs/Subject';
import { Observable } from 'rxjs/Observable';

export abstract class DataSource {

  protected onCreatedSource = new Subject<any>();
  protected onSavedSource = new Subject<any>();
  protected onDeletedSource = new Subject<any>();

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
  onDeleted(): Observable<any> {
    return this.onDeletedSource.asObservable();
  }

  onChanged(): Observable<any> {
    return this.onChangedSource.asObservable();
  }

  up(element: any): Promise<any> {
    this.emitOnDeleted(element);
    this.emitOnChanged('up');
    return Promise.resolve();
  }
  down(element: any): Promise<any> {
    this.emitOnDeleted(element);
    this.emitOnChanged('down');
    return Promise.resolve();
  }

  create(element: any, curr: any): Promise<any> {
    this.emitOnCreated(element);
    this.emitOnChanged('create');
    return Promise.resolve();
  }

  delete(element: any): Promise<any> {
    this.emitOnDeleted(element);
    this.emitOnChanged('delete');
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
  protected emitOnDeleted(element: any) {
    this.onDeletedSource.next(element);
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
