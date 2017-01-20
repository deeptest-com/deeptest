import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../theme/nga.module';

import { routing }       from './account.routing';

import { PaginationModule} from 'ng2-bootstrap';
import { TabsModule } from 'ng2-bootstrap';
import { CollapseModule } from 'ng2-bootstrap';
import { ModalModule } from 'ng2-bootstrap';
import { ButtonsModule } from 'ng2-bootstrap';
import { FileUploadModule } from 'ng2-file-upload';

import { ComponentsModule } from '../components/components.module';
import { PipeModule } from '../../pipe/pipe.module';

import { RouteService } from '../../service/route';
import { RequestService } from '../../service/request';
import { AccountService } from '../../service/account';

import { Account } from './account.component';
import { AccountList } from './account-list';
import { AccountEdit } from './account-edit';

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

    ComponentsModule,
    PipeModule
  ],
  declarations: [
    Account,
    AccountList,
    AccountEdit
  ],
  providers: [
    RouteService,
    RequestService,
    AccountService
  ]
})
export default class AccountModule {}
