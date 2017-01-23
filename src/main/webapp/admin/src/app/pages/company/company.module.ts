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
import { CompanyEdit } from './company-edit';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NgaModule,
    routing,
    ModalModule,
    ButtonsModule,

    ComponentsModule,
    PipeModule
  ],
  declarations: [
    Company,
    CompanyEdit
  ],
  providers: [
    RouteService,
    RequestService,
    CompanyService
  ]
})
export default class CompanyModule {}
