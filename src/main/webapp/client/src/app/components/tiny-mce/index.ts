import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

import { TinyMCEComponent } from './src/tiny-mce.component';

export * from './src/tiny-mce.component';

@NgModule({
  imports: [CommonModule, RouterModule],
  declarations: [TinyMCEComponent],
  exports: [TinyMCEComponent],
  providers: []
})
export class TinyMCEModule {
  static forRoot(): ModuleWithProviders {
    return {
      ngModule: TinyMCEModule,
      providers: []
    };
  }
}
