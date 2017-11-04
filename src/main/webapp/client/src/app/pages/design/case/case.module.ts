import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import {ToastyModule} from 'ng2-toasty';

import { NgbModalModule, NgbPaginationModule, NgbDropdownModule,
  NgbTabsetModule, NgbButtonsModule, NgbCollapseModule } from '@ng-bootstrap/ng-bootstrap';

import { NgaModule } from '../../../theme/nga.module';
import { routing }       from './case.routing';

import { DirectiveModule } from '../../../directive/directive.module';
import { SlimLoadingBarModule } from '../../../components/ng2-loading-bar';
import { ZtreeModule } from '../../../components/ztree';
import { StepsTableModule } from '../../../components/steps-table';
import { CustomFieldModule } from '../../../components/custom-field';
import { TinyMCEModule } from '../../../components/tiny-mce';

import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { DatetimePickerService } from '../../../service/datetime-picker';

import { SuiteService } from '../../../service/suite';
import { CaseService } from '../../../service/case';
import { CaseStepService } from '../../../service/case-step';

import { Case } from './case.component';
import { CaseSuite } from './suite/suite.component';
import { CaseEdit } from './edit/edit.component';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NgaModule,
    routing,

    NgbModalModule, NgbPaginationModule, NgbDropdownModule,
    NgbTabsetModule, NgbButtonsModule, NgbCollapseModule,

    DirectiveModule,
    ToastyModule,
    SlimLoadingBarModule.forRoot(),
    ZtreeModule,
    StepsTableModule,
    CustomFieldModule,
    TinyMCEModule
  ],
  declarations: [
    Case,
    CaseSuite,
    CaseEdit
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    SuiteService,
    CaseService,
    CaseStepService
  ]
})
export class CaseModule {}

