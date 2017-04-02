import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../../theme/nga.module';

import { routing }       from './role.routing';

import { PaginationModule} from 'ng2-bootstrap';
import { TabsModule } from 'ng2-bootstrap';
import { CollapseModule } from 'ng2-bootstrap';
import { ModalModule } from 'ng2-bootstrap';
import { ButtonsModule } from 'ng2-bootstrap';
import { FileUploadModule } from 'ng2-file-upload';

import { DirectiveModule } from '../../../directive/directive.module';
import { PopDialogModule } from '../../../components/pop-dialog';

import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { DatetimePickerService } from '../../../service/datetime-picker';
import { RoleService } from '../../../service/role';

import { Role } from './role.component';
import { RoleList } from './list';
import { RoleEdit } from './edit';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NgaModule,
    routing,

    TabsModule,
    PaginationModule,
    ModalModule,
    ButtonsModule,
    CollapseModule,
    FileUploadModule,

    DirectiveModule,
    PopDialogModule
  ],
  declarations: [
    Role,
    RoleList,
    RoleEdit
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    RoleService,
  ]
})
export default class RoleModule {}

