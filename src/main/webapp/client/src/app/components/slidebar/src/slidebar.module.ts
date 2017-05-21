import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { Slidebar } from './slidebar.component';
import { SlidebarMenu } from './slidebar-menu.component';
import { SlidebarItem } from './slidebar-item.component';

import { SlidebarService } from './slidebar.service';

@NgModule({
  imports: [CommonModule],
  declarations: [Slidebar, SlidebarMenu, SlidebarItem],
  exports: [Slidebar],
  providers: [SlidebarService]
})
export class SlidebarModule {
}
