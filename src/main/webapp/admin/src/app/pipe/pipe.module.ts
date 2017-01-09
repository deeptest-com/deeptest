import { NgModule }      from '@angular/core';

import { ImgPathPipe, ThumbPathPipe } from './img-path';
import { KeysPipe } from './keys';
import { DatePipe } from './date';
import { EventStatusPipe } from './event-status';

@NgModule({
  imports: [

  ],
  declarations: [ImgPathPipe, ThumbPathPipe, KeysPipe, DatePipe,EventStatusPipe],
  exports:      [ImgPathPipe, ThumbPathPipe, KeysPipe, DatePipe,EventStatusPipe],
})
export class PipeModule {
}
