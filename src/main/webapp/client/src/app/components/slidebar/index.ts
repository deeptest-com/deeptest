import { NgModule, ModuleWithProviders } from "@angular/core";
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { FormsModule } from '@angular/forms';

import { Slidebar } from './src/slidebar.component';
import { SlidebarMenu } from './src/slidebar-menu.component';
import { SlidebarItem } from './src/slidebar-item.component';

export * from './src/slidebar.component';

@NgModule({
  imports: [CommonModule, RouterModule, FormsModule],
  declarations: [Slidebar, SlidebarMenu, SlidebarItem],
  exports: [Slidebar, SlidebarMenu, SlidebarItem],
  providers: []
})
export class SlidebarModule {
  static forRoot(): ModuleWithProviders {
    return {
      ngModule: SlidebarModule,
      providers: []
    };
  }
}
