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

import { RouteService } from '../../service/route';
import { RequestService } from '../../service/request';
import { DatetimePickerService } from '../../service/datetime-picker';
import { EventService } from '../../service/event';
import { SessionService } from '../../service/session';
import { ScheduleService } from '../../service/schedule';
import { GuestService } from '../../service/guest';
import { ServiceService } from '../../service/service';
import { BannerService } from '../../service/banner';
import { DocumentService } from '../../service/document';
import { OrganizerService } from '../../service/organizer';

import { Event } from './event.component';
import { EventList } from './event-list';
import { EventEditProperty } from './event-edit/property';
import { EventEditOrganizer } from './event-edit/organizer';
import { EventEditSchedule } from './event-edit/schedule';
import { EventEditGuest } from './event-edit/guest';
import { EventEditService } from './event-edit/service';
import { EventEditBanner } from './event-edit/banner';
import { EventEditDocument } from './event-edit/document';

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
    FileUploadModule
  ],
  declarations: [
    Event,
    EventList,
    EventEditProperty,
    EventEditOrganizer,
    EventEditSchedule,
    EventEditGuest,
    EventEditService,
    EventEditBanner,
    EventEditDocument
  ],
  providers: [
    RouteService,
    RequestService,
    DatetimePickerService,
    EventService,
    OrganizerService,
    SessionService,
    ScheduleService,
    GuestService,
    ServiceService,
    BannerService,
    DocumentService
  ]
})
export default class EventModule {}
