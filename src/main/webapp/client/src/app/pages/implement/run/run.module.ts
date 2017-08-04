import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';

import { NgbModalModule, NgbPaginationModule, NgbDropdownModule,
  NgbTabsetModule, NgbButtonsModule, NgbCollapseModule } from '@ng-bootstrap/ng-bootstrap';
import { NgUploaderModule } from 'ngx-uploader';

import { NgaModule } from '../../../theme/nga.module';
import { routing }       from './run.routing';

import { PipeModule } from '../../../pipe/pipe.module';
import { DirectiveModule } from '../../../directive/directive.module';
import { PopDialogModule } from '../../../components/pop-dialog';

import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { DatetimePickerService } from '../../../service/datetime-picker';
import { RunService } from '../../../service/run';

import { Run } from './run.component';

@NgModule({
  imports: [
    CommonModule,
    FormsModule, ReactiveFormsModule,
    NgaModule,
    routing,

    NgbModalModule, NgbPaginationModule, NgbDropdownModule,
    NgbTabsetModule, NgbButtonsModule, NgbCollapseModule,
    NgUploaderModule,

    DirectiveModule,
    PopDialogModule,
    PipeModule
  ],
  declarations: [
    Run
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    RunService
  ]
})
export class RunModule {}
