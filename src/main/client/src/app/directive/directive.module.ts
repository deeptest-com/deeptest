import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { UserService } from '../service/user';
import { PrivilegeService } from '../service/privilege';

import { ResizeDirective } from './resize/resize.directive';
import { PrivilegeDirective } from './privilege/privilege.directive';

@NgModule({
  imports: [CommonModule],
  declarations: [ResizeDirective, PrivilegeDirective],
  exports: [ResizeDirective, PrivilegeDirective],
  providers: [UserService, PrivilegeService]
})
export class DirectiveModule {

}

