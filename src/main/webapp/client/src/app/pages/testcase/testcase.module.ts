import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../theme/nga.module';

import { routing }       from './testcase.routing';

import { PaginationModule} from 'ng2-bootstrap';
import { TabsModule } from 'ng2-bootstrap';
import { CollapseModule } from 'ng2-bootstrap';
import { ModalModule } from 'ng2-bootstrap';
import { ButtonsModule } from 'ng2-bootstrap';
import { FileUploadModule } from 'ng2-file-upload';

import { TreeModule } from '../components/ng2-tree';

import { RouteService } from '../../service/route';
import { RequestService } from '../../service/request';
import { DatetimePickerService } from '../../service/datetime-picker';
import { TestcaseService } from '../../service/testcase';

import { Testcase } from './testcase.component';
import { TestcaseList } from './testcase-list';

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
    TreeModule
  ],
  declarations: [
    Testcase,
    TestcaseList
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    TestcaseService,
  ]
})
export default class TestcaseModule {}

