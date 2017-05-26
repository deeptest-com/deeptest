import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../../../theme/nga.module';

import { routing }       from './case-exe-status.routing';

import { NgbModalModule, NgbPaginationModule, NgbDropdownModule,
  NgbTabsetModule, NgbButtonsModule, NgbCollapseModule } from '@ng-bootstrap/ng-bootstrap';
import { NgUploaderModule } from 'ngx-uploader';

import { PipeModule } from '../../../../pipe/pipe.module';
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
    FormsModule, ReactiveFormsModule,
    NgaModule,
    routing,

    NgbModalModule, NgbPaginationModule, NgbDropdownModule,
    NgbTabsetModule, NgbButtonsModule, NgbCollapseModule,
    NgUploaderModule,

    PipeModule, DirectiveModule, PopDialogModule
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

