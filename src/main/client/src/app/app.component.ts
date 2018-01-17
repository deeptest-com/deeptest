import { Component, ViewContainerRef } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import * as $ from 'jquery';

import { GlobalState } from './global.state';
import { BaImageLoaderService, BaThemePreloader, BaThemeSpinner } from './theme/services';
import { BaThemeConfig } from './theme/theme.config';
import { layoutPaths } from './theme/theme.constants';

import {CONSTANT} from './utils/constant';
import {Utils} from './utils/utils';

/*
 * App Component
 * Top Level Component
 */
@Component({
  selector: 'app',
  styleUrls: ['./app.component.scss'],
  template: `
    <main baThemeRun>
      <div class="additional-bg"></div>
      <router-outlet></router-outlet>
    </main>
  `
})
export class App {
  eventCode: string = 'App';

  constructor(private _state: GlobalState,
              private _activatedRoute: ActivatedRoute,
              private _imageLoader: BaImageLoaderService,
              private _spinner: BaThemeSpinner,
              private viewContainerRef: ViewContainerRef,
              private themeConfig: BaThemeConfig) {

    themeConfig.config();

    this._loadImages();

    Utils.config();

    CONSTANT.ScreenSize = Utils.getScreenSize();

    this._state.subscribe(CONSTANT.EVENT_LOADING_COMPLETE, this.eventCode, (json) => {
      console.log(CONSTANT.EVENT_LOADING_COMPLETE + ' in ' + this.eventCode, json);
      this._spinner.hide();
    });
  }

  public ngAfterViewInit(): void {
    // hide spinner once all loaders are completed
    BaThemePreloader.load().then((values) => {
      // this._spinner.hide();
    });
  }

  private _loadImages(): void {
    // register some loaders
    BaThemePreloader.registerLoader(this._imageLoader.load('/assets/img/sky-bg.jpg'));
  }

}
