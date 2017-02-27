import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';

import { routing }       from './pages.routing';
import { NgaModule } from '../theme/nga.module';

import { Pages } from './pages.component';

import { RouteService } from '../service/route';
import { UserService } from '../service/user';
import { RequestService } from '../service/request';

@NgModule({
  imports: [
    CommonModule, NgaModule, routing
  ],
  declarations: [Pages],
  providers: [RouteService, UserService, RequestService]
})
export class PagesModule {

}

