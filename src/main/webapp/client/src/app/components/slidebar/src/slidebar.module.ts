import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';

import { Slidebar } from './slidebar.component';
import { SlidebarMenu } from './slidebar-menu.component';
import { SlidebarItem } from './slidebar-item.component';

import { SlidebarService } from './slidebar.service';

@NgModule({
  imports: [CommonModule, FormsModule, ReactiveFormsModule],
  declarations: [Slidebar, SlidebarMenu, SlidebarItem],
  exports: [Slidebar],
  providers: [SlidebarService]
})
export class SlidebarModule {
}
