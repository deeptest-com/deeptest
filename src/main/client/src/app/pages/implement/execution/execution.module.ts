import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';

import { NgbModalModule, NgbPaginationModule, NgbDropdownModule,
  NgbTabsetModule, NgbButtonsModule, NgbCollapseModule } from '@ng-bootstrap/ng-bootstrap';

import { NgaModule } from '../../../theme/nga.module';
import { routing }       from './execution.routing';

import { PipeModule } from '../../../pipe/pipe.module';

import { DirectiveModule } from '../../../directive/directive.module';
import { SlimLoadingBarModule } from '../../../components/ng2-loading-bar';
import { ZtreeModule } from '../../../components/ztree';
import { StepsTableModule } from '../../../components/steps-table';
import { CustomFieldModule } from '../../../components/custom-field';
import { FieldShowModule } from '../../../components/field-show';
import { CaseCommentsModule } from '../../../components/case-comments';

import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { DatetimePickerService } from '../../../service/datetime-picker';

import { RunService } from '../../../service/run';
import { SuiteService } from '../../../service/suite';
import { CaseService } from '../../../service/case';
import { CaseStepService } from '../../../service/case-step';
import { CaseInRunService } from '../../../service/case-in-run';
import { PrivilegeService } from '../../../service/privilege';
import { CaseCommentsService } from '../../../service/case-comments';

import { Execution } from './execution.component';
import { ExecutionSuite } from './suite/suite.component';
import { ExecutionResult } from './result/result.component';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NgaModule,
    routing,

    NgbModalModule, NgbPaginationModule, NgbDropdownModule,
    NgbTabsetModule, NgbButtonsModule, NgbCollapseModule,

    PipeModule,

    DirectiveModule,
    SlimLoadingBarModule.forRoot(),
    ZtreeModule,
    StepsTableModule,
    CustomFieldModule,
    FieldShowModule,
    CaseCommentsModule
  ],
  declarations: [
    Execution,
    ExecutionSuite,
    ExecutionResult
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    RunService,
    SuiteService,
    CaseService,
    CaseStepService,
    CaseInRunService,
    PrivilegeService,
    CaseCommentsService
  ]
})
export class ExecutionModule {}

