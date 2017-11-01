import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../theme/nga.module';

import { routing }       from './personal.routing';

import { NgbModalModule, NgbPaginationModule, NgbDropdownModule,
  NgbTabsetModule, NgbButtonsModule, NgbCollapseModule } from '@ng-bootstrap/ng-bootstrap';
import { NgUploaderModule } from 'ngx-uploader';

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

import { Settings } from './settings';
import { SettingsEdit } from './settings/settings-edit';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NgaModule,
    routing,

    NgbModalModule, NgbPaginationModule, NgbDropdownModule,
    NgbTabsetModule, NgbButtonsModule, NgbCollapseModule,
    NgUploaderModule,
  ],
  declarations: [
    Personal,
    Password,
    PasswordEdit,
    Profile,
    ProfileEdit,
    Settings,
    SettingsEdit
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    AccountService,
    CompanyService
  ]
})
export class PersonalModule {}
