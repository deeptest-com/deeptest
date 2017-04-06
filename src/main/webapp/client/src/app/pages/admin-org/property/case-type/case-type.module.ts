import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../../../theme/nga.module';

import { routing }       from './case-type.routing';

import { TabsModule, ModalModule, PaginationModule, DropdownModule } from 'ng2-bootstrap';

import { DirectiveModule } from '../../../../directive/directive.module';
import { PopDialogModule } from '../../../../components/pop-dialog';

import { RouteService } from '../../../../service/route';
import { RequestService } from '../../../../service/request';
import { DatetimePickerService } from '../../../../service/datetime-picker';
import { CaseTypeService } from '../../../../service/case-type';

import { CaseType } from './case-type.component';
import { CaseTypeList } from './list';
import {CaseTypeEdit } from './edit';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NgaModule,
    routing,

    TabsModule,
    ModalModule,
    PaginationModule,
    DropdownModule,
    DirectiveModule,
    PopDialogModule
  ],
  declarations: [
    CaseType,
    CaseTypeList,
    CaseTypeEdit
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    CaseTypeService
  ]
})
export default class CaseTypeModule {}

