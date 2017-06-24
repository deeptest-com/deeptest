import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { CellModule } from './components/cell/cell.module';
import { TBodyModule } from './components/tbody/tbody.module';

import { StepsTableComponent } from './steps-table.component';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    CellModule,
    TBodyModule
  ],
  declarations: [
    StepsTableComponent,
  ],
  exports: [
    StepsTableComponent,
  ],
})
export class StepsTableModule {
}
