import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../theme/nga.module';

import { routing }       from './personal.routing';

import { PaginationModule} from 'ng2-bootstrap';
import { TabsModule } from 'ng2-bootstrap';
import { CollapseModule } from 'ng2-bootstrap';
import { ModalModule } from 'ng2-bootstrap';
import { ButtonsModule } from 'ng2-bootstrap';
import { FileUploadModule } from 'ng2-file-upload';

import { RouteService } from '../../service/route';
import { RequestService } from '../../service/request';
import { DatetimePickerService } from '../../service/datetime-picker';

import { AccountService } from '../../service/account';
import { CompanyService } from '../../service/company';


import { Personal } from './personal.component';

import { Password } from './password';
import { PasswordEdit } from './password/password-edit';

import { Profile } from './profile';
import { ProfileEdit } from './profile/profile-edit';

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
    FileUploadModule
  ],
  declarations: [
    Personal,
    Password,
    PasswordEdit,
    Profile,
    ProfileEdit
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    AccountService,
    CompanyService
  ]
})
export default class EventModule {}
