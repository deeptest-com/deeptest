import { NgModule }      from '@angular/core';

import { ImgPathPipe, ThumbPathPipe } from './img-path';
import { MapToArrayPipe } from './map-to-array';
import { DatePipe } from './date';
import { ModelStatusPipe } from './model-status';
import { MarkErrorPipe } from './mark-error';

@NgModule({
  imports: [ ],
  declarations: [ImgPathPipe, ThumbPathPipe, MapToArrayPipe, DatePipe, ModelStatusPipe, MarkErrorPipe],
  exports:      [ImgPathPipe, ThumbPathPipe, MapToArrayPipe, DatePipe, ModelStatusPipe, MarkErrorPipe],
})
export class PipeModule {
}
