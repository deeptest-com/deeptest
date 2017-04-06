import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../../../theme/nga.module';

import { routing }       from './custom-field.routing';

import { TabsModule, ModalModule, PaginationModule, DropdownModule } from 'ng2-bootstrap';

import { DirectiveModule } from '../../../../directive/directive.module';
import { PopDialogModule } from '../../../../components/pop-dialog';

import { RouteService } from '../../../../service/route';
import { RequestService } from '../../../../service/request';
import { DatetimePickerService } from '../../../../service/datetime-picker';
import { CustomFieldService } from '../../../../service/custom-field';

import { CustomField } from './custom-field.component';
import { CustomFieldList } from './list';
import { CustomFieldEdit } from './edit';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NgaModule,
    routing,

    TabsModule,
    ModalModule,
    PaginationModule,
    DropdownModule,
    DirectiveModule,
    PopDialogModule
  ],
  declarations: [
    CustomField,
    CustomFieldList,
    CustomFieldEdit
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    CustomFieldService
  ]
})
export default class CustomFieldModule {}

