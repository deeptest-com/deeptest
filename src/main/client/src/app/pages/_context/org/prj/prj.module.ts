import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { RouterModule } from '@angular/router';
import { routing }       from './prj.routing';
import { NgaModule } from '../../../../theme/nga.module';

import { ProjectService } from '../../../../service/project';
import { AccountService } from '../../../../service/account';

import { Prj } from './prj.component';

@NgModule({
  imports: [CommonModule, RouterModule, NgaModule, routing],
  declarations: [Prj],
  providers: [
    AccountService, ProjectService
  ]
})
export class PrjModule {

}
