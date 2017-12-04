import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../../theme/nga.module';

import {NgbModule} from '@ng-bootstrap/ng-bootstrap';
import { PipeModule } from '../../../pipe/pipe.module';

import { routing }       from './report.routing';

import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { ReportService } from '../../../service/report';

import { Report } from './report.component';
import { ReportList } from './list';
import { ReportView } from './view';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NgbModule,
    NgaModule,
    PipeModule,
    routing
  ],
  declarations: [
    Report,
    ReportList,
    ReportView
  ],
  providers: [
    RouteService,
    RequestService,
    ReportService,
  ]
})
export class ReportModule {}

