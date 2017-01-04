import { NgModule }      from '@angular/core';

import { KeysPipe } from './keys';
import { DatePipe } from './date';
import { EventStatusPipe } from './event-status';

@NgModule({
  imports: [
    
  ],
  declarations: [KeysPipe, DatePipe,EventStatusPipe],
  exports:      [ KeysPipe, DatePipe,EventStatusPipe],
})
export class PipeModule {
}
