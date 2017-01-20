import { Routes, RouterModule }  from '@angular/router';

import { Event } from './event.component';
import { EventList } from './event-list/list.component';
import { EventEditProperty } from './event-edit/property/edit.component';
import { EventEditSchedule } from './event-edit/schedule/edit.component';
import { EventEditGuest } from './event-edit/guest/edit.component';
import { EventEditService } from './event-edit/service/edit.component';
import { EventEditBanner } from './event-edit/banner/edit.component';
import { EventEditDocument } from './event-edit/document/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Event,
    children: [
      { path: 'list', component: EventList },
      { path: 'edit/:id/property', component: EventEditProperty },
      { path: 'edit/:id/schedule', component: EventEditSchedule },
      { path: 'edit/:id/guest', component: EventEditGuest },
      { path: 'edit/:id/service', component: EventEditService },
      { path: 'edit/:id/banner', component: EventEditBanner },
      { path: 'edit/:id/document', component: EventEditDocument }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
