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

  @Output() reSearchEvent: EventEmitter<any> = new EventEmitter<any>();
  @Output() renameEvent: EventEmitter<any> = new EventEmitter<any>();
  @Output() removeEvent: EventEmitter<any> = new EventEmitter<any>();
  @Output() moveEvent: EventEmitter<any> = new EventEmitter<any>();

  ztree: any;
  keywordsControl = new FormControl();
  keywords: string = '';
  isExpanded: boolean = true;

  log: any;
  newCount: number = 0;
  className: string = "dark";
  curDragNodes: any[] = [];
  autoExpandNode: any;

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
        // beforeDrag: this.beforeDrag,
        // beforeDrop: this.beforeDrop,
        // beforeDragOpen: this.beforeDragOpen,
        // onDrag: this.onDrag,
        onDrop: this.onMove,
        onExpand: this.onExpand
      }
    };

     _.merge(this.settings, this.treeSettings)
  }

  public ngOnInit(): void {
    this.keywordsControl.valueChanges.debounceTime(CONSTANT.DebounceTime).subscribe(values => this.onChange(values));

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
    this.onChange('');
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

  onMove = (event, treeId, treeNodes, targetNode, moveType, isCopy) => {
    if(!targetNode) {
      return;
    }

    const deferred = new Deferred();
    deferred.promise.then((data) => {
      console.log('success to move', data);
      this._state.notifyDataChanged('case.change', data);
    }).catch((err) => {console.log('err', err);});

    this.moveEvent.emit({
      data: {srcId: treeNodes[0].id, targetId: targetNode.id, moveType: moveType, isCopy: isCopy},
      deferred: deferred,
    });
  }

  // dropPrev = (treeId, nodes, targetNode) => {
  //   var pNode = targetNode.getParentNode();
  //   if (pNode && pNode.dropInner === false) {
  //     return false;
  //   } else {
  //     for (var i=0,l=this.curDragNodes.length; i<l; i++) {
  //       var curPNode = this.curDragNodes[i].getParentNode();
  //       if (curPNode && curPNode !== targetNode.getParentNode() && curPNode.childOuter === false) {
  //         return false;
  //       }
  //     }
  //   }
  //   return true;
  // }
  //
  // dropInner = (treeId, nodes, targetNode) => {
  //   if (targetNode && targetNode.dropInner === false) {
  //     return false;
  //   } else {
  //     for (var i=0,l=this.curDragNodes.length; i<l; i++) {
  //       if (!targetNode && this.curDragNodes[i].dropRoot === false) {
  //         return false;
  //       } else if (this.curDragNodes[i].parentTId && this.curDragNodes[i].getParentNode() !== targetNode && this.curDragNodes[i].getParentNode().childOuter === false) {
  //         return false;
  //       }
  //     }
  //   }
  //   return true;
  // }
  //
  // dropNext = (treeId, nodes, targetNode) => {
  //   var pNode = targetNode.getParentNode();
  //   if (pNode && pNode.dropInner === false) {
  //     return false;
  //   } else {
  //     for (var i=0,l=this.curDragNodes.length; i<l; i++) {
  //       var curPNode = this.curDragNodes[i].getParentNode();
  //       if (curPNode && curPNode !== targetNode.getParentNode() && curPNode.childOuter === false) {
  //         return false;
  //       }
  //     }
  //   }
  //   return true;
  // }

  // beforeDrag = (treeId, treeNodes) => {
  //   this.className = (this.className === "dark" ? "":"dark");
  //   this.showLog("[ "+this.getTime()+" beforeDrag ]&nbsp;&nbsp;&nbsp;&nbsp; drag: " + treeNodes.length + " nodes." );
  //   for (var i=0,l=treeNodes.length; i<l; i++) {
  //     if (treeNodes[i].drag === false) {
  //       this.curDragNodes = null;
  //       console.log('=1=', this.curDragNodes);
  //       return false;
  //     } else if (treeNodes[i].parentTId && treeNodes[i].getParentNode().childDrag === false) {
  //       this.curDragNodes = null;
  //       console.log('=2=', this.curDragNodes);
  //       return false;
  //     }
  //   }
  //   this.curDragNodes = treeNodes;
  //   console.log('=3=', this.curDragNodes);
  //   return true;
  // }
  //
  // beforeDragOpen = (treeId, treeNode) => {
  //   this.autoExpandNode = treeNode;
  //   return true;
  // }
  //
  // beforeDrop = (treeId, treeNodes, targetNode, moveType, isCopy) => {
  //   this.className = (this.className === "dark" ? "":"dark");
  //   this.showLog("[ "+this.getTime()+" beforeDrop ]&nbsp;&nbsp;&nbsp;&nbsp; moveType:" + moveType);
  //   this.showLog("target: " + (targetNode ? targetNode.name : "root") + "  -- is "+ (isCopy==null? "cancel" : isCopy ? "copy" : "move"));
  //   return true;
  // }
  //
  // onDrag = (event, treeId, treeNodes) => {
  //   this.className = (this.className === "dark" ? "":"dark");
  //   this.showLog("[ "+this.getTime()+" onDrag ]&nbsp;&nbsp;&nbsp;&nbsp; drag: " + treeNodes.length + " nodes." );
  // }

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

  onChange(values) {
    this.keywords = values;

    const deferred = new Deferred();
    deferred.promise.then((data) => {
      this.ztree = jQuery.fn.zTree.init($('#tree'), this.settings, data);

      this.ztree.expandNode(this.ztree.getNodes()[0], true, true, true);

      let nodes = this.ztree.getNodesByParam("isHidden", true);
      this.ztree.showNodes(nodes);

      nodes = this.ztree.getNodesByFilter((node) => {
        return this.keywords || node.name.indexOf(this.keywords) < 0;
      });
      this.ztree.hideNodes(nodes);

    }).catch((err) => {console.log('err', err);});

    this.reSearchEvent.emit({
      keywords: this.keywords,
      deferred: deferred,
    });
  }

}
