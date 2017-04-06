import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../../../theme/nga.module';

import { routing }       from './case-exe-status.routing';

import { TabsModule, ModalModule, PaginationModule, DropdownModule } from 'ng2-bootstrap';

import { DirectiveModule } from '../../../../directive/directive.module';
import { PopDialogModule } from '../../../../components/pop-dialog';

import { RouteService } from '../../../../service/route';
import { RequestService } from '../../../../service/request';
import { DatetimePickerService } from '../../../../service/datetime-picker';
import { CaseExeStatusService } from '../../../../service/case-exe-status';

import { CaseExeStatus } from './case-exe-status.component';
import { CaseExeStatusList } from './list';
import {CaseExeStatusEdit } from './edit';

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
    CaseExeStatus,
    CaseExeStatusList,
    CaseExeStatusEdit
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    CaseExeStatusService
  ]
})
export default class CaseExeStatusModule {}

