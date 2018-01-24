import {Component, OnInit, AfterViewInit, OnDestroy} from "@angular/core";

import {Router} from "@angular/router";

import {GlobalState} from "../../../global.state";

import {CONSTANT} from "../../../utils/constant";
import {WS_CONSTANT} from "../../../utils/ws-constant";

import {RouteService} from "../../../service/route";

import {OrgService} from "../../../service/org";
import {AccountService} from "../../../service/account";

@Component({
  selector: 'ba-page-top',
  templateUrl: './baPageTop.html',
  styleUrls: ['./baPageTop.scss']
})
export class BaPageTop implements OnInit, AfterViewInit, OnDestroy {
  eventCode: string = 'BaPageTop';

  orgId: number;
  prjId: number;

  profile: any = {};
  project: any = {};
  projects: any[] = [];

  orgs: any[] = [];
  keywords: string;

  msgs: any[] = [];
  alerts: any[] = [];

  public isScrolled: boolean = false;

  constructor(private _router: Router, private _state: GlobalState, private _routeService: RouteService,
              private orgService: OrgService, private accountService: AccountService) {
    // console.log('==== BaPageTop constructor ');

    this._state.subscribe(WS_CONSTANT.WS_MSG_AND_ALERT_LASTEST, this.eventCode, (json) => {
      console.log(WS_CONSTANT.WS_MSG_AND_ALERT_LASTEST + ' in ' + this.eventCode, json);
      this.alerts = json.alerts;
      this.msgs = json.msgs;
    });

    this._state.subscribe(WS_CONSTANT.WS_USER_SETTINGS, this.eventCode, (json) => {
      console.log(WS_CONSTANT.WS_USER_SETTINGS + ' in ' + this.eventCode, json);

      this.profile = json.profile;
    });

    this._state.subscribe(WS_CONSTANT.WS_MY_ORGS, this.eventCode, (json: any) => {
      console.log(WS_CONSTANT.WS_MY_ORGS + ' in ' + this.eventCode, json);

      this.orgs = json.myOrgs;
    });

    this._state.subscribe(WS_CONSTANT.WS_RECENT_PROJECTS, this.eventCode, (json) => {
      console.log(WS_CONSTANT.WS_RECENT_PROJECTS + ' in ' + this.eventCode, json);

      CONSTANT.CURR_ORG_ID = json.defaultOrgId;
      CONSTANT.CURR_PRJ_ID = json.defaultPrjId;

      this.orgId = json.defaultOrgId;
      this.prjId = json.defaultPrjId;

      this.projects = json.recentProjects;

    });

    this._state.subscribe(WS_CONSTANT.WS_PRJ_SETTINGS, this.eventCode, (json) => {
      console.log(WS_CONSTANT.WS_PRJ_SETTINGS + ' in ' + this.eventCode, json);

      CONSTANT.PRJ_PRIVILEGES = json.prjPrivileges;
    });

  }

  ngOnInit() {
    this.orgId = CONSTANT.CURR_ORG_ID;
    this.prjId = CONSTANT.CURR_PRJ_ID;

    this.profile = CONSTANT.PROFILE;
    this.orgs = CONSTANT.MY_ORGS;
    this.projects = CONSTANT.RECENT_PROJECTS;

    // console.log('==== BaPageTop ngOnInit ', this.orgId, this.prjId, this.profile);
  }
  ngAfterViewInit() {}

  public changeOrg(org: any) {
    this.orgService.setDefault(org.id, {disabled: false}).subscribe((json: any) => {
      this._routeService.navTo('/pages/org/' + org.id + '/prjs');
    });
  }

  public scrolledChanged(isScrolled) {
    this.isScrolled = isScrolled;
  }

  gotoModule(module: string) {
    let url = '';
    if (module == 'design') {
      url = '/pages/org/' + CONSTANT.CURR_ORG_ID + '/prj/' + CONSTANT.CURR_PRJ_ID + '/design/case';
    } else if (module == 'implement') {
      url = '/pages/org/' + CONSTANT.CURR_ORG_ID + '/prj/' + CONSTANT.CURR_PRJ_ID + '/implement/plan/list';
    }

    this._routeService.navTo(url);
  }

  logout() {
    this.accountService.logout();

    this._routeService.navTo('/pages/login');
  }

  onSearchKeywordChanged(e: any) {
    console.log('0-', this._router.url);
    if (this._router.url.indexOf('design/case') < 0) {
      console.log('1-', this.keywords);
      this._routeService.quickJump(this.keywords);
    } else {
      console.log('2-', this.keywords);
      this._state.notifyDataChanged(CONSTANT.EVENT_CASE_JUMP, this._routeService.caseIdForJump(this.keywords));
    }
  }

  ngOnDestroy(): void {
    this._state.unsubscribe(WS_CONSTANT.WS_MSG_AND_ALERT_LASTEST, this.eventCode);
    this._state.unsubscribe(WS_CONSTANT.WS_USER_SETTINGS, this.eventCode);
    this._state.unsubscribe(WS_CONSTANT.WS_MY_ORGS, this.eventCode);
    this._state.unsubscribe(WS_CONSTANT.WS_RECENT_PROJECTS, this.eventCode);
  };

}
