import {Injectable} from '@angular/core';

import {Subject} from 'rxjs/Subject';
import {Observable} from 'rxjs/Observable';

import { FieldType, CustomFieldDefinition } from '../field.definitions';
import { CustomFieldModel } from '../field.models';
import { FieldChangedEvent } from '../field.events';

@Injectable()
export class InputTextService {

    private value: string = '';

    private eventSource: Subject<FieldChangedEvent> = new Subject<FieldChangedEvent>();
    public events: Observable<FieldChangedEvent> = this.eventSource.asObservable();

    constructor() {}

}

