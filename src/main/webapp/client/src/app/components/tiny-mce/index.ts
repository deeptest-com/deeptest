import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule } from '@angular/common';
import {FormsModule} from "@angular/forms";
import { RouterModule } from '@angular/router';

import { TinyMCEComponent } from './src/tiny-mce.component';

export * from './src/tiny-mce.component';

@NgModule({
  imports: [CommonModule, RouterModule, FormsModule],
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
