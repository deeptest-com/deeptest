import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../theme/nga.module';

import { routing }       from './implement.routing';
import { NgbModalModule, NgbPaginationModule, NgbDropdownModule,
  NgbTabsetModule, NgbButtonsModule, NgbCollapseModule } from '@ng-bootstrap/ng-bootstrap';

import { I18n, CustomDatepickerI18n } from '../../service/datepicker-I18n';

import { Implement } from './implement.component';

import { ProjectService } from '../../service/project';
import { AccountService } from '../../service/account';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NgaModule,
    routing,

    NgbModalModule, NgbPaginationModule, NgbDropdownModule,
    NgbTabsetModule, NgbButtonsModule, NgbCollapseModule
  ],
  declarations: [
    Implement
  ],
  providers: [
    I18n, CustomDatepickerI18n,
    AccountService, ProjectService
  ]
})
export class ImplementModule {}


