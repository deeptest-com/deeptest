import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

import { NgUploaderService } from 'ngx-uploader';
import { BaPictureUploader } from './baPictureUploader.component';

export * from './baPictureUploader.component';

@NgModule({
  imports: [CommonModule, RouterModule],
  declarations: [BaPictureUploader],
  exports: [BaPictureUploader],
  providers: []
})
export class NgUploaderModule {
  static forRoot(): ModuleWithProviders {
    return {
      ngModule: NgUploaderModule,
      providers: [NgUploaderService]
    };
  }
}
