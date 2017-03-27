import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../theme/nga.module';

import { routing }       from './admin.routing';

import { CollapseModule } from 'ng2-bootstrap';
import { ModalModule } from 'ng2-bootstrap';
import { ButtonsModule } from 'ng2-bootstrap';
import { FileUploadModule } from 'ng2-file-upload';

import { SlidebarModule } from '../../components/slidebar';

import { Admin } from './admin.component';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NgaModule,
    routing,

    ModalModule,
    ButtonsModule,
    CollapseModule,
    FileUploadModule,

    SlidebarModule
  ],
  declarations: [
    Admin
  ],
  providers: [

  ]
})
export default class AdminModule {}
