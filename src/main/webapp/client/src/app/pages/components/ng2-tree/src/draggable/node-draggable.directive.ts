import { Directive, ElementRef, Input, Inject, Renderer, OnDestroy, OnInit } from '@angular/core';
import { NodeDraggableService } from './node-draggable.service';
import { CapturedNode } from './captured-node';
import { Tree } from '../tree';

@Directive({
  selector: '[nodeDraggable]'
})
export class NodeDraggableDirective implements OnDestroy, OnInit {
  public static DATA_TRANSFER_STUB_DATA: string = 'some browsers enable drag-n-drop only when dataTransfer has data';

  @Input()
  public nodeDraggable: ElementRef;

  @Input()
  public tree: Tree;

    mode: string;
  isCopy: boolean;

  private nodeNativeElement: HTMLElement;
  private disposersForDragListeners: Function[] = [];

  public constructor(@Inject(ElementRef) public element: ElementRef,
                     @Inject(NodeDraggableService) private nodeDraggableService: NodeDraggableService,
                     @Inject(Renderer) private renderer: Renderer) {
    this.nodeNativeElement = element.nativeElement;
  }

  public ngOnInit(): void {
    if (!this.tree.isStatic()) {
      this.renderer.setElementAttribute(this.nodeNativeElement, 'draggable', 'true');
      this.disposersForDragListeners.push(this.renderer.listen(this.nodeNativeElement, 'dragenter', this.handleDragEnter.bind(this)));
      this.disposersForDragListeners.push(this.renderer.listen(this.nodeNativeElement, 'dragover', this.handleDragOver.bind(this)));
      this.disposersForDragListeners.push(this.renderer.listen(this.nodeNativeElement, 'dragstart', this.handleDragStart.bind(this)));
      this.disposersForDragListeners.push(this.renderer.listen(this.nodeNativeElement, 'dragleave', this.handleDragLeave.bind(this)));
      this.disposersForDragListeners.push(this.renderer.listen(this.nodeNativeElement, 'drop', this.handleDrop.bind(this)));
      this.disposersForDragListeners.push(this.renderer.listen(this.nodeNativeElement, 'dragend', this.handleDragEnd.bind(this)));
    }
  }

  public ngOnDestroy(): void {
    /* tslint:disable:typedef */
    this.disposersForDragListeners.forEach(dispose => dispose());
    /* tslint:enable:typedef */
  }

  private handleDragStart(e: DragEvent): any {
    e.stopPropagation();

    console.log('===', this.nodeDraggable, this.tree);
    this.nodeDraggableService.captureNode(new CapturedNode(this.nodeDraggable, this.tree));

    e.dataTransfer.setData('text', NodeDraggableDirective.DATA_TRANSFER_STUB_DATA);
    e.dataTransfer.effectAllowed = 'all';

    e.dataTransfer.setDragImage(this.nodeDraggable.nativeElement.querySelector('.value-container'), '10px', '10px');
  }

  private handleDragOver(e: DragEvent): any {
    e.preventDefault();
    if (e.shiftKey) {
      e.dataTransfer.dropEffect = 'copy';
      this.isCopy = true;
    } else {
      e.dataTransfer.dropEffect = 'move';
      this.isCopy = false;
    }

    let yOfCursor = e.offsetY;
    if (!yOfCursor) return ;
    if (!this.isDropPossible(e)) {
      return false;
    }

      this.removeClassForInsert();

    let tagHeight = this.element.nativeElement.offsetHeight;
    let space;

    let type = this.tree.node.type;
      if (type < 2) { // 文件夹
          space = tagHeight / 3;
          if (yOfCursor < space) {
              this.addClass('over-drop-target-before');
              this.mode = 'before';
          } else if (yOfCursor >= space && yOfCursor <= space * 2) {
              this.addClass('over-drop-target-inner');
              this.mode = 'inner';
          } else {
              this.addClass('over-drop-target-after');
              this.mode = 'after';
          }
      } else {
          space = tagHeight / 2;
          if (yOfCursor < space) {
              this.addClass('over-drop-target-before');
              this.mode = 'before';
          } else {
              this.addClass('over-drop-target-after');
              this.mode = 'after';
          }
      }
  }

  private handleDragEnter(e: DragEvent): any {
    e.preventDefault();
  }

  private handleDragLeave(e: DragEvent): any {
    if (!this.containsElementAt(e)) {
      this.removeClassForInsert();
    }
  }

  private handleDrop(e: DragEvent): any {
    e.preventDefault();
    e.stopPropagation();

    this.removeClassForInsert();

    if (!this.isDropPossible(e)) {
      return false;
    }

    if (this.nodeDraggableService.getCapturedNode()) {
      return this.notifyThatNodeWasDropped();
    }
  }

  private isDropPossible(e: DragEvent): boolean {
    const capturedNode = this.nodeDraggableService.getCapturedNode();
    return capturedNode
      && capturedNode.canBeDroppedAt(this.nodeDraggable)
      && this.containsElementAt(e);
  }

  private handleDragEnd(e: DragEvent): any {
    this.removeClassForInsert();
    this.nodeDraggableService.releaseCapturedNode();
  }

  private containsElementAt(e: DragEvent): boolean {
    const {x = e.clientX, y = e.clientY} = e;
    return this.nodeNativeElement.contains(document.elementFromPoint(x, y));
  }

  private addClass(className: string): void {
    const classList: DOMTokenList = this.nodeNativeElement.classList;
    classList.add(className);
  }
  private removeClassForInsert(): void {
    this.removeClass('over-drop-target-before');
    this.removeClass('over-drop-target-inner');
    this.removeClass('over-drop-target-after');
  }
  private removeClass(className: string): void {
    const classList: DOMTokenList = this.nodeNativeElement.classList;
    classList.remove(className);
  }

  private notifyThatNodeWasDropped(): void {
    this.nodeDraggableService.fireNodeDragged(this.nodeDraggableService.getCapturedNode(), this.nodeDraggable,
        this.mode, this.isCopy);
  }
}
