import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ResizeDirective } from './resize/resize.directive';

@NgModule({
  imports: [CommonModule],
  declarations: [ResizeDirective],
  exports: [ResizeDirective],
  providers: []
})
export class DirectiveModule {

}

