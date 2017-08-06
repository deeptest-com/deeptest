import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';

import { NgbModalModule, NgbPaginationModule, NgbDropdownModule,
  NgbTabsetModule, NgbButtonsModule, NgbCollapseModule } from '@ng-bootstrap/ng-bootstrap';
import { NgUploaderModule } from 'ngx-uploader';

import { NgaModule } from '../../../theme/nga.module';
import { routing }       from './execution.routing';

import { DirectiveModule } from '../../../directive/directive.module';
import { SlimLoadingBarModule } from '../../../components/ng2-loading-bar';
import { TreeModule } from '../../../components/ng2-tree';
import { StepsTableModule } from '../../../components/steps-table';
import { CustomFieldModule } from '../../../components/custom-field';
import { FieldShowModule } from '../../../components/field-show';

import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { DatetimePickerService } from '../../../service/datetime-picker';

import { RunService } from '../../../service/run';
import { SuiteService } from '../../../service/suite';
import { CaseService } from '../../../service/case';
import { CaseStepService } from '../../../service/case-step';

import { Execution } from './execution.component';
import { ExecutionSuite } from './suite/suite.component';
import { ExecutionList } from './list/list.component';
import { ExecutionResult } from './result/result.component';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NgaModule,
    routing,

    NgbModalModule, NgbPaginationModule, NgbDropdownModule,
    NgbTabsetModule, NgbButtonsModule, NgbCollapseModule,
    NgUploaderModule,

    DirectiveModule,
    SlimLoadingBarModule.forRoot(),
    TreeModule,
    StepsTableModule,
    CustomFieldModule,
    FieldShowModule
  ],
  declarations: [
    Execution,
    ExecutionSuite,
    ExecutionList,
    ExecutionResult
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    RunService,
    SuiteService,
    CaseService,
    CaseStepService
  ]
})
export class ExecutionModule {}

