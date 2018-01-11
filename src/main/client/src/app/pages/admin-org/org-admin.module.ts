import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../theme/nga.module';

import { routing }       from './org-admin.routing';
import { NgbModalModule, NgbPaginationModule, NgbDropdownModule,
  NgbTabsetModule, NgbButtonsModule, NgbCollapseModule } from '@ng-bootstrap/ng-bootstrap';

import { SlidebarModule } from '../../components/slidebar';

import { OrgAdmin } from './org-admin.component';

import {RouteService} from "../../service/route";
import {AccountService} from "../../service/account";

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NgaModule,
    routing,

    NgbModalModule, NgbPaginationModule, NgbDropdownModule,
    NgbTabsetModule, NgbButtonsModule, NgbCollapseModule,

    SlidebarModule
  ],
  declarations: [
    OrgAdmin
  ],
  providers: [
    RouteService, AccountService
  ]
})
export class OrgAdminModule {}


