import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { CellModule } from '../cell/cell.module';

import { Ng2SmartTableTbodyComponent } from './tbody.component';
import { TbodyChangeOrderComponent } from './cells/change-order.component';
import { TbodyCreateCancelComponent } from './cells/create-cancel.component';
import { TbodyAddEditDeleteComponent } from './cells/add-edit-delete.component';

const TBODY_COMPONENTS = [
  TbodyChangeOrderComponent,
  TbodyCreateCancelComponent,
  TbodyAddEditDeleteComponent,
  Ng2SmartTableTbodyComponent
];

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    CellModule,
  ],
  declarations: [
    ...TBODY_COMPONENTS,
  ],
  exports: [
    ...TBODY_COMPONENTS,
  ],
})
export class TBodyModule { }
