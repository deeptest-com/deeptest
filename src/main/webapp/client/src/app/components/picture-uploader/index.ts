import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

import { FileUploadModule } from 'ng2-file-upload';
import {PipeModule} from '../../pipe/pipe.module';

import { PictureUploaderComponent } from './src/picture-uploader.component';

export * from './src/picture-uploader.component';

@NgModule({
  imports: [CommonModule, RouterModule, PipeModule, FileUploadModule],
  declarations: [PictureUploaderComponent],
  exports: [PictureUploaderComponent],
  providers: []
})
export class PictureUploaderModule {
  static forRoot(): ModuleWithProviders {
    return {
      ngModule: PictureUploaderModule,
      providers: []
    };
  }
}
