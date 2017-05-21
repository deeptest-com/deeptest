import {

  NodeRemovedEvent,
  NodeRemovedRemoteEvent,

  NodeRenamedEvent,
  NodeRenamedRemoteEvent,

  NodeCreatedEvent,
  NodeCreatedRemoteEvent,

  NodeMovedEvent,
  NodeMovedRemoteEvent,

  NodeSelectedEvent

} from './tree.events';
import { RenamableNode } from './type/renamable.node';
import { Tree } from './tree';
import { Subject, Observable } from 'rxjs/Rx';
import { Injectable, Inject, ElementRef } from '@angular/core';
import { NodeDraggableService } from './draggable/node-draggable.service';
import { NodeDraggableEvent } from './draggable/draggable.events';

@Injectable()
export class TreeService {
  public nodeMoved$: Subject<NodeMovedEvent> = new Subject<NodeMovedEvent>();
  public nodeMovedRemote$: Subject<NodeMovedRemoteEvent> = new Subject<NodeMovedRemoteEvent>();

  public nodeRemoved$: Subject<NodeRemovedEvent> = new Subject<NodeRemovedEvent>();
  public nodeRemovedRemote$: Subject<NodeRemovedRemoteEvent> = new Subject<NodeRemovedRemoteEvent>();

  public nodeRenamed$: Subject<NodeRenamedEvent> = new Subject<NodeRenamedEvent>();
  public nodeRenamedRemote$: Subject<NodeRenamedRemoteEvent> = new Subject<NodeRenamedRemoteEvent>();

  public nodeCreated$: Subject<NodeCreatedEvent> = new Subject<NodeCreatedEvent>();
  public nodeCreatedRemote$: Subject<NodeCreatedRemoteEvent> = new Subject<NodeCreatedRemoteEvent>();

  public nodeSelected$: Subject<NodeSelectedEvent> = new Subject<NodeSelectedEvent>();

  public constructor(@Inject(NodeDraggableService) private nodeDraggableService: NodeDraggableService) {
    this.nodeRemoved$.subscribe((e: NodeRemovedEvent) => {
      e.node.removeItselfFromParent();
      console.log(e, 'NodeRemovedEvent');
    });

    this.nodeMoved$.subscribe((e: NodeMovedEvent) => {
      console.log(e, 'NodeMovedEvent');

      if (e.options.mode === 'inner') {
        this.moveNodeToFolder(e);
      } else {
        this.moveToBeforeOrAfter(e);
      }
    });
  }

  private moveNodeToFolder(e: NodeMovedEvent): void {
    if (!e.options.isCopy) {
        this.fireNodeRemoved(e.srcTree);
    }
    e.node.addChild(e.srcTree);
  }

  private moveToBeforeOrAfter(e: NodeMovedEvent): void {
    if (!e.options.isCopy) {
      this.fireNodeRemoved(e.srcTree);
    }

    if (e.node.hasSibling(e.srcTree)) {
      e.node.swapWithSibling(e.srcTree, e.options.mode, e.options.isCopy);
    } else {
      let positionInParent = e.node.positionInParent;
      if (e.options.mode === 'after') {
          positionInParent++;
      }
      e.node.addSibling(e.srcTree, positionInParent);
    }
  }

  public unselectStream(tree: Tree): Observable<any> {
    return this.nodeSelected$.filter((e: NodeSelectedEvent) => tree !== e.node);
  }

  public fireNodeRemoved(tree: Tree): void {
    this.nodeRemoved$.next(new NodeRemovedEvent(tree));
  }
  public fireNodeRemovedRemote(tree: Tree): void {
    this.nodeRemovedRemote$.next(new NodeRemovedRemoteEvent(tree));
  }

  public fireNodeCreated(tree: Tree): void {
    this.nodeCreated$.next(new NodeCreatedEvent(tree));
  }
  public fireNodeCreatedRemote(tree: Tree): void {
    this.nodeCreatedRemote$.next(new NodeCreatedRemoteEvent(tree));
  }

  public fireNodeRenamed(oldValue: RenamableNode | string, tree: Tree): void {
    this.nodeRenamed$.next(new NodeRenamedEvent(tree, oldValue, tree.value));
  }
  public fireNodeRenamedRemote(oldValue: RenamableNode | string, tree: Tree): void {
    this.nodeRenamedRemote$.next(new NodeRenamedRemoteEvent(tree, oldValue, tree.value));
  }

  public fireNodeMoved(tree: Tree, parent: Tree, options: any): void {
    this.nodeMoved$.next(new NodeMovedEvent(tree, parent, options));
  }
  public fireNodeMovedRemote(targetTree: Tree, srcTree: Tree, options: any): void {
    this.nodeMovedRemote$.next(new NodeMovedRemoteEvent(targetTree, srcTree, options));
  }

  public fireNodeSelected(tree: Tree): void {
    this.nodeSelected$.next(new NodeSelectedEvent(tree));
  }

  public draggedStream(tree: Tree, element: ElementRef): Observable<NodeDraggableEvent> {
    return this.nodeDraggableService.draggableNodeEvents$
      .filter((e: NodeDraggableEvent) => {
        // if (e.target === element) {
        //   console.log('e.target', e.target.nativeElement['children'][0]['innerText']);
        //   console.log('element', element.nativeElement['children'][0]['innerText']);
        // }

        return e.target === element;
      })
      .filter((e: NodeDraggableEvent) => !e.captured.tree.hasChild(tree));
  }
}
