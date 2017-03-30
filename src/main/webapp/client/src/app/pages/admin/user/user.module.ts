import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../../theme/nga.module';

import { routing }       from './user.routing';

import { TabsModule, ModalModule, PaginationModule, DropdownModule } from 'ng2-bootstrap';

import { DirectiveModule } from '../../../directive/directive.module';
import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { DatetimePickerService } from '../../../service/datetime-picker';
import { UserService } from '../../../service/user';

import { User } from './user.component';
import { UserList } from './list';
import {UserEdit, UserEditInfo, UserEditGroups } from './edit';

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
    DirectiveModule
  ],
  declarations: [
    User,
    UserList,

    UserEdit,
    UserEditInfo,
    UserEditGroups
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    UserService,
  ]
})
export default class UserModule {}

