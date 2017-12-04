import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../../theme/nga.module';

import { routing }       from './reset-password.routing';
import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { AccountService } from '../../../service/account';

import { ResetPassword } from './reset-password.component';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    FormsModule,
    NgaModule,
    routing
  ],
  declarations: [
    ResetPassword
  ],
  providers: [
    RouteService,
    RequestService,
    AccountService
  ]
})
export class ResetPasswordModule {}
