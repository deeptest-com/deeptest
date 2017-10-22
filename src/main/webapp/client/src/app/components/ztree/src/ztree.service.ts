import { Subject, Observable } from 'rxjs/Rx';
import { Injectable, Inject, ElementRef } from '@angular/core';

@Injectable()
export class ZtreeService {

  public constructor() {

  }

  selectNode(node: any): any {
    let ztree = jQuery.fn.zTree.getZTreeObj('tree');
    ztree.selectNode(node);
  }

  getNextNode(id: number): any {
    let ztree = jQuery.fn.zTree.getZTreeObj('tree');
    let nodes = ztree.getNodesByParam("id", id, null);
    let curr = nodes[0];

    return this.getNextNodeObject(curr);
  }

  getNextNodeObject(node: any): any {
    if (!node) {
      return null;
    }

    var next = node.getNextNode();
    console.log('next', next);

    if (next != null) {
      if (!next.isParent) {
        return next;
      } else {
        console.log('isParent', true, next);
        if (next.children == null || next.children.length == 0) {
          return null;
        } else {
          return next.children[0];
        }
      }
    } else {
      console.log('node', node);
      let parent = node.getParentNode();
      console.log('parent', parent);

      return this.getNextNodeObject(parent);
    }
  }

}
