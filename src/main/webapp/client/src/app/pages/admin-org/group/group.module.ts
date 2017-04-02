import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../../theme/nga.module';

import { routing }       from './group.routing';

import { ModalModule } from 'ng2-bootstrap';
import { PaginationModule} from 'ng2-bootstrap';
import { DropdownModule } from 'ng2-bootstrap/ng2-bootstrap';

import { DirectiveModule } from '../../../directive/directive.module';
import { PopDialogModule } from '../../../components/pop-dialog';

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

    ModalModule,
    PaginationModule,
    DropdownModule,
    DirectiveModule,
    PopDialogModule
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
export default class GroupModule {}

