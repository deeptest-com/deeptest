import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { NgaModule } from '../../../theme/nga.module';
import { SlimLoadingBarModule } from '../../../components/ng2-loading-bar';

import { routing }       from './register.routing';
import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { AccountService } from '../../../service/account';

import { Register } from './register.component';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    FormsModule,
    NgaModule,
    SlimLoadingBarModule,
    routing
  ],
  declarations: [
    Register
  ],
  providers: [
    RouteService,
    RequestService,
    AccountService
  ]
})
export class RegisterModule {}
