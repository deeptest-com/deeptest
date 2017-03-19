import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

import { TableTreeComponent } from './src/table-tree.component';

export * from './src/table-tree.component';

@NgModule({
  imports: [CommonModule, RouterModule],
  declarations: [TableTreeComponent],
  exports: [TableTreeComponent],
  providers: []
})
export class TableTreeModule {
  static forRoot(): ModuleWithProviders {
    return {
      ngModule: TableTreeModule,
      providers: []
    };
  }
}
