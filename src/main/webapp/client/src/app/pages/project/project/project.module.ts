import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../../theme/nga.module';

import { routing }       from './project.routing';

import { PaginationModule} from 'ng2-bootstrap';
import { TabsModule } from 'ng2-bootstrap';
import { CollapseModule } from 'ng2-bootstrap';
import { ModalModule } from 'ng2-bootstrap';
import { ButtonsModule } from 'ng2-bootstrap';
import { FileUploadModule } from 'ng2-file-upload';

import { DirectiveModule } from '../../../directive/directive.module';
import { TableTreeModule } from '../../../components/table-tree';
import { PopDialogModule } from '../../../components/pop-dialog';

import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { DatetimePickerService } from '../../../service/datetime-picker';
import { ProjectService } from '../../../service/project';

import { Project } from './project.component';
import { ProjectList } from './list/list.component';
import { ProjectEdit } from './edit/edit.component';
import { ProjectView } from './view/view.component';

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

    DirectiveModule,
    TableTreeModule,
    PopDialogModule
  ],
  declarations: [
    Project,
    ProjectList,
    ProjectEdit,
    ProjectView
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    ProjectService,
  ]
})
export default class ProjectModule {}

