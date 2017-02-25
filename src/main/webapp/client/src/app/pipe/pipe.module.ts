import { NgModule }      from '@angular/core';

import { ImgPathPipe, ThumbPathPipe } from './img-path';
import { KeysPipe } from './keys';
import { DatePipe } from './date';
import { ModelStatusPipe } from './model-status';
import { EventStatusPipe } from './event-status';

@NgModule({
  imports: [ ],
  declarations: [ImgPathPipe, ThumbPathPipe, KeysPipe, DatePipe,ModelStatusPipe, EventStatusPipe],
  exports:      [ImgPathPipe, ThumbPathPipe, KeysPipe, DatePipe,ModelStatusPipe, EventStatusPipe],
})
export class PipeModule {
}
