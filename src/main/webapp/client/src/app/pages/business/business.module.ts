import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../theme/nga.module';

import { routing }       from './business.routing';

import { PaginationModule} from 'ng2-bootstrap';
import { TabsModule } from 'ng2-bootstrap';
import { CollapseModule } from 'ng2-bootstrap';
import { ModalModule } from 'ng2-bootstrap';
import { ButtonsModule } from 'ng2-bootstrap';
import { FileUploadModule } from 'ng2-file-upload';

import { RouteService } from '../../service/route';
import { RequestService } from '../../service/request';
import { DatetimePickerService } from '../../service/datetime-picker';

import { UserService } from '../../service/user';
import { AccountService } from '../../service/account';
import { CompanyService } from '../../service/company';


import { Business } from './business.component';

import { Account } from './account';
import { AccountList } from './account/account-list';
import { AccountEdit } from './account/account-edit';

import { Company } from './company';
import { CompanyEdit } from './company/company-edit';

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
    Business,
    Account,
    AccountList,
    AccountEdit,
    Company,
    CompanyEdit
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    UserService,
    AccountService,
    CompanyService
  ]
})
export default class EventModule {}
