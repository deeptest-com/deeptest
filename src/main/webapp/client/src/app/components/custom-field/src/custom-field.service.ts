import { Subject, Observable } from 'rxjs/Rx';
import { Injectable, Inject, ElementRef } from '@angular/core';

import { FieldChangedEvent } from './field.events';

@Injectable()
export class CustomFieldService {
  public nodeMoved$: Subject<FieldChangedEvent> = new Subject<FieldChangedEvent>();

  public constructor() {

  }

}
