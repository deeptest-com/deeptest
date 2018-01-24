import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import {FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../../theme/nga.module';

import { NgbModalModule, NgbPaginationModule, NgbDropdownModule,
  NgbTabsetModule, NgbButtonsModule, NgbCollapseModule } from '@ng-bootstrap/ng-bootstrap';

import { AppTranslationModule } from '../../../app.translation.module';
import { PipeModule } from '../../../pipe/pipe.module';
import { DirectiveModule } from '../../../directive/directive.module';
import { TableTreeModule } from '../../../components/table-tree';
import { PopDialogModule } from '../../../components/pop-dialog';
import { ChartDesignModule } from '../../../components/chart/chart-design';
import { ChartExecutionModule } from '../../../components/chart/chart-execution';
import { SearchSelectModule } from '../../../components/search-select';

import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { DatetimePickerService } from '../../../service/datetime-picker';

import { routing }       from './project.routing';
import { ProjectService } from '../../../service/project';
import { UserService } from '../../../service/user';
import { UserAndGroupService } from '../../../service/userAndGroup';
import { ReportService } from '../../../service/report';

import { Project } from './project.component';
import { ProjectList } from './list/list.component';
import { ProjectEdit } from './edit/edit.component';
import { ProjectView } from './view/view.component';

@NgModule({
  imports: [
    CommonModule,
    FormsModule, ReactiveFormsModule,
    NgaModule,
    routing,

    NgbModalModule, NgbPaginationModule, NgbDropdownModule,
    NgbTabsetModule, NgbButtonsModule, NgbCollapseModule,

    PipeModule,
    DirectiveModule,
    TableTreeModule,
    PopDialogModule,
    ChartDesignModule,
    ChartExecutionModule,
    SearchSelectModule,
    AppTranslationModule
  ],
  declarations: [
    Project,
    ProjectList,
    ProjectEdit,
    ProjectView
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    ProjectService,
    UserService,
    UserAndGroupService,
    ReportService
  ]
})
export class ProjectModule {}

