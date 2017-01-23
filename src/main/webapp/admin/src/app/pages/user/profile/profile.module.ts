import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../../theme/nga.module';

import { routing }       from './profile.routing';

import { ModalModule } from 'ng2-bootstrap';
import { ButtonsModule } from 'ng2-bootstrap';

import { ComponentsModule } from '../../components/components.module';
import { PipeModule } from '../../../pipe/pipe.module';

import { RouteService } from '../../../service/route';
import { RequestService } from '../../../service/request';
import { UserService } from '../../../service/user';

import { Profile } from './profile.component';
import { ProfileEdit } from './profile-edit';

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
    Profile,
    ProfileEdit
  ],
  providers: [
    RouteService,
    RequestService,
    UserService
  ]
})
export default class ProfileModule {}
