import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';

import { NgbModalModule, NgbPaginationModule, NgbDropdownModule,
  NgbTabsetModule, NgbButtonsModule, NgbCollapseModule, NgbDatepickerModule, NgbDateParserFormatter } from '@ng-bootstrap/ng-bootstrap';

import { NgaModule } from '../../../theme/nga.module';
import { routing }       from './plan.routing';

import { PipeModule } from '../../../pipe/pipe.module';
import { DirectiveModule } from '../../../directive/directive.module';
import { PopDialogModule } from '../../../components/pop-dialog';
import { CaseSelectionModule, CaseSelectionComponent } from '../../../components/case-selection';
import { EnvironmentConfigModule, EnvironmentConfigComponent } from '../../../components/environment-config';
import { ExecutionReportModule } from '../../../components/execution-report';
import { ExecutionBarModule } from '../../../components/execution-bar';

import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { DatetimePickerService } from '../../../service/datetime-picker';
import { MyDateParserFormatter } from '../../../service/my-date-parser-formatter';
import { PlanService } from '../../../service/plan';
import { RunService } from '../../../service/run';
import { SuiteService } from '../../../service/suite';
import { CaseService } from '../../../service/case';
import { UserService } from '../../../service/user';

import { ProjectService } from '../../../service/project';
import { AccountService } from '../../../service/account';

import { Plan } from './plan.component';
import { PlanList } from './list/list.component';
import { PlanView } from './view/view.component';
import { PlanEdit } from './edit/edit.component';

export function myDateParserFormatterFactory() {
  return new MyDateParserFormatter("y-MM-dd");
}

@NgModule({
  imports: [
    CommonModule,
    FormsModule, ReactiveFormsModule,
    NgaModule,
    routing,

    NgbModalModule, NgbPaginationModule, NgbDropdownModule,
    NgbTabsetModule, NgbButtonsModule, NgbCollapseModule, NgbDatepickerModule,

    DirectiveModule,
    PipeModule,
    PopDialogModule,

    CaseSelectionModule,
    EnvironmentConfigModule,
    ExecutionReportModule,
    ExecutionBarModule
  ],
  declarations: [
    Plan,
    PlanList,
    PlanView,
    PlanEdit
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    PlanService, RunService, SuiteService, CaseService, UserService,
    AccountService, ProjectService,
    {
      provide: NgbDateParserFormatter,
      useFactory: myDateParserFormatterFactory
    }
  ],
  entryComponents: [
    CaseSelectionComponent,
    EnvironmentConfigComponent
  ]
})
export class PlanModule {}
