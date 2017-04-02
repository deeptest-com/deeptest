import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../../theme/nga.module';

import { routing }       from './org.routing';

import { TabsModule, ModalModule, PaginationModule, DropdownModule } from 'ng2-bootstrap';

import { DirectiveModule } from '../../../directive/directive.module';
import { PopDialogModule } from '../../../components/pop-dialog';

import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { DatetimePickerService } from '../../../service/datetime-picker';
import { OrgService } from '../../../service/org';
import { GroupService } from '../../../service/group';
import { RoleService } from '../../../service/role';

import { Org } from './org.component';
import { OrgList } from './list';
import {OrgEdit } from './edit';

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
    Org,
    OrgList,
    OrgEdit
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    OrgService,
    GroupService,
    RoleService
  ]
})
export default class OrgModule {}

