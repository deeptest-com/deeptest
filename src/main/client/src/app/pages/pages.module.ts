import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';

import 'rxjs/add/operator/debounceTime';
import 'rxjs/add/operator/map';

import { routing }       from './pages.routing';
import { NgaModule } from '../theme/nga.module';
import { AppTranslationModule } from '../app.translation.module';

import { UserService } from '../service/user';

import { Pages } from './pages.component';
import { PagesResolve } from './pages.resolve';

@NgModule({
  imports: [CommonModule, AppTranslationModule, NgaModule, routing],
  declarations: [Pages],
  providers: [
    UserService,
    PagesResolve
  ]
})
export class PagesModule {

}
