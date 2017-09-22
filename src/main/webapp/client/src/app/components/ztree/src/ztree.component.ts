import { Input, Component, OnInit, AfterViewInit, EventEmitter, Output, Inject, OnChanges, SimpleChanges } from '@angular/core';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import * as _ from 'lodash';

import {GlobalState} from "../../../global.state";
import { Deferred, getDeepFromObject } from './helpers';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';

import { ZtreeService } from './ztree.service';

declare var jQuery;

@Component({
  selector: 'ztree',
  templateUrl: './ztree.html',
  styleUrls: ['./styles.scss',
    '../../../../vendor/ztree/css/zTreeStyle/zTreeStyle.css'],
  providers: [ZtreeService]
})
export class ZtreeComponent implements OnInit, AfterViewInit {

  @Input()
  public treeSettings: any;
  public settings: any;

  @Output() renameEvent: EventEmitter<any> = new EventEmitter<any>();
  @Output() removeEvent: EventEmitter<any> = new EventEmitter<any>();
  @Output() moveEvent: EventEmitter<any> = new EventEmitter<any>();

  _treeModel: any;
  ztree: any;
  keywordsControl = new FormControl();
  keywords: string = '';
  isExpanded: boolean = true;
  isDragging: boolean = false;

  log: any;
  newCount: number = 0;
  className: string = "dark";
  curDragNodes: any[] = [];
  autoExpandNode: any;

  @Input() set treeModel(model: any) {
    if(!model) {
      return;
    }

    this._treeModel = model;
    this.ztree = jQuery.fn.zTree.init($('#tree'), this.settings, this._treeModel);
    // this.ztree.reAsyncChildNodes(null, '');
    this.ztree.expandNode(this.ztree.getNodes()[0], true, true, true);
  }

  public constructor(private _state:GlobalState, @Inject(ZtreeService) private ztreeService: ZtreeService) {

    this.settings = {
      view: {
        addHoverDom: this.addHoverDom,
        removeHoverDom: this.removeHoverDom,
        selectedMulti: false
      },
      edit: {
        enable: true,
        editNameSelectAll: true,
        showRemoveBtn: true,
        showRenameBtn: true,
        drag: {
          autoExpandTrigger: true,
          // prev: this.dropPrev,
          // inner: this.dropInner,
          // next: this.dropNext
        }
      },
      data: {
        simpleData: {
          enable: true
        }
      },
      callback: {
        onClick: this.onClick,
        beforeRemove: this.beforeRemove,
        onRemove: this.onRemove,
        onRename: this.onRename,
        onDrag: this.onDrag,
        onDrop: this.onMove,
        onExpand: this.onExpand
      }
    };

     _.merge(this.settings, this.treeSettings)
  }

  public ngOnInit(): void {
    this.keywordsControl.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(values => this.onKeywordsChange(values));

    this._state.subscribe('case.save', (testCase: any) => {
      console.log(testCase);
      if (testCase) {
        var node = this.ztree.getNodeByParam("id", testCase.id, null);

        node.name = testCase.name;
        this.ztree.updateNode(node);
      }
    });
  }

  ngAfterViewInit() {

  }

  expandOrNot() {
    if (!this.isExpanded) {
      this.ztree.expandAll(true);
    } else {
      this.ztree.expandAll(false);
    }

    this.isExpanded = !this.isExpanded;
  }

  onClick = (event, treeId, treeNode) => {
    this._state.notifyDataChanged('case.change', treeNode);
  }

  addHoverDom = (treeId, treeNode) => {
    var sObj = $("#" + treeNode.tId + "_span");
    if (treeNode.editNameFlag || $("#addBtn_"+treeNode.tId).length>0) return;
    var addStr = "<span class='button add' id='addBtn_" + treeNode.tId
      + "' title='add node' onfocus='this.blur();'></span>";
    sObj.after(addStr);

    var btn = jQuery("#addBtn_"+treeNode.tId);
    if (btn) btn.bind("click", () => {
      let newNode = this.ztree.addNodes(treeNode, {id: -1 * this.newCount++, pId:treeNode.id, name:"新用例",
        type: "functional", priority: 2, estimate: undefined});
      this.ztree.editName(newNode[0]);
      return false;
    });
  }
  removeHoverDom = (treeId, treeNode) => {
    $("#addBtn_"+treeNode.tId).unbind().remove();
  }

  onRename = (e, treeId, treeNode, isCancel) => {
    const deferred = new Deferred();
    deferred.promise.then((data) => {
      console.log('success to rename', data);
      treeNode.id = data.id;
      treeNode.ordr = data.ordr;

      treeNode.tm = new Date().getTime();
      this._state.notifyDataChanged('case.change', _.clone(treeNode));
    }).catch((err) => {console.log('err', err);});

    this.renameEvent.emit({
      data: treeNode,
      deferred: deferred,
    });
  }

  beforeRemove = (treeId, treeNode) => {
    this.className = (this.className === "dark" ? "":"dark");
    this.ztree.selectNode(treeNode);
    return confirm('确认删除名为"' + treeNode.name + '"的用例吗？');
  }
  onRemove = (e, treeId, treeNode) => {
    const deferred = new Deferred();
    deferred.promise.then((data) => {
      console.log('success to remove', treeNode);
      this._state.notifyDataChanged('case.change', null);
    }).catch((err) => {console.log('err', err);});

    this.removeEvent.emit({
      data: treeNode,
      deferred: deferred,
    });
  }
  onDrag = (event, treeId, treeNodes) => {
    this.isDragging = true;
  }

  onMove = (event, treeId, treeNodes, targetNode, moveType, isCopy) => {
    this.isDragging = false;
    if(!targetNode) {
      return;
    }

    const deferred = new Deferred();
    deferred.promise.then((data) => {
      console.log('success to move', data);
      treeNodes[0].id = data.id;
      this._state.notifyDataChanged('case.change', data);
    }).catch((err) => {console.log('err', err);});

    this.moveEvent.emit({
      data: {srcId: treeNodes[0].id, targetId: targetNode.id, moveType: moveType, isCopy: isCopy},
      deferred: deferred,
    });
  }

  onExpand = (event, treeId, treeNode) => {
    if (treeNode === this.autoExpandNode) {
      this.className = (this.className === "dark" ? "":"dark");
      this.showLog("[ "+this.getTime()+" onExpand ]&nbsp;&nbsp;&nbsp;&nbsp;" + treeNode.name);
    }
  }

  showLog = (str) => {
    console.log(str);
  }
  getTime = () => {
    var now= new Date(),
      h=now.getHours(),
      m=now.getMinutes(),
      s=now.getSeconds(),
      ms=now.getMilliseconds();
    return (h+":"+m+":"+s+ " " +ms);
  }

  onKeywordsChange(values) {
    this.keywords = values;
    let nodes = this.ztree.getNodesByParam("isHidden", true);
    this.ztree.showNodes(nodes);

    nodes = this.ztree.getNodesByFilter((node) => {
      return this.keywords && !node.isParent && node.name.indexOf(this.keywords) < 0;
    });
    this.ztree.hideNodes(nodes);
  }

}
