import {Component, ViewEncapsulation, OnInit, AfterViewInit} from "@angular/core";

import {
  NodeEvent,
  NodeMovedRemoteEvent,
  NodeRemovedRemoteEvent,
  NodeCreatedEvent,
  NodeRenamedEvent,
  NodeSelectedEvent,
  TreeModel,
  TreeOptions
} from "../../../../components/ng2-tree";

import {GlobalState} from "../../../../global.state";
import {CONSTANT} from "../../../../utils/constant";
import {Utils} from "../../../../utils/utils";
import {RouteService} from "../../../../service/route";
import {SlimLoadingBarService} from "../../../../components/ng2-loading-bar";
import {TreeService} from "../../../../components/ng2-tree/src/tree.service";
import {SuiteService} from "../../../../service/suite";

@Component({
  selector: 'execution-suite',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./suite.scss', '../../../../components/ng2-tree/src/styles.scss'],
  templateUrl: './suite.html'
})
export class ExecutionSuite implements OnInit, AfterViewInit {
  query:any = {keywords: '', status: ''};

  public options: TreeOptions = {
    usage: 'design',
    isExpanded: false,
    nodeName: '用例',
    folderName: '模块'
  }
  public tree:TreeModel;

  constructor(private _routeService:RouteService, private _state:GlobalState,
              private _treeService:TreeService, private _sutieService: SuiteService,
              private slimLoadingBarService:SlimLoadingBarService) {
  }

  ngOnInit() {
    let that = this;
    that.loadData();
  }

  ngAfterViewInit() {

  }

  delete(eventId:string):void {
    let that = this;
    console.log('eventId=' + eventId);
  }

  loadData() {
    let that = this;
    this.startLoading();
    that._sutieService.query(that.query).subscribe((json:any) => {
      that.tree = json.data;
      CONSTANT.CUSTOM_FIELD_FOR_PROJECT = json.customFields;

      that.completeLoading();
      this._state.notifyDataChanged('title.change', '测试用例');
    });
  }

  public onNodeRemovedRemote(e:NodeRemovedRemoteEvent):void {
    let that = this;
    this.logEvent(e, 'NodeRemovedRemoteEvent');

    this.startLoading();
    that._sutieService.delete(e.node.node).subscribe((json:any) => {
      this._treeService.fireNodeRemoved(e.node);
      that.completeLoading();
    });
  }

  public onNodeMovedRemote(e:NodeMovedRemoteEvent):void {
    let that = this;
    this.logEvent(e, 'NodeMovedRemoteEvent');
    this.startLoading();
    that._sutieService.move(e.node.node, e.srcTree.node, e.options).subscribe((json:any) => {
      this._treeService.fireNodeMoved(e.node, e.srcTree, e.options);
      that.completeLoading();
    });
  }

  public onNodeRenamedRemote(e:NodeRenamedEvent):void {
    let that = this;
    this.logEvent(e, 'NodeRenamedEvent');
    this.startLoading();
    that._sutieService.rename(e.node.node).subscribe((json:any) => {
      that.completeLoading();
    });
  }

  public onNodeCreatedRemote(e:NodeCreatedEvent):void {
    let that = this;
    that.logEvent(e, 'NodeCreatedEvent');
    this.startLoading();
    that._sutieService.create(e.node.node).subscribe((json:any) => {
      e.node.node.id = json.data.id;
      that.completeLoading();
    });
  }

  public onNodeSelected(e:NodeSelectedEvent):void {
    this._state.notifyDataChanged('suite.change', e.node.node);
  }

  public logEvent(e:NodeEvent, message:string):void {
    console.log(e, message);
  }

  startLoading() {
    this.slimLoadingBarService.start(() => {
      console.log('Loading complete');
    });
  }

  completeLoading() {
    let that = this;
    setTimeout(function () {
      that.slimLoadingBarService.complete();
    }, 500);
  }

}

