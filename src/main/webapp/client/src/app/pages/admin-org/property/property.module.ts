import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../../theme/nga.module';

import { routing }       from './property.routing';
import { CollapseModule } from 'ng2-bootstrap';
import { ModalModule } from 'ng2-bootstrap';
import { ButtonsModule } from 'ng2-bootstrap';

import { Property } from './property.component';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NgaModule,
    routing,

    ModalModule,
    ButtonsModule,
    CollapseModule
  ],
  declarations: [
    Property
  ],
  providers: [

  ]
})
export default class PropertyModule {}
