import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../theme/nga.module';

import { routing }       from './event.routing';

import { PaginationModule} from 'ng2-bootstrap';
import { TabsModule } from 'ng2-bootstrap';
import { CollapseModule } from 'ng2-bootstrap';
import { ModalModule } from 'ng2-bootstrap';
import { ButtonsModule } from 'ng2-bootstrap';
import { FileUploadModule } from 'ng2-file-upload';

import { PipeModule } from '../../pipe/pipe.module';

import { RequestService } from '../../service/request';
import { DatetimePickerService } from '../../service/datetime-picker';
import { EventService } from '../../service/event';
import { SessionService } from '../../service/session';
import { ScheduleService } from '../../service/schedule';
import { GuestService } from '../../service/guest';
import { ServiceService } from '../../service/service';
import { AroundService } from '../../service/around';

import { Event } from './event.component';
import { EventList } from './event-list';
import { EventEditProperty } from './event-edit/property';
import { EventEditSchedule } from './event-edit/schedule';
import { EventEditGuest } from './event-edit/guest';
import { EventEditService } from './event-edit/service';
import { EventEditAround } from './event-edit/around';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NgaModule,
    routing,

    TabsModule,
    PaginationModule,
    ModalModule,
    ButtonsModule,
    CollapseModule,
    FileUploadModule,

    PipeModule
  ],
  declarations: [
    Event,
    EventList,
    EventEditProperty,
    EventEditSchedule,
    EventEditGuest,
    EventEditService,
    EventEditAround
  ],
  providers: [
    RequestService,
    DatetimePickerService,
    EventService,
    SessionService,
    ScheduleService,
    GuestService,
    ServiceService,
    AroundService
  ]
})
export default class EventModule {}
