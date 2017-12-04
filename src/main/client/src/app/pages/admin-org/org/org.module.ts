import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';

import { NgaModule } from '../../../theme/nga.module';
import { routing }       from './org.routing';

import { NgbModalModule, NgbPaginationModule, NgbDropdownModule,
  NgbTabsetModule, NgbButtonsModule, NgbCollapseModule } from '@ng-bootstrap/ng-bootstrap';

import { PipeModule } from '../../../pipe/pipe.module';
import { DirectiveModule } from '../../../directive/directive.module';
import { PopDialogModule } from '../../../components/pop-dialog';

import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { DatetimePickerService } from '../../../service/datetime-picker';
import { OrgService } from '../../../service/org';
import { ProjectService } from '../../../service/project';
import { GroupService } from '../../../service/group';
import { RoleService } from '../../../service/role';

import { Org } from './org.component';
import { OrgList } from './list';
import {OrgEdit } from './edit';

@NgModule({
  imports: [
    CommonModule,
    FormsModule, ReactiveFormsModule,
    NgaModule,
    routing,

    NgbModalModule, NgbPaginationModule, NgbDropdownModule,
    NgbTabsetModule, NgbButtonsModule, NgbCollapseModule,

    PipeModule, DirectiveModule, PopDialogModule
  ],
  declarations: [
    Org,
    OrgList,
    OrgEdit
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    OrgService,
    ProjectService,
    GroupService,
    RoleService
  ]
})
export class OrgModule {}

