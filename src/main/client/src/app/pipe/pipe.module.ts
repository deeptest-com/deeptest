import { NgModule }      from '@angular/core';

import { MsgReadPipe } from './msg-read';
import { ImgPathPipe, ThumbPathPipe } from './img-path';
import { MapToArrayPipe } from './map-to-array';
import { TimePassedPipe } from './date';
import { ModelStatusPipe } from './model-status';
import { ExeStatusPipe } from './exe-status';
import { PercentPipe } from './percent';
import { MarkErrorPipe } from './mark-error';
import { FieldTypePipe, FieldApplyToPipe, FieldFormatPipe, TrueOrFalsePipe, DisableOrNotPipe } from './field-property';

@NgModule({
  imports: [],
  declarations: [MsgReadPipe, ExeStatusPipe, ImgPathPipe, ThumbPathPipe, MapToArrayPipe, TimePassedPipe, ModelStatusPipe, MarkErrorPipe,
    FieldTypePipe, FieldApplyToPipe, FieldFormatPipe, TrueOrFalsePipe, DisableOrNotPipe, PercentPipe],
  exports:      [MsgReadPipe, ExeStatusPipe, ImgPathPipe, ThumbPathPipe, MapToArrayPipe, TimePassedPipe, ModelStatusPipe, MarkErrorPipe,
    FieldTypePipe, FieldApplyToPipe, FieldFormatPipe, TrueOrFalsePipe, DisableOrNotPipe],
})
export class PipeModule {

}
