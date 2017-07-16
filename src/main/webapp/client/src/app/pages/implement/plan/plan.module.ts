import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';

import { NgbModalModule, NgbPaginationModule, NgbDropdownModule,
  NgbTabsetModule, NgbButtonsModule, NgbCollapseModule, NgbDatepickerModule } from '@ng-bootstrap/ng-bootstrap';
import { NgUploaderModule } from 'ngx-uploader';

import { NgaModule } from '../../../theme/nga.module';
import { routing }       from './plan.routing';

import { PipeModule } from '../../../pipe/pipe.module';
import { DirectiveModule } from '../../../directive/directive.module';
import { PopDialogModule } from '../../../components/pop-dialog';
import { CaseSelectionModule, CaseSelectionComponent } from '../../../components/case-selection';
import { EnvironmentConfigModule, EnvironmentConfigComponent } from '../../../components/environment-config';

import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { DatetimePickerService } from '../../../service/datetime-picker';
import { PlanService } from '../../../service/plan';
import { RunService } from '../../../service/run';

import { Plan } from './plan.component';
import { PlanList } from './list/list.component';
import { PlanEdit } from './edit/edit.component';

@NgModule({
  imports: [
    CommonModule,
    FormsModule, ReactiveFormsModule,
    NgaModule,
    routing,

    NgbModalModule, NgbPaginationModule, NgbDropdownModule,
    NgbTabsetModule, NgbButtonsModule, NgbCollapseModule, NgbDatepickerModule,
    NgUploaderModule,

    DirectiveModule,
    PipeModule,
    PopDialogModule,

    CaseSelectionModule,
    EnvironmentConfigModule
  ],
  declarations: [
    Plan,
    PlanList,
    PlanEdit
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    PlanService, RunService
  ],
  entryComponents: [
    CaseSelectionComponent,
    EnvironmentConfigComponent
  ]
})
export default class PlanModule {}
