import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

import { NgUploaderModule } from 'ngx-uploader';
import {PipeModule} from '../../pipe/pipe.module';

import { BaPictureUploaderComponent } from './src/baPictureUploader.component';

export * from './src/baPictureUploader.component';

@NgModule({
  imports: [CommonModule, RouterModule, PipeModule, NgUploaderModule],
  declarations: [BaPictureUploaderComponent],
  exports: [BaPictureUploaderComponent],
  providers: []
})
export class BaPictureUploaderModule {
  static forRoot(): ModuleWithProviders {
    return {
      ngModule: BaPictureUploaderModule,
      providers: []
    };
  }
}
