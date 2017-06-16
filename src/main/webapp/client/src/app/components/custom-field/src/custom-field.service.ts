import { Subject, Observable } from 'rxjs/Rx';
import { Injectable, Inject, ElementRef } from '@angular/core';

import { InputTextEvent } from './field.events';

@Injectable()
export class CustomFieldService {
  public nodeMoved$: Subject<InputTextEvent> = new Subject<InputTextEvent>();

  public constructor() {

  }

}
