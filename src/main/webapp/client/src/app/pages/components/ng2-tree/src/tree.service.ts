import {
    NodeDeletedEvent,
  NodeRemovedEvent,
  NodeRenamedEvent,
  NodeCreatedEvent,
  NodeSelectedEvent,
  NodeMovedEvent
} from './tree.events';
import { RenamableNode } from './tree.types';
import { Tree } from './tree';
import { Subject, Observable } from 'rxjs/Rx';
import { Injectable, Inject, ElementRef } from '@angular/core';
import { NodeDraggableService } from './draggable/node-draggable.service';
import { NodeDraggableEvent } from './draggable/draggable.events';

@Injectable()
export class TreeService {
  public nodeMoved$: Subject<NodeMovedEvent> = new Subject<NodeMovedEvent>();
  public nodeDeleted$: Subject<NodeDeletedEvent> = new Subject<NodeDeletedEvent>();
  public nodeRemoved$: Subject<NodeRemovedEvent> = new Subject<NodeRemovedEvent>();
  public nodeRenamed$: Subject<NodeRenamedEvent> = new Subject<NodeRenamedEvent>();
  public nodeCreated$: Subject<NodeCreatedEvent> = new Subject<NodeCreatedEvent>();
  public nodeSelected$: Subject<NodeSelectedEvent> = new Subject<NodeSelectedEvent>();

  public constructor(@Inject(NodeDraggableService) private nodeDraggableService: NodeDraggableService) {
    this.nodeRemoved$.subscribe((e: NodeRemovedEvent) => e.node.removeItselfFromParent());
  }

  public unselectStream(tree: Tree): Observable<any> {
    return this.nodeSelected$.filter((e: NodeSelectedEvent) => tree !== e.node);
  }

  public fireNodeRemoved(tree: Tree): void {
    this.nodeRemoved$.next(new NodeRemovedEvent(tree));
  }

  public fireNodeDeleted(tree: Tree): void {
    this.nodeDeleted$.next(new NodeDeletedEvent(tree));
  }

  public fireNodeCreated(tree: Tree): void {
    this.nodeCreated$.next(new NodeCreatedEvent(tree));
  }

  public fireNodeSelected(tree: Tree): void {
    this.nodeSelected$.next(new NodeSelectedEvent(tree));
  }

  public fireNodeRenamed(oldValue: RenamableNode | string, tree: Tree): void {
    this.nodeRenamed$.next(new NodeRenamedEvent(tree, oldValue, tree.value));
  }

  public fireNodeMoved(tree: Tree, parent: Tree): void {
    this.nodeMoved$.next(new NodeMovedEvent(tree, parent));
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
