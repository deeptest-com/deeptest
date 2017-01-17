import { NgModule, ModuleWithProviders }      from '@angular/core';
import { TabsModule } from 'ng2-bootstrap';
import { ButtonsModule } from 'ng2-bootstrap';

import {
  EventNav
} from './components';

const My_COMPONENTS = [
  EventNav
];

@NgModule({
  imports: [
    TabsModule,
    ButtonsModule
  ],
  declarations: [
    ...My_COMPONENTS
  ],
  exports: [
    ...My_COMPONENTS
  ]
})
export class ComponentsModule {
  
}



