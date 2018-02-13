import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule } from '@angular/common';
import {FormsModule} from "@angular/forms";
import { RouterModule } from '@angular/router';

import { DropdownOptionsComponent } from './src/dropdown-options.component';

export * from './src/dropdown-options.component';

@NgModule({
  imports: [CommonModule, RouterModule, FormsModule],
  declarations: [DropdownOptionsComponent],
  exports: [DropdownOptionsComponent],
  providers: []
})
export class DropdownOptionsModule {
  static forRoot(): ModuleWithProviders {
    return {
      ngModule: DropdownOptionsModule,
      providers: []
    };
  }
}
