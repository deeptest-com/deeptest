import {Injectable} from '@angular/core';

import {isPresent} from './slim-loading-bar.utils';
import {Subject} from 'rxjs/Subject';
import {Observable} from 'rxjs/Observable';

import { FieldType, InputTextEvent } from '../field-type';

@Injectable()
export class InputTextService {

    private val: string = '';

    private eventSource: Subject<InputTextEvent> = new Subject<InputTextEvent>();
    public events: Observable<InputTextEvent> = this.eventSource.asObservable();

    constructor() {}

    set val(v:string) {
      this.val = v;
    }

    get val():number {
      return this.val;
    }

}

