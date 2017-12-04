import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../../theme/nga.module';

import { routing }       from './forgot-password.routing';
import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { AccountService } from '../../../service/account';

import { ForgotPassword } from './forgot-password.component';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    FormsModule,
    NgaModule,
    routing
  ],
  declarations: [
    ForgotPassword
  ],
  providers: [
    RouteService,
    RequestService,
    AccountService
  ]
})
export class ForgotPasswordModule {}
