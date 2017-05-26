import { NgModule }      from '@angular/core';

import { ImgPathPipe, ThumbPathPipe } from './img-path';
import { MapToArrayPipe } from './map-to-array';
import { DatePipe } from './date';
import { ModelStatusPipe } from './model-status';
import { MarkErrorPipe } from './mark-error';
import { FieldTypePipe, FieldApplyToPipe, FieldFormatPipe, TrueOrFalsePipe, DisableOrNotPipe } from './field-property';

@NgModule({
  imports: [ ],
  declarations: [ImgPathPipe, ThumbPathPipe, MapToArrayPipe, DatePipe, ModelStatusPipe, MarkErrorPipe,
    FieldTypePipe, FieldApplyToPipe, FieldFormatPipe, TrueOrFalsePipe, DisableOrNotPipe],
  exports:      [ImgPathPipe, ThumbPathPipe, MapToArrayPipe, DatePipe, ModelStatusPipe, MarkErrorPipe,
    FieldTypePipe, FieldApplyToPipe, FieldFormatPipe, TrueOrFalsePipe, DisableOrNotPipe],
})
export class PipeModule {

}
