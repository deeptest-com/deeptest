import { Input, Component, OnInit, ElementRef, Inject } from '@angular/core';
import { Ng2TreeSettings, Ng2TreeOptions } from './tree.types';
import { Subject, Observable } from 'rxjs/Rx';
import { Tree } from './tree';
import {
  NodeRemovedEvent,

  NodeRenamedEvent,

  NodeCreatedEvent,

  NodeMovedEvent

} from './tree.events';

import { NodeMenuService } from './menu/node-menu.service';
import { NodeMenuItemSelectedEvent, NodeMenuItemAction } from './menu/menu.events';
import { NodeEditableEvent, NodeEditableEventAction } from './editable/editable.events';
import { TreeService } from './tree.service';
import * as EventUtils from './utils/event.utils';
import { NodeDraggableEvent } from './draggable/draggable.events';

@Component({
  selector: 'tree-internal',
  template: `
  <ul class="tree" *ngIf="tree" [ngClass]="{rootless: isRootHidden()}">
    <li> 
      <div class="value-container fa"
        [ngClass]="{rootless: isRootHidden()}" 
        (contextmenu)="showMenu($event)" 
        [nodeDraggable]="element"
        [tree]="tree">
            <div class="folding" (click)="tree.switchFoldingType()" [ngClass]="tree.foldingType.cssClass"></div>
            <div class="node-value" 
              *ngIf="!shouldShowInputForTreeValue()" 
              [class.node-selected]="isSelected" 
              (click)="onNodeSelected($event)">{{tree.value}}
            </div>
    
            <input type="text" class="node-value" 
               *ngIf="shouldShowInputForTreeValue()"
               [nodeEditable]="tree.value"
               (valueChanged)="applyNewValue($event)"/>
      </div>

      <node-menu *ngIf="isMenuVisible" (menuItemSelected)="onMenuItemSelected($event)"
            [options]="options"></node-menu>

      <template [ngIf]="tree.isNodeExpanded()">
        <tree-internal *ngFor="let child of tree.children" [tree]="child" [options]="options"></tree-internal>
      </template>
    </li>
  </ul>
  `
})
export class TreeInternalComponent implements OnInit {
  @Input()
  public tree: Tree;

  @Input()
  public settings: Ng2TreeSettings;

    @Input()
    public options: Ng2TreeOptions;

  public isSelected: boolean = false;
  private isMenuVisible: boolean = false;

  public nodeMoved$: Subject<NodeMovedEvent> = new Subject<NodeMovedEvent>();

  public constructor(@Inject(NodeMenuService) private nodeMenuService: NodeMenuService,
                     @Inject(TreeService) private treeService: TreeService,
                     @Inject(ElementRef) public element: ElementRef) {

    this.nodeMoved$.subscribe((e: NodeMovedEvent) => {
      alert(1);
    });
  }

  public ngOnInit(): void {

    this.settings = this.settings || { rootIsVisible: false };

    this.nodeMenuService.hideMenuStream(this.element)
      .subscribe(() => this.isMenuVisible = false);

    this.treeService.unselectStream(this.tree)
      .subscribe(() => this.isSelected = false);

    this.treeService.draggedStream(this.tree, this.element)
      .subscribe((e: NodeDraggableEvent) => {
        this.moveNode(e);
      });
  }

  private moveNode(e: NodeDraggableEvent): void {
    this.treeService.fireNodeMovedRemote(this.tree, e.captured.tree, {mode: e.mode, isCopy: e.isCopy});
  }

  public onNodeSelected(e: MouseEvent): void {
    if (EventUtils.isLeftButtonClicked(e)) {
      this.isSelected = true;
      this.treeService.fireNodeSelected(this.tree);
    }
  }

  public showMenu(e: MouseEvent): void {
    if (this.tree.isStatic()) {
      return;
    }

    if (EventUtils.isRightButtonClicked(e)) {
      this.isMenuVisible = !this.isMenuVisible;
      this.nodeMenuService.hideMenuForAllNodesExcept(this.element);
    }
    e.preventDefault();
  }

  public onMenuItemSelected(e: NodeMenuItemSelectedEvent): void {
    switch (e.nodeMenuItemAction) {
      case NodeMenuItemAction.NewTag:
        this.onNewSelected(e);
        break;
      case NodeMenuItemAction.NewFolder:
        this.onNewSelected(e);
        break;
      case NodeMenuItemAction.Rename:
        this.onRenameSelected();
        break;
      case NodeMenuItemAction.Remove:
        this.onRemoveSelected();
        break;
      default:
        throw new Error(`Chosen menu item doesn't exist`);
    }
  }

  private onNewSelected(e: NodeMenuItemSelectedEvent): void {
    this.tree.createNode(e.nodeMenuItemAction === NodeMenuItemAction.NewFolder);
    this.isMenuVisible = false;
  }

  private onRenameSelected(): void {
    this.tree.markAsBeingRenamed();
    this.isMenuVisible = false;
  }

  private onRemoveSelected(): void {
      this.treeService.fireNodeRemovedRemote(this.tree);
  }

  public applyNewValue(e: NodeEditableEvent): void {
    if ((e.action === NodeEditableEventAction.Cancel || this.tree.isNew()) && Tree.isValueEmpty(e.value)) {
      return this.treeService.fireNodeRemoved(this.tree);
    }

    if (this.tree.isNew()) {
      this.tree.value = e.value;
      this.treeService.fireNodeCreated(this.tree);
    }

    if (this.tree.isBeingRenamed()) {
      const oldValue = this.tree.value;
      this.tree.value = e.value;
      this.treeService.fireNodeRenamed(oldValue, this.tree);
    }

    this.tree.markAsModified();
  }

  public shouldShowInputForTreeValue(): boolean {
    return this.tree.isNew() || this.tree.isBeingRenamed();
  }

  public isRootHidden(): boolean {
    return this.tree.isRoot() && !this.settings.rootIsVisible;
  }
}
