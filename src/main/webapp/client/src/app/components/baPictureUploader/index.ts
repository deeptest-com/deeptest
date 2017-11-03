import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

import { NgUploaderModule } from 'ngx-uploader';
import { NgUploaderService } from 'ngx-uploader';
import {PipeModule} from '../../pipe/pipe.module';

export * from './baPictureUploader.component';

@NgModule({
  imports: [CommonModule, RouterModule, PipeModule, NgUploaderModule],
  declarations: [BaPictureUploader],
  exports: [BaPictureUploader],
  providers: []
})
export class BaPictureUploader {
  static forRoot(): ModuleWithProviders {
    return {
      ngModule: BaPictureUploader,
      providers: [NgUploaderService]
    };
  }
}
