import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { UserService } from '../service/user';

import { ResizeDirective } from './resize/resize.directive';

@NgModule({
  imports: [CommonModule],
  declarations: [ResizeDirective],
  exports: [ResizeDirective],
  providers: [UserService]
})
export class DirectiveModule {

}

