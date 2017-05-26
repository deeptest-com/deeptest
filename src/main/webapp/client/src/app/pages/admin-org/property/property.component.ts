import {Component, OnInit, OnDestroy} from "@angular/core";
import {Router, NavigationEnd} from '@angular/router';
import {Subscription} from 'rxjs/Rx';
import { GlobalState } from '../../../global.state';
import { RouteService } from '../../../service/route';

@Component({
  selector: 'property',
  styleUrls: ['./property.scss'],
  templateUrl: './property.html'
})
export class Property implements OnInit, OnDestroy {
  protected _onRouteChange:Subscription;
  tabs: string[] = ['case-type', 'case-priority', 'case-exe-status', 'custom-field'];
  tab: string = 'case-type';
  status: string = 'list';

  constructor(private _router:Router, private _state: GlobalState, private _routeService: RouteService) {

    this._onRouteChange = this._router.events.subscribe((event) => {
      if (event instanceof NavigationEnd && event.url) {
        let arr = event.url.split('property/')[1].split('/');
        this.tab = arr[0];
        this.status = arr[1];
      }
    });
  }

  ngOnInit() {

  }
  ngOnDestroy(): void {
    this._onRouteChange.unsubscribe();
  }

  tabChange(event: any) {
    let index = event.activeId.replace('ngb-tab-', '');
    this.tab = this.tabs[index];
    this._routeService.navTo("/pages/org-admin/property/" + this.tab + "/list");
  }

  create() {
    this.status = 'edit';
    this._routeService.navTo("/pages/org-admin/property/" + this.tab + "/edit/null");
  }
  back() {
    this.status = 'list';
    this._routeService.navTo("/pages/org-admin/property/" + this.tab + "/list");
  }

}
