import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../theme/nga.module';

import { routing }       from './personal.routing';

import { NgbModalModule, NgbPaginationModule, NgbDropdownModule,
  NgbTabsetModule, NgbButtonsModule, NgbCollapseModule } from '@ng-bootstrap/ng-bootstrap';

import { PictureUploaderModule } from '../../components/picture-uploader';
import { FieldShowModule } from '../../components/field-show';
import {PipeModule} from '../../pipe/pipe.module';

import { RouteService } from '../../service/route';
import { RequestService } from '../../service/request';
import { DatetimePickerService } from '../../service/datetime-picker';

import { AccountService } from '../../service/account';
import { CompanyService } from '../../service/company';

import { Personal } from './personal.component';

import { PasswordEditComponent, PasswordEditPopupComponent } from './password';

import { Profile } from './profile';
import { ProfileEdit } from './profile/profile-edit';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NgaModule,
    routing,

    NgbModalModule, NgbPaginationModule, NgbDropdownModule,
    NgbTabsetModule, NgbButtonsModule, NgbCollapseModule,
    PictureUploaderModule, FieldShowModule, PipeModule
  ],
  declarations: [
    Personal,
    PasswordEditComponent,
    PasswordEditPopupComponent,
    Profile,
    ProfileEdit
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    AccountService,
    CompanyService
  ],
  entryComponents: [
    PasswordEditPopupComponent
  ]
})
export class PersonalModule {}
