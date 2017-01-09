import { NgModule }      from '@angular/core';

import { ImgPathPipe } from './img-path';
import { KeysPipe } from './keys';
import { DatePipe } from './date';
import { EventStatusPipe } from './event-status';

@NgModule({
  imports: [

  ],
  declarations: [ImgPathPipe, KeysPipe, DatePipe,EventStatusPipe],
  exports:      [ImgPathPipe, KeysPipe, DatePipe,EventStatusPipe],
})
export class PipeModule {
}
