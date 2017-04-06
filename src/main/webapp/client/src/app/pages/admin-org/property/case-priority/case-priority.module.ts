import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../../../theme/nga.module';

import { routing }       from './case-priority.routing';

import { TabsModule, ModalModule, PaginationModule, DropdownModule } from 'ng2-bootstrap';

import { DirectiveModule } from '../../../../directive/directive.module';
import { PopDialogModule } from '../../../../components/pop-dialog';

import { RouteService } from '../../../../service/route';
import { RequestService } from '../../../../service/request';
import { DatetimePickerService } from '../../../../service/datetime-picker';
import { CasePriorityService } from '../../../../service/case-priority';

import { CasePriority } from './case-priority.component';
import { CasePriorityList } from './list';
import {CasePriorityEdit } from './edit';

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
    CasePriority,
    CasePriorityList,
    CasePriorityEdit
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    CasePriorityService
  ]
})
export default class CasePriorityModule {}

