import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { NgaModule } from '../../theme/nga.module';

import { routing }       from './autotest.routing';

import { AutoTest } from './autotest.component';

@NgModule({
  imports: [
    CommonModule,
    NgaModule,
    routing

  ],
  declarations: [
    AutoTest
  ],
  providers: [

  ]
})
export class AutoTestModule {}
