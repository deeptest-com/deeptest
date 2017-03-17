import {Component, ViewEncapsulation, OnInit, AfterViewInit} from "@angular/core";
import {Ng2TreeOptions} from "../../../../components/ng2-tree/src/tree.types";
import {
  NodeEvent,
  NodeMovedRemoteEvent,
  NodeRemovedRemoteEvent,
  NodeCreatedEvent,
  NodeRenamedEvent,
  NodeSelectedEvent,
  TreeModel
} from "../../../../components/ng2-tree";
import {GlobalState} from "../../../../global.state";
import {Utils} from "../../../../utils/utils";
import {RouteService} from "../../../../service/route";
import {SlimLoadingBarService} from "../../../../components/ng2-loading-bar";
import {TreeService} from "../../../../components/ng2-tree/src/tree.service";
import {CaseService} from "../../../../service/case";

@Component({
  selector: 'case-list',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./list.scss'), require('../../../../components/ng2-tree/src/styles.scss')],
  template: require('./list.html')
})
export class CaseList implements OnInit, AfterViewInit {
  query:any = {keywords: '', status: ''};

  public options:Ng2TreeOptions = {
    isExpanded: false,
    nodeName: '用例',
    folderName: '模块'
  }
  public tree:TreeModel;

  constructor(private _routeService:RouteService, private _state:GlobalState,
              private _treeService:TreeService, private _caseService:CaseService,
              private slimLoadingBarService:SlimLoadingBarService) {
  }

  ngOnInit() {
    let that = this;
    that.loadData();
  }

  ngAfterViewInit() {
    let that = this;
  }

  create():void {
    let that = this;

    that._routeService.navTo("/pages/event/edit/null/property");
  }

  statusChange(e:any):void {
    let that = this;
    that.query.status = e;
    that.loadData();
  }

  delete(eventId:string):void {
    let that = this;
    console.log('eventId=' + eventId);
  }

  loadData() {
    let that = this;
    this.startLoading();
    that._caseService.query(that.query).subscribe((json:any) => {
      that.tree = json.data;
      that.completeLoading();
      this._state.notifyDataChanged('title.change', '测试用例');
    });
  }

  public onNodeRemovedRemote(e:NodeRemovedRemoteEvent):void {
    let that = this;
    this.logEvent(e, 'NodeRemovedRemoteEvent');

    this.startLoading();
    that._caseService.delete(e.node.node).subscribe((json:any) => {
      this._treeService.fireNodeRemoved(e.node);
      that.completeLoading();
    });
  }

  public onNodeMovedRemote(e:NodeMovedRemoteEvent):void {
    let that = this;
    this.logEvent(e, 'NodeMovedRemoteEvent');
    this.startLoading();
    that._caseService.move(e.node.node, e.srcTree.node, e.options).subscribe((json:any) => {
      this._treeService.fireNodeMoved(e.node, e.srcTree, e.options);
      that.completeLoading();
    });
  }

  public onNodeRenamedRemote(e:NodeRenamedEvent):void {
    let that = this;
    this.logEvent(e, 'NodeRenamedEvent');
    this.startLoading();
    that._caseService.rename(e.node.node).subscribe((json:any) => {
      that.completeLoading();
    });
  }

  public onNodeCreatedRemote(e:NodeCreatedEvent):void {
    let that = this;
    that.logEvent(e, 'NodeCreatedEvent');
    this.startLoading();
    that._caseService.create(e.node.node).subscribe((json:any) => {
      that.completeLoading();
    });
  }

  public onNodeSelected(e:NodeSelectedEvent):void {
    this.logEvent(e, 'Selected');
  }

  public logEvent(e:NodeEvent, message:string):void {
    console.log(e, message);
  }

  startLoading() {
    this.slimLoadingBarService.start(() => {
      console.log('Loading complete');
    });
  }

  stopLoading() {
    this.slimLoadingBarService.stop();
  }

  completeLoading() {
    let that = this;
    setTimeout(function () {
      that.slimLoadingBarService.complete();
    }, 500);
  }

}

