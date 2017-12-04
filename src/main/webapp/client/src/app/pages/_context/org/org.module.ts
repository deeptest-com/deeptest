import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { RouterModule } from '@angular/router';
import { routing }       from './org.routing';
import { NgaModule } from '../../../theme/nga.module';

import { ProjectService } from '../../../service/project';
import { AccountService } from '../../../service/account';

import { Org } from './org.component';

@NgModule({
  imports: [CommonModule, RouterModule, NgaModule, routing],
  declarations: [Org],
  providers: [
    AccountService, ProjectService
  ]
})
export class OrgModule {

}
