import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

import { ModalModule } from 'ng2-bootstrap';

import { PopDialogComponent } from './src/pop-dialog.component';

export * from './src/pop-dialog.component';

@NgModule({
  imports: [CommonModule, RouterModule, ModalModule],
  declarations: [PopDialogComponent],
  exports: [PopDialogComponent],
  providers: []
})
export class PopDialogModule {
  static forRoot(): ModuleWithProviders {
    return {
      ngModule: PopDialogModule,
      providers: []
    };
  }
}
