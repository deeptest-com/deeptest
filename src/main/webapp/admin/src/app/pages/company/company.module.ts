import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../theme/nga.module';

import { routing }       from './company.routing';

import { ModalModule } from 'ng2-bootstrap';
import { ButtonsModule } from 'ng2-bootstrap';

import { ComponentsModule } from '../components/components.module';
import { PipeModule } from '../../pipe/pipe.module';

import { RouteService } from '../../service/route';
import { RequestService } from '../../service/request';
import { CompanyService } from '../../service/company';

import { Company } from './company.component';
import { CompanyList } from './company-list';
import { CompanyEdit } from './company-edit';

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
    Company,
    CompanyList,
    CompanyEdit
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    CompanyService
  ]
})
export default class CompanyModule {}
