import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';

import { NgbModalModule, NgbPaginationModule, NgbDropdownModule,
  NgbTabsetModule, NgbButtonsModule, NgbCollapseModule } from '@ng-bootstrap/ng-bootstrap';

import { NgaModule } from '../../../theme/nga.module';
import { routing }       from './report.routing';

import { DirectiveModule } from '../../../directive/directive.module';
import { SlimLoadingBarModule } from '../../../components/ng2-loading-bar';
import { StepsTableModule } from '../../../components/steps-table';
import { CustomFieldModule } from '../../../components/custom-field';

import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { DatetimePickerService } from '../../../service/datetime-picker';
import { CaseService } from '../../../service/case';
import { CaseStepService } from '../../../service/case-step';

import { Report } from './report.component';
import { ReportList } from './list/list.component';
import { ReportEdit } from './edit/edit.component';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NgaModule,
    routing,

    NgbModalModule, NgbPaginationModule, NgbDropdownModule,
    NgbTabsetModule, NgbButtonsModule, NgbCollapseModule,

    DirectiveModule,
    SlimLoadingBarModule.forRoot(),
    StepsTableModule,
    CustomFieldModule
  ],
  declarations: [
    Report,
    ReportList,
    ReportEdit
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    CaseService,
    CaseStepService
  ]
})
export class ReportModule {}

