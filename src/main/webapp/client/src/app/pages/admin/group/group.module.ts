import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../../theme/nga.module';

import { routing }       from './group.routing';

import { DropdownModule } from 'ng2-bootstrap/ng2-bootstrap';

import { DirectiveModule } from '../../../directive/directive.module';

import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { DatetimePickerService } from '../../../service/datetime-picker';
import { GroupService } from '../../../service/group';

import { Group } from './group.component';
import { GroupList } from './list';
import { GroupEdit } from './edit';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NgaModule,
    routing,

    DropdownModule,
    DirectiveModule
  ],
  declarations: [
    Group,
    GroupList,
    GroupEdit
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    GroupService,
  ]
})
export default class ProjectModule {}

