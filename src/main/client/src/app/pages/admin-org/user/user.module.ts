import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../../theme/nga.module';

import { routing }       from './user.routing';

import { NgbModalModule, NgbPaginationModule, NgbDropdownModule,
  NgbTabsetModule, NgbButtonsModule, NgbCollapseModule } from '@ng-bootstrap/ng-bootstrap';

import { PipeModule } from '../../../pipe/pipe.module';
import { DirectiveModule } from '../../../directive/directive.module';
import { PopDialogModule } from '../../../components/pop-dialog';

import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { DatetimePickerService } from '../../../service/datetime-picker';
import { UserService } from '../../../service/user';
import { GroupService } from '../../../service/group';
import { RoleService } from '../../../service/role';

import { User } from './user.component';
import { UserList } from './list';
import { UserEdit } from './edit';
import { UserInvite } from './invite';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NgaModule,
    routing,

    NgbModalModule, NgbPaginationModule, NgbDropdownModule,
    NgbTabsetModule, NgbButtonsModule, NgbCollapseModule,

    PipeModule, DirectiveModule, PopDialogModule
  ],
  declarations: [
    User,
    UserList,
    UserEdit,
    UserInvite
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    UserService,
    GroupService,
    RoleService
  ]
})
export class UserModule {}

