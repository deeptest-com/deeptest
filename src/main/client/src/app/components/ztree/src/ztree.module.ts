import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';

import {ToastyService, ToastyConfig, ToastOptions, ToastData} from 'ng2-toasty';
import { PrivilegeService } from '../../../service/privilege';

import { ZtreeComponent } from './ztree.component';
import { ZtreeService } from './ztree.service';

@NgModule({
  imports: [CommonModule, FormsModule, ReactiveFormsModule],
  declarations: [ZtreeComponent],
  exports: [ZtreeComponent],
  providers: [ToastyService,ToastyConfig, ZtreeService, PrivilegeService]
})
export class ZtreeModule {

}
