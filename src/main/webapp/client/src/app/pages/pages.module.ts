import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';

import { routing }       from './pages.routing';
import { NgaModule } from '../theme/nga.module';

import { Pages } from './pages.component';

import { RouteService } from '../service/route';
import { AccountService } from '../service/account';
import { RequestService } from '../service/request';

@NgModule({
  imports: [
    CommonModule, NgaModule, routing
  ],
  declarations: [Pages],
  providers: [RouteService, RequestService, AccountService]
})
export class PagesModule {

}

